// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.pinniped.dev/generated/latest/apis/concierge/config/v1alpha1"
	scheme "go.pinniped.dev/generated/latest/client/concierge/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CredentialIssuersGetter has a method to return a CredentialIssuerInterface.
// A group's client should implement this interface.
type CredentialIssuersGetter interface {
	CredentialIssuers() CredentialIssuerInterface
}

// CredentialIssuerInterface has methods to work with CredentialIssuer resources.
type CredentialIssuerInterface interface {
	Create(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.CreateOptions) (*v1alpha1.CredentialIssuer, error)
	Update(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (*v1alpha1.CredentialIssuer, error)
	UpdateStatus(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (*v1alpha1.CredentialIssuer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CredentialIssuer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CredentialIssuerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CredentialIssuer, err error)
	CredentialIssuerExpansion
}

// credentialIssuers implements CredentialIssuerInterface
type credentialIssuers struct {
	client rest.Interface
}

// newCredentialIssuers returns a CredentialIssuers
func newCredentialIssuers(c *ConfigV1alpha1Client) *credentialIssuers {
	return &credentialIssuers{
		client: c.RESTClient(),
	}
}

// Get takes name of the credentialIssuer, and returns the corresponding credentialIssuer object, and an error if there is any.
func (c *credentialIssuers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CredentialIssuer, err error) {
	result = &v1alpha1.CredentialIssuer{}
	err = c.client.Get().
		Resource("credentialissuers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CredentialIssuers that match those selectors.
func (c *credentialIssuers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CredentialIssuerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CredentialIssuerList{}
	err = c.client.Get().
		Resource("credentialissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested credentialIssuers.
func (c *credentialIssuers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("credentialissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a credentialIssuer and creates it.  Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *credentialIssuers) Create(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.CreateOptions) (result *v1alpha1.CredentialIssuer, err error) {
	result = &v1alpha1.CredentialIssuer{}
	err = c.client.Post().
		Resource("credentialissuers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(credentialIssuer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a credentialIssuer and updates it. Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *credentialIssuers) Update(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (result *v1alpha1.CredentialIssuer, err error) {
	result = &v1alpha1.CredentialIssuer{}
	err = c.client.Put().
		Resource("credentialissuers").
		Name(credentialIssuer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(credentialIssuer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *credentialIssuers) UpdateStatus(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (result *v1alpha1.CredentialIssuer, err error) {
	result = &v1alpha1.CredentialIssuer{}
	err = c.client.Put().
		Resource("credentialissuers").
		Name(credentialIssuer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(credentialIssuer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the credentialIssuer and deletes it. Returns an error if one occurs.
func (c *credentialIssuers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("credentialissuers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *credentialIssuers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("credentialissuers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched credentialIssuer.
func (c *credentialIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CredentialIssuer, err error) {
	result = &v1alpha1.CredentialIssuer{}
	err = c.client.Patch(pt).
		Resource("credentialissuers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
