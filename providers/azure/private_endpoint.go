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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
)

type PrivateEndpointGenerator struct {
	AzureService
}

func (az *PrivateEndpointGenerator) listServices() ([]network.PrivateLinkService, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewPrivateLinkServicesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.PrivateLinkServiceListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListBySubscriptionComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []network.PrivateLinkService
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

func (az *PrivateEndpointGenerator) AppendServices(link *network.PrivateLinkService) {
	az.AppendSimpleResource(*link.ID, *link.Name, "azurerm_private_link_service")
}

func (az *PrivateEndpointGenerator) listEndpoints() ([]network.PrivateEndpoint, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewPrivateEndpointsClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.PrivateEndpointListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListBySubscriptionComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []network.PrivateEndpoint
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

func (az *PrivateEndpointGenerator) AppendEndpoint(link *network.PrivateEndpoint) {
	az.AppendSimpleResource(*link.ID, *link.Name, "azurerm_private_endpoint")
}

func (az *PrivateEndpointGenerator) InitResources() error {

	services, err := az.listServices()
	if err != nil {
		return err
	}
	for _, link := range services {
		az.AppendServices(&link)
	}
	endpoints, err := az.listEndpoints()
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints {
		az.AppendEndpoint(&endpoint)
	}
	return nil
}
