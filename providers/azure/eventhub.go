// Copyright 2021 The Terraformer Authors.
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

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
)

type EventHubGenerator struct {
	AzureService
}

func (az *EventHubGenerator) listNamespaces() ([]eventhub.EHNamespace, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := eventhub.NewNamespacesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator eventhub.EHNamespaceListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []eventhub.EHNamespace
	for iterator.NotDone() {
		item := iterator.Value()
		resources = append(resources, item)
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *EventHubGenerator) AppendNamespace(namespace *eventhub.EHNamespace) {
	az.AppendSimpleResource(*namespace.ID, *namespace.Name, "azurerm_eventhub_namespace")
}

func (az *EventHubGenerator) appendEventHubs(namespace *eventhub.EHNamespace, namespaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := eventhub.NewEventHubsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByNamespaceComplete(ctx, namespaceRg.ResourceGroup, *namespace.Name, nil, nil)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()

		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_eventhub")
		err = az.appendConsumerGroups(namespace, namespaceRg, *item.Name)
		if err != nil {
			return err
		}
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *EventHubGenerator) appendConsumerGroups(namespace *eventhub.EHNamespace, namespaceRg *ResourceID, eventHubName string) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := eventhub.NewConsumerGroupsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByEventHubComplete(ctx, namespaceRg.ResourceGroup, *namespace.Name, eventHubName, nil, nil)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_eventhub_consumer_group")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *EventHubGenerator) appendAuthorizationRules(namespace *eventhub.EHNamespace, namespaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := eventhub.NewNamespacesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListAuthorizationRulesComplete(ctx, namespaceRg.ResourceGroup, *namespace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()

		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_eventhub_namespace_authorization_rule")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *EventHubGenerator) InitResources() error {

	namespaces, err := az.listNamespaces()
	if err != nil {
		return err
	}
	for _, namespace := range namespaces {
		az.AppendNamespace(&namespace)
		namespaceRg, err := ParseAzureResourceID(*namespace.ID)
		if err != nil {
			return err
		}
		err = az.appendEventHubs(&namespace, namespaceRg)
		if err != nil {
			return err
		}
		err = az.appendAuthorizationRules(&namespace, namespaceRg)
		if err != nil {
			return err
		}
	}
	return nil
}
