// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "go.pinniped.dev/generated/1.25/client/concierge/clientset/versioned/typed/login/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLoginV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLoginV1alpha1) TokenCredentialRequests() v1alpha1.TokenCredentialRequestInterface {
	return &FakeTokenCredentialRequests{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLoginV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
