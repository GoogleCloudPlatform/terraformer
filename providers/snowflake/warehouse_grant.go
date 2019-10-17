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

type WarehouseGrantGenerator struct {
	SnowflakeService
}

func (g WarehouseGrantGenerator) createResources(warehouseGrantList []warehouseGrant) ([]terraform_utils.Resource, error) {
	groupedResources := map[string]*TfGrant{}
	for _, grant := range warehouseGrantList {
		// TODO(ad): Fix this csv delimited when fixed in the provider. We should use the same functionality.
		id := fmt.Sprintf("%s|||%s", grant.Name.String, grant.Privilege.String)
		_, ok := groupedResources[id]
		if !ok {
			groupedResources[id] = &TfGrant{
				Name:      grant.Name.String,
				Privilege: grant.Privilege.String,
				Roles:     []string{},
				Shares:    []string{},
			}
		}
		tfGrant := groupedResources[id]
		switch grant.GrantedTo.String {
		case "ROLE":
			tfGrant.Roles = append(tfGrant.Roles, grant.GranteeName.String)
		case "SHARE":
			tfGrant.Shares = append(tfGrant.Shares, grant.GranteeName.String)
		default:
			return nil, errors.New(fmt.Sprintf("[ERROR] Unrecognized type of grant: %s", grant.GrantedTo.String))
		}

	}
	var resources []terraform_utils.Resource
	for id, grant := range groupedResources {
		resources = append(resources, terraform_utils.NewResource(
			id,
			fmt.Sprintf("%s_%s", grant.Name, grant.Privilege),
			"snowflake_warehouse_grant",
			"snowflake",
			map[string]string{
				"privilege": grant.Privilege,
			},
			[]string{},
			map[string]interface{}{
				"roles":  grant.Roles,
				"shares": grant.Shares,
			},
		))
	}
	return resources, nil
}

func (g *WarehouseGrantGenerator) InitResources() error {
	db, err := g.generateService()
	if err != nil {
		return err
	}
	warehouses, err := db.ListWarehouses()
	if err != nil {
		return err
	}
	allGrants := []warehouseGrant{}
	for _, warehouse := range warehouses {
		grants, err := db.ListWarehouseGrants(warehouse)
		if err != nil {
			return err
		}
		allGrants = append(allGrants, grants...)
	}
	g.Resources, err = g.createResources(allGrants)
	return err
}
