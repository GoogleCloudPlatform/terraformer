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
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

type GlueGenerator struct {
	AWSService
}

func (g *GlueGenerator) loadGlueCrawlers(svc *glue.Client) error {
	var GlueCrawlerAllowEmptyValues = []string{"tags."}
	p := glue.NewGetCrawlersPaginator(svc.GetCrawlersRequest(&glue.GetCrawlersInput{}))
	for p.Next(context.Background()) {
		for _, crawler := range p.CurrentPage().Crawlers {
			resource := terraformutils.NewSimpleResource(*crawler.Name, *crawler.Name,
				"aws_glue_crawler",
				"aws",
				GlueCrawlerAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return p.Err()
}

func (g *GlueGenerator) loadGlueCatalogDatabase(svc *glue.Client, account *string) (databaseNames []*string, error error) {
	var GlueCatalogDatabaseAllowEmptyValues = []string{"tags."}
	p := glue.NewGetDatabasesPaginator(svc.GetDatabasesRequest(&glue.GetDatabasesInput{}))
	for p.Next(context.Background()) {
		for _, catalogDatabase := range p.CurrentPage().DatabaseList {
			// format of ID is "CATALOG-ID:DATABASE-NAME".
			// CATALOG-ID is AWS Account ID
			// https://docs.aws.amazon.com/cli/latest/reference/glue/create-database.html#options
			id := *account + ":" + *catalogDatabase.Name
			resource := terraformutils.NewSimpleResource(id, *catalogDatabase.Name,
				"aws_glue_catalog_database",
				"aws",
				GlueCatalogDatabaseAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
			databaseNames = append(databaseNames, catalogDatabase.Name)
		}
	}
	return databaseNames, p.Err()
}

func (g *GlueGenerator) loadGlueCatalogTable(svc *glue.Client, account *string, databaseName *string) error {
	// format of ID is "CATALOG-ID:DATABASE-NAME:TABLE-NAME".
	// CATALOG-ID is AWS Account ID
	// https://docs.aws.amazon.com/cli/latest/reference/glue/create-database.html#options
	var GlueCatalogTableAllowEmptyValues = []string{"tags."}
	p := glue.NewGetTablesPaginator(svc.GetTablesRequest(&glue.GetTablesInput{DatabaseName: databaseName}))
	for p.Next(context.Background()) {
		for _, catalogTable := range p.CurrentPage().TableList {
			databaseTable := *databaseName + ":" + *catalogTable.Name
			id := *account + ":" + databaseTable
			resource := terraformutils.NewSimpleResource(id, databaseTable,
				"aws_glue_catalog_table",
				"aws",
				GlueCatalogTableAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return p.Err()
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *GlueGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := glue.New(config)

	account, err := g.getAccountNumber(config)
	if err != nil {
		return err
	}

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
