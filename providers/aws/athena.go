// Copyright 2023 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/athena"
)

var athenaAllowEmptyValues = []string{"tags."}

type AthenaGenerator struct {
	AWSService
}

func (g *AthenaGenerator) InitResources() error {
	config, err := g.generateConfig()
	if err != nil {
		return err
	}
	svc := athena.NewFromConfig(config)
	err = g.createWorkGroups(svc)
	if err != nil {
		return err
	}
	err = g.createDataCatalogs(svc)
	if err != nil {
		return err
	}
	err = g.createDatabase(svc)
	if err != nil {
		return err
	}
	err = g.createNamedQueries(svc)
	if err != nil {
		return err
	}

	return nil
}

func (g *AthenaGenerator) createWorkGroups(svc *athena.Client) error {
	p := athena.NewListWorkGroupsPaginator(svc, &athena.ListWorkGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}

		for _, workGroup := range page.WorkGroups {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*workGroup.Name,
				*workGroup.Name,
				"aws_athena_workgroup",
				"aws",
				athenaAllowEmptyValues))

		}
	}

	return nil
}

func (g *AthenaGenerator) createDataCatalogs(svc *athena.Client) error {
	p := athena.NewListDataCatalogsPaginator(svc, &athena.ListDataCatalogsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.Background())
		if err != nil {
			return err
		}

		for _, dataCatalog := range page.DataCatalogsSummary {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*dataCatalog.CatalogName,
				*dataCatalog.CatalogName,
				"aws_athena_data_catalog",
				"aws",
				athenaAllowEmptyValues))
		}
	}

	return nil
}

func (g *AthenaGenerator) createDatabase(svc *athena.Client) error {
	p := athena.NewListDatabasesPaginator(svc, &athena.ListDatabasesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.Background())
		if err != nil {
			return err
		}

		for _, database := range page.DatabaseList {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*database.Name,
				*database.Name,
				"aws_athena_database",
				"aws",
				athenaAllowEmptyValues))
		}
	}

	return nil
}

func (g *AthenaGenerator) createNamedQueries(svc *athena.Client) error {
	p := athena.NewListNamedQueriesPaginator(svc, &athena.ListNamedQueriesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.Background())
		if err != nil {
			return err
		}

		for _, namedQueryId := range page.NamedQueryIds {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				namedQueryId,
				namedQueryId,
				"aws_athena_named_query",
				"aws",
				athenaAllowEmptyValues))
		}
	}

	return nil
}
