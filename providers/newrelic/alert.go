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

type AlertGenerator struct {
	NewRelicService
}

func (g *AlertGenerator) createAlertChannelResources(client *newrelic.Client) error {
	alertChannels, err := client.ListAlertChannels()
	if err != nil {
		return err
	}

	for _, channel := range alertChannels {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			fmt.Sprintf("%d", channel.ID),
			fmt.Sprintf("%s-%d", normalizeResourceName(channel.Name), channel.ID),
			"newrelic_alert_channel",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}

func (g *AlertGenerator) createAlertConditionResources(client *newrelic.Client) error {
	alertPolicies, err := client.ListAlertPolicies()
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		alertConditions, err := client.ListAlertConditions(alertPolicy.ID)
		if err != nil {
			return err
		}

		for _, alertCondition := range alertConditions {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%d", alertPolicy.ID, alertCondition.ID),
				fmt.Sprintf("%s-%d", normalizeResourceName(alertCondition.Name), alertCondition.ID),
				"newrelic_alert_condition",
				g.ProviderName,
				[]string{}))
		}
	}
	return nil
}

func (g *AlertGenerator) createAlertNrqlConditionResources(client *newrelic.Client) error {
	alertPolicies, err := client.ListAlertPolicies()
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		alertNrqlConditions, err := client.ListAlertNrqlConditions(alertPolicy.ID)
		if err != nil {
			return err
		}

		for _, alertNrqlCondition := range alertNrqlConditions {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%d", alertPolicy.ID, alertNrqlCondition.ID),
				fmt.Sprintf("%s-%d", normalizeResourceName(alertNrqlCondition.Name), alertNrqlCondition.ID),
				"newrelic_nrql_alert_condition",
				g.ProviderName,
				[]string{}))
		}
	}
	return nil
}

func (g *AlertGenerator) createAlertPolicyResources(client *newrelic.Client) error {
	alertPolicies, err := client.ListAlertPolicies()
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			fmt.Sprintf("%d", alertPolicy.ID),
			fmt.Sprintf("%s-%d", normalizeResourceName(alertPolicy.Name), alertPolicy.ID),
			"newrelic_alert_policy",
			g.ProviderName,
			[]string{}))
	}

	return nil
}

func (g *AlertGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*newrelic.Client) error{
		g.createAlertChannelResources,
		g.createAlertConditionResources,
		g.createAlertNrqlConditionResources,
		g.createAlertPolicyResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *AlertGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "newrelic_alert_condition" {
			if resource.Item["violation_close_timer"] == "0" {
				delete(g.Resources[i].Item, "violation_close_timer")
			}
		}
	}

	return nil
}
