// Copyright 2021 The Terraformer Authors.
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

package okta

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type NetworkZoneGenerator struct {
	OktaService
}

func (g NetworkZoneGenerator) createResources(networkZoneList []*okta.NetworkZone) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, networkZone := range networkZoneList {

		resources = append(resources, terraformutils.NewResource(
			networkZone.Id,
			networkZone.Name,
			"okta_network_zone",
			"okta",
			map[string]string{
				"name": networkZone.Name,
				"type": networkZone.Type,
			},
			[]string{},
			attributesNetworkZone(networkZone),
		))
	}
	return resources
}

func (g *NetworkZoneGenerator) InitResources() error {
	ctx, client, err := g.Client()
	if err != nil {
		return err
	}

	output, resp, err := client.NetworkZone.ListNetworkZones(ctx, nil)
	if err != nil {
		return err
	}

	for resp.HasNextPage() {
		var networkZoneSet []*okta.NetworkZone
		resp, _ = resp.Next(ctx, &networkZoneSet)
		output = append(output, networkZoneSet...)
	}

	g.Resources = g.createResources(output)
	return nil
}

func attributesNetworkZone(networkZone *okta.NetworkZone) map[string]interface{} {
	attributes := map[string]interface{}{}
	attributes["usage"] = networkZone.Usage

	if networkZone.Type == "DYNAMIC" {
		if networkZone.Locations != nil {
			attributes["dynamic_locations"] = networkZone.Locations
		}
	} else if networkZone.Type == "IP" {
		switch {
		case networkZone.Proxies != nil && networkZone.Gateways != nil:
			attributes["proxies"] = networkZone.Proxies
			attributes["gateways"] = networkZone.Gateways
		case networkZone.Proxies != nil && networkZone.Gateways == nil:
			attributes["proxies"] = networkZone.Proxies
		case networkZone.Proxies == nil && networkZone.Gateways != nil:
			attributes["gateways"] = networkZone.Gateways
		}
	}

	return attributes
}
