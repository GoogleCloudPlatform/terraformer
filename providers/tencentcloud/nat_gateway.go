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

package tencentcloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type NatGatewayGenerator struct {
	TencentCloudService
}

func (g *NatGatewayGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeNatGatewaysRequest()
	offset := 0
	pageSize := 50
	allNatGateways := make([]*vpc.NatGateway, 0)

	for {
		offsetUint := uint64(offset)
		limitUint := uint64(pageSize)
		request.Offset = &offsetUint
		request.Limit = &limitUint
		response, err := client.DescribeNatGateways(request)
		if err != nil {
			return err
		}

		allNatGateways = append(allNatGateways, response.Response.NatGatewaySet...)
		if len(response.Response.NatGatewaySet) < pageSize {
			break
		}
		offset += pageSize
	}

	for _, natGateway := range allNatGateways {
		resource := terraformutils.NewResource(
			*natGateway.NatGatewayId,
			*natGateway.NatGatewayName+"_"+*natGateway.NatGatewayId,
			"tencentcloud_nat_gateway",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
