// Copyright 2021 The Terraformer Authors.
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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type RoleGenerator struct {
	SumoLogicService
}

func (g *RoleGenerator) createResources(roles []sumologic.RoleModel) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(roles))

	for i, role := range roles {
		name := strcase.ToSnake(replaceSpaceAndDash(role.Name))
		resource := terraformutils.NewSimpleResource(
			role.Id,
			fmt.Sprintf("%s-%s", name, role.Id),
			"sumologic_role",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *RoleGenerator) InitResources() error {
	client := g.Client()
	req := client.RoleManagementApi.ListRoles(g.AuthCtx())
	req = req.Limit(100)
	for _, filter := range g.Filter {
		if filter.IsApplicable("role") && filter.FieldPath == "name" {
			if len(filter.AcceptableValues) == 1 {
				req = req.Name(filter.AcceptableValues[0])
			}
		}
	}

	respBody, _, err := client.RoleManagementApi.ListRolesExecute(req)
	if err != nil {
		return err
	}
	roles := respBody.Data
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.RoleManagementApi.ListRolesExecute(req)
		if err != nil {
			return err
		}
		roles = append(roles, respBody.Data...)
	}

	resources := g.createResources(roles)
	g.Resources = resources
	return nil
}
