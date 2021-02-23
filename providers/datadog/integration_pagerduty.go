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
	// IntegrationPagerdutyAllowEmptyValues ...
	IntegrationPagerdutyAllowEmptyValues = []string{"tags."}
)

// IntegrationPagerdutyGenerator ...
type IntegrationPagerdutyGenerator struct {
	DatadogService
}

func (g *IntegrationPagerdutyGenerator) createResources(pdSubdomain string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	resources = append(resources, g.createResource(pdSubdomain))

	return resources
}

func (g *IntegrationPagerdutyGenerator) createResource(serviceName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		serviceName,
		fmt.Sprintf("integration_pagerduty_%s", serviceName),
		"datadog_integration_pagerduty",
		"datadog",
		map[string]string{
			"individual_services": "true",
		},
		IntegrationPagerdutyAllowEmptyValues,
		map[string]interface{}{},
	)
	// Ignore services in favor of individual_services
	resource.IgnoreKeys = append(resource.IgnoreKeys, "^services$")

	return resource
}

// InitResources Generate TerraformResources from Datadog API,
// from PD Service create 1 TerraformResource.
// Need IntegrationPagerduty Subdomain as ID for terraform resource
func (g *IntegrationPagerdutyGenerator) InitResources() error {
	client := datadogCommunity.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))

	integration, err := client.GetIntegrationPD()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(integration.GetSubdomain())
	return nil
}
