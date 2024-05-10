// Copyright 2021-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package deploymentref

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	kubefake "k8s.io/client-go/kubernetes/fake"
	kubetesting "k8s.io/client-go/testing"

	"go.pinniped.dev/internal/downward"
)

func TestNew(t *testing.T) {
	troo := true
	goodDeployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "some-namespace",
			Name:      "some-name",
		},
	}
	goodPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "some-namespace",
			Name:      "some-name-rsname-podhash",
			OwnerReferences: []metav1.OwnerReference{
				{
					Controller: &troo,
					Name:       "some-name-rsname",
				},
			},
		},
	}
	tests := []struct {
		name            string
		apiObjects      []runtime.Object
		client          func(*kubefake.Clientset)
		createClientErr error
		podInfo         *downward.PodInfo
		wantDeployment  *appsv1.Deployment
		wantPod         *corev1.Pod
		wantError       string
	}{
		{
			name: "happy",
			apiObjects: []runtime.Object{
				goodDeployment,
				&appsv1.ReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "some-namespace",
						Name:      "some-name-rsname",
						OwnerReferences: []metav1.OwnerReference{
							{
								Controller: &troo,
								Name:       "some-name",
							},
						},
					},
				},
				goodPod,
			},
			podInfo: &downward.PodInfo{
				Namespace: "some-namespace",
				Name:      "some-name-rsname-podhash",
			},
			wantDeployment: goodDeployment,
			wantPod:        goodPod,
		},
		{
			name:            "failed to create client",
			createClientErr: errors.New("some create error"),
			podInfo: &downward.PodInfo{
				Namespace: "some-namespace",
				Name:      "some-name-rsname-podhash",
			},
			wantError: "cannot create temp client: some create error",
		},
		{
			name: "failed to talk to api",
			client: func(c *kubefake.Clientset) {
				c.PrependReactor(
					"get",
					"pods",
					func(_ kubetesting.Action) (bool, runtime.Object, error) {
						return true, nil, errors.New("get failed")
					},
				)
			},
			podInfo: &downward.PodInfo{
				Namespace: "some-namespace",
				Name:      "some-name-rsname-podhash",
			},
			wantError: "cannot get deployment: could not get pod: get failed",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := kubefake.NewSimpleClientset(test.apiObjects...)
			if test.client != nil {
				test.client(client)
			}

			getTempClient = func() (kubernetes.Interface, error) {
				return client, test.createClientErr
			}

			_, d, p, err := New(test.podInfo)
			if test.wantError != "" {
				require.EqualError(t, err, test.wantError)
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.wantDeployment, d)
			require.Equal(t, test.wantPod, p)
		})
	}
}
