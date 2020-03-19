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
	"errors"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

type OpenIDClientGenerator struct {
	KeycloakService
}

var OpenIDClientAllowEmptyValues = []string{"web_origins"}
var OpenIDClientAdditionalFields = map[string]interface{}{}

func (g OpenIDClientGenerator) createResources(openIDClients []*keycloak.OpenidClient) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, openIDClient := range openIDClients {
		resources = append(resources, terraform_utils.NewResource(
			openIDClient.Id,
			"openid_client_"+normalizeResourceName(openIDClient.RealmId)+"_"+normalizeResourceName(openIDClient.ClientId),
			"keycloak_openid_client",
			"keycloak",
			map[string]string{
				"realm_id": openIDClient.RealmId,
			},
			OpenIDClientAllowEmptyValues,
			OpenIDClientAdditionalFields,
		))
	}
	return resources
}

func (g OpenIDClientGenerator) createServiceAccountClientRolesResources(realmId string, clientRoles []*keycloak.Role, usersInRole []keycloak.UsersInRole, mapServiceAccountIds map[string]map[string]string, mapClientIDs map[string]string) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, role := range clientRoles {
		for _, users := range usersInRole {
			if len(*users.Users) == 0 || role.Id != users.Role.Id {
				continue
			}
			for _, user := range *users.Users {
				// Test if role is mapped to a User, and not a ServiceAccountUser
				if mapServiceAccountIds[user.Id] == nil {
					continue
				}
				resources = append(resources, terraform_utils.NewResource(
					realmId+"/"+user.Id+"/"+role.ClientId+"/"+role.Name,
					"openid_client_service_account_role_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(mapServiceAccountIds[user.Id]["ClientId"])+"_"+normalizeResourceName(mapClientIDs[role.ClientId])+"_"+normalizeResourceName(role.Name),
					"keycloak_openid_client_service_account_role",
					"keycloak",
					map[string]string{
						"realm_id":                realmId,
						"service_account_user_id": user.Id,
						"client_id":               role.ClientId,
						"role":                    role.Name,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}

	return resources
}

func (g OpenIDClientGenerator) createOpenIdGenericProtocolMapperResource(protocolMapperType, protocolMapperId, protocolMapperName, realmId, clientId, clientName string) terraform_utils.Resource {
	return terraform_utils.NewResource(
		protocolMapperId,
		"openid_"+protocolMapperType+"_protocol_mapper_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(clientName)+"_"+normalizeResourceName(protocolMapperName),
		"keycloak_openid_"+protocolMapperType+"_protocol_mapper",
		"keycloak",
		map[string]string{
			"realm_id":  realmId,
			"client_id": clientId,
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g OpenIDClientGenerator) createOpenIdProtocolMapperResources(clientId string, openidClient *keycloak.OpenidClientWithGenericClientProtocolMappers) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, protocolMapper := range openidClient.ProtocolMappers {
		switch protocolMapper.ProtocolMapper {
		case "oidc-audience-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("audience", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-full-name-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("full_name", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-group-membership-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("group_membership", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-hardcoded-claim-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("hardcoded_claim", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-hardcoded-role-mapper":
			// Not supported for the moment
			// Only works with client roles
			//resources = append(resources, g.createOpenIdGenericProtocolMapperResource("hardcoded_role", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
			continue
		case "oidc-usermodel-attribute-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("user_attribute", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-usermodel-property-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("user_property", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
		case "oidc-usermodel-realm-role-mapper":
			// Not supported for the moment
			//resources = append(resources, g.createOpenIdGenericProtocolMapperResource("user_realm_role", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
			continue
		}
	}
	return resources
}

func (g *OpenIDClientGenerator) InitResources() error {
	var openIDClientsFull []*keycloak.OpenidClient
	mapClientIDs := map[string]string{}
	client, err := keycloak.NewKeycloakClient(g.Args["url"].(string), g.Args["client_id"].(string), g.Args["client_secret"].(string), g.Args["realm"].(string), "", "", true, 5)
	if err != nil {
		return errors.New("keycloak: could not connect to Keycloak")
	}
	var realms []*keycloak.Realm
	if g.Args["target"].(string) == "" {
		realms, err = client.GetRealms()
		if err != nil {
			return err
		}
	} else {
		realm, err := client.GetRealm(g.Args["target"].(string))
		if err != nil {
			return err
		}
		realms = append(realms, realm)
	}
	for _, realm := range realms {
		openIDClients, err := client.GetOpenidClients(realm.Id, true)
		if err != nil {
			return err
		}
		mapServiceAccountIds := map[string]map[string]string{}
		for _, openIDClient := range openIDClients {
			mapClientIDs[openIDClient.Id] = openIDClient.ClientId
			openidClientWithGenericClientProtocolMappers, err := client.GetGenericClientProtocolMappers(realm.Id, openIDClient.Id)
			if err != nil {
				return err
			}
			g.Resources = append(g.Resources, g.createOpenIdProtocolMapperResources(openIDClient.ClientId, openidClientWithGenericClientProtocolMappers)...)

			if !openIDClient.ServiceAccountsEnabled {
				continue
			}
			serviceAccountUser, err := client.GetOpenidClientServiceAccountUserId(realm.Id, openIDClient.Id)
			if err != nil {
				return err
			}
			mapServiceAccountIds[serviceAccountUser.Id] = map[string]string{}
			mapServiceAccountIds[serviceAccountUser.Id]["Id"] = openIDClient.Id
			mapServiceAccountIds[serviceAccountUser.Id]["ClientId"] = openIDClient.ClientId
		}
		openIDClientsFull = append(openIDClientsFull, openIDClients...)

		clientRoles, err := client.GetClientRoles(realm.Id, openIDClients)
		if err != nil {
			return err
		}

		usersInRole, err := client.GetClientRoleUsers(realm.Id, clientRoles)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, g.createServiceAccountClientRolesResources(realm.Id, clientRoles, *usersInRole, mapServiceAccountIds, mapClientIDs)...)
	}
	g.Resources = append(g.Resources, g.createResources(openIDClientsFull)...)
	return nil
}

func (g *OpenIDClientGenerator) PostConvertHook() error {
	mapClientIDs := map[string]string{}
	mapClientClientIDs := map[string]string{}
	mapServiceAccountUserIDs := map[string]string{}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_openid_client" {
			continue
		}
		mapClientIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		mapClientClientIDs[r.InstanceState.Attributes["client_id"]] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".client_id}"
		if _, exist := r.InstanceState.Attributes["service_account_user_id"]; exist {
			mapServiceAccountUserIDs[r.InstanceState.Attributes["service_account_user_id"]] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".service_account_user_id}"
		}
	}
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_openid_client" &&
			r.InstanceInfo.Type != "keycloak_openid_client_service_account_role" &&
			r.InstanceInfo.Type != "keycloak_openid_audience_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_full_name_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_group_membership_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_hardcoded_claim_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_hardcoded_role_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_user_attribute_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_user_property_protocol_mapper" &&
			r.InstanceInfo.Type != "keycloak_openid_user_realm_role_protocol_mapper" {
			continue
		}
		if _, exist := r.Item["client_id"]; exist && (r.InstanceInfo.Type == "keycloak_openid_client_service_account_role" ||
			r.InstanceInfo.Type == "keycloak_openid_audience_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_full_name_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_group_membership_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_hardcoded_claim_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_hardcoded_role_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_attribute_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_property_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_realm_role_protocol_mapper") {
			r.Item["client_id"] = mapClientIDs[r.Item["client_id"].(string)]
		}
		if r.InstanceInfo.Type == "keycloak_openid_client" {
			if _, exist := r.Item["valid_redirect_uris"]; exist {
				sortedValidRedirectUris := []string{}
				for _, v := range r.Item["valid_redirect_uris"].([]interface{}) {
					sortedValidRedirectUris = append(sortedValidRedirectUris, v.(string))
				}
				sort.Strings(sortedValidRedirectUris)
				r.Item["valid_redirect_uris"] = sortedValidRedirectUris
			}

			if _, exist := r.Item["web_origins"]; exist {
				sortedWebOrigins := []string{}
				for _, v := range r.Item["web_origins"].([]interface{}) {
					sortedWebOrigins = append(sortedWebOrigins, v.(string))
				}
				sort.Strings(sortedWebOrigins)
				r.Item["web_origins"] = sortedWebOrigins
			}
		}
		if _, exist := r.Item["included_client_audience"]; exist && r.InstanceInfo.Type == "keycloak_openid_audience_protocol_mapper" {
			r.Item["included_client_audience"] = mapClientClientIDs[r.Item["included_client_audience"].(string)]
		}
		if _, exist := r.Item["service_account_user_id"]; exist && r.InstanceInfo.Type == "keycloak_openid_client_service_account_role" {
			r.Item["service_account_user_id"] = mapServiceAccountUserIDs[r.Item["service_account_user_id"].(string)]
		}
		if strings.Contains(r.InstanceState.Attributes["name"], "$") {
			g.Resources[i].Item["name"] = strings.ReplaceAll(r.InstanceState.Attributes["name"], "$", "$$")
		}
	}
	return nil
}
