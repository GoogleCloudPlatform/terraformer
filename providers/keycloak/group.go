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
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

type GroupGenerator struct {
	KeycloakService
}

var GroupAllowEmptyValues = []string{}
var GroupAdditionalFields = map[string]interface{}{}

func (g GroupGenerator) createResources(groups []*keycloak.Group) []terraform_utils.Resource {
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
			GroupAllowEmptyValues,
			GroupAdditionalFields,
		))
	}
	return resources
}

func (g GroupGenerator) createDefaultGroupResource(realmId string) terraform_utils.Resource {
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

func (g GroupGenerator) createGroupMembershipsResource(realmId, groupId, groupName string, members []string) terraform_utils.Resource {
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

func (g GroupGenerator) createGroupRolesResource(realmId, groupId, groupName string, roles []string) terraform_utils.Resource {
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

func (g *GroupGenerator) flattenGroups(groups []*keycloak.Group, realmId, parentId string) []*keycloak.Group {
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

func (g *GroupGenerator) InitResources() error {
	var groupsFull []*keycloak.Group
	var groupMembers = []string{}
	var groupRoles = []string{}
	client, err := keycloak.NewKeycloakClient(g.Args["url"].(string), g.Args["client_id"].(string), g.Args["client_secret"].(string), g.Args["realm"].(string), "", "", true, 5)
	if err != nil {
		return err
	}
	realms, err := client.GetRealms()
	if err != nil {
		return err
	}
	for _, realm := range realms {
		groups, err := client.GetGroups(realm.Id)
		if err != nil {
			return err
		}
		groupsFull = append(groupsFull, groups...)
		g.Resources = append(g.Resources, g.createDefaultGroupResource(realm.Id))
	}
	flattenedGroups := g.flattenGroups(groupsFull, "", "")
	for _, group := range flattenedGroups {
		members, err := client.GetGroupMembers(group.RealmId, group.Id)
		if err != nil {
			return err
		}
		groupMembers = []string{}
		for _, member := range members {
			groupMembers = append(groupMembers, member.Username)
		}
		if len(groupMembers) > 0 {
			g.Resources = append(g.Resources, g.createGroupMembershipsResource(group.RealmId, group.Id, group.Name, groupMembers))
		}

		groupDetails, err := client.GetGroup(group.RealmId, group.Id)
		if err != nil {
			return err
		}
		groupRoles = []string{}
		if len(groupDetails.RealmRoles) > 0 {
			for _, realmRole := range groupDetails.RealmRoles {
				groupRoles = append(groupRoles, realmRole)
			}
		}
		if len(groupDetails.ClientRoles) > 0 {
			for _, clientRoles := range groupDetails.ClientRoles {
				for _, clientRole := range clientRoles {
					groupRoles = append(groupRoles, clientRole)
				}
			}
		}
		if len(groupRoles) > 0 {
			g.Resources = append(g.Resources, g.createGroupRolesResource(group.RealmId, group.Id, group.Name, groupRoles))
		}
	}
	g.Resources = append(g.Resources, g.createResources(flattenedGroups)...)
	return nil
}

func (g *GroupGenerator) PostConvertHook() error {
	mapGroupIDs := map[string]string{}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_group" {
			continue
		}
		mapGroupIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
	}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_group" && r.InstanceInfo.Type != "keycloak_default_groups" && r.InstanceInfo.Type != "keycloak_group_memberships" && r.InstanceInfo.Type != "keycloak_group_roles" {
			continue
		}
		if _, exist := r.Item["parent_id"]; exist && r.InstanceInfo.Type == "keycloak_group" {
			r.Item["parent_id"] = mapGroupIDs[r.Item["parent_id"].(string)]
		}
		if r.InstanceInfo.Type == "keycloak_group_memberships" || r.InstanceInfo.Type == "keycloak_group_roles" {
			r.Item["group_id"] = mapGroupIDs[r.Item["group_id"].(string)]
		}
		if r.InstanceInfo.Type == "keycloak_default_groups" {
			if _, exist := r.Item["group_ids"]; exist {
				renamedGroupIDs := []string{}
				for _, v := range r.Item["group_ids"].([]interface{}) {
					renamedGroupIDs = append(renamedGroupIDs, mapGroupIDs[v.(string)])
				}
				sort.Strings(renamedGroupIDs)
				r.Item["group_ids"] = renamedGroupIDs
			} else {
				r.Item["group_ids"] = []string{}
			}
		}
		if r.InstanceInfo.Type == "keycloak_group_memberships" {
			sortedMembers := []string{}
			for _, v := range r.Item["members"].([]interface{}) {
				sortedMembers = append(sortedMembers, v.(string))
			}
			sort.Strings(sortedMembers)
			r.Item["members"] = sortedMembers
		}

		if r.InstanceInfo.Type == "keycloak_group_roles" {
			sortedRoles := []string{}
			for _, v := range r.Item["role_ids"].([]interface{}) {
				sortedRoles = append(sortedRoles, v.(string))
			}
			sort.Strings(sortedRoles)
			r.Item["role_ids"] = sortedRoles
		}
	}
	return nil
}
