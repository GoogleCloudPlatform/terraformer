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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/mrparkers/terraform-provider-keycloak/keycloak"
)

type UserGenerator struct {
	KeycloakService
}

var UserAllowEmptyValues = []string{}
var UserAdditionalFields = map[string]interface{}{}

func (g UserGenerator) createResources(users []*keycloak.User) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, user := range users {
		resources = append(resources, terraform_utils.NewResource(
			user.Id,
			"user_"+normalizeResourceName(user.RealmId)+"_"+normalizeResourceName(user.Username),
			"keycloak_user",
			"keycloak",
			map[string]string{
				"realm_id": user.RealmId,
			},
			UserAllowEmptyValues,
			UserAdditionalFields,
		))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	var usersFull []*keycloak.User
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
		users, err := client.GetUsers(realm.Id)
		if err != nil {
			return err
		}
		usersFull = append(usersFull, users...)
	}
	g.Resources = g.createResources(usersFull)
	return nil
}
