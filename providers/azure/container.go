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
	"log"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type ContainerGenerator struct {
	AzureService
}

// func (g *ContainerGenerator) listSQLDatabasesAndContainersBehind(resourceGroupName string, accountName string) ([]terraformutils.Resource, []terraformutils.Resource, error) {
// 	var resourcesDatabase []terraformutils.Resource
// 	var resourcesContainer []terraformutils.Resource
// 	ctx := context.Background()
// 	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
// 	SQLResourcesClient := documentdb.NewSQLResourcesClient(subscriptionID, subscriptionID)
// 	SQLResourcesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

// 	sqlDatabases, err := SQLResourcesClient.ListSQLDatabases(ctx, resourceGroupName, accountName)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	for _, sqlDatabase := range *sqlDatabases.Value {
// 		// NOTE:
// 		// For a similar reason as
// 		// https://github.com/terraform-providers/terraform-provider-azurerm/issues/7472#issuecomment-650684349
// 		// The cosmosdb resource format change is NOT yet addressed in terraform provider
// 		// This line is a workaround to convert to old format, and might be removed if they deprecate the old format
// 		sqlDatabaseIDInOldFormat := strings.Replace(*sqlDatabase.ID, "sqlDatabases", "databases", 1)
// 		resourcesDatabase = append(resourcesDatabase, terraformutils.NewSimpleResource(
// 			sqlDatabaseIDInOldFormat,
// 			*sqlDatabase.Name,
// 			"azurerm_cosmosdb_sql_database",
// 			g.ProviderName,
// 			[]string{}))

// 		sqlContainers, err := SQLResourcesClient.ListSQLContainers(ctx, resourceGroupName, accountName, *sqlDatabase.Name)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		for _, sqlContainer := range *sqlContainers.Value {
// 			// NOTE:
// 			// For a similar reason as
// 			// https://github.com/terraform-providers/terraform-provider-azurerm/issues/7472#issuecomment-650684349
// 			// The cosmosdb resource format change is NOT yet addressed in terraform provider
// 			// This line is a workaround to convert to old format, and might be removed if they deprecate the old format
// 			sqlContainerIDInOldFormat := strings.Replace(*sqlContainer.ID, "sqlDatabases", "databases", 1)
// 			resourcesContainer = append(resourcesContainer, terraformutils.NewSimpleResource(
// 				sqlContainerIDInOldFormat,
// 				*sqlContainer.Name,
// 				"azurerm_cosmosdb_sql_container",
// 				g.ProviderName,
// 				[]string{}))
// 		}
// 	}

// 	return resourcesDatabase, resourcesContainer, nil
// }

func (g *ContainerGenerator) listRegistryWebhooks(resourceGroupName string, registryName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	WebhooksClient := containerregistry.NewWebhooksClient(subscriptionID)
	WebhooksClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	webhookIterator, err := WebhooksClient.ListComplete(ctx, resourceGroupName, registryName)
	if err != nil {
		return nil, err
	}
	for webhookIterator.NotDone() {
		webhook := webhookIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*webhook.ID,
			*webhook.Name,
			"azurerm_container_registry_webhook",
			g.ProviderName,
			[]string{}))
		if err := webhookIterator.Next(); err != nil {
			log.Println(err)
			break
		}

	}
	return resources, nil
}

func (g *ContainerGenerator) listAndAddForContainerRegistry() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	ContainerRegistriesClient := containerregistry.NewRegistriesClient(subscriptionID)
	ContainerRegistriesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	containerRegistryIterator, err := ContainerRegistriesClient.ListComplete(ctx)
	if err != nil {
		return nil, err
	}
	for containerRegistryIterator.NotDone() {
		containerRegistry := containerRegistryIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*containerRegistry.ID,
			*containerRegistry.Name,
			"azurerm_container_registry",
			g.ProviderName,
			[]string{}))

		id, err := ParseAzureResourceID(*containerRegistry.ID)
		if err != nil {
			return nil, err
		}

		webhooks, err := g.listRegistryWebhooks(id.ResourceGroup, *containerRegistry.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, webhooks...)

		if err := containerRegistryIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

func (g *ContainerGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listAndAddForContainerRegistry,
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
