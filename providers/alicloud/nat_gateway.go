// Copyright 2018 The Terraformer Authors.
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

package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

// NatGatewayGenerator Struct for generating AliCloud Elastic Compute Service
type NatGatewayGenerator struct {
	AliCloudService
}

func resourceFromNatGatewayResponse(natGateway vpc.NatGateway) terraform_utils.Resource {
	return terraform_utils.NewResource(
		natGateway.NatGatewayId,                      // id
		natGateway.NatGatewayId+"__"+natGateway.Name, // name
		"alicloud_nat_gateway",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all natgateway NatGateway ids and generates resources
func (g *NatGatewayGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allNatGateways := make([]vpc.NatGateway, 0)

	for remaining > 0 {
		raw, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			request := vpc.CreateDescribeNatGatewaysRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return vpcClient.DescribeNatGateways(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*vpc.DescribeNatGatewaysResponse)
		for _, NatGateway := range response.NatGateways.NatGateway {
			allNatGateways = append(allNatGateways, NatGateway)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	for _, NatGateway := range allNatGateways {
		resource := resourceFromNatGatewayResponse(NatGateway)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
