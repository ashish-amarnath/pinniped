// Copyright 2020-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package upstreamprovider

import (
	"context"
	"net/url"

	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/types"

	idpv1alpha1 "go.pinniped.dev/generated/latest/apis/supervisor/idp/v1alpha1"
	"go.pinniped.dev/internal/authenticators"
	"go.pinniped.dev/internal/setutil"
	"go.pinniped.dev/pkg/oidcclient/nonce"
	"go.pinniped.dev/pkg/oidcclient/oidctypes"
	"go.pinniped.dev/pkg/oidcclient/pkce"
)

type RevocableTokenType string

// These strings correspond to the token types defined by https://datatracker.ietf.org/doc/html/rfc7009#section-2.1
const (
	RefreshTokenType RevocableTokenType = "refresh_token"
	AccessTokenType  RevocableTokenType = "access_token"
)

// LDAPRefreshAttributes contains information about the user from the original login request
// and previous refreshes to be used during an LDAP session refresh.
type LDAPRefreshAttributes struct {
	Username             string
	Subject              string
	DN                   string
	Groups               []string
	AdditionalAttributes map[string]string
}

// UpstreamIdentityProviderI includes the interface functions that are common to all upstream identity provider types.
// These represent the identity provider resources, i.e. OIDCIdentityProvider, etc.
type UpstreamIdentityProviderI interface {
	// GetResourceName returns a name for this upstream provider. The controller watching the identity provider resources will
	// set this to be the Name of the CR from its metadata. Note that this is different from the DisplayName configured
	// in each FederationDomain that uses this provider, so this name is for internal use only, not for interacting
	// with clients. Clients should not expect to see this name or send this name.
	GetResourceName() string

	// GetResourceUID returns the Kubernetes resource ID
	GetResourceUID() types.UID
}

type UpstreamOIDCIdentityProviderI interface {
	UpstreamIdentityProviderI

	// GetClientID returns the OAuth client ID registered with the upstream provider to be used in the authorization code flow.
	GetClientID() string

	// GetAuthorizationURL returns the Authorization Endpoint fetched from discovery.
	GetAuthorizationURL() *url.URL

	// HasUserInfoURL returns whether there is a non-empty value for userinfo_endpoint fetched from discovery.
	HasUserInfoURL() bool

	// GetScopes returns the scopes to request in authorization (authcode or password grant) flow.
	GetScopes() []string

	// GetUsernameClaim returns the ID Token username claim name. May return empty string, in which case we
	// will use some reasonable defaults.
	GetUsernameClaim() string

	// GetGroupsClaim returns the ID Token groups claim name. May return empty string, in which case we won't
	// try to read groups from the upstream provider.
	GetGroupsClaim() string

	// AllowsPasswordGrant returns true if a client should be allowed to use the resource owner password credentials grant
	// flow with this upstream provider. When false, it should not be allowed.
	AllowsPasswordGrant() bool

	// GetAdditionalAuthcodeParams returns additional params to be sent on authcode requests.
	GetAdditionalAuthcodeParams() map[string]string

	// GetAdditionalClaimMappings returns additional claims to be mapped from the upstream ID token.
	GetAdditionalClaimMappings() map[string]string

	// PasswordCredentialsGrantAndValidateTokens performs upstream OIDC resource owner password credentials grant and
	// token validation. Returns the validated raw tokens as well as the parsed claims of the ID token.
	PasswordCredentialsGrantAndValidateTokens(ctx context.Context, username, password string) (*oidctypes.Token, error)

	// ExchangeAuthcodeAndValidateTokens performs upstream OIDC authorization code exchange and token validation.
	// Returns the validated raw tokens as well as the parsed claims of the ID token.
	ExchangeAuthcodeAndValidateTokens(
		ctx context.Context,
		authcode string,
		pkceCodeVerifier pkce.Code,
		expectedIDTokenNonce nonce.Nonce,
		redirectURI string,
	) (*oidctypes.Token, error)

	// PerformRefresh will call the provider's token endpoint to perform a refresh grant. The provider may or may not
	// return a new ID or refresh token in the response. If it returns an ID token, then use ValidateToken to
	// validate the ID token.
	PerformRefresh(ctx context.Context, refreshToken string) (*oauth2.Token, error)

	// RevokeToken will attempt to revoke the given token, if the provider has a revocation endpoint.
	// It may return an error wrapped by a RetryableRevocationError, which is an error indicating that it may
	// be worth trying to revoke the same token again later. Any other error returned should be assumed to
	// represent an error such that it is not worth retrying revocation later, even though revocation failed.
	RevokeToken(ctx context.Context, token string, tokenType RevocableTokenType) error

	// ValidateTokenAndMergeWithUserInfo will validate the ID token. It will also merge the claims from the userinfo endpoint response
	// into the ID token's claims, if the provider offers the userinfo endpoint. It returns the validated/updated
	// tokens, or an error.
	ValidateTokenAndMergeWithUserInfo(ctx context.Context, tok *oauth2.Token, expectedIDTokenNonce nonce.Nonce, requireIDToken bool, requireUserInfo bool) (*oidctypes.Token, error)
}

type UpstreamLDAPIdentityProviderI interface {
	UpstreamIdentityProviderI

	// GetURL returns a URL which uniquely identifies this LDAP provider, e.g. "ldaps://host.example.com:1234".
	// This URL is not used for connecting to the provider, but rather is used for creating a globally unique user
	// identifier by being combined with the user's UID, since user UIDs are only unique within one provider.
	GetURL() *url.URL

	// UserAuthenticator adds an interface method for performing user authentication against the upstream LDAP provider.
	authenticators.UserAuthenticator

	// PerformRefresh performs a refresh against the upstream LDAP identity provider
	PerformRefresh(ctx context.Context, storedRefreshAttributes LDAPRefreshAttributes, idpDisplayName string) (groups []string, err error)
}

type GitHubUser struct {
	Username          string   // could be login name, id, or login:id
	Groups            []string // could be names or slugs
	DownstreamSubject string   // the whole downstream subject URI
}

// GitHubLoginDeniedError can be returned by UpstreamGithubIdentityProviderI GetUser() when a policy
// configured on GitHubIdentityProvider should prevent this user from completing authentication.
type GitHubLoginDeniedError struct {
	message string
}

func NewGitHubLoginDeniedError(message string) GitHubLoginDeniedError {
	return GitHubLoginDeniedError{message: message}
}

func (g GitHubLoginDeniedError) Error() string {
	return g.message
}

var _ error = &GitHubLoginDeniedError{}

type UpstreamGithubIdentityProviderI interface {
	UpstreamIdentityProviderI

	// GetClientID returns the OAuth client ID registered with the upstream provider to be used in the authorization code flow.
	GetClientID() string

	// GetScopes returns the scopes to request in authorization (authcode or password grant) flow.
	GetScopes() []string

	// GetUsernameAttribute returns the attribute from the GitHub API user response to use for the downstream username.
	// See https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user.
	// Note that this is a constructed value - do not expect that the result will exactly match one of the JSON fields.
	GetUsernameAttribute() idpv1alpha1.GitHubUsernameAttribute

	// GetGroupNameAttribute returns the attribute from the GitHub API team response to use for the downstream group names.
	// See https://docs.github.com/en/rest/teams/teams?apiVersion=2022-11-28#list-teams-for-the-authenticated-user.
	// Note that this is a constructed value - do not expect that the result will exactly match one of the JSON fields.
	GetGroupNameAttribute() idpv1alpha1.GitHubGroupNameAttribute

	// GetAllowedOrganizations returns a list of organizations configured to allow authentication.
	// If this list has contents, a user must have membership in at least one of these organizations to log in,
	// and only teams from the listed organizations should be represented as groups for the downstream token.
	// If this list is empty, then any user can log in regardless of org membership, and any observable
	// teams memberships should be represented as groups for the downstream token.
	GetAllowedOrganizations() *setutil.CaseInsensitiveSet

	// GetAuthorizationURL returns the authorization URL for the configured GitHub. This will look like:
	// https://<spec.githubAPI.host>/login/oauth/authorize
	// It will not include any query parameters or fragment. Any subdomains or port will come from <spec.githubAPI.host>.
	// It will never include a username or password in the authority section.
	GetAuthorizationURL() string

	// ExchangeAuthcode performs an upstream GitHub authorization code exchange.
	// Returns the raw access token. The access token expiry is not known.
	ExchangeAuthcode(ctx context.Context, authcode string, redirectURI string) (string, error)

	// GetUser calls the user, orgs, and teams APIs of GitHub using the accessToken.
	// It validates any required org memberships. It returns a User or an error.
	// The IDP display name is passed to aid in building a suitable downstream subject string.
	GetUser(ctx context.Context, accessToken string, idpDisplayName string) (*GitHubUser, error)
}
