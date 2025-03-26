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
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

type RealmGenerator struct {
	KeycloakService
}

func (g *RealmGenerator) InitResources() error {
	var realms []*keycloak.Realm
	var realmsGroups []*keycloak.Group
	ctx := context.TODO()

	// Connect to keycloak instance
	userAgent := "GoogleCloudPlatform Terraformer/0.8.22 (+https://github.com/GoogleCloudPlatform/terraformer) Terraform Plugin SDK/2.10.1"
	kck, err := keycloak.NewKeycloakClient(ctx, g.GetArgs()["url"].(string), g.GetArgs()["base_path"].(string), g.GetArgs()["client_id"].(string), g.GetArgs()["client_secret"].(string), g.GetArgs()["realm"].(string), "", "", true, g.GetArgs()["client_timeout"].(int), g.GetArgs()["root_ca_certificate"].(string), g.GetArgs()["tls_insecure_skip_verify"].(bool), userAgent, g.GetArgs()["red_hat_sso"].(bool), make(map[string]string))
	if err != nil {
		return errors.New("keycloak: could not connect to Keycloak")
	}

	// Get realm resources
	target := g.GetArgs()["target"].(string)
	if target == "" {
		realms, err = kck.GetRealms(ctx)
		if err != nil {
			return errors.New("keycloak: could not get realms attributes in Keycloak")
		}
	} else {
		realm, err := kck.GetRealm(ctx, target)
		if err != nil {
			return errors.New("keycloak: could not get " + target + " realm attributes in Keycloak")
		}
		realms = append(realms, realm)
	}
	g.Resources = append(g.Resources, g.createRealmResources(realms)...)

	// For each realm, get resources
	for _, realm := range realms {
		// Get required actions resources
		requiredActions, err := kck.GetRequiredActions(ctx, realm.Realm)
		if err != nil {
			return fmt.Errorf("keycloak: could not get required actions of realm %s in Keycloak. err: %w", realm.Realm, err)
		}
		g.Resources = append(g.Resources, g.createRequiredActionResources(requiredActions)...)

		// Get top-level authentication flows resources
		authenticationFlows, err := kck.ListAuthenticationFlows(ctx, realm.Realm)
		if err != nil {
			return fmt.Errorf("keycloak: could not get authentication flows of realm %s in Keycloak. err: %w", realm.Realm, err)
		}
		g.Resources = append(g.Resources, g.createAuthenticationFlowResources(authenticationFlows)...)

		// For each authentication flow, get subFlow, execution and execution config resources
		for _, topLevelAuthenticationFlow := range authenticationFlows {
			authenticationSubFlowOrExecutions, err := kck.ListAuthenticationExecutions(ctx, realm.Realm, topLevelAuthenticationFlow.Alias)
			if err != nil {
				return fmt.Errorf("keycloak: could not get authentication executions of authentication flow %s of realm %s in Keycloak. err: %w",
					topLevelAuthenticationFlow.Alias, realm.Realm, err)
			}

			var stack []*keycloak.AuthenticationExecutionInfo
			parentFlowAlias := topLevelAuthenticationFlow.Alias

			for _, authenticationSubFlowOrExecution := range authenticationSubFlowOrExecutions {

				// Find the parent flow alias
				if len(stack) > 0 {
					previous := stack[len(stack)-1]
					if authenticationSubFlowOrExecution.Level < previous.Level {
						// Find the last sub flow/execution for the current level
						stack = stack[:authenticationSubFlowOrExecution.Level+1]
						previous = stack[len(stack)-1]
					}
					if authenticationSubFlowOrExecution.Level == previous.Level {
						// Same level sub flow/execution, it means that the sub flow/execution has same parent flow of the last sub flow/execution
						parentFlowAlias = previous.ParentFlowAlias

					} else if authenticationSubFlowOrExecution.Level > previous.Level {
						// Deep level sub flow/execution, it means that the parent flow is the last sub flow/execution
						if previous.AuthenticationFlow {
							parentFlowAlias = previous.Alias
						} else {
							return errors.New("keycloak: invalid parent sub flow, it should be a sub flow but it's an execution")
						}
					}
				}

				var resource terraformutils.Resource

				switch authenticationSubFlowOrExecution.AuthenticationFlow {
				case true:
					authenticationSubFlow, err := kck.GetAuthenticationSubFlow(ctx, realm.Realm, parentFlowAlias, authenticationSubFlowOrExecution.FlowId)
					if err != nil {
						return fmt.Errorf("keycloak: could not get authentication subflow %s of realm %s in Keycloak. err: %w",
							authenticationSubFlowOrExecution.FlowId, realm.Realm, err)
					}

					// Need to store the alias and parent flow alias
					authenticationSubFlowOrExecution.Alias = authenticationSubFlow.Alias
					authenticationSubFlowOrExecution.ParentFlowAlias = parentFlowAlias

					resource = g.createAuthenticationSubFlowResource(authenticationSubFlow)
					g.Resources = append(g.Resources, resource)

				case false:
					authenticationExecution, err := kck.GetAuthenticationExecution(ctx, realm.Realm, parentFlowAlias, authenticationSubFlowOrExecution.Id)
					if err != nil {
						return fmt.Errorf("keycloak: could not get authentication execution %s of realm %s in Keycloak. err: %w",
							authenticationSubFlowOrExecution.Id, realm.Realm, err)
					}

					// Need to store the parent flow alias
					authenticationSubFlowOrExecution.ParentFlowAlias = parentFlowAlias

					resource = g.createAuthenticationExecutionResource(authenticationExecution)
					g.Resources = append(g.Resources, resource)

					if authenticationSubFlowOrExecution.AuthenticationConfig != "" {
						authenticationExecutionConfig := &keycloak.AuthenticationExecutionConfig{
							RealmId:     realm.Realm,
							Id:          authenticationSubFlowOrExecution.AuthenticationConfig,
							ExecutionId: authenticationSubFlowOrExecution.Id,
						}
						err := kck.GetAuthenticationExecutionConfig(ctx, authenticationExecutionConfig)
						if err != nil {
							return fmt.Errorf("keycloak: could not get authentication execution config %s of realm %s in Keycloak. err: %w",
								authenticationExecutionConfig.Id, realm.Realm, err)
						}

						g.Resources = append(g.Resources, g.createAuthenticationExecutionConfigResource(authenticationExecutionConfig))
					}
				}

				if len(stack) > 0 && authenticationSubFlowOrExecution.Index > 0 {
					previous := stack[len(stack)-1]
					var resourceType string
					var resourceName string
					if previous.AuthenticationFlow {
						resourceType = "keycloak_authentication_subflow"
						resourceName = "authentication_subflow_" +
							normalizeResourceName(realm.Realm) + "_" + normalizeResourceName(previous.FlowId)
					} else {
						resourceType = "keycloak_authentication_execution"
						resourceName = "authentication_execution_" +
							normalizeResourceName(realm.Realm) + "_" + normalizeResourceName(previous.Id)
					}
					resource.AdditionalFields["depends_on"] = []string{resourceType + "." + terraformutils.TfSanitize(resourceName)}
				}

				// Stack the current sub flow/execution
				if len(stack) > 0 && stack[len(stack)-1].Level == authenticationSubFlowOrExecution.Level {
					// Replace it if it's same level
					stack[len(stack)-1] = authenticationSubFlowOrExecution
				} else {
					stack = append(stack, authenticationSubFlowOrExecution)
				}
			}
		}

		// Get custom federations resources
		// TODO: support kerberos user federation
		customUserFederations, err := kck.GetCustomUserFederations(ctx, realm.Realm, realm.Id)
		if err != nil {
			return errors.New("keycloak: could not get custom user federations of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createCustomUserFederationResources(customUserFederations)...)

		// For each custom federation, get mappers resources
		for _, customUserFederation := range *customUserFederations {
			if customUserFederation.ProviderId == "ldap" {
				mappers, err := kck.GetLdapUserFederationMappers(ctx, realm.Realm, customUserFederation.Id)
				if err != nil {
					return errors.New("keycloak: could not get mappers of ldap user federation " + customUserFederation.Name + " of realm " + realm.Realm + " in Keycloak")
				}
				g.Resources = append(g.Resources, g.createLdapMapperResources(realm.Realm, customUserFederation.Name, mappers)...)
			}
		}

		// Get groups tree and default groups resources
		realmGroups, err := kck.GetGroups(ctx, realm.Realm)
		if err != nil {
			return errors.New("keycloak: could not get groups of realm " + realm.Realm + " in Keycloak")
		}
		realmsGroups = append(realmsGroups, realmGroups...)
		g.Resources = append(g.Resources, g.createDefaultGroupResource(realm.Realm))

		// Get users resources
		realmUsers, err := kck.GetUsers(ctx, realm.Realm)
		if err != nil {
			return errors.New("keycloak: could not get users of realm " + realm.Realm + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createUserResources(realmUsers)...)

		// Get realm open id client scopes resources
		realmScopes, err := kck.ListOpenidClientScopesWithFilter(ctx, realm.Realm, func(scope *keycloak.OpenidClientScope) bool { return true })
		if err != nil {
			return errors.New("keycloak: could not get realm scopes of realm " + realm.Realm + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createScopeResources(realm.Realm, realmScopes)...)

		// Get open id clients
		realmClients, err := kck.GetOpenidClients(ctx, realm.Realm, true)
		if err != nil {
			return errors.New("keycloak: could not get open id clients of realm " + realm.Realm + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createOpenIDClientResources(realmClients)...)

		// For each open id client, get resources
		mapServiceAccountIds := map[string]map[string]string{}
		mapContainerIDs := map[string]string{}
		mapClientIDs := map[string]string{}
		for _, client := range realmClients {
			mapClientIDs[client.Id] = client.ClientId
			mapContainerIDs[client.Id] = "_" + client.ClientId

			// Get open id client protocol mappers resources
			clientMappers, err := kck.GetGenericProtocolMappers(ctx, realm.Realm, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get protocol mappers of open id client " + client.ClientId + " of realm " + realm.Realm + " in Keycloak")
			}
			g.Resources = append(g.Resources, g.createOpenIDProtocolMapperResources(client.ClientId, clientMappers)...)

			// Get open id client default scopes resources
			clientScopes, err := kck.GetOpenidDefaultClientScopes(ctx, realm.Realm, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get default client scopes of open id client " + client.ClientId + " of realm " + realm.Realm + " in Keycloak")
			}
			if len(*clientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Realm, client.Id, client.ClientId, "default", clientScopes))
			}

			// Get open id client optional scopes resources
			clientScopes, err = kck.GetOpenidOptionalClientScopes(ctx, realm.Realm, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get optional client scopes of open id client " + client.ClientId + " of realm " + realm.Realm + " in Keycloak")
			}
			if len(*clientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Realm, client.Id, client.ClientId, "optional", clientScopes))
			}

			// Prepare a slice to be able to link roles associated to service account roles to be associated to the open id client, only if service accounts are enabled
			if !client.ServiceAccountsEnabled {
				continue
			}
			serviceAccountUser, err := kck.GetOpenidClientServiceAccountUserId(ctx, realm.Realm, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get service account user associated to open id client " + client.ClientId + " of realm " + realm.Realm + " in Keycloak")
			}
			mapServiceAccountIds[serviceAccountUser.Id] = map[string]string{}
			mapServiceAccountIds[serviceAccountUser.Id]["Id"] = client.Id
			mapServiceAccountIds[serviceAccountUser.Id]["ClientId"] = client.ClientId
		}

		// Get open id client roles
		clientRoles, err := kck.GetClientRoles(ctx, realm.Realm, realmClients)
		if err != nil {
			return errors.New("keycloak: could not get open id clients roles of realm " + realm.Realm + " in Keycloak")
		}

		// Get roles
		realmRoles, err := kck.GetRealmRoles(ctx, realm.Realm)
		if err != nil {
			return errors.New("keycloak: could not get realm roles of realm " + realm.Realm + " in Keycloak")
		}

		// Set ContainerId of the roles, for realm = "", for open id clients = "_" + client.ClientId
		// and get roles resources
		mapContainerIDs[realm.Realm] = ""
		roles := append(clientRoles, realmRoles...)
		for _, role := range roles {
			role.ContainerId = mapContainerIDs[role.ContainerId]
		}
		g.Resources = append(g.Resources, g.createRoleResources(roles)...)

		// Get service account roles resources
		usersInRole, err := kck.GetClientRoleUsers(ctx, realm.Realm, clientRoles)
		if err != nil {
			return errors.New("keycloak: could not get users roles of realm " + realm.Realm + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createServiceAccountClientRolesResources(realm.Realm, clientRoles, *usersInRole, mapServiceAccountIds, mapClientIDs)...)
	}

	// Parse the groups trees, and get all the groups
	// Get groups resources
	groups := g.flattenGroups(realmsGroups, "")
	g.Resources = append(g.Resources, g.createGroupResources(groups)...)

	// For each group, get group memberships and roles resources
	for _, group := range groups {
		// Get group members resources
		members, err := kck.GetGroupMembers(ctx, group.RealmId, group.Id)
		if err != nil {
			return errors.New("keycloak: could not get group members of group " + group.Name + " in Keycloak")
		}
		if len(members) > 0 {
			groupMembers := make([]string, len(members))
			for k, member := range members {
				groupMembers[k] = member.Username
			}
			g.Resources = append(g.Resources, g.createGroupMembershipsResource(group.RealmId, group.Id, group.Name, groupMembers))
		}

		// Get group roles resources
		// For realm roles and open id clients roles
		groupDetails, err := kck.GetGroup(ctx, group.RealmId, group.Id)
		if err != nil {
			return errors.New("keycloak: could not get details about group " + group.Name + " in Keycloak")
		}
		groupRoles := []string{}
		if len(groupDetails.RealmRoles) > 0 {
			groupRoles = append(groupRoles, groupDetails.RealmRoles...)
		}
		if len(groupDetails.ClientRoles) > 0 {
			for _, clientRoles := range groupDetails.ClientRoles {
				groupRoles = append(groupRoles, clientRoles...)
			}
		}
		if len(groupRoles) > 0 {
			g.Resources = append(g.Resources, g.createGroupRolesResource(group.RealmId, group.Id, group.Name, groupRoles))
		}
	}

	return nil
}

func (g *RealmGenerator) PostConvertHook() error {
	mapRealmIDs := map[string]string{}
	mapUserFederationIDs := map[string]string{}
	mapGroupIDs := map[string]string{}
	mapClientIDs := map[string]string{}
	mapClientNames := map[string]string{}
	mapClientClientIDs := map[string]string{}
	mapClientClientNames := map[string]string{}
	mapServiceAccountUserIDs := map[string]string{}
	mapRoleIDs := map[string]string{}
	mapClientRoleNames := map[string]string{}
	mapClientRoleShortNames := map[string]string{}
	mapScopeNames := map[string]string{}
	mapUserNames := map[string]string{}
	mapGroupNames := map[string]string{}
	mapAuthenticationFlowAliases := map[string]string{}
	mapAuthenticationExecutionIDs := map[string]string{}

	// Set slices to be able to map IDs with Terraform variables
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_realm" &&
			r.InstanceInfo.Type != "keycloak_ldap_user_federation" &&
			r.InstanceInfo.Type != "keycloak_group" &&
			r.InstanceInfo.Type != "keycloak_openid_client" &&
			r.InstanceInfo.Type != "keycloak_role" &&
			r.InstanceInfo.Type != "keycloak_openid_client_scope" &&
			r.InstanceInfo.Type != "keycloak_user" &&
			r.InstanceInfo.Type != "keycloak_authentication_flow" &&
			r.InstanceInfo.Type != "keycloak_authentication_subflow" &&
			r.InstanceInfo.Type != "keycloak_authentication_execution" {
			continue
		}
		if r.InstanceInfo.Type == "keycloak_realm" {
			mapRealmIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		}
		if r.InstanceInfo.Type == "keycloak_ldap_user_federation" {
			mapUserFederationIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		}
		if r.InstanceInfo.Type == "keycloak_group" {
			mapGroupIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
			mapGroupNames[r.Item["realm_id"].(string)+"_"+r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
		}
		if r.InstanceInfo.Type == "keycloak_openid_client" {
			mapClientIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
			mapClientNames[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = r.Item["client_id"].(string)
			mapClientClientNames[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".client_id}"
			mapClientClientIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.Attributes["client_id"]] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".client_id}"
			if _, exist := r.InstanceState.Attributes["service_account_user_id"]; exist {
				mapServiceAccountUserIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.Attributes["service_account_user_id"]] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".service_account_user_id}"
			}
		}
		if r.InstanceInfo.Type == "keycloak_openid_client_scope" {
			mapScopeNames[r.Item["realm_id"].(string)+"_"+r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
		}
		if r.InstanceInfo.Type == "keycloak_user" {
			mapUserNames[r.Item["realm_id"].(string)+"_"+r.Item["username"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".username}"
		}
		if r.InstanceInfo.Type == "keycloak_authentication_flow" || r.InstanceInfo.Type == "keycloak_authentication_subflow" {
			mapAuthenticationFlowAliases[r.Item["realm_id"].(string)+"_"+r.Item["alias"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".alias}"
		}
		if r.InstanceInfo.Type == "keycloak_authentication_execution" {
			mapAuthenticationExecutionIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		}
	}

	// Set slices to be able to map IDs with Terraform variables
	// Separate loop for roles to avoid fetching with a key that is not present in the map
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "keycloak_role" {
			mapRoleIDs[r.Item["realm_id"].(string)+"_"+r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
			if _, exist := r.Item["client_id"]; exist {
				mapClientRoleNames[r.Item["realm_id"].(string)+"_"+mapClientNames[r.Item["realm_id"].(string)+"_"+r.Item["client_id"].(string)]+"."+r.Item["name"].(string)] = mapClientClientNames[r.Item["realm_id"].(string)+"_"+r.Item["client_id"].(string)] + ".${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
				mapClientRoleShortNames[r.Item["realm_id"].(string)+"_"+mapClientNames[r.Item["realm_id"].(string)+"_"+r.Item["client_id"].(string)]+"."+r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			} else {
				mapClientRoleNames[r.Item["realm_id"].(string)+"_"+r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
			}
		}
	}

	// For each resource, modify import if needed...
	for i, r := range g.Resources {
		// Escape keycloak text inputs not to get unpredictable results or errors when Terraform will try to interpret variables ($ vs $$)
		// TODO: ensure that we escape all existing fields
		if strings.Contains(r.InstanceState.Attributes["consent_screen_text"], "$") {
			g.Resources[i].Item["consent_screen_text"] = strings.ReplaceAll(r.InstanceState.Attributes["consent_screen_text"], "$", "$$")
		}
		if strings.Contains(r.InstanceState.Attributes["name"], "$") {
			g.Resources[i].Item["name"] = strings.ReplaceAll(r.InstanceState.Attributes["name"], "$", "$$")
		}
		if strings.Contains(r.InstanceState.Attributes["description"], "$") {
			g.Resources[i].Item["description"] = strings.ReplaceAll(r.InstanceState.Attributes["description"], "$", "$$")
		}
		if strings.Contains(r.InstanceState.Attributes["root_url"], "$") {
			g.Resources[i].Item["root_url"] = strings.ReplaceAll(r.InstanceState.Attributes["root_url"], "$", "$$")
		}

		// Sort supported_locales to get reproducible results for keycloak_realm resources
		if r.InstanceInfo.Type == "keycloak_realm" {
			if _, exist := r.Item["internationalization"]; exist {
				for _, v := range r.Item["internationalization"].([]interface{}) {
					sortedSupportedLocales := make([]string, len(v.(map[string]interface{})["supported_locales"].([]interface{})))
					for k, vv := range v.(map[string]interface{})["supported_locales"].([]interface{}) {
						sortedSupportedLocales[k] = vv.(string)
					}
					sort.Strings(sortedSupportedLocales)
					v.(map[string]interface{})["supported_locales"] = sortedSupportedLocales
				}
			}
		}

		// Sort group_ids to get reproducible results for keycloak_default_groups resources
		// Set an empty string slice if the attribute doesn't exist as it is mandatory
		if r.InstanceInfo.Type == "keycloak_default_groups" {
			if _, exist := r.Item["group_ids"]; exist {
				renamedGroupIDs := make([]string, len(r.Item["group_ids"].([]interface{})))
				for k, v := range r.Item["group_ids"].([]interface{}) {
					renamedGroupIDs[k] = mapGroupIDs[r.Item["realm_id"].(string)+"_"+v.(string)]
				}
				sort.Strings(renamedGroupIDs)
				g.Resources[i].Item["group_ids"] = renamedGroupIDs
			} else {
				g.Resources[i].Item["group_ids"] = []string{}
			}
		}

		// Sort valid_redirect_uris and web_origins to get reproducible results for keycloak_openid_client resources
		if r.InstanceInfo.Type == "keycloak_openid_client" {
			if _, exist := r.Item["valid_redirect_uris"]; exist {
				sortedValidRedirectUris := make([]string, len(r.Item["valid_redirect_uris"].([]interface{})))
				for k, v := range r.Item["valid_redirect_uris"].([]interface{}) {
					sortedValidRedirectUris[k] = v.(string)
				}
				sort.Strings(sortedValidRedirectUris)
				g.Resources[i].Item["valid_redirect_uris"] = sortedValidRedirectUris
			}

			if _, exist := r.Item["web_origins"]; exist {
				sortedWebOrigins := make([]string, len(r.Item["web_origins"].([]interface{})))
				for k, v := range r.Item["web_origins"].([]interface{}) {
					sortedWebOrigins[k] = v.(string)
				}
				sort.Strings(sortedWebOrigins)
				g.Resources[i].Item["web_origins"] = sortedWebOrigins
			}
		}

		// Sort composite_roles to get reproducible results for keycloak_role resources
		if _, exist := r.Item["composite_roles"]; exist && r.InstanceInfo.Type == "keycloak_role" {
			renamedCompositeRoles := make([]string, len(r.Item["composite_roles"].([]interface{})))
			for k, v := range r.Item["composite_roles"].([]interface{}) {
				renamedCompositeRoles[k] = mapRoleIDs[r.Item["realm_id"].(string)+"_"+v.(string)]
			}
			sort.Strings(renamedCompositeRoles)
			g.Resources[i].Item["composite_roles"] = renamedCompositeRoles
		}

		// Sort default_scopes to get reproducible results for keycloak_openid_client_default_scopes resources
		if _, exist := r.Item["default_scopes"]; exist && r.InstanceInfo.Type == "keycloak_openid_client_default_scopes" {
			renamedScopes := make([]string, len(r.Item["default_scopes"].([]interface{})))
			for k, v := range r.Item["default_scopes"].([]interface{}) {
				renamedScopes[k] = mapScopeNames[r.Item["realm_id"].(string)+"_"+v.(string)]
			}
			sort.Strings(renamedScopes)
			g.Resources[i].Item["default_scopes"] = renamedScopes
		}

		// Sort optional_scopes to get reproducible results for keycloak_openid_client_optional_scopes resources
		if _, exist := r.Item["optional_scopes"]; exist && r.InstanceInfo.Type == "keycloak_openid_client_optional_scopes" {
			renamedScopes := make([]string, len(r.Item["optional_scopes"].([]interface{})))
			for k, v := range r.Item["optional_scopes"].([]interface{}) {
				renamedScopes[k] = mapScopeNames[r.Item["realm_id"].(string)+"_"+v.(string)]
			}
			sort.Strings(renamedScopes)
			g.Resources[i].Item["optional_scopes"] = renamedScopes
		}

		// Sort role_ids to get reproducible results for keycloak_group_roles resources
		if r.InstanceInfo.Type == "keycloak_group_roles" {
			if roleIDs, ok := r.Item["role_ids"].([]interface{}); ok {
				sortedRoles := make([]string, len(roleIDs))
				for k, v := range roleIDs {
					sortedRoles[k] = mapRoleIDs[r.Item["realm_id"].(string)+"_"+v.(string)]
				}
				sort.Strings(sortedRoles)
				g.Resources[i].Item["role_ids"] = sortedRoles
			} else {
				g.Resources[i].Item["role_ids"] = []string{}
			}
		}

		// Sort members to get reproducible results for keycloak_group_memberships resources
		// Map members to keycloak_user.foo.username Terraform variables
		if r.InstanceInfo.Type == "keycloak_group_memberships" {
			sortedMembers := make([]string, len(r.Item["members"].([]interface{})))
			for k, v := range r.Item["members"].([]interface{}) {
				if mapUserNames[r.Item["realm_id"].(string)+"_"+v.(string)] != "" {
					sortedMembers[k] = mapUserNames[r.Item["realm_id"].(string)+"_"+v.(string)]
				} else {
					sortedMembers[k] = v.(string)
				}
			}
			sort.Strings(sortedMembers)
			g.Resources[i].Item["members"] = sortedMembers
		}

		// Map ldap_user_federation_id attributes to keycloak_ldap_user_federation.foo.id Terraform variables for ldap mappers resources
		if r.InstanceInfo.Type == "keycloak_ldap_full_name_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_group_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_role_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_hardcoded_group_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_hardcoded_role_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_msad_lds_user_account_control_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_msad_user_account_control_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_user_attribute_mapper" {
			g.Resources[i].Item["ldap_user_federation_id"] = mapUserFederationIDs[r.Item["realm_id"].(string)+"_"+g.Resources[i].Item["ldap_user_federation_id"].(string)]
		}

		// Map group to keycloak_group.foo.name Terraform variables for ldap hardcoded group mapper resources
		if r.InstanceInfo.Type == "keycloak_ldap_hardcoded_group_mapper" {
			g.Resources[i].Item["group"] = mapGroupNames[r.Item["realm_id"].(string)+"_"+r.Item["group"].(string)]
		}

		// Map role to Terraform variables for ldap hardcoded role mapper resources
		if r.InstanceInfo.Type == "keycloak_ldap_hardcoded_role_mapper" {
			g.Resources[i].Item["role"] = mapClientRoleNames[r.Item["realm_id"].(string)+"_"+r.Item["role"].(string)]
		}

		// Map parent_id to keycloak_group.foo.id Terraform variables for keycloak_group resources
		if _, exist := r.Item["parent_id"]; exist && r.InstanceInfo.Type == "keycloak_group" {
			g.Resources[i].Item["parent_id"] = mapGroupIDs[r.Item["realm_id"].(string)+"_"+r.Item["parent_id"].(string)]
		}

		// Map group_id to keycloak_group.foo.id Terraform variables for keycloak_group_memberships and keycloak_group_roles resources
		if r.InstanceInfo.Type == "keycloak_group_memberships" || r.InstanceInfo.Type == "keycloak_group_roles" {
			g.Resources[i].Item["group_id"] = mapGroupIDs[r.Item["realm_id"].(string)+"_"+r.Item["group_id"].(string)]
		}

		// Map service_account_user_id to keycloak_openid_client.foo.service_account_user_id Terraform variables for service account role resources
		if r.InstanceInfo.Type == "keycloak_openid_client_service_account_role" {
			g.Resources[i].Item["service_account_user_id"] = mapServiceAccountUserIDs[r.Item["realm_id"].(string)+"_"+r.Item["service_account_user_id"].(string)]
			g.Resources[i].Item["role"] = mapClientRoleShortNames[r.Item["realm_id"].(string)+"_"+mapClientNames[r.Item["realm_id"].(string)+"_"+r.Item["client_id"].(string)]+"."+r.Item["role"].(string)]
		}

		// Map client_id attributes to keycloak_openid_client.foo.id Terraform variables for open id mappers resources
		if _, exist := r.Item["client_id"]; exist && (r.InstanceInfo.Type == "keycloak_openid_client_service_account_role" ||
			r.InstanceInfo.Type == "keycloak_openid_audience_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_audience_resolve_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_full_name_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_group_membership_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_hardcoded_claim_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_hardcoded_role_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_script_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_attribute_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_client_role_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_property_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_realm_role_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_user_session_note_protocol_mapper" ||
			r.InstanceInfo.Type == "keycloak_openid_client_default_scopes" ||
			r.InstanceInfo.Type == "keycloak_openid_client_optional_scopes" ||
			r.InstanceInfo.Type == "keycloak_role") {
			g.Resources[i].Item["client_id"] = mapClientIDs[r.Item["realm_id"].(string)+"_"+r.Item["client_id"].(string)]
		}

		// Map included_client_audience to keycloak_openid_client.foo.client_id Terraform variables for open id audience mapper resources
		if _, exist := r.Item["included_client_audience"]; exist && r.InstanceInfo.Type == "keycloak_openid_audience_protocol_mapper" {
			g.Resources[i].Item["included_client_audience"] = mapClientClientIDs[r.Item["realm_id"].(string)+"_"+r.Item["included_client_audience"].(string)]
		}

		// Map parent_flow_alias attributes to keycloak_authentication_(sub)flow.foo.alias Terraform variables for authentication subflow and execution resources
		if r.InstanceInfo.Type == "keycloak_authentication_subflow" || r.InstanceInfo.Type == "keycloak_authentication_execution" {
			g.Resources[i].Item["parent_flow_alias"] = mapAuthenticationFlowAliases[r.Item["realm_id"].(string)+"_"+r.Item["parent_flow_alias"].(string)]
		}

		// Map execution_id attributes to keycloak_authentication_execution_config.foo.execution_id Terraform variables for authentication execution config resources
		if r.InstanceInfo.Type == "keycloak_authentication_execution_config" {
			g.Resources[i].Item["execution_id"] = mapAuthenticationExecutionIDs[r.Item["realm_id"].(string)+"_"+r.Item["execution_id"].(string)]
		}

		// Map realm_id attributes to keycloak_realm.foo.id Terraform variables for all the resources (almost all resources have this attribute)
		if _, exist := r.Item["realm_id"]; exist {
			g.Resources[i].Item["realm_id"] = mapRealmIDs[r.Item["realm_id"].(string)]
		}
	}
	return nil
}
