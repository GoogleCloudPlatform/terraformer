// Copyright 2022 The Terraformer Authors.
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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"gopkg.in/auth0.v5/management"
)

var (
	EmailAllowEmptyValues = []string{}
)

type EmailGenerator struct {
	Auth0Service
}

func (g EmailGenerator) createResources(email *management.Email) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	resourceName := *email.Name
	resources = append(resources, terraformutils.NewSimpleResource(
		resourceName,
		resourceName,
		"auth0_email",
		"auth0",
		EmailAllowEmptyValues,
	))
	return resources
}

func (g *EmailGenerator) InitResources() error {
	m := g.generateClient()
	Email, err := m.Email.Read()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(Email)
	return nil
}
