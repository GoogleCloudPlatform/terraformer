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

package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type PermissionsGenerator struct {
	RBTService
}

type Permissions struct {
	User  string `json:"user"`
	Vhost string `json:"vhost"`
}

type AllPermissions []Permissions

var PermissionsAllowEmptyValues = []string{"configure", "write", "read"}
var PermissionsAdditionalFields = map[string]interface{}{}

func (g PermissionsGenerator) createResources(allPermissions AllPermissions) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, permissions := range allPermissions {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", permissions.User, permissions.Vhost),
			fmt.Sprintf("permissions_%s_%s", normalizeResourceName(permissions.User), normalizeResourceName(permissions.Vhost)),
			"rabbitmq_permissions",
			"rabbitmq",
			map[string]string{
				"user":  permissions.User,
				"vhost": permissions.Vhost,
			},
			PermissionsAllowEmptyValues,
			PermissionsAdditionalFields,
		))
	}
	return resources
}

func (g *PermissionsGenerator) InitResources() error {
	body, err := g.generateRequest("/api/permissions?columns=user,vhost")
	if err != nil {
		return err
	}
	var permissions AllPermissions
	err = json.Unmarshal(body, &permissions)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(permissions)
	return nil
}
