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

type ScopeGenerator struct {
	KeycloakService
}

var ScopeAllowEmptyValues = []string{}
var ScopeAdditionalFields = map[string]interface{}{}

func (g ScopeGenerator) createResources(realmId string, openidClientScopes []*keycloak.OpenidClientScope) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, openidClientScope := range openidClientScopes {
		resources = append(resources, terraform_utils.NewResource(
			openidClientScope.Id,
			"openid_client_scope_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(openidClientScope.Name),
			"keycloak_openid_client_scope",
			"keycloak",
			map[string]string{
				"realm_id": realmId,
			},
			ScopeAllowEmptyValues,
			ScopeAdditionalFields,
		))
	}
	return resources
}

func (g ScopeGenerator) createOpenidClientScopesResources(realmId, clientId, clientClientId, t string, openidClientScopes *[]keycloak.OpenidClientScope) terraform_utils.Resource {
	var scopes []string
	for _, openidClientScope := range *openidClientScopes {
		scopes = append(scopes, openidClientScope.Name)
	}
	return terraform_utils.NewResource(
		realmId+"/"+clientId,
		"openid_client_"+t+"_scopes_"+normalizeResourceName(realmId)+"_"+normalizeResourceName(clientClientId),
		"keycloak_openid_client_"+t+"_scopes",
		"keycloak",
		map[string]string{
			"realm_id":    realmId,
			"client_id":   clientId,
			t + "_scopes": strings.Join(scopes, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}

func (g *ScopeGenerator) InitResources() error {
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
		// Scopes at Realm level
		openidClientScopes, err := client.ListOpenidClientScopesWithFilter(realm.Id, func(scope *keycloak.OpenidClientScope) bool { return true })
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createResources(realm.Id, openidClientScopes)...)
		openIDClients, err := client.GetOpenidClients(realm.Id, true)
		if err != nil {
			return err
		}
		// Scopes at OpenId Client level
		for _, openIDClient := range openIDClients {
			openidClientScopes, err := client.GetOpenidDefaultClientScopes(realm.Id, openIDClient.Id)
			if err != nil {
				return err
			}
			if len(*openidClientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Id, openIDClient.Id, openIDClient.ClientId, "default", openidClientScopes))
			}
			openidClientScopes, err = client.GetOpenidOptionalClientScopes(realm.Id, openIDClient.Id)
			if err != nil {
				return err
			}
			if len(*openidClientScopes) > 0 {
				g.Resources = append(g.Resources, g.createOpenidClientScopesResources(realm.Id, openIDClient.Id, openIDClient.ClientId, "optional", openidClientScopes))
			}
		}
	}
	return nil
}

func (g *ScopeGenerator) PostConvertHook() error {
	mapScopeNames := map[string]string{}
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_openid_client_scope" {
			continue
		}
		mapScopeNames[r.Item["realm_id"].(string)+"_"+r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
	}
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "keycloak_openid_client_scope" && r.InstanceInfo.Type != "keycloak_openid_client_default_scopes" && r.InstanceInfo.Type != "keycloak_openid_client_optional_scopes" {
			continue
		}
		if strings.Contains(r.InstanceState.Attributes["consent_screen_text"], "$") {
			g.Resources[i].Item["consent_screen_text"] = strings.ReplaceAll(r.InstanceState.Attributes["consent_screen_text"], "$", "$$")
		}
		if _, exist := r.Item["default_scopes"]; exist {
			renamedScopes := []string{}
			for _, v := range r.Item["default_scopes"].([]interface{}) {
				renamedScopes = append(renamedScopes, mapScopeNames[r.Item["realm_id"].(string)+"_"+v.(string)])
			}
			sort.Strings(renamedScopes)
			r.Item["default_scopes"] = renamedScopes
		}
		if _, exist := r.Item["optional_scopes"]; exist {
			renamedScopes := []string{}
			for _, v := range r.Item["optional_scopes"].([]interface{}) {
				renamedScopes = append(renamedScopes, mapScopeNames[r.Item["realm_id"].(string)+"_"+v.(string)])
			}
			sort.Strings(renamedScopes)
			r.Item["optional_scopes"] = renamedScopes
		}
	}
	return nil
}
