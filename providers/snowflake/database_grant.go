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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
)

type DatabaseGrantGenerator struct {
	SnowflakeService
}

func (g DatabaseGrantGenerator) createResources(databaseGrantList []databaseGrant) ([]terraformutils.Resource, error) {
	groupedResources := map[string]*TfGrant{}
	for _, grant := range databaseGrantList {
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
	var resources []terraformutils.Resource
	for id, grant := range groupedResources {
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			fmt.Sprintf("%s_%s", grant.Name, grant.Privilege),
			"snowflake_database_grant",
			"snowflake",
			[]string{},
		))
	}
	return resources, nil
}

func (g *DatabaseGrantGenerator) InitResources() error {
	db, err := g.generateService()
	if err != nil {
		return err
	}
	databases, err := db.ListDatabases()
	if err != nil {
		return err
	}
	allGrants := []databaseGrant{}
	for _, database := range databases {
		if database.Origin.String != "" {
			// Provider does not support grants on imported databases yet
			continue
		}
		grants, err := db.ListDatabaseGrants(database)
		if err != nil {
			return err
		}
		allGrants = append(allGrants, grants...)
	}
	g.Resources, err = g.createResources(allGrants)
	return err
}
