// Copyright 2021-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package clientcertissuer

import (
	"fmt"
	"strings"
	"time"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"go.pinniped.dev/internal/constable"
)

const defaultCertIssuerErr = constable.Error("failed to issue cert")

type ClientCertIssuer interface {
	Name() string
	IssueClientCertPEM(username string, groups []string, ttl time.Duration) (certPEM, keyPEM []byte, err error)
}

var _ ClientCertIssuer = ClientCertIssuers{}

type ClientCertIssuers []ClientCertIssuer

func (c ClientCertIssuers) Name() string {
	if len(c) == 0 {
		return "empty-client-cert-issuers"
	}

	names := make([]string, 0, len(c))
	for _, issuer := range c {
		names = append(names, issuer.Name())
	}

	return strings.Join(names, ",")
}

func (c ClientCertIssuers) IssueClientCertPEM(username string, groups []string, ttl time.Duration) ([]byte, []byte, error) {
	errs := make([]error, 0, len(c))

	for _, issuer := range c {
		certPEM, keyPEM, err := issuer.IssueClientCertPEM(username, groups, ttl)
		if err == nil {
			return certPEM, keyPEM, nil
		}
		errs = append(errs, fmt.Errorf("%s failed to issue client cert: %w", issuer.Name(), err))
	}

	if err := utilerrors.NewAggregate(errs); err != nil {
		return nil, nil, err
	}

	return nil, nil, defaultCertIssuerErr
}
