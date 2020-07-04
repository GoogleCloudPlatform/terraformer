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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createScopeResources(realmID string, openidClientScopes []*keycloak.OpenidClientScope) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, openidClientScope := range openidClientScopes {
		resources = append(resources, terraformutils.NewResource(
			openidClientScope.Id,
			"openid_client_scope_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(openidClientScope.Name),
			"keycloak_openid_client_scope",
			"keycloak",
			map[string]string{
				"realm_id": realmID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g RealmGenerator) createOpenidClientScopesResources(realmID, clientID, clientClientID, t string, openidClientScopes *[]keycloak.OpenidClientScope) terraformutils.Resource {
	var scopes []string
	for _, openidClientScope := range *openidClientScopes {
		scopes = append(scopes, openidClientScope.Name)
	}
	return terraformutils.NewResource(
		realmID+"/"+clientID,
		"openid_client_"+t+"_scopes_"+normalizeResourceName(realmID)+"_"+normalizeResourceName(clientClientID),
		"keycloak_openid_client_"+t+"_scopes",
		"keycloak",
		map[string]string{
			"realm_id":    realmID,
			"client_id":   clientID,
			t + "_scopes": strings.Join(scopes, ","),
		},
		[]string{},
		map[string]interface{}{},
	)
}
