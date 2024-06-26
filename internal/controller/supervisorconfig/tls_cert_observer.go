// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package supervisorconfig

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1informers "k8s.io/client-go/informers/core/v1"

	"go.pinniped.dev/generated/latest/client/supervisor/informers/externalversions/config/v1alpha1"
	pinnipedcontroller "go.pinniped.dev/internal/controller"
	"go.pinniped.dev/internal/controllerlib"
	"go.pinniped.dev/internal/plog"
)

type tlsCertObserverController struct {
	issuerTLSCertSetter             IssuerTLSCertSetter
	defaultTLSCertificateSecretName string
	federationDomainInformer        v1alpha1.FederationDomainInformer
	secretInformer                  corev1informers.SecretInformer
}

type IssuerTLSCertSetter interface {
	SetIssuerHostToTLSCertMap(issuerHostToTLSCertMap map[string]*tls.Certificate)
	SetDefaultTLSCert(certificate *tls.Certificate)
}

func NewTLSCertObserverController(
	issuerTLSCertSetter IssuerTLSCertSetter,
	defaultTLSCertificateSecretName string,
	secretInformer corev1informers.SecretInformer,
	federationDomainInformer v1alpha1.FederationDomainInformer,
	withInformer pinnipedcontroller.WithInformerOptionFunc,
) controllerlib.Controller {
	return controllerlib.New(
		controllerlib.Config{
			Name: "tls-certs-observer-controller",
			Syncer: &tlsCertObserverController{
				issuerTLSCertSetter:             issuerTLSCertSetter,
				defaultTLSCertificateSecretName: defaultTLSCertificateSecretName,
				federationDomainInformer:        federationDomainInformer,
				secretInformer:                  secretInformer,
			},
		},
		withInformer(
			secretInformer,
			pinnipedcontroller.MatchAnySecretOfTypeFilter(corev1.SecretTypeTLS, nil),
			controllerlib.InformerOption{},
		),
		withInformer(
			federationDomainInformer,
			pinnipedcontroller.MatchAnythingFilter(nil),
			controllerlib.InformerOption{},
		),
	)
}

func (c *tlsCertObserverController) Sync(ctx controllerlib.Context) error {
	ns := ctx.Key.Namespace
	allProviders, err := c.federationDomainInformer.Lister().FederationDomains(ns).List(labels.Everything())
	if err != nil {
		return fmt.Errorf("failed to list FederationDomains: %w", err)
	}

	// Rebuild the whole map on any change to any Secret or FederationDomain, because either can have changes that
	// can cause the map to need to be updated.
	issuerHostToTLSCertMap := map[string]*tls.Certificate{}

	for _, provider := range allProviders {
		issuerURL, err := url.Parse(provider.Spec.Issuer)
		if err != nil {
			plog.Debug("tlsCertObserverController Sync found an invalid issuer URL", "namespace", ns, "issuer", provider.Spec.Issuer)
			continue
		}

		secretName := ""
		if provider.Spec.TLS != nil {
			secretName = provider.Spec.TLS.SecretName
		}
		if secretName == "" {
			// No secret name provided, so no need to try to load any Secret.
			continue
		}

		certFromSecret, err := c.certFromSecret(ns, secretName)
		if err != nil {
			// The user configured a TLS secret on the FederationDomain but it could not be loaded,
			// so log a message which is visible at the default log level. Any error here indicates a problem,
			// including "not found" errors.
			plog.Error("error loading TLS certificate from Secret for FederationDomain.spec.tls.secretName", err,
				"FederationDomain.metadata.name", provider.Name,
				"FederationDomain.spec.tls.secretName", provider.Spec.TLS.SecretName,
			)
			continue
		}

		// Lowercase the host part of the URL because hostnames should be treated as case-insensitive.
		issuerHostToTLSCertMap[lowercaseHostWithoutPort(issuerURL)] = certFromSecret
	}

	plog.Debug("tlsCertObserverController Sync updated the TLS cert cache", "issuerHostCount", len(issuerHostToTLSCertMap))
	c.issuerTLSCertSetter.SetIssuerHostToTLSCertMap(issuerHostToTLSCertMap)

	defaultCert, err := c.certFromSecret(ns, c.defaultTLSCertificateSecretName)
	if err != nil {
		c.issuerTLSCertSetter.SetDefaultTLSCert(nil)
		// It's okay if the default TLS cert Secret is not found (it is not required).
		if !apierrors.IsNotFound(err) {
			// For any other error, log a message which is visible at the default log level.
			plog.Error("error loading TLS certificate from Secret for Supervisor default TLS cert", err,
				"defaultCertSecretName", c.defaultTLSCertificateSecretName,
			)
		}
	} else {
		c.issuerTLSCertSetter.SetDefaultTLSCert(defaultCert)
	}

	return nil
}

func (c *tlsCertObserverController) certFromSecret(ns string, secretName string) (*tls.Certificate, error) {
	tlsSecret, err := c.secretInformer.Lister().Secrets(ns).Get(secretName)
	if err != nil {
		plog.Debug("tlsCertObserverController Sync could not find TLS cert secret", "namespace", ns, "secretName", secretName)
		return nil, err
	}
	certFromSecret, err := tls.X509KeyPair(tlsSecret.Data["tls.crt"], tlsSecret.Data["tls.key"])
	if err != nil {
		plog.Debug("tlsCertObserverController Sync found a TLS secret with Data in an unexpected format", "namespace", ns, "secretName", secretName)
		return nil, err
	}
	return &certFromSecret, nil
}

func lowercaseHostWithoutPort(issuerURL *url.URL) string {
	lowercaseHost := strings.ToLower(issuerURL.Hostname())
	return lowercaseHost
}
