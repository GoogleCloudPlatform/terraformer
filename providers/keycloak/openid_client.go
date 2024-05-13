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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createOpenIDClientResources(openIDClients []*keycloak.OpenidClient) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, openIDClient := range openIDClients {
		resources = append(resources, terraformutils.NewResource(
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

func (g RealmGenerator) createServiceAccountClientRolesResources(realmID string, clientRoles []*keycloak.Role, usersInRole []keycloak.UsersInRole, mapServiceAccountIds map[string]map[string]string, mapClientIDs map[string]string) []terraformutils.Resource {
	var resources []terraformutils.Resource
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
				resources = append(resources, terraformutils.NewResource(
					realmID+"/"+user.Id+"/"+role.ClientId+"/"+role.Name,
					"openid_client_service_account_role_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(mapServiceAccountIds[user.Id]["ClientId"])+"_"+normalizeResourceName(mapClientIDs[role.ClientId])+"_"+normalizeResourceName(role.Name),
					"keycloak_openid_client_service_account_role",
					"keycloak",
					map[string]string{
						"realm_id":                realmID,
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

func (g RealmGenerator) createOpenIDGenericProtocolMapperResource(protocolMapperType, protocolMapperID, protocolMapperName, realmID, clientID, clientName string) terraformutils.Resource {
	return terraformutils.NewResource(
		protocolMapperID,
		"openid_"+protocolMapperType+"_protocol_mapper_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(clientName)+"_"+normalizeResourceName(protocolMapperName),
		"keycloak_openid_"+protocolMapperType+"_protocol_mapper",
		"keycloak",
		map[string]string{
			"realm_id":  realmID,
			"client_id": clientID,
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g RealmGenerator) createOpenIDProtocolMapperResources(clientID string, openidClient *keycloak.OpenidClientWithGenericProtocolMappers) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, protocolMapper := range openidClient.ProtocolMappers {
		switch protocolMapper.ProtocolMapper {
		case "oidc-audience-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("audience", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-audience-resolve-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("audience_resolve", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-full-name-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("full_name", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-group-membership-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("group_membership", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-hardcoded-claim-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("hardcoded_claim", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-hardcoded-role-mapper":
			// Only works with client roles
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("hardcoded_role", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-script-based-protocol-mapper":
			// Support for this protocol mapper was removed in Keycloak 18
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("script", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-usermodel-attribute-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("user_attribute", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-usermodel-property-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("user_property", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-usermodel-realm-role-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("user_realm_role", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-usermodel-client-role-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("user_client_role", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-usersessionmodel-note-mapper":
			resources = append(resources, g.createOpenIDGenericProtocolMapperResource("user_session_note", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
		case "oidc-address-mapper":
			// Not supported for the moment
			// resources = append(resources, g.createOpenIDGenericProtocolMapperResource("address", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
			continue
		case "oidc-role-name-mapper":
			// Not supported for the moment
			// resources = append(resources, g.createOpenIDGenericProtocolMapperResource("role_name", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
			continue
		case "oidc-sha256-pairwise-sub-mapper":
			// Not supported for the moment
			// resources = append(resources, g.createOpenIDGenericProtocolMapperResource("pairwise_subject_identifier", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
			continue
		case "oidc-allowed-origins-mapper":
			// Not supported for the moment
			// resources = append(resources, g.createOpenIDGenericProtocolMapperResource("allowed_web_origins", protocolMapper.Id, protocolMapper.Name, openidClient.RealmId, openidClient.ClientId, clientID))
			continue
		}
	}
	return resources
}
