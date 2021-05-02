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
	"fmt"
	"github.com/zclconf/go-cty/cty"
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

	// Connect to keycloak instance
	kck, err := keycloak.NewKeycloakClient(g.GetArgs()["url"].(string), g.GetArgs()["client_id"].(string), g.GetArgs()["client_secret"].(string), g.GetArgs()["realm"].(string), "", "", true, g.GetArgs()["client_timeout"].(int), g.GetArgs()["root_ca_certificate"].(string), g.GetArgs()["tls_insecure_skip_verify"].(bool))
	if err != nil {
		return errors.New("keycloak: could not connect to Keycloak")
	}

	// Get realm resources
	target := g.GetArgs()["target"].(string)
	if target == "" {
		realms, err = kck.GetRealms()
		if err != nil {
			return errors.New("keycloak: could not get realms attributes in Keycloak")
		}
	} else {
		realm, err := kck.GetRealm(target)
		if err != nil {
			return errors.New("keycloak: could not get " + target + " realm attributes in Keycloak")
		}
		realms = append(realms, realm)
	}
	g.Resources = append(g.Resources, g.createRealmResources(realms)...)

	// For each realm, get resources
	for _, realm := range realms {
		// Get required actions resources
		requiredActions, err := kck.GetRequiredActions(realm.Id)
		if err != nil {
			return fmt.Errorf("keycloak: could not get required actions of realm %s in Keycloak. err: %w", realm.Id, err)
		}
		g.Resources = append(g.Resources, g.createRequiredActionResources(requiredActions)...)

		// Get top-level authentication flows resources
		authenticationFlows, err := kck.ListAuthenticationFlows(realm.Id)
		if err != nil {
			return fmt.Errorf("keycloak: could not get authentication flows of realm %s in Keycloak. err: %w", realm.Id, err)
		}
		g.Resources = append(g.Resources, g.createAuthenticationFlowResources(authenticationFlows)...)

		// For each authentication flow, get subFlow, execution and execution config resources
		for _, topLevelAuthenticationFlow := range authenticationFlows {
			authenticationSubFlowOrExecutions, err := kck.ListAuthenticationExecutions(realm.Id, topLevelAuthenticationFlow.Alias)
			if err != nil {
				return fmt.Errorf("keycloak: could not get authentication executions of authentication flow %s of realm %s in Keycloak. err: %w",
					topLevelAuthenticationFlow.Alias, realm.Id, err)
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
					authenticationSubFlow, err := kck.GetAuthenticationSubFlow(realm.Id, parentFlowAlias, authenticationSubFlowOrExecution.FlowId)
					if err != nil {
						return fmt.Errorf("keycloak: could not get authentication subflow %s of realm %s in Keycloak. err: %w",
							authenticationSubFlowOrExecution.FlowId, realm.Id, err)
					}

					// Need to store the alias and parent flow alias
					authenticationSubFlowOrExecution.Alias = authenticationSubFlow.Alias
					authenticationSubFlowOrExecution.ParentFlowAlias = parentFlowAlias

					resource = g.createAuthenticationSubFlowResource(authenticationSubFlow)
					g.Resources = append(g.Resources, resource)

				case false:
					authenticationExecution, err := kck.GetAuthenticationExecution(realm.Id, parentFlowAlias, authenticationSubFlowOrExecution.Id)
					if err != nil {
						return fmt.Errorf("keycloak: could not get authentication execution %s of realm %s in Keycloak. err: %w",
							authenticationSubFlowOrExecution.Id, realm.Id, err)
					}

					// Need to store the parent flow alias
					authenticationSubFlowOrExecution.ParentFlowAlias = parentFlowAlias

					resource = g.createAuthenticationExecutionResource(authenticationExecution)
					g.Resources = append(g.Resources, resource)

					if authenticationSubFlowOrExecution.AuthenticationConfig != "" {
						authenticationExecutionConfig := &keycloak.AuthenticationExecutionConfig{
							RealmId:     realm.Id,
							Id:          authenticationSubFlowOrExecution.AuthenticationConfig,
							ExecutionId: authenticationSubFlowOrExecution.Id,
						}
						err := kck.GetAuthenticationExecutionConfig(authenticationExecutionConfig)
						if err != nil {
							return fmt.Errorf("keycloak: could not get authentication execution config %s of realm %s in Keycloak. err: %w",
								authenticationExecutionConfig.Id, realm.Id, err)
						}

						g.Resources = append(g.Resources, g.createAuthenticationExecutionConfigResource(authenticationExecutionConfig))
					}
				}

				if len(stack) > 0 && authenticationSubFlowOrExecution.Index > 0 {
					previous := stack[len(stack)-1]
					var resouceType string
					var resouceName string
					if previous.AuthenticationFlow {
						resouceType = "keycloak_authentication_subflow"
						resouceName = "authentication_subflow_" +
							normalizeResourceName(realm.Id) + "_" + normalizeResourceName(previous.FlowId)
					} else {
						resouceType = "keycloak_authentication_execution"
						resouceName = "authentication_execution_" +
							normalizeResourceName(realm.Id) + "_" + normalizeResourceName(previous.Id)
					}
					resource.AdditionalFields["depends_on"] = []string{resouceType + "." + terraformutils.TfSanitize(resouceName)}
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
		customUserFederations, err := kck.GetCustomUserFederations(realm.Id)
		if err != nil {
			return errors.New("keycloak: could not get custom user federations of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createCustomUserFederationResources(customUserFederations)...)

		// For each custom federation, get mappers resources
		for _, customUserFederation := range *customUserFederations {
			if customUserFederation.ProviderId == "ldap" {
				mappers, err := kck.GetLdapUserFederationMappers(realm.Id, customUserFederation.Id)
				if err != nil {
					return errors.New("keycloak: could not get mappers of ldap user federation " + customUserFederation.Name + " of realm " + realm.Id + " in Keycloak")
				}
				g.Resources = append(g.Resources, g.createLdapMapperResources(realm.Id, customUserFederation.Name, mappers)...)
			}
		}

		// Get groups tree and default groups resources
		realmGroups, err := kck.GetGroups(realm.Id)
		if err != nil {
			return errors.New("keycloak: could not get groups of realm " + realm.Id + " in Keycloak")
		}
		realmsGroups = append(realmsGroups, realmGroups...)
		g.Resources = append(g.Resources, g.createDefaultGroupResource(realm.Id))

		// Get users resources
		realmUsers, err := kck.GetUsers(realm.Id)
		if err != nil {
			return errors.New("keycloak: could not get users of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createUserResources(realmUsers)...)

		// Get realm open id client scopes resources
		realmScopes, err := kck.ListOpenidClientScopesWithFilter(realm.Id, func(scope *keycloak.OpenidClientScope) bool { return true })
		if err != nil {
			return errors.New("keycloak: could not get realm scopes of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createScopeResources(realm.Id, realmScopes)...)

		// Get open id clients
		realmClients, err := kck.GetOpenidClients(realm.Id, true)
		if err != nil {
			return errors.New("keycloak: could not get open id clients of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createOpenIDClientResources(realmClients)...)

		// For earch open id client, get resources
		mapServiceAccountIds := map[string]map[string]string{}
		mapContainerIDs := map[string]string{}
		mapClientIDs := map[string]string{}
		for _, client := range realmClients {
			mapClientIDs[client.Id] = client.ClientId
			mapContainerIDs[client.Id] = "_" + client.ClientId

			// Get open id client protocol mappers resources
			clientMappers, err := kck.GetGenericClientProtocolMappers(realm.Id, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get protocol mappers of open id client " + client.ClientId + " of realm " + realm.Id + " in Keycloak")
			}
			g.Resources = append(g.Resources, g.createOpenIDProtocolMapperResources(client.ClientId, clientMappers)...)

			// Get open id client default scopes resources
			clientScopes, err := kck.GetOpenidDefaultClientScopes(realm.Id, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get default client scopes of open id client " + client.ClientId + " of realm " + realm.Id + " in Keycloak")
			}
			if len(*clientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Id, client.Id, client.ClientId, "default", clientScopes))
			}

			// Get open id client optional scopes resources
			clientScopes, err = kck.GetOpenidOptionalClientScopes(realm.Id, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get optional client scopes of open id client " + client.ClientId + " of realm " + realm.Id + " in Keycloak")
			}
			if len(*clientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Id, client.Id, client.ClientId, "optional", clientScopes))
			}

			// Prepare a slice to be able to link roles associated to service account roles to be associated to the open id client, only if service accounts are enabled
			if !client.ServiceAccountsEnabled {
				continue
			}
			serviceAccountUser, err := kck.GetOpenidClientServiceAccountUserId(realm.Id, client.Id)
			if err != nil {
				return errors.New("keycloak: could not get service account user associated to open id client " + client.ClientId + " of realm " + realm.Id + " in Keycloak")
			}
			mapServiceAccountIds[serviceAccountUser.Id] = map[string]string{}
			mapServiceAccountIds[serviceAccountUser.Id]["Id"] = client.Id
			mapServiceAccountIds[serviceAccountUser.Id]["ClientId"] = client.ClientId
		}

		// Get open id client roles
		clientRoles, err := kck.GetClientRoles(realm.Id, realmClients)
		if err != nil {
			return errors.New("keycloak: could not get open id clients roles of realm " + realm.Id + " in Keycloak")
		}

		// Get roles
		realmRoles, err := kck.GetRealmRoles(realm.Id)
		if err != nil {
			return errors.New("keycloak: could not get realm roles of realm " + realm.Id + " in Keycloak")
		}

		// Set ContainerId of the roles, for realm = "", for open id clients = "_" + client.ClientId
		// and get roles resources
		mapContainerIDs[realm.Id] = ""
		roles := append(clientRoles, realmRoles...)
		for _, role := range roles {
			role.ContainerId = mapContainerIDs[role.ContainerId]
		}
		g.Resources = append(g.Resources, g.createRoleResources(roles)...)

		// Get service account roles resources
		usersInRole, err := kck.GetClientRoleUsers(realm.Id, clientRoles)
		if err != nil {
			return errors.New("keycloak: could not get users roles of realm " + realm.Id + " in Keycloak")
		}
		g.Resources = append(g.Resources, g.createServiceAccountClientRolesResources(realm.Id, clientRoles, *usersInRole, mapServiceAccountIds, mapClientIDs)...)
	}

	// Parse the groups trees, and get all the groups
	// Get groups resources
	groups := g.flattenGroups(realmsGroups, "")
	g.Resources = append(g.Resources, g.createGroupResources(groups)...)

	// For each group, get group memberships and roles resources
	for _, group := range groups {
		// Get group members resources
		members, err := kck.GetGroupMembers(group.RealmId, group.Id)
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
		groupDetails, err := kck.GetGroup(group.RealmId, group.Id)
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
		if r.Address.Type != "keycloak_realm" &&
			r.Address.Type != "keycloak_ldap_user_federation" &&
			r.Address.Type != "keycloak_group" &&
			r.Address.Type != "keycloak_openid_client" &&
			r.Address.Type != "keycloak_role" &&
			r.Address.Type != "keycloak_openid_client_scope" &&
			r.Address.Type != "keycloak_user" &&
			r.Address.Type != "keycloak_authentication_flow" &&
			r.Address.Type != "keycloak_authentication_subflow" &&
			r.Address.Type != "keycloak_authentication_execution" {
			continue
		}
		if r.Address.Type == "keycloak_realm" {
			mapRealmIDs[r.ImportID] = "${" + r.Address.String() + ".id}"
		}
		if r.Address.Type == "keycloak_ldap_user_federation" {
			mapUserFederationIDs[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".id}"
		}
		if r.Address.Type == "keycloak_group" {
			mapGroupIDs[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".id}"
			mapGroupNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("name")] = "${" + r.Address.String() + ".name}"
		}
		if r.Address.Type == "keycloak_openid_client" {
			mapClientIDs[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".id}"
			mapClientNames[r.GetStateAttr("realm_id")+"_"+r.ImportID] = r.GetStateAttr("client_id")
			mapClientClientNames[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".client_id}"
			mapClientClientIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")] = "${" + r.Address.String() + ".client_id}"
			if r.HasStateAttr("service_account_user_id") {
				mapServiceAccountUserIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("service_account_user_id")] = "${" + r.Address.String() + ".service_account_user_id}"
			}
		}
		if r.Address.Type == "keycloak_role" {
			mapRoleIDs[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".id}"
			if r.HasStateAttr("client_id") {
				mapClientRoleNames[r.GetStateAttr("realm_id")+"_"+mapClientNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")]+"."+r.GetStateAttr("name")] = mapClientClientNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")] + ".${" + r.Address.String() + ".name}"
				mapClientRoleShortNames[r.GetStateAttr("realm_id")+"_"+mapClientNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")]+"."+r.GetStateAttr("name")] = "${" + r.Address.String() + ".name}"
			} else {
				mapClientRoleNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("name")] = "${" + r.Address.String() + ".name}"
			}
		}
		if r.Address.Type == "keycloak_openid_client_scope" {
			mapScopeNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("name")] = "${" + r.Address.String() + ".name}"
		}
		if r.Address.Type == "keycloak_user" {
			mapUserNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("username")] = "${" + r.Address.String() + ".username}"
		}
		if r.Address.Type == "keycloak_authentication_flow" || r.Address.Type == "keycloak_authentication_subflow" {
			mapAuthenticationFlowAliases[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("alias")] = "${" + r.Address.String() + ".alias}"
		}
		if r.Address.Type == "keycloak_authentication_execution" {
			mapAuthenticationExecutionIDs[r.GetStateAttr("realm_id")+"_"+r.ImportID] = "${" + r.Address.String() + ".id}"
		}
	}

	// For each resources, modify import if needed...
	for _, r := range g.Resources {
		// Escape keycloak text inputs not to get unpredictable results or errors when Terraform will try to interpret variables ($ vs $$)
		// TODO: ensure that we escape all existing fields
		if strings.Contains(r.GetStateAttr("consent_screen_text"), "$") {
			r.SetStateAttr("consent_screen_text", cty.StringVal(strings.ReplaceAll(r.GetStateAttr("consent_screen_text"), "$", "$$")))
		}
		if strings.Contains(r.GetStateAttr("name"), "$") {
			r.SetStateAttr("name", cty.StringVal(strings.ReplaceAll(r.GetStateAttr("name"), "$", "$$")))
		}
		if strings.Contains(r.GetStateAttr("description"), "$") {
			r.SetStateAttr("description", cty.StringVal(strings.ReplaceAll(r.GetStateAttr("description"), "$", "$$")))
		}
		if strings.Contains(r.GetStateAttr("root_url"), "$") {
			r.SetStateAttr("root_url", cty.StringVal(strings.ReplaceAll(r.GetStateAttr("root_url"), "$", "$$")))
		}

		// Sort supported_locales to get reproducible results for keycloak_realm resources
		if r.Address.Type == "keycloak_realm" {
			r.SortStateAttrEachAttrStringSlice("internationalization", "supported_locales")
		}

		// Sort group_ids to get reproducible results for keycloak_default_groups resources
		// Set an empty string slice if the attribute doesn't exist as it is mandatory
		if r.Address.Type == "keycloak_default_groups" {
			if r.HasStateAttr("group_ids") {
				g.sortAttrWithRealmId(r, "group_ids", mapGroupIDs)
			} else {
				r.SetStateAttr("group_ids", cty.ListVal([]cty.Value{}))
			}
		}

		// Sort valid_redirect_uris and web_origins to get reproducible results for keycloak_openid_client resources
		if r.Address.Type == "keycloak_openid_client" {
			r.SortStateAttrStringSlice("valid_redirect_uris")
			r.SortStateAttrStringSlice("web_origins")
		}

		// Sort composite_roles to get reproducible results for keycloak_role resources
		if r.Address.Type == "keycloak_role" && r.HasStateAttr("composite_roles") {
			g.sortAttrWithRealmId(r, "composite_roles", mapGroupIDs)
		}

		// Sort default_scopes to get reproducible results for keycloak_openid_client_default_scopes resources
		if r.Address.Type == "keycloak_openid_client_default_scopes" && r.HasStateAttr("default_scopes") {
			g.sortAttrWithRealmId(r, "default_scopes", mapGroupIDs)
		}

		// Sort optional_scopes to get reproducible results for keycloak_openid_client_optional_scopes resources
		if r.Address.Type == "keycloak_openid_client_optional_scopes" && r.HasStateAttr("optional_scopes") {
			g.sortAttrWithRealmId(r, "optional_scopes", mapGroupIDs)
		}

		// Sort role_ids to get reproducible results for keycloak_group_roles resources
		if r.Address.Type == "keycloak_group_roles" {
			g.sortAttrWithRealmId(r, "role_ids", mapGroupIDs)
		}

		// Sort members to get reproducible results for keycloak_group_memberships resources
		// Map members to keycloak_user.foo.username Terraform variables
		if r.Address.Type == "keycloak_group_memberships" {
			var sortedMembers []string
			for _, v := range r.GetStateAttrSlice("members") {
				if mapUserNames[r.GetStateAttr("realm_id")+"_"+v.AsString()] != "" {
					sortedMembers = append(sortedMembers, mapUserNames[r.GetStateAttr("realm_id")+"_"+v.AsString()])
				} else {
					sortedMembers = append(sortedMembers, v.AsString())
				}
			}
			sort.Strings(sortedMembers)
			var sortedValues []cty.Value
			for _, v := range sortedMembers {
				sortedValues = append(sortedValues, cty.StringVal(v))
			}
			r.SetStateAttr("members", cty.ListVal(sortedValues))
		}

		// Map ldap_user_federation_id attributes to keycloak_ldap_user_federation.foo.id Terraform variables for ldap mappers resources
		if r.Address.Type == "keycloak_ldap_full_name_mapper" ||
			r.Address.Type == "keycloak_ldap_group_mapper" ||
			r.Address.Type == "keycloak_ldap_role_mapper" ||
			r.Address.Type == "keycloak_ldap_hardcoded_group_mapper" ||
			r.Address.Type == "keycloak_ldap_hardcoded_role_mapper" ||
			r.Address.Type == "keycloak_ldap_msad_lds_user_account_control_mapper" ||
			r.Address.Type == "keycloak_ldap_msad_user_account_control_mapper" ||
			r.Address.Type == "keycloak_ldap_user_attribute_mapper" {
			r.SetStateAttr("ldap_user_federation_id", cty.StringVal(mapUserFederationIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("ldap_user_federation_id")]))
		}

		// Map group to keycloak_group.foo.name Terraform variables for ldap hardcoded group mapper resources
		if r.Address.Type == "keycloak_ldap_hardcoded_group_mapper" {
			r.SetStateAttr("group", cty.StringVal(mapGroupNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("group")]))

		}

		// Map role to Terraform variables for ldap hardcoded role mapper resources
		if r.Address.Type == "keycloak_ldap_hardcoded_role_mapper" {
			r.SetStateAttr("role", cty.StringVal(mapClientRoleNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("role")]))

		}

		// Map parent_id to keycloak_group.foo.id Terraform variables for keycloak_group resources
		if r.Address.Type == "keycloak_group" && r.HasStateAttr("parent_id") {
			r.SetStateAttr("parent_id", cty.StringVal(mapGroupIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("parent_id")]))

		}

		// Map group_id to keycloak_group.foo.id Terraform variables for keycloak_group_memberships and keycloak_group_roles resources
		if r.Address.Type == "keycloak_group_memberships" || r.Address.Type == "keycloak_group_roles" {
			r.SetStateAttr("group_id", cty.StringVal(mapGroupIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("group_id")]))
		}

		// Map service_account_user_id to keycloak_openid_client.foo.service_account_user_id Terraform variables for service account role resources
		if r.Address.Type == "keycloak_openid_client_service_account_role" {
			r.SetStateAttr("service_account_user_id", cty.StringVal(mapServiceAccountUserIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("service_account_user_id")]))
			r.SetStateAttr("role", cty.StringVal(mapClientRoleShortNames[r.GetStateAttr("realm_id")+"_"+mapClientNames[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")]+"."+r.GetStateAttr("role")]))
		}

		// Map client_id attributes to keycloak_openid_client.foo.id Terraform variables for open id mappers resources
		if r.HasStateAttr("client_id") && (r.Address.Type == "keycloak_openid_client_service_account_role" ||
			r.Address.Type == "keycloak_openid_audience_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_full_name_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_group_membership_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_hardcoded_claim_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_hardcoded_group_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_hardcoded_role_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_user_attribute_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_user_property_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_user_realm_role_protocol_mapper" ||
			r.Address.Type == "keycloak_openid_client_default_scopes" ||
			r.Address.Type == "keycloak_openid_client_optional_scopes" ||
			r.Address.Type == "keycloak_role") {
			r.SetStateAttr("client_id", cty.StringVal(mapClientIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("client_id")]))
		}

		// Map included_client_audience to keycloak_openid_client.foo.client_id Terraform variables for open id audience mapper resources
		if r.Address.Type == "keycloak_openid_audience_protocol_mapper" && r.HasStateAttr("included_client_audience") {
			r.SetStateAttr("included_client_audience", cty.StringVal(mapClientClientIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("included_client_audience")]))
		}

		// Map parent_flow_alias attributes to keycloak_authentication_(sub)flow.foo.alias Terraform variables for authentication subflow and execution resources
		if r.Address.Type == "keycloak_authentication_subflow" || r.Address.Type == "keycloak_authentication_execution" {
			r.SetStateAttr("parent_flow_alias", cty.StringVal(mapAuthenticationFlowAliases[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("parent_flow_alias")]))
		}

		// Map execution_id attributes to keycloak_authentication_execution_config.foo.execution_id Terraform variables for authentication execution config resources
		if r.Address.Type == "keycloak_authentication_execution_config" {
			r.SetStateAttr("execution_id", cty.StringVal(mapAuthenticationExecutionIDs[r.GetStateAttr("realm_id")+"_"+r.GetStateAttr("execution_id")]))
		}

		// Map realm_id attributes to keycloak_realm.foo.id Terraform variables for all the resources (almost all resources have this attribute)
		if r.HasStateAttr("realm_id") {
			r.SetStateAttr("realm_id", cty.StringVal(mapRealmIDs[r.GetStateAttr("realm_id")]))
		}
	}
	return nil
}

func (g *RealmGenerator) sortAttrWithRealmId(r terraformutils.Resource, attr string, mapGroupIDs map[string]string) {
	var renamedStrings []string
	for _, v := range r.GetStateAttrSlice(attr) {
		renamedStrings = append(renamedStrings, mapGroupIDs[r.GetStateAttr("realm_id")+"_"+v.AsString()])
	}
	sort.Strings(renamedStrings)
	var renamedValues []cty.Value
	for _, v := range renamedStrings {
		renamedValues = append(renamedValues, cty.StringVal(v))
	}
	r.SetStateAttr(attr, cty.ListVal(renamedValues))
}
