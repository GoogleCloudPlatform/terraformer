// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth0

import (
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

type Auth0Service struct { //nolint
	terraformutils.Service
}

func (s *Auth0Service) generateClient() *management.Management {
	m, err := management.New(
		s.Args["domain"].(string),
		management.WithClientCredentials(
			s.Args["client_id"].(string),
			s.Args["client_secret"].(string),
		),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

	c := &management.Client{
		Name:        auth0.String("Auth0 Management Client"),
		Description: auth0.String("Client used by Terraformer"),
	}

	err = m.Client.Create(c)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return m
}
