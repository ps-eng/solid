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
	"log"
	"net/http"

	"go.zenithar.org/solid/pkg/rfcerrors"

	"github.com/golang/protobuf/ptypes/wrappers"
	corev1 "go.zenithar.org/solid/api/gen/go/oidc/core/v1"
	"go.zenithar.org/solid/pkg/authorizationserver"
)

// Authorization handles authorization HTTP requests.
func Authorization(as authorizationserver.AuthorizationServer) http.Handler {

	type response struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only GET verb
		if r.Method != http.MethodGet {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		}

		// Parameters
		var (
			ctx        = r.Context()
			q          = r.URL.Query()
			requestURI = q.Get("request_uri")
		)

		// Send request to reactor
		res, err := as.Do(ctx, &corev1.AuthorizationRequest{
			RequestUri: &wrappers.StringValue{
				Value: requestURI,
			},
		})
		authRes, ok := res.(*corev1.AuthorizationResponse)
		if !ok {
			withJSON(w, r, http.StatusInternalServerError, rfcerrors.ServerError(""))
			return
		}
		if err != nil {
			log.Println("unable to process authorization request:", err)
			withJSON(w, r, http.StatusBadRequest, authRes.Error)
			return
		}

		// Return JSON response
		withJSON(w, r, http.StatusOK, &response{
			Code:  authRes.Code,
			State: authRes.State,
		})
	})
}
