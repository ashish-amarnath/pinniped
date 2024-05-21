// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by MockGen. DO NOT EDIT.
// Source: go.pinniped.dev/internal/registry/credentialrequest (interfaces: TokenCredentialRequestAuthenticator)
//
// Generated by this command:
//
//	mockgen -destination=mockcredentialrequest.go -package=mockcredentialrequest -copyright_file=../../../hack/header.txt go.pinniped.dev/internal/registry/credentialrequest TokenCredentialRequestAuthenticator
//

// Package mockcredentialrequest is a generated GoMock package.
package mockcredentialrequest

import (
	context "context"
	reflect "reflect"

	login "go.pinniped.dev/generated/latest/apis/concierge/login"
	gomock "go.uber.org/mock/gomock"
	user "k8s.io/apiserver/pkg/authentication/user"
)

// MockTokenCredentialRequestAuthenticator is a mock of TokenCredentialRequestAuthenticator interface.
type MockTokenCredentialRequestAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCredentialRequestAuthenticatorMockRecorder
}

// MockTokenCredentialRequestAuthenticatorMockRecorder is the mock recorder for MockTokenCredentialRequestAuthenticator.
type MockTokenCredentialRequestAuthenticatorMockRecorder struct {
	mock *MockTokenCredentialRequestAuthenticator
}

// NewMockTokenCredentialRequestAuthenticator creates a new mock instance.
func NewMockTokenCredentialRequestAuthenticator(ctrl *gomock.Controller) *MockTokenCredentialRequestAuthenticator {
	mock := &MockTokenCredentialRequestAuthenticator{ctrl: ctrl}
	mock.recorder = &MockTokenCredentialRequestAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenCredentialRequestAuthenticator) EXPECT() *MockTokenCredentialRequestAuthenticatorMockRecorder {
	return m.recorder
}

// AuthenticateTokenCredentialRequest mocks base method.
func (m *MockTokenCredentialRequestAuthenticator) AuthenticateTokenCredentialRequest(arg0 context.Context, arg1 *login.TokenCredentialRequest) (user.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateTokenCredentialRequest", arg0, arg1)
	ret0, _ := ret[0].(user.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateTokenCredentialRequest indicates an expected call of AuthenticateTokenCredentialRequest.
func (mr *MockTokenCredentialRequestAuthenticatorMockRecorder) AuthenticateTokenCredentialRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateTokenCredentialRequest", reflect.TypeOf((*MockTokenCredentialRequestAuthenticator)(nil).AuthenticateTokenCredentialRequest), arg0, arg1)
}
