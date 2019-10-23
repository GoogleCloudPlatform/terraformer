// Copyright 2019 The Terraformer Authors.
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
package snowflake

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/pkg/errors"
)

type RoleGrantGenerator struct {
	SnowflakeService
}

func (g RoleGrantGenerator) createResources(roleGrantList []roleGrant) ([]terraform_utils.Resource, error) {
	groupedResources := map[string]*TfGrant{}
	for _, grant := range roleGrantList {
		id := grant.Name.String
		_, ok := groupedResources[id]
		if !ok {
			groupedResources[id] = &TfGrant{
				Name:      grant.Name.String,
				Privilege: grant.Privilege.String,
				Roles:     []string{},
			}
		}
		tfGrant := groupedResources[id]
		switch grant.GrantedTo.String {
		case "ROLE":
			tfGrant.Roles = append(tfGrant.Roles, grant.GranteeName.String)
		default:
			return nil, errors.New(fmt.Sprintf("[ERROR] Unrecognized type of grant: %s", grant.GrantedTo.String))
		}
	}
	var resources []terraform_utils.Resource
	for id, grant := range groupedResources {
		resources = append(resources, terraform_utils.NewResource(
			id,
			grant.Name,
			"snowflake_role_grants",
			"snowflake",
			map[string]string{
				"privilege": grant.Privilege,
			},
			[]string{},
			map[string]interface{}{
				"roles": grant.Roles,
			},
		))
	}
	return resources, nil
}

func (g *RoleGrantGenerator) InitResources() error {
	db, err := g.generateService()
	if err != nil {
		return err
	}
	roles, err := db.ListRoles()
	if err != nil {
		return err
	}

	allGrants := []roleGrant{}
	for _, role := range roles {
		grants, err := db.ListRoleGrants(role)
		if err != nil {
			return err
		}
		allGrants = append(allGrants, grants...)
	}
	g.Resources, err = g.createResources(allGrants)
	return err
}
