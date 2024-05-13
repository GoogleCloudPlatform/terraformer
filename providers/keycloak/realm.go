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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createRealmResources(realms []*keycloak.Realm) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, realm := range realms {
		resources = append(resources, terraformutils.NewSimpleResource(
			realm.Realm,
			"realm_"+normalizeResourceName(realm.Realm),
			"keycloak_realm",
			"keycloak",
			[]string{},
		))
	}
	return resources
}

func (g RealmGenerator) createRequiredActionResources(requiredActions []*keycloak.RequiredAction) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, requiredAction := range requiredActions {
		resources = append(resources, terraformutils.NewResource(
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

func (g RealmGenerator) createCustomUserFederationResources(customUserFederations *[]keycloak.CustomUserFederation) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, customUserFederation := range *customUserFederations {
		if customUserFederation.ProviderId == "ldap" {
			if customUserFederation.Config["bindCredential"][0] != "" {
				var bindDn string
				for _, i := range strings.Split(customUserFederation.Config["bindDn"][0], ",") {
					attrib := strings.Split(i, "=")
					if strings.ToLower(attrib[0]) == "cn" {
						bindDn = attrib[1]
					}
				}
				resources = append(resources, terraformutils.NewResource(
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
				resources = append(resources, terraformutils.NewResource(
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

func (g RealmGenerator) createLdapMapperResources(realmID, providerName string, mappers *[]interface{}) []terraformutils.Resource {
	var resources []terraformutils.Resource
	var providerID string
	var mapperID string
	var mapperName string
	var mapperType string
	var name string
	mapperNames := make(map[string]int)
	for _, mapper := range *mappers {
		switch reflect.TypeOf(mapper).String() {
		case "*keycloak.LdapFullNameMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapFullNameMapper).Name
			mapperType = "full_name"
		case "*keycloak.LdapGroupMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapGroupMapper).Name
			mapperType = "group"
		case "*keycloak.LdapRoleMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapRoleMapper).Name
			mapperType = "role"
		case "*keycloak.LdapHardcodedGroupMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedGroupMapper).Name
			mapperType = "hardcoded_group"
		case "*keycloak.LdapHardcodedRoleMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapHardcodedRoleMapper).Name
			mapperType = "hardcoded_role"
		case "*keycloak.LdapMsadLdsUserAccountControlMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadLdsUserAccountControlMapper).Name
			mapperType = "msad_lds_user_account_control"
		case "*keycloak.LdapMsadUserAccountControlMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapMsadUserAccountControlMapper).Name
			mapperType = "msad_user_account_control"
		case "*keycloak.LdapUserAttributeMapper":
			providerID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).LdapUserFederationId
			mapperID = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).Id
			mapperName = reflect.ValueOf(mapper).Interface().(*keycloak.LdapUserAttributeMapper).Name
			mapperType = "user_attribute"
		default:
			continue
		}
		name = "ldap_" + mapperType + "_mapper_" + normalizeResourceName(realmID) + "_" + normalizeResourceName(providerName) + "_" + normalizeResourceName(mapperName)
		for k, v := range mapperNames {
			if k == name {
				v++
				name += strconv.Itoa(v)
			}
		}
		if name == "ldap_"+mapperType+"_mapper_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(providerName)+"_"+normalizeResourceName(mapperName) {
			mapperNames[name] = 1
		}
		resources = append(resources, terraformutils.NewResource(
			mapperID,
			name,
			"keycloak_ldap_"+mapperType+"_mapper",
			"keycloak",
			map[string]string{
				"realm_id":    realmID,
				"provider_id": providerID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}
