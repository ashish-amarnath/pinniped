// Copyright 2021-2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tlsserver

import (
	"context"
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/apimachinery/pkg/util/sets"

	"go.pinniped.dev/internal/crypto/ptls"
)

type ctxKey int

const (
	mapKey ctxKey = iota + 1
	helloKey
)

// TLSTestIPv6Server returns a TLS-required server that listens at an IPv6 loopback
// TODO: Rename since the package name already includes TLS.
func TLSTestIPv6Server(t *testing.T, handler http.Handler, f func(*httptest.Server)) *httptest.Server {
	t.Helper()

	listener, err := net.Listen("tcp6", "[::1]:0")
	require.NoError(t, err, "TLSTestIPv6Server: failed to listen on a port")

	server := &httptest.Server{
		Listener: listener,
		Config:   &http.Server{Handler: handler}, //nolint:gosec //ReadHeaderTimeout is not needed for a localhost listener
	}

	return testServer(t, server, f)
}

// TLSTestServer returns a TLS-required server that listens at an IPv4 loopback.
// TODO: Rename since the package name already includes TLS.
func TLSTestServer(t *testing.T, handler http.Handler, f func(*httptest.Server)) *httptest.Server {
	t.Helper()

	server := httptest.NewUnstartedServer(handler)
	return testServer(t, server, f)
}

func testServer(t *testing.T, server *httptest.Server, f func(*httptest.Server)) *httptest.Server {
	t.Helper()

	server.TLS = ptls.Default(nil) // mimic API server config
	if f != nil {
		f(server)
	}
	server.StartTLS()
	t.Cleanup(server.Close)
	return server
}

func TLSTestServerCA(server *httptest.Server) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: server.Certificate().Raw,
	})
}

func RecordTLSHello(server *httptest.Server) {
	server.Config.ConnContext = func(ctx context.Context, _ net.Conn) context.Context {
		return context.WithValue(ctx, mapKey, &sync.Map{})
	}

	server.TLS.GetConfigForClient = func(info *tls.ClientHelloInfo) (*tls.Config, error) {
		m, ok := getCtxMap(info.Context())
		if !ok {
			return nil, fmt.Errorf("could not find ctx map")
		}
		if actual, loaded := m.LoadOrStore(helloKey, info); loaded && !reflect.DeepEqual(info, actual) {
			return nil, fmt.Errorf("different client hello seen")
		}
		return nil, nil
	}
}

func AssertTLS(t *testing.T, r *http.Request, clientTLSConfigFunc ptls.ConfigFunc) {
	t.Helper()

	m, ok := getCtxMap(r.Context())
	require.True(t, ok)

	h, ok := m.Load(helloKey)
	require.True(t, ok)

	actualClientHello, ok := h.(*tls.ClientHelloInfo)
	require.True(t, ok)

	clientTLSConfig := clientTLSConfigFunc(nil)

	var wantClientSupportedVersions []uint16
	var wantClientSupportedCiphers []uint16

	switch {
	// When the provided config only supports TLS 1.2, then set up the expected values for TLS 1.2.
	case clientTLSConfig.MinVersion == tls.VersionTLS12 && clientTLSConfig.MaxVersion == tls.VersionTLS12:
		wantClientSupportedVersions = []uint16{tls.VersionTLS12}
		wantClientSupportedCiphers = clientTLSConfig.CipherSuites
	// When the provided config only supports TLS 1.3, then set up the expected values for TLS 1.3.
	case clientTLSConfig.MinVersion == tls.VersionTLS13:
		wantClientSupportedVersions = []uint16{tls.VersionTLS13}
		wantClientSupportedCiphers = GetExpectedTLS13Ciphers()
	// When the provided config supports both TLS 1.2 and 1.3, then set up the expected values for both.
	case clientTLSConfig.MinVersion == tls.VersionTLS12 && (clientTLSConfig.MaxVersion == 0 || clientTLSConfig.MaxVersion == tls.VersionTLS13):
		wantClientSupportedVersions = []uint16{tls.VersionTLS13, tls.VersionTLS12}
		wantClientSupportedCiphers = appendIfNotAlreadyIncluded(clientTLSConfig.CipherSuites, GetExpectedTLS13Ciphers())
	default:
		require.Fail(t, "incorrect test setup: clientTLSConfig supports an unexpected combination of TLS versions")
	}

	wantClientProtos := clientTLSConfig.NextProtos
	if httpstream.IsUpgradeRequest(r) {
		wantClientProtos = clientTLSConfig.NextProtos[1:]
	}

	// use assert instead of require to not break the http.Handler with a panic
	ok1 := assert.Equal(t, wantClientSupportedVersions, actualClientHello.SupportedVersions)
	ok2 := assert.Equal(t, cipherSuiteIDsToStrings(wantClientSupportedCiphers), cipherSuiteIDsToStrings(actualClientHello.CipherSuites))
	ok3 := assert.Equal(t, wantClientProtos, actualClientHello.SupportedProtos)

	if all := ok1 && ok2 && ok3; !all {
		t.Errorf("insecure TLS detected for %q %q %q upgrade=%v wantClientSupportedVersions=%v wantClientSupportedCiphers=%v wantClientProtos=%v",
			r.Proto, r.Method, r.URL.String(), httpstream.IsUpgradeRequest(r), ok1, ok2, ok3)
	}
}

// appendIfNotAlreadyIncluded only adds the newItems to the list if they are not already included
// in this list. It returns the potentially updated list.
func appendIfNotAlreadyIncluded(list []uint16, newItems []uint16) []uint16 {
	originals := sets.New(list...)
	for _, newItem := range newItems {
		if !originals.Has(newItem) {
			list = append(list, newItem)
		}
	}
	return list
}

func cipherSuiteIDsToStrings(ids []uint16) []string {
	cipherSuites := make([]string, 0, len(ids))
	for _, id := range ids {
		cipherSuites = append(cipherSuites, tls.CipherSuiteName(id))
	}
	return cipherSuites
}

func getCtxMap(ctx context.Context) (*sync.Map, bool) {
	m, ok := ctx.Value(mapKey).(*sync.Map)
	return m, ok
}
