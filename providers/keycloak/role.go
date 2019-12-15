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

type RoleGenerator struct {
	KeycloakService
}

var RoleAllowEmptyValues = []string{}
var RoleAdditionalFields = map[string]interface{}{}

func (g RoleGenerator) createResources(roles []*keycloak.Role) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, role := range roles {
		resources = append(resources, terraform_utils.NewResource(
			role.Id,
			"role_"+normalizeResourceName(role.RealmId)+normalizeResourceName(role.ContainerId)+"_"+normalizeResourceName(role.Name),
			"keycloak_role",
			"keycloak",
			map[string]string{
				"realm_id": role.RealmId,
			},
			RoleAllowEmptyValues,
			RoleAdditionalFields,
		))
	}
	return resources
}

func (g *RoleGenerator) InitResources() error {
	var rolesFull []*keycloak.Role
	client, err := keycloak.NewKeycloakClient(g.Args["url"].(string), g.Args["client_id"].(string), g.Args["client_secret"].(string), g.Args["realm"].(string), "", "", true, 5)
	if err != nil {
		return err
	}
	realms, err := client.GetRealms()
	if err != nil {
		return err
	}
	for _, realm := range realms {
		roles, err := client.GetRealmRoles(realm.Id)
		if err != nil {
			return err
		}
		openIDClients, err := client.GetOpenidClients(realm.Id, false)
		if err != nil {
			return err
		}
		mapContainerIDs := map[string]string{}
		for _, v := range openIDClients {
			mapContainerIDs[v.Id] = "_" + v.ClientId
		}
		mapContainerIDs[realm.Id] = ""
		clientRoles, err := client.GetClientRoles(realm.Id, openIDClients)
		if err != nil {
			return err
		}
		roles = append(roles, clientRoles...)
		for _, role := range roles {
			role.ContainerId = mapContainerIDs[role.ContainerId]
		}
		rolesFull = append(rolesFull, roles...)
	}
	g.Resources = g.createResources(rolesFull)
	return nil
}

func (g *RoleGenerator) PostConvertHook() error {
	mapRoleIDs := map[string]string{}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_role" {
			continue
		}
		mapRoleIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
	}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_role" {
			continue
		}
		if strings.Contains(r.InstanceState.Attributes["description"], "$") {
			r.Item["description"] = strings.ReplaceAll(r.InstanceState.Attributes["description"], "$", "$$")
		}
		if _, exist := r.Item["composite_roles"]; !exist {
			continue
		}
		renamedCompositeRoles := []string{}
		for _, v := range r.Item["composite_roles"].([]interface{}) {
			renamedCompositeRoles = append(renamedCompositeRoles, mapRoleIDs[v.(string)])
		}
		r.Item["composite_roles"] = renamedCompositeRoles
	}
	return nil
}
