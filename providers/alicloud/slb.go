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
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

// SlbGenerator Struct for generating AliCloud Elastic Compute Service
type SlbGenerator struct {
	AliCloudService
}

func resourceFromSlbResponse(loadBalancer slb.LoadBalancer) terraform_utils.Resource {
	return terraform_utils.NewResource(
		loadBalancer.LoadBalancerId,                                    // id
		loadBalancer.LoadBalancerId+"__"+loadBalancer.LoadBalancerName, // name
		"alicloud_slb",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all slb loadBalancer ids and generates resources
func (g *SlbGenerator) InitResources() error {
	client, err := LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allLoadBalancers := make([]slb.LoadBalancer, 1)

	for remaining > 0 {
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			request := slb.CreateDescribeLoadBalancersRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return slbClient.DescribeLoadBalancers(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*slb.DescribeLoadBalancersResponse)
		for _, loadBalancer := range response.LoadBalancers.LoadBalancer {
			allLoadBalancers = append(allLoadBalancers, loadBalancer)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	for _, loadBalancer := range allLoadBalancers {
		resource := resourceFromSlbResponse(loadBalancer)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
