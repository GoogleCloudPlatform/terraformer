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

type ServiceMetadataGenerator struct {
	serviceName string
	MackerelService
}

func (g *ServiceMetadataGenerator) createServiceMetadataGeneratorResources(client *mackerel.Client) error {
	services, err := client.FindServices()
	if err != nil {
		return err
	}

	for _, service := range services {
		if service.Name != g.serviceName {
			continue
		}

		namespaces, err := client.GetServiceMetaDataNameSpaces(g.serviceName)
		if err != nil {
			return err
		}

		for _, namespace := range namespaces {
			metadata, err := client.GetServiceMetaData(g.serviceName, namespace)
			if err != nil {
				return err
			}

			b, err := json.Marshal(metadata.ServiceMetaData)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, terraformutils.NewResource(
				g.serviceName+"."+namespace,
				fmt.Sprintf("service_metadata_%s", service.Name),
				"mackerel_service_metadata",
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
// from each service metadata create 1 TerraformResource.
// Need ServiceMetadata Name as ID for terraform resource
func (g *ServiceMetadataGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createServiceMetadataGeneratorResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
