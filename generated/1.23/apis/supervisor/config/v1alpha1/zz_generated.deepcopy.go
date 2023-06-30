//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020-2023 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomain) DeepCopyInto(out *FederationDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomain.
func (in *FederationDomain) DeepCopy() *FederationDomain {
	if in == nil {
		return nil
	}
	out := new(FederationDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederationDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainIdentityProvider) DeepCopyInto(out *FederationDomainIdentityProvider) {
	*out = *in
	in.ObjectRef.DeepCopyInto(&out.ObjectRef)
	in.Transforms.DeepCopyInto(&out.Transforms)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainIdentityProvider.
func (in *FederationDomainIdentityProvider) DeepCopy() *FederationDomainIdentityProvider {
	if in == nil {
		return nil
	}
	out := new(FederationDomainIdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainList) DeepCopyInto(out *FederationDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederationDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainList.
func (in *FederationDomainList) DeepCopy() *FederationDomainList {
	if in == nil {
		return nil
	}
	out := new(FederationDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederationDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainSecrets) DeepCopyInto(out *FederationDomainSecrets) {
	*out = *in
	out.JWKS = in.JWKS
	out.TokenSigningKey = in.TokenSigningKey
	out.StateSigningKey = in.StateSigningKey
	out.StateEncryptionKey = in.StateEncryptionKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainSecrets.
func (in *FederationDomainSecrets) DeepCopy() *FederationDomainSecrets {
	if in == nil {
		return nil
	}
	out := new(FederationDomainSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainSpec) DeepCopyInto(out *FederationDomainSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(FederationDomainTLSSpec)
		**out = **in
	}
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]FederationDomainIdentityProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainSpec.
func (in *FederationDomainSpec) DeepCopy() *FederationDomainSpec {
	if in == nil {
		return nil
	}
	out := new(FederationDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainStatus) DeepCopyInto(out *FederationDomainStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Secrets = in.Secrets
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainStatus.
func (in *FederationDomainStatus) DeepCopy() *FederationDomainStatus {
	if in == nil {
		return nil
	}
	out := new(FederationDomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTLSSpec) DeepCopyInto(out *FederationDomainTLSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTLSSpec.
func (in *FederationDomainTLSSpec) DeepCopy() *FederationDomainTLSSpec {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTransforms) DeepCopyInto(out *FederationDomainTransforms) {
	*out = *in
	if in.Constants != nil {
		in, out := &in.Constants, &out.Constants
		*out = make([]FederationDomainTransformsConstant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Expressions != nil {
		in, out := &in.Expressions, &out.Expressions
		*out = make([]FederationDomainTransformsExpression, len(*in))
		copy(*out, *in)
	}
	if in.Examples != nil {
		in, out := &in.Examples, &out.Examples
		*out = make([]FederationDomainTransformsExample, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTransforms.
func (in *FederationDomainTransforms) DeepCopy() *FederationDomainTransforms {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTransforms)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTransformsConstant) DeepCopyInto(out *FederationDomainTransformsConstant) {
	*out = *in
	if in.StringListValue != nil {
		in, out := &in.StringListValue, &out.StringListValue
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTransformsConstant.
func (in *FederationDomainTransformsConstant) DeepCopy() *FederationDomainTransformsConstant {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTransformsConstant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTransformsExample) DeepCopyInto(out *FederationDomainTransformsExample) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Expects.DeepCopyInto(&out.Expects)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTransformsExample.
func (in *FederationDomainTransformsExample) DeepCopy() *FederationDomainTransformsExample {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTransformsExample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTransformsExampleExpects) DeepCopyInto(out *FederationDomainTransformsExampleExpects) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTransformsExampleExpects.
func (in *FederationDomainTransformsExampleExpects) DeepCopy() *FederationDomainTransformsExampleExpects {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTransformsExampleExpects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederationDomainTransformsExpression) DeepCopyInto(out *FederationDomainTransformsExpression) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederationDomainTransformsExpression.
func (in *FederationDomainTransformsExpression) DeepCopy() *FederationDomainTransformsExpression {
	if in == nil {
		return nil
	}
	out := new(FederationDomainTransformsExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCClient) DeepCopyInto(out *OIDCClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCClient.
func (in *OIDCClient) DeepCopy() *OIDCClient {
	if in == nil {
		return nil
	}
	out := new(OIDCClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OIDCClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCClientList) DeepCopyInto(out *OIDCClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OIDCClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCClientList.
func (in *OIDCClientList) DeepCopy() *OIDCClientList {
	if in == nil {
		return nil
	}
	out := new(OIDCClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OIDCClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCClientSpec) DeepCopyInto(out *OIDCClientSpec) {
	*out = *in
	if in.AllowedRedirectURIs != nil {
		in, out := &in.AllowedRedirectURIs, &out.AllowedRedirectURIs
		*out = make([]RedirectURI, len(*in))
		copy(*out, *in)
	}
	if in.AllowedGrantTypes != nil {
		in, out := &in.AllowedGrantTypes, &out.AllowedGrantTypes
		*out = make([]GrantType, len(*in))
		copy(*out, *in)
	}
	if in.AllowedScopes != nil {
		in, out := &in.AllowedScopes, &out.AllowedScopes
		*out = make([]Scope, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCClientSpec.
func (in *OIDCClientSpec) DeepCopy() *OIDCClientSpec {
	if in == nil {
		return nil
	}
	out := new(OIDCClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCClientStatus) DeepCopyInto(out *OIDCClientStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCClientStatus.
func (in *OIDCClientStatus) DeepCopy() *OIDCClientStatus {
	if in == nil {
		return nil
	}
	out := new(OIDCClientStatus)
	in.DeepCopyInto(out)
	return out
}
