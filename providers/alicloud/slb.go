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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

// SlbGenerator Struct for generating AliCloud Elastic Compute Service
type SlbGenerator struct {
	AliCloudService
}

func resourceFromSlbListener(loadBalancer slb.LoadBalancer, suffix string) terraform_utils.Resource {
	id := loadBalancer.LoadBalancerId + ":" + suffix
	return terraform_utils.NewResource(
		id, // id
		id+"__"+loadBalancer.LoadBalancerName, // name
		"alicloud_slb_listener",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
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

func resourceFromVServerGroupResponse(vServerGroup slb.VServerGroup) terraform_utils.Resource {
	return terraform_utils.NewResource(
		vServerGroup.VServerGroupId,                                    // id
		vServerGroup.VServerGroupId+"__"+vServerGroup.VServerGroupName, // name
		"alicloud_slb_server_group",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initSlb(client *connectivity.AliyunClient) ([]slb.LoadBalancer, error) {
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allLoadBalancers := make([]slb.LoadBalancer, 0)

	for remaining > 0 {
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			request := slb.CreateDescribeLoadBalancersRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return slbClient.DescribeLoadBalancers(request)
		})
		if err != nil {
			return nil, err
		}

		response := raw.(*slb.DescribeLoadBalancersResponse)
		for _, loadBalancer := range response.LoadBalancers.LoadBalancer {
			allLoadBalancers = append(allLoadBalancers, loadBalancer)
		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	return allLoadBalancers, nil
}

func initVServerGroups(client *connectivity.AliyunClient, allLoadBalancers []slb.LoadBalancer) ([]slb.VServerGroup, error) {
	allVserverGroups := make([]slb.VServerGroup, 0)
	for _, loadBalancer := range allLoadBalancers {
		if loadBalancer.LoadBalancerId == "" {
			continue
		}
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			request := slb.CreateDescribeVServerGroupsRequest()
			request.RegionId = client.RegionId
			request.LoadBalancerId = loadBalancer.LoadBalancerId
			return slbClient.DescribeVServerGroups(request)
		})
		if err != nil {
			return nil, err
		}
		response := raw.(*slb.DescribeVServerGroupsResponse)
		for _, vServerGroup := range response.VServerGroups.VServerGroup {
			allVserverGroups = append(allVserverGroups, vServerGroup)
		}
	}

	return allVserverGroups, nil
}

func initSlbListeners(client *connectivity.AliyunClient, allLoadBalancers []slb.LoadBalancer) ([]slb.LoadBalancer, []string, error) {
	alignedLoadBalancers := make([]slb.LoadBalancer, 0)
	suffixes := make([]string, 0)
	for _, loadBalancer := range allLoadBalancers {
		if loadBalancer.LoadBalancerId == "" {
			continue
		}
		raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			request := slb.CreateDescribeLoadBalancerAttributeRequest()
			request.RegionId = client.RegionId
			request.LoadBalancerId = loadBalancer.LoadBalancerId
			return slbClient.DescribeLoadBalancerAttribute(request)
		})
		if err != nil {
			return nil, nil, err
		}
		response := raw.(*slb.DescribeLoadBalancerAttributeResponse)
		for _, listenerPortAndProtocol := range response.ListenerPortsAndProtocol.ListenerPortAndProtocol {
			suffix := fmt.Sprintf("%s:%d", listenerPortAndProtocol.ListenerProtocol, listenerPortAndProtocol.ListenerPort)
			suffixes = append(suffixes, suffix)

			alignedLoadBalancers = append(alignedLoadBalancers, loadBalancer)
		}
	}

	return alignedLoadBalancers, suffixes, nil
}

// InitResources Gets the list of all slb loadBalancer ids and generates resources
func (g *SlbGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}

	allLoadBalancers, err := initSlb(client)
	if err != nil {
		return err
	}
	allVserverGroups, err := initVServerGroups(client, allLoadBalancers)
	if err != nil {
		return err
	}
	alignedLoadBalancers, suffixes, err := initSlbListeners(client, allLoadBalancers)
	if err != nil {
		return err
	}

	for _, loadBalancer := range allLoadBalancers {
		resource := resourceFromSlbResponse(loadBalancer)
		g.Resources = append(g.Resources, resource)
	}

	for _, vServerGroup := range allVserverGroups {
		resource := resourceFromVServerGroupResponse(vServerGroup)
		g.Resources = append(g.Resources, resource)
	}

	for i, alignedSlb := range alignedLoadBalancers {
		resource := resourceFromSlbListener(alignedSlb, suffixes[i])
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

// PostConvertHook Runs before HCL files are generated
func (g *SlbGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_slb" {
			// internet is deprecrated
			// https://www.terraform.io/docs/providers/alicloud/r/slb.html#internet
			delete(r.Item, "internet")

			// https://www.terraform.io/docs/providers/alicloud/r/slb.html#bandwidth
			if r.Item["internet_charge_type"] == "PayByTraffic" {
				delete(r.Item, "bandwidth")
			}
		}
	}

	return nil
}
