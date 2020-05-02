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

func (g RealmGenerator) createUserResources(users []*keycloak.User) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		resources = append(resources, terraformutils.NewResource(
			user.Id,
			"user_"+normalizeResourceName(user.RealmId)+"_"+normalizeResourceName(user.Username),
			"keycloak_user",
			"keycloak",
			map[string]string{
				"realm_id": user.RealmId,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}
