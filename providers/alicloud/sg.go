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
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// SgGenerator Struct for generating AliCloud Security group
type SgGenerator struct {
	AliCloudService
}

func resourceFromSecurityGroup(securitygroup ecs.SecurityGroup) terraform_utils.Resource {
	return terraform_utils.NewResource(
		securitygroup.SecurityGroupId,                                      // id
		securitygroup.SecurityGroupId+"__"+securitygroup.SecurityGroupName, // name
		"alicloud_security_group",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all security group ids and generates resources
func (g *SgGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allSecurityGroups := make([]ecs.SecurityGroup, 0)

	for remaining > 0 {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			request := ecs.CreateDescribeSecurityGroupsRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return ecsClient.DescribeSecurityGroups(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*ecs.DescribeSecurityGroupsResponse)
		for _, securitygroup := range response.SecurityGroups.SecurityGroup {
			allSecurityGroups = append(allSecurityGroups, securitygroup)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	for _, securitygroup := range allSecurityGroups {
		resource := resourceFromSecurityGroup(securitygroup)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
