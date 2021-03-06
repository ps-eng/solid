// Licensed to SolID under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. SolID licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package authorizationserver

import (
	"zntr.io/solid/pkg/sdk/generator"
	"zntr.io/solid/pkg/sdk/token"
	"zntr.io/solid/pkg/server/storage"
)

// Builder options holder
type options struct {
	authorizationCodeGenerator      generator.AuthorizationCode
	accessTokenGenerator            token.Generator
	refreshTokenGenerator           token.Generator
	clientReader                    storage.ClientReader
	clientWriter                    storage.ClientWriter
	authorizationRequestManager     storage.AuthorizationRequest
	authorizationCodeSessionManager storage.AuthorizationCodeSession
	deviceCodeSessionManager        storage.DeviceCodeSession
	tokenManager                    storage.Token
	resourceReader                  storage.ResourceReader
}

// Option defines functional pattern function type contract.
type Option func(*options)

// ClientReader defines the implementation for retrieving client details.
func ClientReader(store storage.ClientReader) Option {
	return func(opts *options) {
		opts.clientReader = store
	}
}

// AuthorizationRequestManager defines the implementation for managing authorization requests.
func AuthorizationRequestManager(store storage.AuthorizationRequest) Option {
	return func(opts *options) {
		opts.authorizationRequestManager = store
	}
}

// AuthorizationCodeSessionManager defines the implementation for managing authorization code sessions.
func AuthorizationCodeSessionManager(store storage.AuthorizationCodeSession) Option {
	return func(opts *options) {
		opts.authorizationCodeSessionManager = store
	}
}

// DeviceCodeSessionManager defines the implementation for managing device code sessions.
func DeviceCodeSessionManager(store storage.DeviceCodeSession) Option {
	return func(opts *options) {
		opts.deviceCodeSessionManager = store
	}
}

// TokenManager defines the implementation for managing tokens.
func TokenManager(store storage.Token) Option {
	return func(opts *options) {
		opts.tokenManager = store
	}
}

// AccessTokenGenerator defines the implementation used to generate access tokens.
func AccessTokenGenerator(g token.Generator) Option {
	return func(opts *options) {
		opts.accessTokenGenerator = g
	}
}

// RefreshTokenGenerator defines the implementation used to generate refresh tokens.
func RefreshTokenGenerator(g token.Generator) Option {
	return func(opts *options) {
		opts.refreshTokenGenerator = g
	}
}

// ClientManager defines the client manager instance to use.
func ClientManager(store storage.Client) Option {
	return func(opts *options) {
		opts.clientWriter = store
		opts.clientReader = store
	}
}

// ResourceReader defines the resource reader instance to use.
func ResourceReader(store storage.ResourceReader) Option {
	return func(opts *options) {
		opts.resourceReader = store
	}
}
