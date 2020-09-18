// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.17/apis/idp/v1alpha1"
	scheme "go.pinniped.dev/generated/1.17/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WebhookIdentityProvidersGetter has a method to return a WebhookIdentityProviderInterface.
// A group's client should implement this interface.
type WebhookIdentityProvidersGetter interface {
	WebhookIdentityProviders(namespace string) WebhookIdentityProviderInterface
}

// WebhookIdentityProviderInterface has methods to work with WebhookIdentityProvider resources.
type WebhookIdentityProviderInterface interface {
	Create(*v1alpha1.WebhookIdentityProvider) (*v1alpha1.WebhookIdentityProvider, error)
	Update(*v1alpha1.WebhookIdentityProvider) (*v1alpha1.WebhookIdentityProvider, error)
	UpdateStatus(*v1alpha1.WebhookIdentityProvider) (*v1alpha1.WebhookIdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WebhookIdentityProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.WebhookIdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WebhookIdentityProvider, err error)
	WebhookIdentityProviderExpansion
}

// webhookIdentityProviders implements WebhookIdentityProviderInterface
type webhookIdentityProviders struct {
	client rest.Interface
	ns     string
}

// newWebhookIdentityProviders returns a WebhookIdentityProviders
func newWebhookIdentityProviders(c *IDPV1alpha1Client, namespace string) *webhookIdentityProviders {
	return &webhookIdentityProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the webhookIdentityProvider, and returns the corresponding webhookIdentityProvider object, and an error if there is any.
func (c *webhookIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.WebhookIdentityProvider, err error) {
	result = &v1alpha1.WebhookIdentityProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WebhookIdentityProviders that match those selectors.
func (c *webhookIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.WebhookIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WebhookIdentityProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested webhookIdentityProviders.
func (c *webhookIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a webhookIdentityProvider and creates it.  Returns the server's representation of the webhookIdentityProvider, and an error, if there is any.
func (c *webhookIdentityProviders) Create(webhookIdentityProvider *v1alpha1.WebhookIdentityProvider) (result *v1alpha1.WebhookIdentityProvider, err error) {
	result = &v1alpha1.WebhookIdentityProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		Body(webhookIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a webhookIdentityProvider and updates it. Returns the server's representation of the webhookIdentityProvider, and an error, if there is any.
func (c *webhookIdentityProviders) Update(webhookIdentityProvider *v1alpha1.WebhookIdentityProvider) (result *v1alpha1.WebhookIdentityProvider, err error) {
	result = &v1alpha1.WebhookIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		Name(webhookIdentityProvider.Name).
		Body(webhookIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *webhookIdentityProviders) UpdateStatus(webhookIdentityProvider *v1alpha1.WebhookIdentityProvider) (result *v1alpha1.WebhookIdentityProvider, err error) {
	result = &v1alpha1.WebhookIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		Name(webhookIdentityProvider.Name).
		SubResource("status").
		Body(webhookIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the webhookIdentityProvider and deletes it. Returns an error if one occurs.
func (c *webhookIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *webhookIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched webhookIdentityProvider.
func (c *webhookIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WebhookIdentityProvider, err error) {
	result = &v1alpha1.WebhookIdentityProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("webhookidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}