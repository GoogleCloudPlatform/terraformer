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

// VpcGenerator Struct for generating AliCloud Elastic Compute Service
type VpcGenerator struct {
	AliCloudService
}

func resourceFromVpcResponse(Vpc vpc.Vpc) terraform_utils.Resource {
	return terraform_utils.NewResource(
		Vpc.VpcId,                  // id
		Vpc.VpcId+"__"+Vpc.VpcName, // name
		"alicloud_vpc",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all vpc Vpc ids and generates resources
func (g *VpcGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allVpcs := make([]vpc.Vpc, 0)

	for remaining > 0 {
		raw, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			request := vpc.CreateDescribeVpcsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return vpcClient.DescribeVpcs(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*vpc.DescribeVpcsResponse)
		for _, Vpc := range response.Vpcs.Vpc {
			allVpcs = append(allVpcs, Vpc)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	for _, Vpc := range allVpcs {
		resource := resourceFromVpcResponse(Vpc)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
