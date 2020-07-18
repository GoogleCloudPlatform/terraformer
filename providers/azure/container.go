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

	"github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2018-10-01/containerinstance"
	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type ContainerGenerator struct {
	AzureService
}

func (g *ContainerGenerator) listAndAddForContainerGroup() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	ContainerGroupsClient := containerinstance.NewContainerGroupsClient(subscriptionID)
	ContainerGroupsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	containerGroupIterator, err := ContainerGroupsClient.ListComplete(ctx)
	if err != nil {
		return nil, err
	}
	for containerGroupIterator.NotDone() {
		containerGroup := containerGroupIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*containerGroup.ID,
			*containerGroup.Name,
			"azurerm_container_group",
			g.ProviderName,
			[]string{}))

		if err := containerGroupIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

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
		g.listAndAddForContainerGroup,
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
