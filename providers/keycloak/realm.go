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
	"reflect"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createRealmResources(realms []*keycloak.Realm) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, realm := range realms {
		resources = append(resources, terraform_utils.NewSimpleResource(
			realm.Id,
			"realm_"+normalizeResourceName(realm.Realm),
			"keycloak_realm",
			"keycloak",
			[]string{},
		))
	}
	return resources
}

func (g RealmGenerator) createRequiredActionResources(requiredActions []*keycloak.RequiredAction) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, requiredAction := range requiredActions {
		resources = append(resources, terraform_utils.NewResource(
			requiredAction.RealmId+"/"+requiredAction.Alias,
			"required_action_"+normalizeResourceName(requiredAction.RealmId)+"_"+normalizeResourceName(requiredAction.Alias),
			"keycloak_required_action",
			"keycloak",
			map[string]string{
				"realm_id": requiredAction.RealmId,
				"alias":    requiredAction.Alias,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g RealmGenerator) createCustomUserFederationResources(customUserFederations *[]keycloak.CustomUserFederation) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, customUserFederation := range *customUserFederations {
		switch customUserFederation.ProviderId {
		case "ldap":
			if customUserFederation.Config["bindCredential"][0] != "" {
				var bindDn string
				for _, i := range strings.Split(customUserFederation.Config["bindDn"][0], ",") {
					attrib := strings.Split(i, "=")
					if strings.ToLower(attrib[0]) == "cn" {
						bindDn = attrib[1]
					}
				}
				resources = append(resources, terraform_utils.NewResource(
					customUserFederation.Id,
					"ldap_user_federation_"+normalizeResourceName(customUserFederation.RealmId)+"_"+normalizeResourceName(customUserFederation.Name)+"_"+normalizeResourceName(bindDn),
					"keycloak_ldap_user_federation",
					"keycloak",
					map[string]string{
						"realm_id":    customUserFederation.RealmId,
						"provider_id": customUserFederation.ProviderId,
						"bind_dn":     bindDn,
					},
					[]string{},
					map[string]interface{}{},
				))
			} else {
				resources = append(resources, terraform_utils.NewResource(
					customUserFederation.Id,
					"ldap_user_federation_"+normalizeResourceName(customUserFederation.RealmId)+"_"+normalizeResourceName(customUserFederation.Name),
					"keycloak_ldap_user_federation",
					"keycloak",
					map[string]string{
						"realm_id":    customUserFederation.RealmId,
						"provider_id": customUserFederation.ProviderId,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}
	return resources
}

func (g RealmGenerator) createLdapMapperResources(realmId, providerName string, mappers *[]interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	var providerId string
	var mapperId string
	var mapperName string
	var mapperType string
	var name string
	mapperNames := make(map[string]int)
	for _, mapper := range *mappers {
		switch reflect.TypeOf(mapper).String() {
		case "*keycloak.LdapFullNameMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).Name
			mapperType = "full_name"
		case "*keycloak.LdapGroupMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).Name
			mapperType = "group"
		case "*keycloak.LdapRoleMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).Name
			mapperType = "role"
		case "*keycloak.LdapHardcodedGroupMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).Name
			mapperType = "hardcoded_group"
		case "*keycloak.LdapHardcodedRoleMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).Name
			mapperType = "hardcoded_role"
		case "*keycloak.LdapMsadLdsUserAccountControlMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).Name
			mapperType = "msad_lds_user_account_control"
		case "*keycloak.LdapMsadUserAccountControlMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).Name
			mapperType = "msad_user_account_control"
		case "*keycloak.LdapUserAttributeMapper":
			providerId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).LdapUserFederationId
			mapperId = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).Name
			mapperType = "user_attribute"
		default:
			continue
		}
		name = "ldap_" + mapperType + "_mapper_" + normalizeResourceName(realmId) + "_" + normalizeResourceName(providerName) + "_" + normalizeResourceName(mapperName)
		for k, v := range mapperNames {
			if k == name {
				v++
				name = name + strconv.Itoa(v)
			}
		}
		if name == "ldap_"+mapperType+"_mapper_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(providerName)+"_"+normalizeResourceName(mapperName) {
			mapperNames[name] = 1
		}
		resources = append(resources, terraform_utils.NewResource(
			mapperId,
			name,
			"keycloak_ldap_"+mapperType+"_mapper",
			"keycloak",
			map[string]string{
				"realm_id":    realmId,
				"provider_id": providerId,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}
