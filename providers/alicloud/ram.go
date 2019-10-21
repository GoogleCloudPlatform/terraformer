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
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
)

// RAMGenerator Struct for generating AliCloud Elastic Compute Service
type RAMGenerator struct {
	AliCloudService
}

func resourceFromRAMRole(role ram.Role) terraform_utils.Resource {
	return terraform_utils.NewResource(
		role.RoleName,                  // id
		role.RoleId+"__"+role.RoleName, // name
		"alicloud_ram_role",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func resourceFromRAMPolicy(policy ram.Policy, roleName string) terraform_utils.Resource {
	// https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/alicloud/resource_alicloud_ram_role_policy_attachment.go#L93
	id := strings.Join([]string{"role", policy.PolicyName, policy.PolicyType, roleName}, ":")

	return terraform_utils.NewResource(
		id, // id
		id+"__"+roleName+"_"+policy.PolicyName, // name
		"alicloud_ram_role_policy_attachment",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initRoles(client *connectivity.AliyunClient) ([]ram.Role, error) {
	allRoles := make([]ram.Role, 0)

	raw, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
		request := ram.CreateListRolesRequest()
		request.RegionId = client.RegionId
		return ramClient.ListRoles(request)
	})
	if err != nil {
		return nil, err
	}

	response := raw.(*ram.ListRolesResponse)
	for _, role := range response.Roles.Role {
		allRoles = append(allRoles, role)

	}

	return allRoles, nil
}

func initRAMPolicyAttachment(client *connectivity.AliyunClient, allRoles []ram.Role) ([]ram.Policy, []string, error) {
	allRAMPolicies := make([]ram.Policy, 0)
	roleNames := make([]string, 0)

	for _, role := range allRoles {
		raw, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
			request := ram.CreateListPoliciesForRoleRequest()
			request.RegionId = client.RegionId
			request.RoleName = role.RoleName
			return ramClient.ListPoliciesForRole(request)
		})
		if err != nil {
			return nil, nil, err
		}

		response := raw.(*ram.ListPoliciesForRoleResponse)
		for _, policy := range response.Policies.Policy {
			allRAMPolicies = append(allRAMPolicies, policy)
			roleNames = append(roleNames, role.RoleName)
		}
	}

	return allRAMPolicies, roleNames, nil
}

// InitResources Gets the list of all ram role ids and generates resources
func (g *RAMGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}

	allRoles, err := initRoles(client)
	if err != nil {
		return err
	}

	allRAMPolicyAttachment, roleNames, err := initRAMPolicyAttachment(client, allRoles)
	if err != nil {
		return err
	}

	for _, role := range allRoles {
		resource := resourceFromRAMRole(role)
		g.Resources = append(g.Resources, resource)
	}

	for i, ramPolicy := range allRAMPolicyAttachment {
		resource := resourceFromRAMPolicy(ramPolicy, roleNames[i])
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

// PostConvertHook Runs before HCL files are generated
func (g *RAMGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_ram_role" {
			// https://www.terraform.io/docs/providers/alicloud/r/ram_role.html
			delete(r.Item, "services")  // deprecated
			delete(r.Item, "ram_users") // deprecated
			delete(r.Item, "version")   // deprecated
		}
	}

	return nil
}
