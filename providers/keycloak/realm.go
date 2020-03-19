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
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

type RealmGenerator struct {
	KeycloakService
}

var RealmAllowEmptyValues = []string{}

func (g RealmGenerator) createResources(realms []*keycloak.Realm) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, realm := range realms {
		resources = append(resources, terraform_utils.NewSimpleResource(
			realm.Id,
			"realm_"+normalizeResourceName(realm.Realm),
			"keycloak_realm",
			"keycloak",
			RealmAllowEmptyValues,
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

func (g RealmGenerator) createLdapUserFederationResources(customUserFederations *[]keycloak.CustomUserFederation) []terraform_utils.Resource {
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
			// the role mapper is not supported for the moment
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

func (g *RealmGenerator) InitResources() error {
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
	g.Resources = g.createResources(realms)
	for _, realm := range realms {
		requiredActions, err := client.GetRequiredActions(realm.Id)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createRequiredActionResources(requiredActions)...)
		customUserFederations, err := client.GetCustomUserFederations(realm.Id)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createLdapUserFederationResources(customUserFederations)...)
		for _, customUserFederation := range *customUserFederations {
			switch customUserFederation.ProviderId {
			case "ldap":
				mappers, err := client.GetLdapUserFederationMappers(realm.Id, customUserFederation.Id)
				if err != nil {
					return err
				}
				g.Resources = append(g.Resources, g.createLdapMapperResources(realm.Id, customUserFederation.Name, mappers)...)
			}
		}
		/* Not supported by the provider for the moment
		openidClientScopes, err = client.GetRealmDefaultClientScopes(realm.Id)
		if err != nil {
			return err
		}
		fmt.Printf("RealmDefaultClientScopes: %+v\n", openidClientScopes)
		openidClientScopes, err = client.GetRealmOptionalClientScopes(realm.Id)
		if err != nil {
			return err
		}
		fmt.Printf("RealmOptionalClientScopes: %+v\n", openidClientScopes)
		*/
	}
	return nil
}

func (g *RealmGenerator) PostConvertHook() error {
	mapRealmIDs := map[string]string{}
	mapUserFederationIDs := map[string]string{}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_realm" && r.InstanceInfo.Type != "keycloak_ldap_user_federation" {
			continue
		}
		if r.InstanceInfo.Type == "keycloak_realm" {
			mapRealmIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		}
		if r.InstanceInfo.Type == "keycloak_ldap_user_federation" {
			mapUserFederationIDs[r.InstanceState.ID] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".id}"
		}
	}
	for i, r := range g.Resources {
		if r.InstanceInfo.Type == "keycloak_realm" {
			if _, exist := r.Item["internationalization"]; exist {
				for _, v := range r.Item["internationalization"].([]interface{}) {
					sortedSupportedLocales := []string{}
					for _, vv := range v.(map[string]interface{})["supported_locales"].([]interface{}) {
						sortedSupportedLocales = append(sortedSupportedLocales, vv.(string))
					}
					sort.Strings(sortedSupportedLocales)
					v.(map[string]interface{})["supported_locales"] = sortedSupportedLocales
				}
			}
		}
		if r.InstanceInfo.Type != "keycloak_required_action" &&
			r.InstanceInfo.Type != "keycloak_ldap_user_federation" &&
			r.InstanceInfo.Type != "keycloak_ldap_full_name_mapper" &&
			r.InstanceInfo.Type != "keycloak_ldap_group_mapper" &&
			r.InstanceInfo.Type != "keycloak_ldap_msad_user_account_control_mapper" &&
			r.InstanceInfo.Type != "keycloak_ldap_user_attribute_mapper" {
			continue
		}
		r.Item["realm_id"] = mapRealmIDs[r.Item["realm_id"].(string)]
		if r.InstanceInfo.Type == "keycloak_ldap_full_name_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_group_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_msad_user_account_control_mapper" ||
			r.InstanceInfo.Type == "keycloak_ldap_user_attribute_mapper" {
			g.Resources[i].Item["ldap_user_federation_id"] = mapUserFederationIDs[g.Resources[i].Item["ldap_user_federation_id"].(string)]
		}
		if r.InstanceInfo.Type == "keycloak_ldap_user_federation" {
			sortedUserObjectClasses := []string{}
			for _, v := range r.Item["user_object_classes"].([]interface{}) {
				sortedUserObjectClasses = append(sortedUserObjectClasses, v.(string))
			}
			sort.Strings(sortedUserObjectClasses)
			r.Item["user_object_classes"] = sortedUserObjectClasses
		}
	}
	return nil
}
