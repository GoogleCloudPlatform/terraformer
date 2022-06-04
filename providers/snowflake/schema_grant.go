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
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
)

type SchemaGrantGenerator struct {
	SnowflakeService
}

var ValidSchemaPrivileges = []string{
	"ALL",
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

func (g SchemaGrantGenerator) filterALLGrants(schemaGrantList []schemaGrant) []schemaGrant {
	// For each database_schema_role, figure out if the grant has all of the privileges
	type databaseSchemaRole struct {
		Name        sql.NullString
		GrantedTo   sql.NullString
		GranteeName sql.NullString
	}
	groupedByRole := map[databaseSchemaRole]map[string]struct{}{}
	for _, schemaGrant := range schemaGrantList {
		id := databaseSchemaRole{
			Name:        schemaGrant.Name,
			GrantedTo:   schemaGrant.GrantedTo,
			GranteeName: schemaGrant.GranteeName,
		}
		if _, ok := groupedByRole[id]; !ok {
			groupedByRole[id] = map[string]struct{}{}
		}
		groupedByRole[id][schemaGrant.Privilege.String] = struct{}{}
	}
	for databaseSchemaRole, privs := range groupedByRole {
		for _, p := range ValidSchemaPrivileges {
			if p == "ALL" || p == "OWNERSHIP" || p == "CREATE STREAM" {
				continue
			}
			if _, ok := privs[p]; !ok {
				delete(groupedByRole, databaseSchemaRole)
				break
			}
		}
	}
	filteredSchemaGrants := []schemaGrant{}

	// Roles with the "ALL" privilege
	for databaseSchemaRole := range groupedByRole {
		filteredSchemaGrants = append(filteredSchemaGrants, schemaGrant{
			Name:        databaseSchemaRole.Name,
			Privilege:   sql.NullString{String: "ALL"},
			GrantedTo:   databaseSchemaRole.GrantedTo,
			GranteeName: databaseSchemaRole.GranteeName,
		})
	}

	for _, schemaGrant := range schemaGrantList {
		id := databaseSchemaRole{
			Name:        schemaGrant.Name,
			GrantedTo:   schemaGrant.GrantedTo,
			GranteeName: schemaGrant.GranteeName,
		}
		// Already added it with the "ALL" privilege, so skip
		if _, ok := groupedByRole[id]; ok {
			continue
		}
		filteredSchemaGrants = append(filteredSchemaGrants, schemaGrant)
	}
	return filteredSchemaGrants
}

func (g SchemaGrantGenerator) createResources(schemaGrantList []schemaGrant) ([]terraformutils.Resource, error) {
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
	var resources []terraformutils.Resource

	for id, grant := range groupedResources {
		DB := strings.Split(grant.Name, ".")[0]
		Schema := strings.Split(grant.Name, ".")[1]

		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			fmt.Sprintf("%s__%s__%s", DB, Schema, grant.Privilege),
			"snowflake_schema_grant",
			"snowflake",
			[]string{},
		))
	}
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].GetIDKey() < resources[j].GetIDKey()
	})
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

	g.Resources, err = g.createResources(g.filterALLGrants(allGrants))
	return err
}
