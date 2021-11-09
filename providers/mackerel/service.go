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

// ServiceGenerator ...
type ServiceGenerator struct {
	MackerelService
}

func (g *ServiceGenerator) createResources(services []*mackerel.Service) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, service := range services {
		resources = append(resources, g.createResource(service.Name))
	}
	return resources
}

func (g *ServiceGenerator) createResource(serviceName string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		serviceName,
		fmt.Sprintf("service_%s", serviceName),
		"mackerel_service",
		"mackerel",
		[]string{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each service create 1 TerraformResource.
// Need Service Name as ID for terraform resource
func (g *ServiceGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)
	services, err := client.FindServices()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(services)...)
	return nil
}
