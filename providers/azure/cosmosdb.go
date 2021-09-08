// Copyright 2020 The Terraformer Authors.
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

package azure

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-06-15/documentdb"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type CosmosDBGenerator struct {
	AzureService
}

func (g *CosmosDBGenerator) listSQLDatabasesAndContainersBehind(resourceGroupName string, accountName string) ([]terraformutils.Resource, []terraformutils.Resource, error) {
	var resourcesDatabase []terraformutils.Resource
	var resourcesContainer []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	SQLResourcesClient := documentdb.NewSQLResourcesClient(subscriptionID)
	SQLResourcesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	sqlDatabases, err := SQLResourcesClient.ListSQLDatabases(ctx, resourceGroupName, accountName)
	if err != nil {
		return nil, nil, err
	}
	for _, sqlDatabase := range *sqlDatabases.Value {
		// NOTE:
		// For a similar reason as
		// https://github.com/terraform-providers/terraform-provider-azurerm/issues/7472#issuecomment-650684349
		// The cosmosdb resource format change is NOT yet addressed in terraform provider
		// This line is a workaround to convert to old format, and might be removed if they deprecate the old format
		sqlDatabaseIDInOldFormat := strings.Replace(*sqlDatabase.ID, "sqlDatabases", "databases", 1)
		resourcesDatabase = append(resourcesDatabase, terraformutils.NewSimpleResource(
			sqlDatabaseIDInOldFormat,
			*sqlDatabase.Name,
			"azurerm_cosmosdb_sql_database",
			g.ProviderName,
			[]string{}))

		sqlContainers, err := SQLResourcesClient.ListSQLContainers(ctx, resourceGroupName, accountName, *sqlDatabase.Name)
		if err != nil {
			return nil, nil, err
		}
		for _, sqlContainer := range *sqlContainers.Value {
			// NOTE:
			// For a similar reason as
			// https://github.com/terraform-providers/terraform-provider-azurerm/issues/7472#issuecomment-650684349
			// The cosmosdb resource format change is NOT yet addressed in terraform provider
			// This line is a workaround to convert to old format, and might be removed if they deprecate the old format
			sqlContainerIDInOldFormat := strings.Replace(*sqlContainer.ID, "sqlDatabases", "databases", 1)
			resourcesContainer = append(resourcesContainer, terraformutils.NewSimpleResource(
				sqlContainerIDInOldFormat,
				*sqlContainer.Name,
				"azurerm_cosmosdb_sql_container",
				g.ProviderName,
				[]string{}))
		}
	}

	return resourcesDatabase, resourcesContainer, nil
}

func (g *CosmosDBGenerator) listTables(resourceGroupName string, accountName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	TableResourcesClient := documentdb.NewTableResourcesClient(subscriptionID)
	TableResourcesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	tables, err := TableResourcesClient.ListTables(ctx, resourceGroupName, accountName)
	if err != nil {
		return nil, err
	}
	for _, table := range *tables.Value {
		resources = append(resources, terraformutils.NewSimpleResource(
			*table.ID,
			*table.Name,
			"azurerm_cosmosdb_table",
			g.ProviderName,
			[]string{}))
	}

	return resources, nil
}

func (g *CosmosDBGenerator) listAndAddForDatabaseAccounts() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	DatabaseAccountsClient := documentdb.NewDatabaseAccountsClient(subscriptionID)
	DatabaseAccountsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var (
		accounts documentdb.DatabaseAccountsListResult
		err      error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		accounts, err = DatabaseAccountsClient.ListByResourceGroup(ctx, rg)
	} else {
		accounts, err = DatabaseAccountsClient.List(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, account := range *accounts.Value {
		resources = append(resources, terraformutils.NewSimpleResource(
			*account.ID,
			*account.Name,
			"azurerm_cosmosdb_account",
			g.ProviderName,
			[]string{}))

		id, err := ParseAzureResourceID(*account.ID)
		if err != nil {
			return nil, err
		}

		tables, err := g.listTables(id.ResourceGroup, *account.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, tables...)

		sqlDatabases, sqlContainers, err := g.listSQLDatabasesAndContainersBehind(id.ResourceGroup, *account.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, sqlDatabases...)
		resources = append(resources, sqlContainers...)
	}

	return resources, nil
}

func (g *CosmosDBGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listAndAddForDatabaseAccounts,
	}

	for _, f := range functions {
		resources, err := f()
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}
