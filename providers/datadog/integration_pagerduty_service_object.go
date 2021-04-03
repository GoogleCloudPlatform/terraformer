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

package datadog

import (
	"fmt"

	datadogCommunity "github.com/zorkian/go-datadog-api"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// IntegrationPagerdutyServiceObjectAllowEmptyValues ...
	IntegrationPagerdutyServiceObjectAllowEmptyValues = []string{"tags."}
)

// IntegrationPagerdutyServiceObjectGenerator ...
type IntegrationPagerdutyServiceObjectGenerator struct {
	DatadogService
}

func (g *IntegrationPagerdutyServiceObjectGenerator) createResources(serviceNames []string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, name := range serviceNames {
		resourceName := name
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *IntegrationPagerdutyServiceObjectGenerator) createResource(serviceName string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		serviceName,
		fmt.Sprintf("integration_pagerduty_service_object_%s", serviceName),
		"datadog_integration_pagerduty_service_object",
		"datadog",
		IntegrationPagerdutyServiceObjectAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each PD Service create 1 TerraformResource.
// Need IntegrationPagerdutyServiceObject ServiceName as ID for terraform resource
func (g *IntegrationPagerdutyServiceObjectGenerator) InitResources() error {
	client := datadogCommunity.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))

	pdIntegration, err := client.GetIntegrationPD()
	if err != nil {
		return err
	}

	var serviceNames []string
	for _, service := range pdIntegration.Services {
		serviceNames = append(serviceNames, *service.ServiceName)
	}

	g.Resources = g.createResources(serviceNames)
	return nil
}
