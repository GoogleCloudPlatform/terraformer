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
)

type SchemaGenerator struct {
	SnowflakeService
}

func (g SchemaGenerator) createResources(schemaList []schema) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, schema := range schemaList {
		name := schema.Name.String
		// These schemas are always managed by Snowflake, skip them
		if name == "INFORMATION_SCHEMA" || name == "PUBLIC" {
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			// Schemas need to be namespaced by their database with two underscores
			fmt.Sprintf("%s|%s", schema.DatabaseName.String, schema.Name.String),
			fmt.Sprintf("%s__%s", schema.DatabaseName.String, schema.Name.String),
			"snowflake_schema",
			"snowflake",
			[]string{}))
	}
	return resources
}

func (g *SchemaGenerator) InitResources() error {
	db, err := g.generateService()
	if err != nil {
		return err
	}
	output, err := db.ListSchemas(nil)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
