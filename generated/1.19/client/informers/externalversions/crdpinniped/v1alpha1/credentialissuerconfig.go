// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	crdpinnipedv1alpha1 "github.com/vmware-tanzu/pinniped/generated/1.19/apis/crdpinniped/v1alpha1"
	versioned "github.com/vmware-tanzu/pinniped/generated/1.19/client/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/pinniped/generated/1.19/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware-tanzu/pinniped/generated/1.19/client/listers/crdpinniped/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CredentialIssuerConfigInformer provides access to a shared informer and lister for
// CredentialIssuerConfigs.
type CredentialIssuerConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CredentialIssuerConfigLister
}

type credentialIssuerConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCredentialIssuerConfigInformer constructs a new informer for CredentialIssuerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCredentialIssuerConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCredentialIssuerConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCredentialIssuerConfigInformer constructs a new informer for CredentialIssuerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCredentialIssuerConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().CredentialIssuerConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().CredentialIssuerConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&crdpinnipedv1alpha1.CredentialIssuerConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *credentialIssuerConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCredentialIssuerConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *credentialIssuerConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdpinnipedv1alpha1.CredentialIssuerConfig{}, f.defaultInformer)
}

func (f *credentialIssuerConfigInformer) Lister() v1alpha1.CredentialIssuerConfigLister {
	return v1alpha1.NewCredentialIssuerConfigLister(f.Informer().GetIndexer())
}
