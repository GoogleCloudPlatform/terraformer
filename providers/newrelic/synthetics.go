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
	synthetics "github.com/dollarshaveclub/new-relic-synthetics-go"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

type SyntheticsGenerator struct {
	NewRelicService
}

func (g *SyntheticsGenerator) createSyntheticsAlertConditionResources(client *newrelic.Client) error {
	alertPolicies, err := client.ListAlertPolicies()
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		syntheticsAlertConditions, err := client.ListAlertSyntheticsConditions(alertPolicy.ID)
		if err != nil {
			return err
		}
		for _, syntheticsAlertCondition := range syntheticsAlertConditions {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%d", alertPolicy.ID, syntheticsAlertCondition.ID),
				fmt.Sprintf("%s-%d", normalizeResourceName(alertPolicy.Name), alertPolicy.ID),
				"newrelic_synthetics_alert_condition",
				g.ProviderName,
				[]string{}))
		}
	}

	return nil
}

func (g *SyntheticsGenerator) createSyntheticsMonitorResources(client *synthetics.Client) error {
	var offset uint
	offset = 0

	allMonitors, err := client.GetAllMonitors(offset, 100)
	if err != nil {
		return err
	}

	for allMonitors.Count > 0 {
		for _, monitor := range allMonitors.Monitors {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprint(monitor.ID),
				fmt.Sprintf("%s-%s", normalizeResourceName(monitor.Name), monitor.ID),
				"newrelic_synthetics_monitor",
				g.ProviderName,
				[]string{}))
		}

		offset += allMonitors.Count
		allMonitors, err = client.GetAllMonitors(offset, 100)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *SyntheticsGenerator) InitResources() error {
	synctheticClient, err := g.SyntheticsClient()
	if err != nil {
		return err
	}
	client, err := g.Client()
	if err != nil {
		return err
	}

	err = g.createSyntheticsMonitorResources(synctheticClient)
	if err != nil {
		return err
	}

	err = g.createSyntheticsAlertConditionResources(client)
	if err != nil {
		return err
	}

	return nil
}
