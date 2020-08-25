/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Package server is the command line entry point for pinniped-server.
package server

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/plugin/pkg/authenticator/token/webhook"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog/v2"

	"github.com/suzerain-io/pinniped/internal/apiserver"
	"github.com/suzerain-io/pinniped/internal/certauthority/kubecertauthority"
	"github.com/suzerain-io/pinniped/internal/controller/issuerconfig"
	"github.com/suzerain-io/pinniped/internal/controllermanager"
	"github.com/suzerain-io/pinniped/internal/downward"
	"github.com/suzerain-io/pinniped/internal/provider"
	"github.com/suzerain-io/pinniped/internal/registry/credentialrequest"
	crdpinnipedv1alpha1 "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/crdpinniped/v1alpha1"
	pinnipedv1alpha1 "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/pinniped/v1alpha1"
	pinnipedclientset "github.com/suzerain-io/pinniped/kubernetes/1.19/client-go/clientset/versioned"
	"github.com/suzerain-io/pinniped/pkg/config"
)

// App is an object that represents the pinniped-server application.
type App struct {
	cmd *cobra.Command

	// CLI flags
	configPath      string
	downwardAPIPath string
}

// This is ignored for now because we turn off etcd storage below, but this is
// the right prefix in case we turn it back on.
const defaultEtcdPathPrefix = "/registry/" + pinnipedv1alpha1.GroupName

// New constructs a new App with command line args, stdout and stderr.
func New(ctx context.Context, args []string, stdout, stderr io.Writer) *App {
	app := &App{}
	app.addServerCommand(ctx, args, stdout, stderr)
	return app
}

// Run the server.
func (a *App) Run() error {
	return a.cmd.Execute()
}

// Create the server command and save it into the App.
func (a *App) addServerCommand(ctx context.Context, args []string, stdout, stderr io.Writer) {
	cmd := &cobra.Command{
		Use: `pinniped-server`,
		Long: "pinniped-server provides a generic API for mapping an external\n" +
			"credential from somewhere to an internal credential to be used for\n" +
			"authenticating to the Kubernetes API.",
		RunE: func(cmd *cobra.Command, args []string) error { return a.runServer(ctx) },
		Args: cobra.NoArgs,
	}

	cmd.SetArgs(args)
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	addCommandlineFlagsToCommand(cmd, a)

	a.cmd = cmd
}

// Define the app's commandline flags.
func addCommandlineFlagsToCommand(cmd *cobra.Command, app *App) {
	cmd.Flags().StringVarP(
		&app.configPath,
		"config",
		"c",
		"pinniped.yaml",
		"path to configuration file",
	)

	cmd.Flags().StringVar(
		&app.downwardAPIPath,
		"downward-api-path",
		"/etc/podinfo",
		"path to Downward API volume mount",
	)
}

// Boot the aggregated API server, which will in turn boot the controllers.
func (a *App) runServer(ctx context.Context) error {
	// Read the server config file.
	cfg, err := config.FromPath(a.configPath)
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}

	// Discover in which namespace we are installed.
	podInfo, err := downward.Load(a.downwardAPIPath)
	if err != nil {
		return fmt.Errorf("could not read pod metadata: %w", err)
	}
	serverInstallationNamespace := podInfo.Namespace

	// Load the Kubernetes cluster signing CA.
	k8sClusterCA, shutdownCA, err := getClusterCASigner(ctx, serverInstallationNamespace)
	if err != nil {
		return err
	}
	defer shutdownCA()

	// Create a WebhookTokenAuthenticator.
	webhookTokenAuthenticator, err := config.NewWebhook(cfg.WebhookConfig)
	if err != nil {
		return fmt.Errorf("could not create webhook client: %w", err)
	}

	// This cert provider will provide certs to the API server and will
	// be mutated by a controller to keep the certs up to date with what
	// is stored in a k8s Secret. Therefore it also effectively acting as
	// an in-memory cache of what is stored in the k8s Secret, helping to
	// keep incoming requests fast.
	dynamicCertProvider := provider.NewDynamicTLSServingCertProvider()

	// Prepare to start the controllers, but defer actually starting them until the
	// post start hook of the aggregated API server.
	startControllersFunc, err := controllermanager.PrepareControllers(
		serverInstallationNamespace,
		cfg.DiscoveryInfo.URL,
		dynamicCertProvider,
		time.Duration(*cfg.APIConfig.ServingCertificateConfig.DurationSeconds)*time.Second,
		time.Duration(*cfg.APIConfig.ServingCertificateConfig.RenewBeforeSeconds)*time.Second,
	)
	if err != nil {
		return fmt.Errorf("could not prepare controllers: %w", err)
	}

	// Get the aggregated API server config.
	aggregatedAPIServerConfig, err := getAggregatedAPIServerConfig(
		dynamicCertProvider,
		webhookTokenAuthenticator,
		k8sClusterCA,
		startControllersFunc,
	)
	if err != nil {
		return fmt.Errorf("could not configure aggregated API server: %w", err)
	}

	// Complete the aggregated API server config and make a server instance.
	server, err := aggregatedAPIServerConfig.Complete().New()
	if err != nil {
		return fmt.Errorf("could not create aggregated API server: %w", err)
	}

	// Run the server. Its post-start hook will start the controllers.
	return server.GenericAPIServer.PrepareRun().Run(ctx.Done())
}

func getClusterCASigner(ctx context.Context, serverInstallationNamespace string) (*kubecertauthority.CA, kubecertauthority.ShutdownFunc, error) {
	// Load the Kubernetes client configuration.
	kubeConfig, err := restclient.InClusterConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("could not load in-cluster configuration: %w", err)
	}

	// Connect to the core Kubernetes API.
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialize Kubernetes client: %w", err)
	}

	// Connect to the pinniped API.
	pinnipedClient, err := pinnipedclientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialize pinniped client: %w", err)
	}

	// Make a clock tick that triggers a periodic refresh.
	ticker := time.NewTicker(5 * time.Minute)

	// Make a CA which uses the Kubernetes cluster API server's signing certs.
	k8sClusterCA, shutdownCA, err := kubecertauthority.New(
		kubeClient,
		kubecertauthority.NewPodCommandExecutor(kubeConfig, kubeClient),
		ticker.C,
	)

	if err != nil {
		ticker.Stop()

		if updateErr := issuerconfig.CreateOrUpdateCredentialIssuerConfig(
			ctx,
			serverInstallationNamespace,
			pinnipedClient,
			func(configToUpdate *crdpinnipedv1alpha1.CredentialIssuerConfig) {
				configToUpdate.Status.Strategies = []crdpinnipedv1alpha1.CredentialIssuerConfigStrategy{
					{
						Type:           crdpinnipedv1alpha1.KubeClusterSigningCertificateStrategyType,
						Status:         crdpinnipedv1alpha1.ErrorStrategyStatus,
						Reason:         crdpinnipedv1alpha1.CouldNotFetchKeyStrategyReason,
						Message:        err.Error(),
						LastUpdateTime: metav1.Now(),
					},
				}
			},
		); updateErr != nil {
			klog.Errorf("error performing create or update on CredentialIssuerConfig to add strategy error: %s", updateErr.Error())
		}

		return nil, nil, fmt.Errorf("could not load cluster signing CA: %w", err)
	}

	updateErr := issuerconfig.CreateOrUpdateCredentialIssuerConfig(
		ctx,
		serverInstallationNamespace,
		pinnipedClient,
		func(configToUpdate *crdpinnipedv1alpha1.CredentialIssuerConfig) {
			configToUpdate.Status.Strategies = []crdpinnipedv1alpha1.CredentialIssuerConfigStrategy{
				{
					Type:           crdpinnipedv1alpha1.KubeClusterSigningCertificateStrategyType,
					Status:         crdpinnipedv1alpha1.SuccessStrategyStatus,
					Reason:         crdpinnipedv1alpha1.FetchedKeyStrategyReason,
					Message:        "Key was fetched successfully",
					LastUpdateTime: metav1.Now(),
				},
			}
		},
	)
	if updateErr != nil {
		//nolint:goerr113
		return nil, nil, fmt.Errorf("error performing create or update on CredentialIssuerConfig to add strategy success: %w", updateErr)
	}

	return k8sClusterCA, func() { shutdownCA(); ticker.Stop() }, nil
}

// Create a configuration for the aggregated API server.
func getAggregatedAPIServerConfig(
	dynamicCertProvider provider.DynamicTLSServingCertProvider,
	webhookTokenAuthenticator *webhook.WebhookTokenAuthenticator,
	ca credentialrequest.CertIssuer,
	startControllersPostStartHook func(context.Context),
) (*apiserver.Config, error) {
	recommendedOptions := genericoptions.NewRecommendedOptions(
		defaultEtcdPathPrefix,
		apiserver.Codecs.LegacyCodec(pinnipedv1alpha1.SchemeGroupVersion),
		// TODO we should check to see if all the other default settings are acceptable for us
	)
	recommendedOptions.Etcd = nil // turn off etcd storage because we don't need it yet
	recommendedOptions.SecureServing.ServerCert.GeneratedCert = dynamicCertProvider

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)
	// Note that among other things, this ApplyTo() function copies
	// `recommendedOptions.SecureServing.ServerCert.GeneratedCert` into
	// `serverConfig.SecureServing.Cert` thus making `dynamicCertProvider`
	// the cert provider for the running server. The provider will be called
	// by the API machinery periodically. When the provider returns nil certs,
	// the API server will return "the server is currently unable to
	// handle the request" error responses for all incoming requests.
	// If the provider later starts returning certs, then the API server
	// will use them to handle the incoming requests successfully.
	if err := recommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	apiServerConfig := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig: apiserver.ExtraConfig{
			Webhook:                       webhookTokenAuthenticator,
			Issuer:                        ca,
			StartControllersPostStartHook: startControllersPostStartHook,
		},
	}
	return apiServerConfig, nil
}
