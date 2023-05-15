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
	RoleAllowEmptyValues = []string{}
)

type RoleGenerator struct {
	Auth0Service
}

func (g RoleGenerator) createResources(roles []*management.Role) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, role := range roles {
		resourceName := *role.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName+"_"+*role.Name,
			"auth0_role",
			"auth0",
			RoleAllowEmptyValues,
		))
	}
	return resources
}

func (g *RoleGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.Role{}

	var page int
	for {
		l, err := m.Role.List(management.Page(page))
		if err != nil {
			return err
		}
		list = append(list, l.Roles...)
		if !l.HasNext() {
			break
		}
		page++
	}

	g.Resources = g.createResources(list)
	return nil
}
