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

package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

// RoleGenerator ...
type RoleGenerator struct {
	MackerelService
}

func (g *RoleGenerator) createResources(serviceName string, roles []*mackerel.Role) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, role := range roles {
		resources = append(resources, g.createResource(serviceName, role.Name))
	}
	return resources
}

func (g *RoleGenerator) createResource(serviceName string, roleName string) terraformutils.Resource {
	return terraformutils.NewResource(
		fmt.Sprintf("%s:%s", serviceName, roleName),
		fmt.Sprintf("role_%s_%s", serviceName, roleName),
		"mackerel_role",
		"mackerel",
		map[string]string{
			"service": serviceName,
			"name":    roleName,
		},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each role create 1 TerraformResource.
// Need Service Name And Role Name as ID for terraform resource
func (g *RoleGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)

	services, err := client.FindServices()
	if err != nil {
		return err
	}

	for _, service := range services {
		roles, err := client.FindRoles(service.Name)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createResources(service.Name, roles)...)
	}
	return nil
}
