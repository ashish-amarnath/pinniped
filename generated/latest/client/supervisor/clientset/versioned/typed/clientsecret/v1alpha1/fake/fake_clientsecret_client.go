// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "go.pinniped.dev/generated/latest/client/supervisor/clientset/versioned/typed/clientsecret/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeClientsecretV1alpha1 struct {
	*testing.Fake
}

func (c *FakeClientsecretV1alpha1) OIDCClientSecretRequests(namespace string) v1alpha1.OIDCClientSecretRequestInterface {
	return &FakeOIDCClientSecretRequests{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeClientsecretV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
