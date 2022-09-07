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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// IntegrationGCPAllowEmptyValues ...
	IntegrationGCPAllowEmptyValues = []string{}
)

// IntegrationGCPGenerator ...
type IntegrationGCPGenerator struct {
	DatadogService
}

func (g *IntegrationGCPGenerator) createResources(gcpAccounts []datadogV1.GCPAccount) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, account := range gcpAccounts {
		resourceID := account.GetProjectId()
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *IntegrationGCPGenerator) createResource(resourceID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		resourceID,
		fmt.Sprintf("integration_gcp_%s", resourceID),
		"datadog_integration_gcp",
		"datadog",
		IntegrationGCPAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need IntegrationGCP ID formatted as '<tenant_name>:<client_id>' as ID for terraform resource
func (g *IntegrationGCPGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewGCPIntegrationApi(datadogClient)

	integrations, _, err := api.ListGCPIntegration(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(integrations)
	return nil
}
