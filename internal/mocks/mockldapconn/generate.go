// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package mockldapconn

//go:generate go run -v github.com/golang/mock/mockgen  -destination=mockldapconn.go -package=mockldapconn -copyright_file=../../../hack/header.txt go.pinniped.dev/internal/upstreamldap Conn
