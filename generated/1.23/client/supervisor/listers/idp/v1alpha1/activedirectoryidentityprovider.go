// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.23/apis/supervisor/idp/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ActiveDirectoryIdentityProviderLister helps list ActiveDirectoryIdentityProviders.
// All objects returned here must be treated as read-only.
type ActiveDirectoryIdentityProviderLister interface {
	// List lists all ActiveDirectoryIdentityProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActiveDirectoryIdentityProvider, err error)
	// ActiveDirectoryIdentityProviders returns an object that can list and get ActiveDirectoryIdentityProviders.
	ActiveDirectoryIdentityProviders(namespace string) ActiveDirectoryIdentityProviderNamespaceLister
	ActiveDirectoryIdentityProviderListerExpansion
}

// activeDirectoryIdentityProviderLister implements the ActiveDirectoryIdentityProviderLister interface.
type activeDirectoryIdentityProviderLister struct {
	indexer cache.Indexer
}

// NewActiveDirectoryIdentityProviderLister returns a new ActiveDirectoryIdentityProviderLister.
func NewActiveDirectoryIdentityProviderLister(indexer cache.Indexer) ActiveDirectoryIdentityProviderLister {
	return &activeDirectoryIdentityProviderLister{indexer: indexer}
}

// List lists all ActiveDirectoryIdentityProviders in the indexer.
func (s *activeDirectoryIdentityProviderLister) List(selector labels.Selector) (ret []*v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActiveDirectoryIdentityProvider))
	})
	return ret, err
}

// ActiveDirectoryIdentityProviders returns an object that can list and get ActiveDirectoryIdentityProviders.
func (s *activeDirectoryIdentityProviderLister) ActiveDirectoryIdentityProviders(namespace string) ActiveDirectoryIdentityProviderNamespaceLister {
	return activeDirectoryIdentityProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActiveDirectoryIdentityProviderNamespaceLister helps list and get ActiveDirectoryIdentityProviders.
// All objects returned here must be treated as read-only.
type ActiveDirectoryIdentityProviderNamespaceLister interface {
	// List lists all ActiveDirectoryIdentityProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActiveDirectoryIdentityProvider, err error)
	// Get retrieves the ActiveDirectoryIdentityProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ActiveDirectoryIdentityProvider, error)
	ActiveDirectoryIdentityProviderNamespaceListerExpansion
}

// activeDirectoryIdentityProviderNamespaceLister implements the ActiveDirectoryIdentityProviderNamespaceLister
// interface.
type activeDirectoryIdentityProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ActiveDirectoryIdentityProviders in the indexer for a given namespace.
func (s activeDirectoryIdentityProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ActiveDirectoryIdentityProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActiveDirectoryIdentityProvider))
	})
	return ret, err
}

// Get retrieves the ActiveDirectoryIdentityProvider from the indexer for a given namespace and name.
func (s activeDirectoryIdentityProviderNamespaceLister) Get(name string) (*v1alpha1.ActiveDirectoryIdentityProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("activedirectoryidentityprovider"), name)
	}
	return obj.(*v1alpha1.ActiveDirectoryIdentityProvider), nil
}
