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
	// IntegrationAWSLogCollectionAllowEmptyValues ...
	IntegrationAWSLogCollectionAllowEmptyValues = []string{"services"}
)

// IntegrationAWSLogCollectionGenerator ...
type IntegrationAWSLogCollectionGenerator struct {
	DatadogService
}

func (g *IntegrationAWSLogCollectionGenerator) createResources(logCollections []datadogV1.AWSLogsListResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logCollection := range logCollections {
		resourceID := logCollection.GetAccountId()
		resources = append(resources, g.createResource(resourceID))
	}
	return resources
}

func (g *IntegrationAWSLogCollectionGenerator) createResource(resourceID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		resourceID,
		fmt.Sprintf("integration_aws_log_collection_%s", resourceID),
		"datadog_integration_aws_log_collection",
		"datadog",
		IntegrationAWSLogCollectionAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need IntegrationAWSLogCollection ID formatted as '<account_id>:<role_name>' as ID for terraform resource
func (g *IntegrationAWSLogCollectionGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewAWSLogsIntegrationApi(datadogClient)

	logCollections, _, err := api.ListAWSLogsIntegrations(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logCollections)
	return nil
}

func (g *IntegrationAWSLogCollectionGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		// services is a required attribute but can be empty. This ensures we append an empty list
		if r.Item["services"] == nil {
			r.Item["services"] = []string{}
		}
	}
	return nil
}
