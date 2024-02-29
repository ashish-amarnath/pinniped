// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.28/apis/concierge/authentication/v1alpha1"
	scheme "go.pinniped.dev/generated/1.28/client/concierge/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JWTAuthenticatorsGetter has a method to return a JWTAuthenticatorInterface.
// A group's client should implement this interface.
type JWTAuthenticatorsGetter interface {
	JWTAuthenticators() JWTAuthenticatorInterface
}

// JWTAuthenticatorInterface has methods to work with JWTAuthenticator resources.
type JWTAuthenticatorInterface interface {
	Create(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.CreateOptions) (*v1alpha1.JWTAuthenticator, error)
	Update(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.UpdateOptions) (*v1alpha1.JWTAuthenticator, error)
	UpdateStatus(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.UpdateOptions) (*v1alpha1.JWTAuthenticator, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JWTAuthenticator, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JWTAuthenticatorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JWTAuthenticator, err error)
	JWTAuthenticatorExpansion
}

// jWTAuthenticators implements JWTAuthenticatorInterface
type jWTAuthenticators struct {
	client rest.Interface
}

// newJWTAuthenticators returns a JWTAuthenticators
func newJWTAuthenticators(c *AuthenticationV1alpha1Client) *jWTAuthenticators {
	return &jWTAuthenticators{
		client: c.RESTClient(),
	}
}

// Get takes name of the jWTAuthenticator, and returns the corresponding jWTAuthenticator object, and an error if there is any.
func (c *jWTAuthenticators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Get().
		Resource("jwtauthenticators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JWTAuthenticators that match those selectors.
func (c *jWTAuthenticators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JWTAuthenticatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.JWTAuthenticatorList{}
	err = c.client.Get().
		Resource("jwtauthenticators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jWTAuthenticators.
func (c *jWTAuthenticators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("jwtauthenticators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jWTAuthenticator and creates it.  Returns the server's representation of the jWTAuthenticator, and an error, if there is any.
func (c *jWTAuthenticators) Create(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.CreateOptions) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Post().
		Resource("jwtauthenticators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jWTAuthenticator).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jWTAuthenticator and updates it. Returns the server's representation of the jWTAuthenticator, and an error, if there is any.
func (c *jWTAuthenticators) Update(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.UpdateOptions) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Put().
		Resource("jwtauthenticators").
		Name(jWTAuthenticator.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jWTAuthenticator).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jWTAuthenticators) UpdateStatus(ctx context.Context, jWTAuthenticator *v1alpha1.JWTAuthenticator, opts v1.UpdateOptions) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Put().
		Resource("jwtauthenticators").
		Name(jWTAuthenticator.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jWTAuthenticator).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jWTAuthenticator and deletes it. Returns an error if one occurs.
func (c *jWTAuthenticators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("jwtauthenticators").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jWTAuthenticators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("jwtauthenticators").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jWTAuthenticator.
func (c *jWTAuthenticators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Patch(pt).
		Resource("jwtauthenticators").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
