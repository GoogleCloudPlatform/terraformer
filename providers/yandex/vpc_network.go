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

package yandex

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type NetworkGenerator struct {
	YandexService
}

func (g *NetworkGenerator) loadNetworks(sdk *ycsdk.SDK, folderID string) ([]*vpc.Network, error) {
	networks := []*vpc.Network{}
	pageToken := ""
	for {
		resp, err := sdk.VPC().Network().List(context.Background(), &vpc.ListNetworksRequest{
			FolderId:  folderID,
			PageSize:  defaultPageSize,
			PageToken: pageToken,
		})

		if err != nil {
			return nil, err
		}

		networks = append(networks, resp.GetNetworks()...)

		if resp.GetNextPageToken() == "" {
			break
		}

	}
	return networks, nil

}

func (g *NetworkGenerator) InitResources() error {
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(g.Args["token"].(string)),
	})
	if err != nil {
		return err
	}

	result, err := g.loadNetworks(sdk, g.Args["folder_id"].(string))
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result)

	return nil
}

func (g *NetworkGenerator) createResources(networks []*vpc.Network) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, network := range networks {
		resources = append(resources, terraformutils.NewSimpleResource(
			network.GetId(),
			network.GetId(),
			"yandex_vpc_network",
			"yandex",
			[]string{}))
	}
	return resources
}
