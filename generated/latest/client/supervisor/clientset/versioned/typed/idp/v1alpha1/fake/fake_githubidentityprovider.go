// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/latest/apis/supervisor/idp/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGitHubIdentityProviders implements GitHubIdentityProviderInterface
type FakeGitHubIdentityProviders struct {
	Fake *FakeIDPV1alpha1
	ns   string
}

var githubidentityprovidersResource = v1alpha1.SchemeGroupVersion.WithResource("githubidentityproviders")

var githubidentityprovidersKind = v1alpha1.SchemeGroupVersion.WithKind("GitHubIdentityProvider")

// Get takes name of the gitHubIdentityProvider, and returns the corresponding gitHubIdentityProvider object, and an error if there is any.
func (c *FakeGitHubIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GitHubIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(githubidentityprovidersResource, c.ns, name), &v1alpha1.GitHubIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubIdentityProvider), err
}

// List takes label and field selectors, and returns the list of GitHubIdentityProviders that match those selectors.
func (c *FakeGitHubIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GitHubIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(githubidentityprovidersResource, githubidentityprovidersKind, c.ns, opts), &v1alpha1.GitHubIdentityProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GitHubIdentityProviderList{ListMeta: obj.(*v1alpha1.GitHubIdentityProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.GitHubIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitHubIdentityProviders.
func (c *FakeGitHubIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(githubidentityprovidersResource, c.ns, opts))

}

// Create takes the representation of a gitHubIdentityProvider and creates it.  Returns the server's representation of the gitHubIdentityProvider, and an error, if there is any.
func (c *FakeGitHubIdentityProviders) Create(ctx context.Context, gitHubIdentityProvider *v1alpha1.GitHubIdentityProvider, opts v1.CreateOptions) (result *v1alpha1.GitHubIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(githubidentityprovidersResource, c.ns, gitHubIdentityProvider), &v1alpha1.GitHubIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubIdentityProvider), err
}

// Update takes the representation of a gitHubIdentityProvider and updates it. Returns the server's representation of the gitHubIdentityProvider, and an error, if there is any.
func (c *FakeGitHubIdentityProviders) Update(ctx context.Context, gitHubIdentityProvider *v1alpha1.GitHubIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.GitHubIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(githubidentityprovidersResource, c.ns, gitHubIdentityProvider), &v1alpha1.GitHubIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitHubIdentityProviders) UpdateStatus(ctx context.Context, gitHubIdentityProvider *v1alpha1.GitHubIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.GitHubIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(githubidentityprovidersResource, "status", c.ns, gitHubIdentityProvider), &v1alpha1.GitHubIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubIdentityProvider), err
}

// Delete takes name of the gitHubIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeGitHubIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(githubidentityprovidersResource, c.ns, name, opts), &v1alpha1.GitHubIdentityProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitHubIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(githubidentityprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GitHubIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched gitHubIdentityProvider.
func (c *FakeGitHubIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GitHubIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(githubidentityprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.GitHubIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubIdentityProvider), err
}
