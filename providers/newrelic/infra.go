// Copyright 2019 The Terraformer Authors.
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

package newrelic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

type InfraGenerator struct {
	NewRelicService
}

func (g *InfraGenerator) createAlertInfraConditionResources(client *newrelic.Client, infraClient *newrelic.InfraClient) error {
	alertPolicies, err := client.ListAlertPolicies()
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		alertInfraConditions, err := infraClient.ListAlertInfraConditions(alertPolicy.ID)
		if err != nil {
			return err
		}
		for _, alertInfraCondition := range alertInfraConditions {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				fmt.Sprintf("%d:%d", alertPolicy.ID, alertInfraCondition.ID),
				fmt.Sprintf("%s-%d", normalizeResourceName(alertInfraCondition.Name), alertInfraCondition.ID),
				"newrelic_infra_alert_condition",
				g.ProviderName,
				map[string]string{
					"type": alertInfraCondition.Type,
				},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}

func (g *InfraGenerator) InitResources() error {
	infraClient, err := g.InfraClient()
	if err != nil {
		return err
	}

	client, err := g.Client()
	if err != nil {
		return err
	}

	err = g.createAlertInfraConditionResources(client, infraClient)
	if err != nil {
		return err
	}

	return nil
}
