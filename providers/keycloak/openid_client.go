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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createOpenIdClientResources(openIDClients []*keycloak.OpenidClient) []terraform_utils.Resource {
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
			[]string{"web_origins"},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g RealmGenerator) createServiceAccountClientRolesResources(realmId string, clientRoles []*keycloak.Role, usersInRole []keycloak.UsersInRole, mapServiceAccountIds map[string]map[string]string, mapClientIDs map[string]string) []terraform_utils.Resource {
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

func (g RealmGenerator) createOpenIdGenericProtocolMapperResource(protocolMapperType, protocolMapperId, protocolMapperName, realmId, clientId, clientName string) terraform_utils.Resource {
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

func (g RealmGenerator) createOpenIdProtocolMapperResources(clientId string, openidClient *keycloak.OpenidClientWithGenericClientProtocolMappers) []terraform_utils.Resource {
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
		case "oidc-hardcoded-group-mapper":
			resources = append(resources, g.createOpenIdGenericProtocolMapperResource("hardcoded_group", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientId))
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
