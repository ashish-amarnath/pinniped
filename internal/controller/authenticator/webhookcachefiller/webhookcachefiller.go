// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package webhookcachefiller implements a controller for filling an authncache.Cache with each added/updated WebhookAuthenticator.
package webhookcachefiller

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/url"
	"time"

	k8sauthv1beta1 "k8s.io/api/authentication/v1beta1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	errorsutil "k8s.io/apimachinery/pkg/util/errors"
	k8snetutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/plugin/pkg/authenticator/token/webhook"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"

	auth1alpha1 "go.pinniped.dev/generated/latest/apis/concierge/authentication/v1alpha1"
	conciergeclientset "go.pinniped.dev/generated/latest/client/concierge/clientset/versioned"
	authinformers "go.pinniped.dev/generated/latest/client/concierge/informers/externalversions/authentication/v1alpha1"
	pinnipedcontroller "go.pinniped.dev/internal/controller"
	"go.pinniped.dev/internal/controller/authenticator/authncache"
	"go.pinniped.dev/internal/controller/conditionsutil"
	"go.pinniped.dev/internal/controllerlib"
	"go.pinniped.dev/internal/crypto/ptls"
	"go.pinniped.dev/internal/endpointaddr"
	"go.pinniped.dev/internal/kubeclient"
	"go.pinniped.dev/internal/plog"
)

const (
	controllerName                   = "webhookcachefiller-controller"
	typeReady                        = "Ready"
	typeTLSConfigurationValid        = "TLSConfigurationValid"
	typeWebhookConnectionValid       = "WebhookConnectionValid"
	typeEndpointURLValid             = "EndpointURLValid"
	typeAuthenticatorValid           = "AuthenticatorValid"
	reasonSuccess                    = "Success"
	reasonNotReady                   = "NotReady"
	reasonUnableToValidate           = "UnableToValidate"
	reasonUnableToCreateClient       = "UnableToCreateClient"
	reasonUnableToInstantiateWebhook = "UnableToInstantiateWebhook"
	reasonInvalidTLSConfiguration    = "InvalidTLSConfiguration"
	reasonInvalidEndpointURL         = "InvalidEndpointURL"
	reasonInvalidEndpointURLScheme   = "InvalidEndpointURLScheme"
	reasonUnableToDialServer         = "UnableToDialServer"
	msgUnableToValidate              = "unable to validate; see other conditions for details"
)

// New instantiates a new controllerlib.Controller which will populate the provided authncache.Cache.
func New(
	cache *authncache.Cache,
	client conciergeclientset.Interface,
	webhooks authinformers.WebhookAuthenticatorInformer,
	clock clock.Clock,
	log plog.Logger,
) controllerlib.Controller {
	return controllerlib.New(
		controllerlib.Config{
			Name: controllerName,
			Syncer: &webhookCacheFillerController{
				cache:    cache,
				client:   client,
				webhooks: webhooks,
				clock:    clock,
				log:      log.WithName(controllerName),
			},
		},
		controllerlib.WithInformer(
			webhooks,
			pinnipedcontroller.MatchAnythingFilter(nil), // nil parent func is fine because each event is distinct
			controllerlib.InformerOption{},
		),
	)
}

type webhookCacheFillerController struct {
	cache    *authncache.Cache
	webhooks authinformers.WebhookAuthenticatorInformer
	client   conciergeclientset.Interface
	clock    clock.Clock
	log      plog.Logger
}

// Sync implements controllerlib.Syncer.
func (c *webhookCacheFillerController) Sync(ctx controllerlib.Context) error {
	obj, err := c.webhooks.Lister().Get(ctx.Key.Name)
	if err != nil && errors.IsNotFound(err) {
		c.log.Info("Sync() found that the WebhookAuthenticator does not exist yet or was deleted")
		return nil
	}
	if err != nil {
		// no unit test for this failure
		return fmt.Errorf("failed to get WebhookAuthenticator %s/%s: %w", ctx.Key.Namespace, ctx.Key.Name, err)
	}

	conditions := make([]*metav1.Condition, 0)
	var errs []error

	certPool, pemBytes, conditions, tlsBundleOk := c.validateTLSBundle(obj.Spec.TLS, conditions)
	endpointHostPort, conditions, endpointOk := c.validateEndpoint(obj.Spec.Endpoint, conditions)
	okSoFar := tlsBundleOk && endpointOk
	conditions, tlsNegotiateErr := c.validateConnection(certPool, endpointHostPort, conditions, okSoFar)
	errs = append(errs, tlsNegotiateErr)
	okSoFar = okSoFar && tlsNegotiateErr == nil

	webhookAuthenticator, conditions, err := newWebhookAuthenticator(
		// Note that we use the whole URL when constructing the webhook client,
		// not just the host and port that ew validated above. We need the path, etc.
		obj.Spec.Endpoint,
		pemBytes,
		conditions,
		okSoFar,
	)
	errs = append(errs, err)

	if !conditionsutil.HadErrorCondition(conditions) {
		c.cache.Store(authncache.Key{
			APIGroup: auth1alpha1.GroupName,
			Kind:     "WebhookAuthenticator",
			Name:     ctx.Key.Name,
		}, webhookAuthenticator)
		c.log.WithValues("webhook", klog.KObj(obj), "endpoint", obj.Spec.Endpoint).Info("added new webhook authenticator")
	}

	err = c.updateStatus(ctx.Context, obj, conditions)
	errs = append(errs, err)

	// sync loop errors:
	// - should not be configuration errors. config errors a user must correct belong on the .Status
	//   object. The controller simply must wait for a user to correct before running again.
	// - other errors, such as networking errors, etc. are the types of errors that should return here
	//   and signal the controller to retry the sync loop. These may be corrected by machines.
	return errorsutil.NewAggregate(errs)
}

// newWebhookAuthenticator creates a webhook from the provided API server url and caBundle
// used to validate TLS connections.
func newWebhookAuthenticator(
	endpointURL string,
	pemBytes []byte,
	conditions []*metav1.Condition,
	prereqOk bool,
) (*webhook.WebhookTokenAuthenticator, []*metav1.Condition, error) {
	if !prereqOk {
		conditions = append(conditions, &metav1.Condition{
			Type:    typeAuthenticatorValid,
			Status:  metav1.ConditionUnknown,
			Reason:  reasonUnableToValidate,
			Message: msgUnableToValidate,
		})
		return nil, conditions, nil
	}

	// We use v1beta1 instead of v1 since v1beta1 is more prevalent in our desired
	// integration points.
	version := k8sauthv1beta1.SchemeGroupVersion.Version

	// At the current time, we don't provide any audiences because we simply don't
	// have any requirements to do so. This can be changed in the future as
	// requirements change.
	var implicitAuds authenticator.Audiences

	// We set this to nil because we would only need this to support some of the
	// custom proxy stuff used by the API server.
	var customDial k8snetutil.DialFunc

	restConfig := &rest.Config{
		Host:            endpointURL,
		TLSClientConfig: rest.TLSClientConfig{CAData: pemBytes},

		// The remainder of these settings are copied from webhookutil.LoadKubeconfig in k8s.io/apiserver/pkg/util/webhook.
		Dial:    customDial,
		Timeout: 30 * time.Second,
		QPS:     -1,
	}

	client, err := kubeclient.New(kubeclient.WithConfig(restConfig), kubeclient.WithTLSConfigFunc(ptls.Default))
	if err != nil {
		errText := "unable to create client for this webhook"
		msg := fmt.Sprintf("%s: %s", errText, err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeAuthenticatorValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonUnableToCreateClient,
			Message: msg,
		})
		return nil, conditions, fmt.Errorf("%s: %w", errText, err)
	}

	webhookAuthenticator, err := webhook.New(client.JSONConfig, version, implicitAuds, *webhook.DefaultRetryBackoff())
	if err != nil {
		// no unit test for this failure.
		errText := "unable to instantiate webhook"
		msg := fmt.Sprintf("%s: %s", errText, err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeAuthenticatorValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonUnableToInstantiateWebhook,
			Message: msg,
		})
		return nil, conditions, fmt.Errorf("%s: %w", errText, err)
	}

	msg := "authenticator initialized"
	conditions = append(conditions, &metav1.Condition{
		Type:    typeAuthenticatorValid,
		Status:  metav1.ConditionTrue,
		Reason:  reasonSuccess,
		Message: msg,
	})

	return webhookAuthenticator, conditions, nil
}

func (c *webhookCacheFillerController) validateConnection(certPool *x509.CertPool, endpointHostPort *endpointaddr.HostPort, conditions []*metav1.Condition, prereqOk bool) ([]*metav1.Condition, error) {
	if !prereqOk {
		conditions = append(conditions, &metav1.Condition{
			Type:    typeWebhookConnectionValid,
			Status:  metav1.ConditionUnknown,
			Reason:  reasonUnableToValidate,
			Message: msgUnableToValidate,
		})
		return conditions, nil
	}

	conn, err := tls.Dial("tcp", endpointHostPort.Endpoint(), ptls.Default(certPool))

	if err != nil {
		errText := "cannot dial server"
		msg := fmt.Sprintf("%s: %s", errText, err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeWebhookConnectionValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonUnableToDialServer,
			Message: msg,
		})
		return conditions, fmt.Errorf("%s: %w", errText, err)
	}

	// this error should never be significant
	err = conn.Close()
	if err != nil {
		// no unit test for this failure
		c.log.Error("error closing dialer", err)
	}

	conditions = append(conditions, &metav1.Condition{
		Type:    typeWebhookConnectionValid,
		Status:  metav1.ConditionTrue,
		Reason:  reasonSuccess,
		Message: "successfully dialed webhook server",
	})
	return conditions, nil
}

func (c *webhookCacheFillerController) validateTLSBundle(tlsSpec *auth1alpha1.TLSSpec, conditions []*metav1.Condition) (*x509.CertPool, []byte, []*metav1.Condition, bool) {
	rootCAs, pemBytes, err := pinnipedcontroller.BuildCertPoolAuth(tlsSpec)
	if err != nil {
		msg := fmt.Sprintf("%s: %s", "invalid TLS configuration", err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeTLSConfigurationValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonInvalidTLSConfiguration,
			Message: msg,
		})
		return rootCAs, pemBytes, conditions, false
	}
	msg := "successfully parsed specified CA bundle"
	if rootCAs == nil {
		msg = "no CA bundle specified"
	}
	conditions = append(conditions, &metav1.Condition{
		Type:    typeTLSConfigurationValid,
		Status:  metav1.ConditionTrue,
		Reason:  reasonSuccess,
		Message: msg,
	})
	return rootCAs, pemBytes, conditions, true
}

func (c *webhookCacheFillerController) validateEndpoint(endpoint string, conditions []*metav1.Condition) (*endpointaddr.HostPort, []*metav1.Condition, bool) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		msg := fmt.Sprintf("%s: %s", "spec.endpoint URL cannot be parsed", err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeEndpointURLValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonInvalidEndpointURL,
			Message: msg,
		})
		return nil, conditions, false
	}

	// handles empty string and other issues as well.
	if endpointURL.Scheme != "https" {
		msg := fmt.Sprintf("spec.endpoint URL %s has invalid scheme, require 'https'", endpoint)
		conditions = append(conditions, &metav1.Condition{
			Type:    typeEndpointURLValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonInvalidEndpointURLScheme,
			Message: msg,
		})
		return nil, conditions, false
	}

	endpointHostPort, err := endpointaddr.ParseFromURL(endpointURL, 443)
	if err != nil {
		msg := fmt.Sprintf("%s: %s", "spec.endpoint URL is not valid", err.Error())
		conditions = append(conditions, &metav1.Condition{
			Type:    typeEndpointURLValid,
			Status:  metav1.ConditionFalse,
			Reason:  reasonInvalidEndpointURL,
			Message: msg,
		})
		return nil, conditions, false
	}

	conditions = append(conditions, &metav1.Condition{
		Type:    typeEndpointURLValid,
		Status:  metav1.ConditionTrue,
		Reason:  reasonSuccess,
		Message: "spec.endpoint is a valid URL",
	})
	return &endpointHostPort, conditions, true
}

func (c *webhookCacheFillerController) updateStatus(
	ctx context.Context,
	original *auth1alpha1.WebhookAuthenticator,
	conditions []*metav1.Condition,
) error {
	updated := original.DeepCopy()

	if conditionsutil.HadErrorCondition(conditions) {
		updated.Status.Phase = auth1alpha1.WebhookAuthenticatorPhaseError
		conditions = append(conditions, &metav1.Condition{
			Type:    typeReady,
			Status:  metav1.ConditionFalse,
			Reason:  reasonNotReady,
			Message: "the WebhookAuthenticator is not ready: see other conditions for details",
		})
	} else {
		updated.Status.Phase = auth1alpha1.WebhookAuthenticatorPhaseReady
		conditions = append(conditions, &metav1.Condition{
			Type:    typeReady,
			Status:  metav1.ConditionTrue,
			Reason:  reasonSuccess,
			Message: "the WebhookAuthenticator is ready",
		})
	}

	_ = conditionsutil.MergeConditions(
		conditions,
		original.Generation,
		&updated.Status.Conditions,
		plog.New().WithName(controllerName),
		metav1.NewTime(c.clock.Now()),
	)

	if equality.Semantic.DeepEqual(original, updated) {
		return nil
	}

	_, err := c.client.AuthenticationV1alpha1().WebhookAuthenticators().UpdateStatus(ctx, updated, metav1.UpdateOptions{})
	return err
}
