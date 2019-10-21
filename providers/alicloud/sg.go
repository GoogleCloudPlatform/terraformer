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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
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

func resourceFromSecurityGroupAttribute(permission ecs.Permission, securityGroup ecs.SecurityGroup) terraform_utils.Resource {
	// https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/alicloud/resource_alicloud_security_group_rule.go#L153
	// sgId + ":" + direction + ":" + ptl + ":" + port + ":" + nicType + ":" + cidr_ip + ":" + policy + ":" + strconv.Itoa(priority)
	id := strings.Join([]string{
		securityGroup.SecurityGroupId,
		permission.Direction,
		permission.IpProtocol,
		permission.PortRange,
		permission.NicType,
		permission.SourceCidrIp,
		permission.Policy,
		permission.Priority,
	}, ":")
	id = strings.ToLower(id)

	return terraform_utils.NewResource(
		id, // id
		id+"__"+securityGroup.SecurityGroupName, // name
		"alicloud_security_group_rule",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initSecurityGroupRules(client *connectivity.AliyunClient, securityGroups []ecs.SecurityGroup) ([]ecs.Permission, []ecs.SecurityGroup, error) {
	allPermissions := make([]ecs.Permission, 0)
	alignedSecurityGroups := make([]ecs.SecurityGroup, 0)

	for _, securityGroup := range securityGroups {
		if securityGroup.SecurityGroupId == "" {
			continue
		}
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			request := ecs.CreateDescribeSecurityGroupAttributeRequest()
			request.RegionId = client.RegionId
			request.SecurityGroupId = securityGroup.SecurityGroupId
			return ecsClient.DescribeSecurityGroupAttribute(request)
		})
		if err != nil {
			return nil, nil, err
		}

		response := raw.(*ecs.DescribeSecurityGroupAttributeResponse)
		for _, zoneRecord := range response.Permissions.Permission {
			allPermissions = append(allPermissions, zoneRecord)
			alignedSecurityGroups = append(alignedSecurityGroups, securityGroup)
		}

	}
	return allPermissions, alignedSecurityGroups, nil
}

func initSecurityGroups(client *connectivity.AliyunClient) ([]ecs.SecurityGroup, error) {
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
			return nil, err
		}

		response := raw.(*ecs.DescribeSecurityGroupsResponse)
		for _, securitygroup := range response.SecurityGroups.SecurityGroup {
			allSecurityGroups = append(allSecurityGroups, securitygroup)

		}
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	return allSecurityGroups, nil
}

// InitResources Gets the list of all security group ids and generates resources
func (g *SgGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}

	allSecurityGroups, err := initSecurityGroups(client)
	if err != nil {
		return err
	}

	allSecurityGroupRules, alignedSecurityGroups, err := initSecurityGroupRules(client, allSecurityGroups)

	for _, securitygroup := range allSecurityGroups {
		resource := resourceFromSecurityGroup(securitygroup)
		g.Resources = append(g.Resources, resource)
	}

	for i, permission := range allSecurityGroupRules {
		resource := resourceFromSecurityGroupAttribute(permission, alignedSecurityGroups[i])
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

// PostConvertHook Runs before HCL files are generated
func (g *SgGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_security_group" {
			// inner_access is deprecrated
			// https://www.terraform.io/docs/providers/alicloud/r/security_group.html#inner_access
			delete(r.Item, "inner_access")
		}
	}

	return nil
}
