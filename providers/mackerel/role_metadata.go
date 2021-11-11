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
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

type RoleMetadataGenerator struct {
	serviceName string
	MackerelService
}

func (g *RoleMetadataGenerator) createRoleMetadataGeneratorResources(client *mackerel.Client) error {
	roles, err := client.FindRoles(g.serviceName)
	if err != nil {
		return err
	}

	for _, role := range roles {
		namespaces, err := client.GetRoleMetaDataNameSpaces(g.serviceName, role.Name)
		if err != nil {
			return err
		}

		for _, namespace := range namespaces {
			metadata, err := client.GetRoleMetaData(g.serviceName, role.Name, namespace)
			if err != nil {
				return err
			}

			b, err := json.Marshal(metadata)
			if err != nil {
				return err
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				role.Name+"."+namespace,
				fmt.Sprintf("role_metadata_%s", role.Name),
				"mackerel_role_metadata",
				g.ProviderName,
				map[string]string{
					"metadata_json": string(b),
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}
	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each role metadata create 1 TerraformResource.
// Need RoleMetadata Name as ID for terraform resource
func (g *RoleMetadataGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createRoleMetadataGeneratorResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
