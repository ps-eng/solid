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

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	corev1 "zntr.io/solid/api/gen/go/oidc/core/v1"
	"zntr.io/solid/pkg/sdk/rfcerrors"
	"zntr.io/solid/pkg/sdk/token"
	"zntr.io/solid/pkg/server/authorizationserver"
	"zntr.io/solid/pkg/server/clientauthentication"
)

// TokenIntrospection handles token introspection HTTP requests.
func TokenIntrospection(as authorizationserver.AuthorizationServer, introspectionEncoder token.Generator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var (
			ctx              = r.Context()
			tokenRaw         = r.FormValue("token")
			tokenTypeHintRaw = r.FormValue("token_type_hint")
		)

		// Retrieve client front context
		client, ok := clientauthentication.FromContext(ctx)
		if client == nil || !ok {
			withError(w, r, http.StatusUnauthorized, rfcerrors.InvalidClient().Build())
			return
		}

		// Prepare msg
		msg := &corev1.TokenIntrospectionRequest{
			Client:        client,
			Token:         tokenRaw,
			TokenTypeHint: optionalString(tokenTypeHintRaw),
		}

		// Send request to reactor
		res, err := as.Do(ctx, msg)
		introRes, ok := res.(*corev1.TokenIntrospectionResponse)
		if !ok {
			withJSON(w, r, http.StatusInternalServerError, rfcerrors.ServerError().Build())
			return
		}
		if err != nil {
			log.Println("unable to process introspection request: %w", err)
			withError(w, r, http.StatusBadRequest, introRes.Error)
			return
		}

		acceptRaw := r.Header.Get("Accept")
		if strings.Contains(acceptRaw, "application/token-introspection+jwt") {
			out, err := introspectionEncoder.Generate(r.Context(), introRes.Token)
			if err != nil {
				log.Println("unable to sign introspection response: %w", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Add("Content-Type", "application/token-introspection+jwt")
			fmt.Fprintln(w, out)

		} else {
			active := (introRes.Token.Status == corev1.TokenStatus_TOKEN_STATUS_ACTIVE) && token.IsUsable(introRes.Token)
			resp := map[string]interface{}{
				"active": active,
			}
			if active {
				resp["token_type"] = "Bearer"
				resp["scope"] = introRes.Token.Metadata.Scope
				resp["client_id"] = introRes.Token.Metadata.ClientId
				resp["exp"] = introRes.Token.Metadata.ExpiresAt
				resp["iat"] = introRes.Token.Metadata.IssuedAt
				resp["nbf"] = introRes.Token.Metadata.NotBefore
				resp["sub"] = introRes.Token.Metadata.Subject
				resp["aud"] = introRes.Token.Metadata.Audience
				resp["iss"] = introRes.Token.Metadata.Issuer
				resp["jti"] = introRes.Token.TokenId
			}

			// Send json reponse
			withJSON(w, r, http.StatusOK, resp)
		}

	})
}
