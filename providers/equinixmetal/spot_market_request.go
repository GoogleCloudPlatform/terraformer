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

package equinixmetal

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/packethost/packngo"
)

type SpotMarketRequestGenerator struct {
	EquinixMetalService
}

func (g SpotMarketRequestGenerator) listSpotMarketRequests(client *packngo.Client) ([]packngo.SpotMarketRequest, error) {
	spot_market_requests, _, err := client.SpotMarketRequests.List(g.GetArgs()["project_id"].(string), nil)
	if err != nil {
		return nil, err
	}

	return spot_market_requests, nil
}

func (g SpotMarketRequestGenerator) createResources(spot_market_requestList []packngo.SpotMarketRequest) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, spot_market_request := range spot_market_requestList {
		resources = append(resources, terraformutils.NewSimpleResource(
			spot_market_request.ID,
			spot_market_request.ID,
			"metal_spot_market_request",
			"equinixmetal",
			[]string{}))
	}
	return resources
}

func (g *SpotMarketRequestGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listSpotMarketRequests(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
