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

package pagerduty

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
)

type ServiceGenerator struct {
	PagerDutyService
}

func (g *ServiceGenerator) createServiceResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListServicesOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Services.List(&options)
		if err != nil {
			return err
		}

		for _, service := range resp.Services {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				service.ID,
				fmt.Sprintf("service_%s", service.Name),
				"pagerduty_service",
				g.ProviderName,
				[]string{},
			))
		}

		if !resp.More {
			break
		}
		offset += resp.Limit
	}

	return nil
}

func (g *ServiceGenerator) createServiceEventRuleResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListServicesOptions{}
	optionsEventRules := pagerduty.ListServiceEventRuleOptions{}
	for {
		options.Offset = offset
		optionsEventRules.Offset = offset
		resp, _, err := client.Services.List(&options)
		if err != nil {
			return err
		}

		for _, service := range resp.Services {
			rules, _, err := client.Services.ListEventRules(service.ID, &optionsEventRules)

			if err != nil {
				return err
			}

			for _, rule := range rules.EventRules {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					rule.ID,
					fmt.Sprintf("%s_%s", service.Name, rule.ID),
					"pagerduty_service_event_rule",
					g.ProviderName,
					map[string]string{
						"service": service.ID,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}

		if !resp.More {
			break
		}
		offset += resp.Limit
	}
	return nil
}

func (g *ServiceGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createServiceResources,
		g.createServiceEventRuleResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
