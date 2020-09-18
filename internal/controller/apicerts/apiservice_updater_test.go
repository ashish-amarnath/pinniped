// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package apicerts

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubeinformers "k8s.io/client-go/informers"
	kubernetesfake "k8s.io/client-go/kubernetes/fake"
	coretesting "k8s.io/client-go/testing"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	aggregatorfake "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/fake"

	loginv1alpha1 "go.pinniped.dev/generated/1.19/apis/login/v1alpha1"
	"go.pinniped.dev/internal/controllerlib"
	"go.pinniped.dev/internal/testutil"
)

func TestAPIServiceUpdaterControllerOptions(t *testing.T) {
	spec.Run(t, "options", func(t *testing.T, when spec.G, it spec.S) {
		const installedInNamespace = "some-namespace"

		var r *require.Assertions
		var observableWithInformerOption *testutil.ObservableWithInformerOption
		var secretsInformerFilter controllerlib.Filter

		it.Before(func() {
			r = require.New(t)
			observableWithInformerOption = testutil.NewObservableWithInformerOption()
			secretsInformer := kubeinformers.NewSharedInformerFactory(nil, 0).Core().V1().Secrets()
			_ = NewAPIServiceUpdaterController(
				installedInNamespace,
				loginv1alpha1.SchemeGroupVersion.Version+"."+loginv1alpha1.GroupName,
				nil,
				secretsInformer,
				observableWithInformerOption.WithInformer, // make it possible to observe the behavior of the Filters
			)
			secretsInformerFilter = observableWithInformerOption.GetFilterForInformer(secretsInformer)
		})

		when("watching Secret objects", func() {
			var subject controllerlib.Filter
			var target, wrongNamespace, wrongName, unrelated *corev1.Secret

			it.Before(func() {
				subject = secretsInformerFilter
				target = &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "api-serving-cert", Namespace: installedInNamespace}}
				wrongNamespace = &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "api-serving-cert", Namespace: "wrong-namespace"}}
				wrongName = &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "wrong-name", Namespace: installedInNamespace}}
				unrelated = &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "wrong-name", Namespace: "wrong-namespace"}}
			})

			when("the target Secret changes", func() {
				it("returns true to trigger the sync method", func() {
					r.True(subject.Add(target))
					r.True(subject.Update(target, unrelated))
					r.True(subject.Update(unrelated, target))
					r.True(subject.Delete(target))
				})
			})

			when("a Secret from another namespace changes", func() {
				it("returns false to avoid triggering the sync method", func() {
					r.False(subject.Add(wrongNamespace))
					r.False(subject.Update(wrongNamespace, unrelated))
					r.False(subject.Update(unrelated, wrongNamespace))
					r.False(subject.Delete(wrongNamespace))
				})
			})

			when("a Secret with a different name changes", func() {
				it("returns false to avoid triggering the sync method", func() {
					r.False(subject.Add(wrongName))
					r.False(subject.Update(wrongName, unrelated))
					r.False(subject.Update(unrelated, wrongName))
					r.False(subject.Delete(wrongName))
				})
			})

			when("a Secret with a different name and a different namespace changes", func() {
				it("returns false to avoid triggering the sync method", func() {
					r.False(subject.Add(unrelated))
					r.False(subject.Update(unrelated, unrelated))
					r.False(subject.Delete(unrelated))
				})
			})
		})
	}, spec.Parallel(), spec.Report(report.Terminal{}))
}

func TestAPIServiceUpdaterControllerSync(t *testing.T) {
	spec.Run(t, "Sync", func(t *testing.T, when spec.G, it spec.S) {
		const installedInNamespace = "some-namespace"

		var r *require.Assertions

		var subject controllerlib.Controller
		var aggregatorAPIClient *aggregatorfake.Clientset
		var kubeInformerClient *kubernetesfake.Clientset
		var kubeInformers kubeinformers.SharedInformerFactory
		var timeoutContext context.Context
		var timeoutContextCancel context.CancelFunc
		var syncContext *controllerlib.Context

		// Defer starting the informers until the last possible moment so that the
		// nested Before's can keep adding things to the informer caches.
		var startInformersAndController = func() {
			// Set this at the last second to allow for injection of server override.
			subject = NewAPIServiceUpdaterController(
				installedInNamespace,
				loginv1alpha1.SchemeGroupVersion.Version+"."+loginv1alpha1.GroupName,
				aggregatorAPIClient,
				kubeInformers.Core().V1().Secrets(),
				controllerlib.WithInformer,
			)

			// Set this at the last second to support calling subject.Name().
			syncContext = &controllerlib.Context{
				Context: timeoutContext,
				Name:    subject.Name(),
				Key: controllerlib.Key{
					Namespace: installedInNamespace,
					Name:      "api-serving-cert",
				},
			}

			// Must start informers before calling TestRunSynchronously()
			kubeInformers.Start(timeoutContext.Done())
			controllerlib.TestRunSynchronously(t, subject)
		}

		it.Before(func() {
			r = require.New(t)

			timeoutContext, timeoutContextCancel = context.WithTimeout(context.Background(), time.Second*3)

			kubeInformerClient = kubernetesfake.NewSimpleClientset()
			kubeInformers = kubeinformers.NewSharedInformerFactory(kubeInformerClient, 0)
			aggregatorAPIClient = aggregatorfake.NewSimpleClientset()
		})

		it.After(func() {
			timeoutContextCancel()
		})

		when("there is not yet an api-serving-cert Secret in the installation namespace or it was deleted", func() {
			it.Before(func() {
				unrelatedSecret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "some other secret",
						Namespace: installedInNamespace,
					},
				}
				err := kubeInformerClient.Tracker().Add(unrelatedSecret)
				r.NoError(err)
			})

			it("does not need to make any API calls with its API client", func() {
				startInformersAndController()
				err := controllerlib.TestSync(t, subject, *syncContext)
				r.NoError(err)
				r.Empty(aggregatorAPIClient.Actions())
			})
		})

		when("there is an api-serving-cert Secret already in the installation namespace", func() {
			it.Before(func() {
				apiServingCertSecret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "api-serving-cert",
						Namespace: installedInNamespace,
					},
					Data: map[string][]byte{
						"caCertificate":       []byte("fake CA cert"),
						"tlsPrivateKey":       []byte("fake private key"),
						"tlsCertificateChain": []byte("fake cert chain"),
					},
				}
				err := kubeInformerClient.Tracker().Add(apiServingCertSecret)
				r.NoError(err)
			})

			when("the APIService exists", func() {
				it.Before(func() {
					apiService := &apiregistrationv1.APIService{
						ObjectMeta: metav1.ObjectMeta{
							Name: loginv1alpha1.SchemeGroupVersion.Version + "." + loginv1alpha1.GroupName,
						},
						Spec: apiregistrationv1.APIServiceSpec{
							CABundle:        nil,
							VersionPriority: 1234,
						},
					}
					err := aggregatorAPIClient.Tracker().Add(apiService)
					r.NoError(err)
				})

				it("updates the APIService's ca bundle", func() {
					startInformersAndController()
					err := controllerlib.TestSync(t, subject, *syncContext)
					r.NoError(err)

					// Make sure we updated the APIService caBundle and left it otherwise unchanged
					r.Len(aggregatorAPIClient.Actions(), 2)
					r.Equal("get", aggregatorAPIClient.Actions()[0].GetVerb())
					expectedAPIServiceName := loginv1alpha1.SchemeGroupVersion.Version + "." + loginv1alpha1.GroupName
					expectedUpdateAction := coretesting.NewUpdateAction(
						schema.GroupVersionResource{
							Group:    apiregistrationv1.GroupName,
							Version:  "v1",
							Resource: "apiservices",
						},
						"",
						&apiregistrationv1.APIService{
							ObjectMeta: metav1.ObjectMeta{
								Name:      expectedAPIServiceName,
								Namespace: "",
							},
							Spec: apiregistrationv1.APIServiceSpec{
								VersionPriority: 1234, // only the CABundle is updated, this other field is left unchanged
								CABundle:        []byte("fake CA cert"),
							},
						},
					)
					r.Equal(expectedUpdateAction, aggregatorAPIClient.Actions()[1])
				})

				when("updating the APIService fails", func() {
					it.Before(func() {
						aggregatorAPIClient.PrependReactor(
							"update",
							"apiservices",
							func(_ coretesting.Action) (bool, runtime.Object, error) {
								return true, nil, errors.New("update failed")
							},
						)
					})

					it("returns the update error", func() {
						startInformersAndController()
						err := controllerlib.TestSync(t, subject, *syncContext)
						r.EqualError(err, "could not update the API service: could not update API service: update failed")
					})
				})
			})

			when("the APIService does not exist", func() {
				it.Before(func() {
					unrelatedAPIService := &apiregistrationv1.APIService{
						ObjectMeta: metav1.ObjectMeta{Name: "some other api service"},
						Spec:       apiregistrationv1.APIServiceSpec{},
					}
					err := aggregatorAPIClient.Tracker().Add(unrelatedAPIService)
					r.NoError(err)
				})

				it("returns an error", func() {
					startInformersAndController()
					err := controllerlib.TestSync(t, subject, *syncContext)
					r.Error(err)
					r.Regexp("could not get existing version of API service: .* not found", err.Error())
				})
			})
		})
	}, spec.Parallel(), spec.Report(report.Terminal{}))
}