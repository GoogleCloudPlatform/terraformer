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
	// IntegrationAWSLambdaARNAllowEmptyValues ...
	IntegrationAWSLambdaARNAllowEmptyValues = []string{}
)

// IntegrationAWSLambdaARNGenerator ...
type IntegrationAWSLambdaARNGenerator struct {
	DatadogService
}

func (g *IntegrationAWSLambdaARNGenerator) createResources(logCollections []datadogV1.AWSLogsListResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logCollection := range logCollections {
		for _, logCollectionLambdaArn := range logCollection.GetLambdas() {
			accountID := logCollection.GetAccountId()
			if v, ok := logCollectionLambdaArn.GetArnOk(); ok {
				resourceID := fmt.Sprintf("%s %s", accountID, *v)
				resources = append(resources, g.createResource(resourceID))
			}
		}
	}
	return resources
}

func (g *IntegrationAWSLambdaARNGenerator) createResource(resourceID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		resourceID,
		fmt.Sprintf("integration_aws_lambda_arn_%s", resourceID),
		"datadog_integration_aws_lambda_arn",
		"datadog",
		IntegrationAWSLambdaARNAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need IntegrationAWSLambdaARN ID formatted as '<account_id>:<role_name>' as ID for terraform resource
func (g *IntegrationAWSLambdaARNGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	logCollections, _, err := datadogClientV1.AWSLogsIntegrationApi.ListAWSLogsIntegrations(authV1).Execute()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logCollections)
	return nil
}
