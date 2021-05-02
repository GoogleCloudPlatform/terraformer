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
	p := glue.NewGetCrawlersPaginator(svc, &glue.GetCrawlersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, crawler := range page.Crawlers {
			resource := terraformutils.NewSimpleResource(*crawler.Name, *crawler.Name,
				"aws_glue_crawler",
				"aws",
				GlueCrawlerAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}

func (g *GlueGenerator) loadGlueCatalogDatabase(svc *glue.Client, account *string) (databaseNames []*string, error error) {
	var GlueCatalogDatabaseAllowEmptyValues = []string{"tags."}
	p := glue.NewGetDatabasesPaginator(svc, &glue.GetDatabasesInput{})
	for p.HasMorePages() {
		page, error := p.NextPage(context.TODO())
		if error != nil {
			return databaseNames, error
		}
		for _, catalogDatabase := range page.DatabaseList {
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
	return databaseNames, nil
}

func (g *GlueGenerator) loadGlueCatalogTable(svc *glue.Client, account *string, databaseName *string) error {
	// format of ID is "CATALOG-ID:DATABASE-NAME:TABLE-NAME".
	// CATALOG-ID is AWS Account ID
	// https://docs.aws.amazon.com/cli/latest/reference/glue/create-database.html#options
	var GlueCatalogTableAllowEmptyValues = []string{"tags."}
	p := glue.NewGetTablesPaginator(svc, &glue.GetTablesInput{DatabaseName: databaseName})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, catalogTable := range page.TableList {
			databaseTable := *databaseName + ":" + *catalogTable.Name
			id := *account + ":" + databaseTable
			resource := terraformutils.NewSimpleResource(id, databaseTable,
				"aws_glue_catalog_table",
				"aws",
				GlueCatalogTableAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}

func (g *GlueGenerator) loadGlueJobs(svc *glue.Client) error {
	var GlueJobAllowEmptyValues = []string{"tags."}
	p := glue.NewGetJobsPaginator(svc, &glue.GetJobsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, job := range page.Jobs {
			resource := terraformutils.NewSimpleResource(*job.Name, *job.Name,
				"aws_glue_job",
				"aws",
				GlueJobAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}

func (g *GlueGenerator) loadGlueTriggers(svc *glue.Client) error {
	var GlueTriggerAllowEmptyValues = []string{"tags."}
	p := glue.NewGetTriggersPaginator(svc, &glue.GetTriggersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, trigger := range page.Triggers {
			resource := terraformutils.NewSimpleResource(*trigger.Name, *trigger.Name,
				"aws_glue_trigger",
				"aws",
				GlueTriggerAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
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
	svc := glue.NewFromConfig(config)

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

	if err := g.loadGlueJobs(svc); err != nil {
		return err
	}

	if err := g.loadGlueTriggers(svc); err != nil {
		return err
	}

	return nil
}
