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
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_route_table")
}

func (az *RouteTableGenerator) appendRoutes(parent *network.RouteTable, resourceGroupId *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := network.NewRoutesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, resourceGroupId.ResourceGroup, *parent.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_route")
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
		resourceGroupId, err := ParseAzureResourceID(*resource.ID)
		if err != nil {
			return err
		}
		err = az.appendRoutes(&resource, resourceGroupId)
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
