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
	// IntegrationAWSAllowEmptyValues ...
	IntegrationAWSAllowEmptyValues = []string{}
)

// IntegrationAWSGenerator ...
type IntegrationAWSGenerator struct {
	DatadogService
}

func (g *IntegrationAWSGenerator) createResources(awsAccounts []datadogV1.AWSAccount) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, account := range awsAccounts {
		resourceID := fmt.Sprintf("%s:%s", account.GetAccountId(), account.GetRoleName())
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *IntegrationAWSGenerator) createResource(resourceID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		resourceID,
		fmt.Sprintf("integration_aws_%s", resourceID),
		"datadog_integration_aws",
		"datadog",
		IntegrationAWSAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need IntegrationAWS ID formatted as '<account_id>:<role_name>' as ID for terraform resource
func (g *IntegrationAWSGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewAWSIntegrationApi(datadogClient)

	integrations, _, err := api.ListAWSAccounts(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(integrations.GetAccounts())
	return nil
}
