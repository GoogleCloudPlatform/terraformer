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
	"log"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// IntegrationSlackChannelAllowEmptyValues ...
	IntegrationSlackChannelAllowEmptyValues = []string{}
)

// IntegrationSlackChannelGenerator ...
type IntegrationSlackChannelGenerator struct {
	DatadogService
}

func (g *IntegrationSlackChannelGenerator) createResources(accountID string, slackChannels []datadogV1.SlackIntegrationChannel) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, slackChannel := range slackChannels {
		id := fmt.Sprintf("%s:%s", accountID, slackChannel.GetName())
		resources = append(resources, g.createResource(id))
	}

	return resources
}

func (g *IntegrationSlackChannelGenerator) createResource(id string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		id,
		fmt.Sprintf("integration_slack_channel_%s", id),
		"datadog_integration_slack_channel",
		"datadog",
		IntegrationSlackChannelAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each slack channel create 1 TerraformResource.
func (g *IntegrationSlackChannelGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewSlackIntegrationApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "account_name" && filter.IsApplicable("integration_slack_channel") {
			for _, value := range filter.AcceptableValues {
				slackChannels, _, err := api.GetSlackIntegrationChannels(auth, value)
				if err != nil {
					return err
				}

				resources = g.createResources(value, slackChannels)
			}
		}
		if filter.FieldPath == "id" && filter.IsApplicable("integration_slack_channel") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	if len(resources) == 0 {
		log.Print("Filter(account_name or resource id) is required to import datadog_integration_slack_channel resource")
		return nil
	}
	g.Resources = resources
	return nil
}
