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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createGroupResources(groups []*keycloak.Group) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, group := range groups {
		resources = append(resources, terraformutils.NewResource(
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

func (g RealmGenerator) createDefaultGroupResource(realmID string) terraformutils.Resource {
	return terraformutils.NewResource(
		realmID+"/default-groups",
		"default_groups_"+normalizeResourceName(realmID),
		"keycloak_default_groups",
		"keycloak",
		map[string]string{
			"realm_id": realmID,
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g RealmGenerator) createGroupMembershipsResource(realmID, groupID, groupName string, members []string) terraformutils.Resource {
	return terraformutils.NewResource(
		realmID+"/group-memberships/"+groupID,
		"group_memberships_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(groupName),
		"keycloak_group_memberships",
		"keycloak",
		map[string]string{
			"realm_id": realmID,
			"group_id": groupID,
			"members":  strings.Join(members, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g RealmGenerator) createGroupRolesResource(realmID, groupID, groupName string, roles []string) terraformutils.Resource {
	return terraformutils.NewResource(
		realmID+"/"+groupID,
		"group_roles_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(groupName),
		"keycloak_group_roles",
		"keycloak",
		map[string]string{
			"realm_id":  realmID,
			"group_id":  groupID,
			"roles_ids": strings.Join(roles, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g *RealmGenerator) flattenGroups(groups []*keycloak.Group, realmID string) []*keycloak.Group {
	var flattenedGroups []*keycloak.Group
	for _, group := range groups {
		if realmID != "" {
			group.RealmId = realmID
		}
		flattenedGroups = append(flattenedGroups, group)
		if len(group.SubGroups) > 0 {
			subGroups := g.flattenGroups(group.SubGroups, group.RealmId)
			flattenedGroups = append(flattenedGroups, subGroups...)
		}
	}
	return flattenedGroups
}
