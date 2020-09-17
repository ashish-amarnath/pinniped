// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package issuerconfig

import (
	"encoding/base64"
	"fmt"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	crdpinnipedv1alpha1 "github.com/vmware-tanzu/pinniped/generated/1.19/apis/crdpinniped/v1alpha1"
	pinnipedclientset "github.com/vmware-tanzu/pinniped/generated/1.19/client/clientset/versioned"
	crdpinnipedv1alpha1informers "github.com/vmware-tanzu/pinniped/generated/1.19/client/informers/externalversions/crdpinniped/v1alpha1"
	pinnipedcontroller "github.com/vmware-tanzu/pinniped/internal/controller"
	"github.com/vmware-tanzu/pinniped/internal/controllerlib"
)

const (
	ClusterInfoNamespace = "kube-public"

	ConfigName = "pinniped-config"

	clusterInfoName         = "cluster-info"
	clusterInfoConfigMapKey = "kubeconfig"
)

type publisherController struct {
	namespace                      string
	serverOverride                 *string
	pinnipedClient                 pinnipedclientset.Interface
	configMapInformer              corev1informers.ConfigMapInformer
	credentialIssuerConfigInformer crdpinnipedv1alpha1informers.CredentialIssuerConfigInformer
}

func NewPublisherController(
	namespace string,
	serverOverride *string,
	pinnipedClient pinnipedclientset.Interface,
	configMapInformer corev1informers.ConfigMapInformer,
	credentialIssuerConfigInformer crdpinnipedv1alpha1informers.CredentialIssuerConfigInformer,
	withInformer pinnipedcontroller.WithInformerOptionFunc,
) controllerlib.Controller {
	return controllerlib.New(
		controllerlib.Config{
			Name: "publisher-controller",
			Syncer: &publisherController{
				namespace:                      namespace,
				serverOverride:                 serverOverride,
				pinnipedClient:                 pinnipedClient,
				configMapInformer:              configMapInformer,
				credentialIssuerConfigInformer: credentialIssuerConfigInformer,
			},
		},
		withInformer(
			configMapInformer,
			pinnipedcontroller.NameAndNamespaceExactMatchFilterFactory(clusterInfoName, ClusterInfoNamespace),
			controllerlib.InformerOption{},
		),
		withInformer(
			credentialIssuerConfigInformer,
			pinnipedcontroller.NameAndNamespaceExactMatchFilterFactory(ConfigName, namespace),
			controllerlib.InformerOption{},
		),
	)
}

func (c *publisherController) Sync(ctx controllerlib.Context) error {
	configMap, err := c.configMapInformer.
		Lister().
		ConfigMaps(ClusterInfoNamespace).
		Get(clusterInfoName)
	notFound := k8serrors.IsNotFound(err)
	if err != nil && !notFound {
		return fmt.Errorf("failed to get %s configmap: %w", clusterInfoName, err)
	}
	if notFound {
		klog.InfoS(
			"could not find config map",
			"configmap",
			klog.KRef(ClusterInfoNamespace, clusterInfoName),
		)
		return nil
	}

	kubeConfig, kubeConfigPresent := configMap.Data[clusterInfoConfigMapKey]
	if !kubeConfigPresent {
		klog.InfoS("could not find kubeconfig configmap key")
		return nil
	}

	config, err := clientcmd.Load([]byte(kubeConfig))
	if err != nil {
		klog.InfoS("could not load kubeconfig configmap key")
		return nil
	}

	var certificateAuthorityData, server string
	for _, v := range config.Clusters {
		certificateAuthorityData = base64.StdEncoding.EncodeToString(v.CertificateAuthorityData)
		server = v.Server
		break
	}

	if c.serverOverride != nil {
		server = *c.serverOverride
	}

	existingCredentialIssuerConfigFromInformerCache, err := c.credentialIssuerConfigInformer.
		Lister().
		CredentialIssuerConfigs(c.namespace).
		Get(ConfigName)
	notFound = k8serrors.IsNotFound(err)
	if err != nil && !notFound {
		return fmt.Errorf("could not get credentialissuerconfig: %w", err)
	}

	updateServerAndCAFunc := func(c *crdpinnipedv1alpha1.CredentialIssuerConfig) {
		c.Status.KubeConfigInfo = &crdpinnipedv1alpha1.CredentialIssuerConfigKubeConfigInfo{
			Server:                   server,
			CertificateAuthorityData: certificateAuthorityData,
		}
	}

	err = createOrUpdateCredentialIssuerConfig(
		ctx.Context,
		existingCredentialIssuerConfigFromInformerCache,
		notFound,
		ConfigName,
		c.namespace,
		c.pinnipedClient,
		updateServerAndCAFunc)

	if err != nil {
		return fmt.Errorf("could not create or update credentialissuerconfig: %w", err)
	}
	return nil
}
