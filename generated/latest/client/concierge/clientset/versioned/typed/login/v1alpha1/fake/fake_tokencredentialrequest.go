// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/latest/apis/concierge/login/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
)

// FakeTokenCredentialRequests implements TokenCredentialRequestInterface
type FakeTokenCredentialRequests struct {
	Fake *FakeLoginV1alpha1
}

var tokencredentialrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("tokencredentialrequests")

var tokencredentialrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("TokenCredentialRequest")

// Create takes the representation of a tokenCredentialRequest and creates it.  Returns the server's representation of the tokenCredentialRequest, and an error, if there is any.
func (c *FakeTokenCredentialRequests) Create(ctx context.Context, tokenCredentialRequest *v1alpha1.TokenCredentialRequest, opts v1.CreateOptions) (result *v1alpha1.TokenCredentialRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tokencredentialrequestsResource, tokenCredentialRequest), &v1alpha1.TokenCredentialRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TokenCredentialRequest), err
}
