// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configv1alpha1 "go.pinniped.dev/generated/1.23/apis/concierge/config/v1alpha1"
	versioned "go.pinniped.dev/generated/1.23/client/concierge/clientset/versioned"
	internalinterfaces "go.pinniped.dev/generated/1.23/client/concierge/informers/externalversions/internalinterfaces"
	v1alpha1 "go.pinniped.dev/generated/1.23/client/concierge/listers/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CredentialIssuerInformer provides access to a shared informer and lister for
// CredentialIssuers.
type CredentialIssuerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CredentialIssuerLister
}

type credentialIssuerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCredentialIssuerInformer constructs a new informer for CredentialIssuer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCredentialIssuerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCredentialIssuerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCredentialIssuerInformer constructs a new informer for CredentialIssuer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCredentialIssuerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().CredentialIssuers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().CredentialIssuers().Watch(context.TODO(), options)
			},
		},
		&configv1alpha1.CredentialIssuer{},
		resyncPeriod,
		indexers,
	)
}

func (f *credentialIssuerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCredentialIssuerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *credentialIssuerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha1.CredentialIssuer{}, f.defaultInformer)
}

func (f *credentialIssuerInformer) Lister() v1alpha1.CredentialIssuerLister {
	return v1alpha1.NewCredentialIssuerLister(f.Informer().GetIndexer())
}
