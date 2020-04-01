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

package keycloak

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createGroupResources(groups []*keycloak.Group) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, group := range groups {
		resources = append(resources, terraform_utils.NewResource(
			group.Id,
			"group_"+normalizeResourceName(group.RealmId)+"_"+normalizeResourceName(group.Name),
			"keycloak_group",
			"keycloak",
			map[string]string{
				"realm_id": group.RealmId,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g RealmGenerator) createDefaultGroupResource(realmId string) terraform_utils.Resource {
	return terraform_utils.NewResource(
		realmId+"/default-groups",
		"default_groups_"+normalizeResourceName(realmId),
		"keycloak_default_groups",
		"keycloak",
		map[string]string{
			"realm_id": realmId,
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g RealmGenerator) createGroupMembershipsResource(realmId, groupId, groupName string, members []string) terraform_utils.Resource {
	return terraform_utils.NewResource(
		realmId+"/group-memberships/"+groupId,
		"group_memberships_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(groupName),
		"keycloak_group_memberships",
		"keycloak",
		map[string]string{
			"realm_id": realmId,
			"group_id": groupId,
			"members":  strings.Join(members, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g RealmGenerator) createGroupRolesResource(realmId, groupId, groupName string, roles []string) terraform_utils.Resource {
	return terraform_utils.NewResource(
		realmId+"/"+groupId,
		"group_roles_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(groupName),
		"keycloak_group_roles",
		"keycloak",
		map[string]string{
			"realm_id":  realmId,
			"group_id":  groupId,
			"roles_ids": strings.Join(roles, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g *RealmGenerator) flattenGroups(groups []*keycloak.Group, realmId, parentId string) []*keycloak.Group {
	var flattenedGroups []*keycloak.Group
	for _, group := range groups {
		if realmId != "" {
			group.RealmId = realmId
		}
		flattenedGroups = append(flattenedGroups, group)
		if len(group.SubGroups) > 0 {
			subGroups := g.flattenGroups(group.SubGroups, group.RealmId, group.Id)
			flattenedGroups = append(flattenedGroups, subGroups...)
		}
	}
	return flattenedGroups
}
