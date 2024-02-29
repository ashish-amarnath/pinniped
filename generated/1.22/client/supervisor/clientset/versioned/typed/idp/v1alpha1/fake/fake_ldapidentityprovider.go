// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.22/apis/supervisor/idp/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLDAPIdentityProviders implements LDAPIdentityProviderInterface
type FakeLDAPIdentityProviders struct {
	Fake *FakeIDPV1alpha1
	ns   string
}

var ldapidentityprovidersResource = schema.GroupVersionResource{Group: "idp.supervisor.pinniped.dev", Version: "v1alpha1", Resource: "ldapidentityproviders"}

var ldapidentityprovidersKind = schema.GroupVersionKind{Group: "idp.supervisor.pinniped.dev", Version: "v1alpha1", Kind: "LDAPIdentityProvider"}

// Get takes name of the lDAPIdentityProvider, and returns the corresponding lDAPIdentityProvider object, and an error if there is any.
func (c *FakeLDAPIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LDAPIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ldapidentityprovidersResource, c.ns, name), &v1alpha1.LDAPIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), err
}

// List takes label and field selectors, and returns the list of LDAPIdentityProviders that match those selectors.
func (c *FakeLDAPIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LDAPIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ldapidentityprovidersResource, ldapidentityprovidersKind, c.ns, opts), &v1alpha1.LDAPIdentityProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LDAPIdentityProviderList{ListMeta: obj.(*v1alpha1.LDAPIdentityProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.LDAPIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lDAPIdentityProviders.
func (c *FakeLDAPIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ldapidentityprovidersResource, c.ns, opts))

}

// Create takes the representation of a lDAPIdentityProvider and creates it.  Returns the server's representation of the lDAPIdentityProvider, and an error, if there is any.
func (c *FakeLDAPIdentityProviders) Create(ctx context.Context, lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider, opts v1.CreateOptions) (result *v1alpha1.LDAPIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ldapidentityprovidersResource, c.ns, lDAPIdentityProvider), &v1alpha1.LDAPIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), err
}

// Update takes the representation of a lDAPIdentityProvider and updates it. Returns the server's representation of the lDAPIdentityProvider, and an error, if there is any.
func (c *FakeLDAPIdentityProviders) Update(ctx context.Context, lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.LDAPIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ldapidentityprovidersResource, c.ns, lDAPIdentityProvider), &v1alpha1.LDAPIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLDAPIdentityProviders) UpdateStatus(ctx context.Context, lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.LDAPIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ldapidentityprovidersResource, "status", c.ns, lDAPIdentityProvider), &v1alpha1.LDAPIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), err
}

// Delete takes name of the lDAPIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeLDAPIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ldapidentityprovidersResource, c.ns, name), &v1alpha1.LDAPIdentityProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLDAPIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ldapidentityprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LDAPIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched lDAPIdentityProvider.
func (c *FakeLDAPIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LDAPIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ldapidentityprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.LDAPIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LDAPIdentityProvider), err
}
