// Copyright 2022 The Terraformer Authors.
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
	as "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as/v20180419"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type AsGenerator struct {
	TencentCloudService
}

func (g *AsGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := as.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err := g.loadScalingGroups(client); err != nil {
		return err
	}
	if err := g.loadScalingConfigs(client); err != nil {
		return err
	}

	return nil
}

func (g *AsGenerator) loadScalingGroups(client *as.Client) error {
	request := as.NewDescribeAutoScalingGroupsRequest()

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*as.AutoScalingGroup, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeAutoScalingGroups(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.AutoScalingGroupSet...)
		if len(response.Response.AutoScalingGroupSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.AutoScalingGroupId,
			*instance.AutoScalingGroupName+"_"+*instance.AutoScalingGroupId,
			"tencentcloud_as_scaling_group",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *AsGenerator) loadScalingConfigs(client *as.Client) error {
	request := as.NewDescribeLaunchConfigurationsRequest()

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*as.LaunchConfiguration, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeLaunchConfigurations(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.LaunchConfigurationSet...)
		if len(response.Response.LaunchConfigurationSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.LaunchConfigurationId,
			*instance.LaunchConfigurationName+"_"+*instance.LaunchConfigurationId,
			"tencentcloud_as_scaling_config",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *AsGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type != "tencentcloud_as_scaling_group" {
			continue
		}
		if configID, exist := resource.InstanceState.Attributes["configuration_id"]; exist {
			for _, r := range g.Resources {
				if r.InstanceInfo.Type != "tencentcloud_as_scaling_config" {
					continue
				}
				if configID == r.InstanceState.Attributes["id"] {
					g.Resources[i].Item["configuration_id"] = "${tencentcloud_as_scaling_config." + r.ResourceName + ".id}"
				}
			}
		}
	}

	return nil
}
