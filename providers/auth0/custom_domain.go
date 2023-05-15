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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"gopkg.in/auth0.v5/management"
)

var (
	CustomDomainAllowEmptyValues = []string{}
)

type CustomDomainGenerator struct {
	Auth0Service
}

func (g CustomDomainGenerator) createResources(customDomains []*management.CustomDomain) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, CustomDomain := range customDomains {
		resourceName := *CustomDomain.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName+"_"+*CustomDomain.Domain,
			"auth0_custom_domain",
			"auth0",
			CustomDomainAllowEmptyValues,
		))
	}
	return resources
}

func (g *CustomDomainGenerator) InitResources() error {
	m := g.generateClient()
	list, err := m.CustomDomain.List()
	if err != nil {
		return err
	}

	g.Resources = g.createResources(list)
	return nil
}
