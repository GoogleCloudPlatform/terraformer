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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
	"log"
	"strconv"
)

type ServiceIntegrationGenerator struct {
	PagerDutyService
}

func (g *ServiceIntegrationGenerator) createServiceIntegrationResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListServicesOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Services.List(&options)
		if err != nil {
			return err
		}
		log.Printf("Offset is at: %s", strconv.Itoa(offset))
		for _, service := range resp.Services {
			for _, serviceIntegration := range service.Integrations {
				respServiceIntegration, _, err := client.Services.GetIntegration(service.ID, serviceIntegration.ID, &pagerduty.GetIntegrationOptions{})
				if err != nil {
					return err
				}
				g.Resources = append(g.Resources, terraformutils.NewResource(
					respServiceIntegration.ID,
					fmt.Sprintf("&s_%s", strings.Replace(service.Name, " ", "_", -1), strings.Replace(respServiceIntegration.Name, " ", "_", -1)),
					"pagerduty_service_integration",
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

func (g *ServiceIntegrationGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createServiceIntegrationResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
