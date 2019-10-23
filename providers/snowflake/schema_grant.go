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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/pkg/errors"
)

type SchemaGrantGenerator struct {
	SnowflakeService
}

var ValidSchemaPrivileges = []string{
	"MODIFY",
	"MONITOR",
	"OWNERSHIP",
	"USAGE",
	"CREATE TABLE",
	"CREATE VIEW",
	"CREATE FILE FORMAT",
	"CREATE STAGE",
	"CREATE PIPE",
	"CREATE STREAM",
	"CREATE TASK",
	"CREATE SEQUENCE",
	"CREATE FUNCTION",
	"CREATE PROCEDURE",
}

func stringArrayContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (g SchemaGrantGenerator) createResources(schemaGrantList []schemaGrant) ([]terraform_utils.Resource, error) {
	groupedResources := map[string]*TfGrant{}
	for _, grant := range schemaGrantList {
		// TODO(ad): Fix this csv delimited when fixed in the provider. We should use the same functionality.
		// Valid Schema Privilege check
		if !stringArrayContains(ValidSchemaPrivileges, grant.Privilege.String) {
			continue
		}
		DB := strings.Split(grant.Name.String, ".")[0]
		Schema := strings.Split(grant.Name.String, ".")[1]
		id := fmt.Sprintf("%s|%s||%s", DB, Schema, grant.Privilege.String)
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
		DB := strings.Split(grant.Name, ".")[0]
		Schema := strings.Split(grant.Name, ".")[1]
		resources = append(resources, terraform_utils.NewResource(
			id,
			fmt.Sprintf("%s__%s__%s", DB, Schema, grant.Privilege),
			"snowflake_schema_grant",
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

func (g *SchemaGrantGenerator) InitResources() error {
	db, err := g.generateService()
	if err != nil {
		return err
	}
	databases, err := db.ListDatabases()
	if err != nil {
		return err
	}

	allGrants := []schemaGrant{}
	for _, database := range databases {
		if database.Origin.String != "" {
			// Provider does not support grants on imported databases yet
			continue
		}
		schemas, err := db.ListSchemas(&database)
		if err != nil {
			return err
		}
		for _, schema := range schemas {
			grants, err := db.ListSchemaGrants(database, schema)
			if err != nil {
				return err
			}
			allGrants = append(allGrants, grants...)
		}
	}
	g.Resources, err = g.createResources(allGrants)
	return err
}
