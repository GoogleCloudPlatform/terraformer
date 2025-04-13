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
	newrelic "github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/common"
)

type TagsGenerator struct {
	NewRelicService
}

func (g *TagsGenerator) createSyntheticsMonitorTagResources(client *newrelic.NewRelic) error {
	allMonitors, err := client.Synthetics.ListMonitors()
	if err != nil {
		return err
	}

	for _, monitor := range allMonitors {
		allTags, err := client.Entities.GetTagsForEntityMutable(common.EntityGUID(monitor.ID))
		if err != nil {
			return err
		}

		for range allTags {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprint(monitor.ID),
				fmt.Sprintf("%s-%s", normalizeResourceName(monitor.Name), monitor.ID),
				"newrelic_entity_tags",
				g.ProviderName,
				[]string{}))
		}
	}

	return nil
}

func (g *TagsGenerator) createAlertConditionTagResources(client *newrelic.NewRelic) error {
	alertPolicies, err := client.Alerts.ListPolicies(nil)
	if err != nil {
		return err
	}

	for _, alertPolicy := range alertPolicies {
		alertConditions, err := client.Alerts.ListConditions(alertPolicy.ID)
		if err != nil {
			return err
		}

		nrqlConditions, err := client.Alerts.ListNrqlConditions(alertPolicy.ID)
		if err != nil {
			return err
		}

		for _, alertCondition := range alertConditions {
			allAlertConditionTags, err := client.Entities.GetTagsForEntityMutable(common.EntityGUID(fmt.Sprint(alertCondition.ID)))
			if err != nil {
				return err
			}
			for range allAlertConditionTags {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					fmt.Sprintf("%d:%d", alertPolicy.ID, alertCondition.ID),
					fmt.Sprintf("%s-%d", normalizeResourceName(alertCondition.Name), alertCondition.ID),
					"newrelic_entity_tags",
					g.ProviderName,
					[]string{}))
			}
		}

		for _, nrqlCondition := range nrqlConditions {
			allNRQLConditionTags, err := client.Entities.GetTagsForEntityMutable(common.EntityGUID(fmt.Sprint(nrqlCondition.ID)))
			if err != nil {
				return err
			}
			for range allNRQLConditionTags {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					fmt.Sprintf("%d:%d", alertPolicy.ID, nrqlCondition.ID),
					fmt.Sprintf("%s-%d", normalizeResourceName(nrqlCondition.Name), nrqlCondition.ID),
					"newrelic_entity_tags",
					g.ProviderName,
					[]string{}))
			}
		}
	}

	return nil
}

func (g *TagsGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*newrelic.NewRelic) error{
		g.createSyntheticsMonitorTagResources,
		g.createAlertConditionTagResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
