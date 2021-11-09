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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
)

type RouteTableGenerator struct {
	AzureService
}

func (az *RouteTableGenerator) listResources() ([]network.RouteTable, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewRouteTablesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.RouteTableListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListAllComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []network.RouteTable
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

func (az *RouteTableGenerator) appendResource(resource *network.RouteTable) {
	az.AppendSimpleResourceWithDuplicateCheck(*resource.ID, *resource.Name, "azurerm_route_table")
}

func (az *RouteTableGenerator) appendRoutes(parent *network.RouteTable, resourceGroupID *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := network.NewRoutesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, resourceGroupID.ResourceGroup, *parent.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResourceWithDuplicateCheck(*item.ID, *item.Name, "azurerm_route")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *RouteTableGenerator) listRouteFilters() ([]network.RouteFilter, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewRouteFiltersClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.RouteFilterListResultIterator
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
	var resources []network.RouteFilter
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

func (az *RouteTableGenerator) appendRouteFilters(resource *network.RouteFilter) {
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_route_filter")
}

func (az *RouteTableGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
		resourceGroupID, err := ParseAzureResourceID(*resource.ID)
		if err != nil {
			return err
		}
		err = az.appendRoutes(&resource, resourceGroupID)
		if err != nil {
			return err
		}
	}

	filters, err := az.listRouteFilters()
	if err != nil {
		return err
	}
	for _, resource := range filters {
		az.appendRouteFilters(&resource)
		if err != nil {
			return err
		}
	}
	return nil
}
