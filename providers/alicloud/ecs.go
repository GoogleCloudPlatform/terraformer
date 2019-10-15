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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type EcsGenerator struct {
	AliCloudService
}

func resourceFromInstance(instance ecs.Instance) terraform_utils.Resource {
	return terraform_utils.NewResource(
		instance.InstanceId,   // id
		instance.InstanceName, // name
		"alicloud_instance",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func handleDuplicates(instances []ecs.Instance) []ecs.Instance {
	m := make(map[string]bool)
	for index, instance := range instances {
		if m[instance.InstanceName] {
			i := 0
			originalID := instance.InstanceName
			instance.InstanceName = fmt.Sprintf("%s_%d", originalID, i)

			for m[instance.InstanceName] == true {
				i++
				instance.InstanceName = fmt.Sprintf("%s_%d", originalID, i)
			}
		}
		instances[index].InstanceName = instance.InstanceName
		m[instance.InstanceName] = true
	}

	return instances
}

func (g *EcsGenerator) InitResources() error {
	client, err := LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allInstances := make([]ecs.Instance, 1)

	for remaining > 0 {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			request := ecs.CreateDescribeInstancesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return ecsClient.DescribeInstances(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*ecs.DescribeInstancesResponse)
		for _, instance := range response.Instances.Instance {
			allInstances = append(allInstances, instance)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	// AliCloud allows instances with same names but terraform doesnt
	allInstances = handleDuplicates(allInstances)

	for _, instance := range allInstances {
		resource := resourceFromInstance(instance)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *EcsGenerator) PostConvertHook() error {
	fmt.Println("PostConvertHook: TODO: NOT IMPLEMENTED")
	return nil
}

func (g *EcsGenerator) ParseFilter(rawFilter []string) {
	fmt.Println("ParseFilter: TODO: NOT IMPLEMENTED")
}
