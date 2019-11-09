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

package aws

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/sts"
)

type GlueGenerator struct {
	AWSService
}

func (g *GlueGenerator) loadGlueCrawlers(svc *glue.Glue) error {
	var GlueCrawlerAllowEmptyValues = []string{"tags."}
	crawlers, err := svc.GetCrawlers(&glue.GetCrawlersInput{})
	if err != nil {
		return err
	}

	for _, crawler := range crawlers.Crawlers {
		resource := terraform_utils.NewSimpleResource(*crawler.Name, *crawler.Name,
			"aws_glue_crawler",
			"aws",
			GlueCrawlerAllowEmptyValues)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *GlueGenerator) loadGlueCatalogDatabase(svc *glue.Glue, account *string) (DatabaseNames []*string, error error) {
	var GlueCatalogDatabaseAllowEmptyValues = []string{"tags."}
	catalogDatabases, err := svc.GetDatabases(&glue.GetDatabasesInput{})
	if err != nil {
		return nil, err
	}

	for _, catalogDatabase := range catalogDatabases.DatabaseList {
		// format of ID is "CATALOG-ID:DATABASE-NAME".
		// CATALOG-ID is AWS Account ID
		// https://docs.aws.amazon.com/cli/latest/reference/glue/create-database.html#options
		id := *account + ":" + *catalogDatabase.Name
		resource := terraform_utils.NewSimpleResource(id, *catalogDatabase.Name,
			"aws_glue_catalog_database",
			"aws",
			GlueCatalogDatabaseAllowEmptyValues)
		g.Resources = append(g.Resources, resource)
		DatabaseNames = append(DatabaseNames, catalogDatabase.Name)
	}
	return DatabaseNames, nil
}

func (g *GlueGenerator) loadGlueCatalogTable(svc *glue.Glue, account *string, databaseName *string) error {
	// format of ID is "CATALOG-ID:DATABASE-NAME:TABLE-NAME".
	// CATALOG-ID is AWS Account ID
	// https://docs.aws.amazon.com/cli/latest/reference/glue/create-database.html#options

	catalogTables, err := svc.GetTables(&glue.GetTablesInput{DatabaseName: databaseName})
	if err != nil {
		return err
	}
	var GlueCatalogTableAllowEmptyValues = []string{"tags."}

	for _, catalogTable := range catalogTables.TableList {
		id := *account + ":" + *databaseName + ":" + *catalogTable.Name
		resource := terraform_utils.NewSimpleResource(id, *catalogTable.Name,
			"aws_glue_catalog_table",
			"aws",
			GlueCatalogTableAllowEmptyValues)
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *GlueGenerator) InitResources() error {
	sess := g.generateSession()
	svc := glue.New(sess)

	stsSvc := sts.New(sess)
	identity, err := stsSvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}
	account := identity.Account

	if err := g.loadGlueCrawlers(svc); err != nil {
		return err
	}
	var DatabaseNames []*string
	if DatabaseNames, err = g.loadGlueCatalogDatabase(svc, account); err != nil {
		return err
	}
	for _, DatabaseName := range DatabaseNames {
		if err := g.loadGlueCatalogTable(svc, account, DatabaseName); err != nil {
			return err
		}
	}

	return nil
}
