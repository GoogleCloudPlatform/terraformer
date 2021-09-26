package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/locks"
)

type ManagementLockGenerator struct {
	AzureService
}

func (az *ManagementLockGenerator) listResources() ([]locks.ManagementLockObject, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := locks.NewManagementLocksClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator locks.ManagementLockListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListAtResourceGroupLevelComplete(ctx, resourceGroup, "")
	} else {
		iterator, err = client.ListAtSubscriptionLevelComplete(ctx, "")
	}
	if err != nil {
		return nil, err
	}
	var resources []locks.ManagementLockObject
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

func (az *ManagementLockGenerator) appendResource(resource *locks.ManagementLockObject) {
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_management_lock")
}

func (az *ManagementLockGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
	}
	return nil
}
