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

package dpop

import (
	"fmt"
	"net/url"
	"time"

	"github.com/dchest/uniuri"
	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
)

// {
// 	"typ":"dpop+jwt",
// 	"alg":"ES256",
// 	"jwk": {
// 	  "kty":"EC",
// 	  "x":"l8tFrhx-34tV3hRICRDY9zCkDlpBhF42UQUfWVAWBFs",
// 	  "y":"9VE4jf_Ok_o64zbTTlcuNJajHmt6v9TDVrU0CdvGRDA",
// 	  "crv":"P-256"
// 	}
// }.{
//   "jti":"-BwC3ESc6acc2lTc",
//   "htm":"POST",
//   "htu":"https://server.example.com/token",
//   "iat":1562262616
// }

// Proof while generate a dpop proof JWT encoded.
func Proof(privateKey *jose.JSONWebKey, htm string, htu *url.URL) (string, error) {
	// Check parameters
	if privateKey == nil {
		return "", fmt.Errorf("unable to generate proof of nil key")
	}
	if privateKey.IsPublic() {
		return "", fmt.Errorf("unable to generate proof for a public key")
	}
	if htm == "" {
		return "", fmt.Errorf("htm must not be blank")
	}
	if htu == nil {
		return "", fmt.Errorf("htu must not be nil")
	}

	// Create proof claims
	claims := map[string]interface{}{
		"jti": uniuri.NewLen(16),
		"htm": htm,
		"htu": fmt.Sprintf("%s://%s%s", htu.Scheme, htu.Host, htu.Path),
		"iat": time.Now().UTC().Unix(),
	}

	// Prepare signer options
	options := (&jose.SignerOptions{}).WithType("dpop+jwt")
	options.EmbedJWK = true

	// Prepare a signer
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.ES256, Key: privateKey}, options)
	if err != nil {
		return "", fmt.Errorf("unable to prepare signer: %w", err)
	}

	// Generate the final proof
	raw, err := jwt.Signed(sig).Claims(claims).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("unable to generate DPoP: %w", err)
	}

	// No error
	return raw, nil
}
