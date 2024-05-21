// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by MockGen. DO NOT EDIT.
// Source: go.pinniped.dev/internal/controller/kubecertagent (interfaces: PodCommandExecutor)
//
// Generated by this command:
//
//	mockgen -destination=mockpodcommandexecutor.go -package=mocks -copyright_file=../../../hack/header.txt go.pinniped.dev/internal/controller/kubecertagent PodCommandExecutor
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPodCommandExecutor is a mock of PodCommandExecutor interface.
type MockPodCommandExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockPodCommandExecutorMockRecorder
}

// MockPodCommandExecutorMockRecorder is the mock recorder for MockPodCommandExecutor.
type MockPodCommandExecutorMockRecorder struct {
	mock *MockPodCommandExecutor
}

// NewMockPodCommandExecutor creates a new mock instance.
func NewMockPodCommandExecutor(ctrl *gomock.Controller) *MockPodCommandExecutor {
	mock := &MockPodCommandExecutor{ctrl: ctrl}
	mock.recorder = &MockPodCommandExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodCommandExecutor) EXPECT() *MockPodCommandExecutorMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockPodCommandExecutor) Exec(arg0 context.Context, arg1, arg2, arg3 string, arg4 ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPodCommandExecutorMockRecorder) Exec(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPodCommandExecutor)(nil).Exec), varargs...)
}
