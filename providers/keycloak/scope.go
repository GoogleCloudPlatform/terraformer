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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

func (g RealmGenerator) createScopeResources(realmId string, openidClientScopes []*keycloak.OpenidClientScope) []terraform_utils.Resource {
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
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g RealmGenerator) createOpenidClientScopesResources(realmId, clientId, clientClientId, t string, openidClientScopes *[]keycloak.OpenidClientScope) terraform_utils.Resource {
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
