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
	"context"
	"fmt"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// IntegrationAzureAllowEmptyValues ...
	IntegrationAzureAllowEmptyValues = []string{}
)

// IntegrationAzureGenerator ...
type IntegrationAzureGenerator struct {
	DatadogService
}

func (g *IntegrationAzureGenerator) createResources(azureAccounts []datadogV1.AzureAccount) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, account := range azureAccounts {
		resourceID := fmt.Sprintf("%s:%s", account.GetTenantName(), account.GetClientId())
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *IntegrationAzureGenerator) createResource(resourceID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		resourceID,
		fmt.Sprintf("integration_azure_%s", resourceID),
		"datadog_integration_azure",
		"datadog",
		IntegrationAzureAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need IntegrationAzure ID formatted as '<tenant_name>:<client_id>' as ID for terraform resource
func (g *IntegrationAzureGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	integrations, _, err := datadogClientV1.AzureIntegrationApi.ListAzureIntegration(authV1).Execute()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(integrations)
	return nil
}
