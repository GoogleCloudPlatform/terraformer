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

type SubnetGenerator struct {
	YandexService
}

func (g *SubnetGenerator) loadSubnets(sdk *ycsdk.SDK, folderID string) ([]*vpc.Subnet, error) {
	subnets := []*vpc.Subnet{}
	pageToken := ""
	for {
		resp, err := sdk.VPC().Subnet().List(context.Background(), &vpc.ListSubnetsRequest{
			FolderId:  folderID,
			PageSize:  defaultPageSize,
			PageToken: pageToken,
		})

		if err != nil {
			return nil, err
		}

		subnets = append(subnets, resp.GetSubnets()...)

		if resp.GetNextPageToken() == "" {
			break
		}

	}
	return subnets, nil

}

func (g *SubnetGenerator) InitResources() error {
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(g.Args["token"].(string)),
	})
	if err != nil {
		return err
	}

	result, err := g.loadSubnets(sdk, g.Args["folder_id"].(string))
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result)

	return nil
}

func (g *SubnetGenerator) createResources(subnets []*vpc.Subnet) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, subnet := range subnets {
		resources = append(resources, terraformutils.NewSimpleResource(
			subnet.GetId(),
			subnet.GetId(),
			"yandex_vpc_subnet",
			"yandex",
			[]string{}))
	}
	return resources
}
