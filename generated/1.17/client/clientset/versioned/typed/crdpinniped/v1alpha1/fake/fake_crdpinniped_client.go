// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/vmware-tanzu/pinniped/generated/1.17/client/clientset/versioned/typed/crdpinniped/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrdV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCrdV1alpha1) CredentialIssuerConfigs(namespace string) v1alpha1.CredentialIssuerConfigInterface {
	return &FakeCredentialIssuerConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrdV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
