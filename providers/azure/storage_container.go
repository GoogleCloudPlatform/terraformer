package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

const (
	containerIDFormat = "https://%s.blob.core.windows.net/%s"
)

type StorageContainerGenerator struct {
	AzureService
}

func NewStorageContainerGenerator(subscriptionID string, authorizer autorest.Authorizer) *StorageContainerGenerator {
	storageContainerGenerator := new(StorageContainerGenerator)
	storageContainerGenerator.Args = map[string]interface{}{}
	storageContainerGenerator.Args["config"] = authentication.Config{SubscriptionID: subscriptionID}
	storageContainerGenerator.Args["authorizer"] = authorizer

	return storageContainerGenerator
}

func (g StorageContainerGenerator) ListBlobContainers() ([]terraformutils.Resource, error) {
	var containerResources []terraformutils.Resource
	blobContainersClient := storage.NewBlobContainersClient(g.Args["config"].(authentication.Config).SubscriptionID)
	blobContainersClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	ctx := context.Background()

	accountListResultIterator, err := g.getStorageAccountsIterator()
	if err != nil {
		return containerResources, err
	}

	for accountListResultIterator.NotDone() {
		storageAccount := accountListResultIterator.Value()
		parsedStorageAccountResourceID, err := ParseAzureResourceID(*storageAccount.ID)
		if err != nil {
			break
		}
		containerItemsIterator, err := blobContainersClient.ListComplete(ctx, parsedStorageAccountResourceID.ResourceGroup, *storageAccount.Name, "", "", "")
		if err != nil {
			return containerResources, err
		}

		for containerItemsIterator.NotDone() {
			containerItem := containerItemsIterator.Value()
			containerResources = append(containerResources,
				terraformutils.NewResource(
					fmt.Sprintf(containerIDFormat, *storageAccount.Name, *containerItem.Name),
					*containerItem.Name,
					"azurerm_storage_container",
					"azurerm",
					map[string]string{
						"storage_account_name": *storageAccount.Name,
						"name":                 *containerItem.Name,
					},
					[]string{},
					map[string]interface{}{}))

			if err := containerItemsIterator.NextWithContext(ctx); err != nil {
				return containerResources, err
			}
		}

		if err := accountListResultIterator.NextWithContext(ctx); err != nil {
			return containerResources, err
		}
	}

	return containerResources, nil
}

func (g *StorageContainerGenerator) getStorageAccountsIterator() (storage.AccountListResultIterator, error) {
	ctx := context.Background()
	accountsClient := storage.NewAccountsClient(g.Args["config"].(authentication.Config).SubscriptionID)

	accountsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	accountsIterator, err := accountsClient.ListComplete(ctx)

	return accountsIterator, err
}

func (g *StorageContainerGenerator) InitResources() error {
	storageAccounts, err := g.ListBlobContainers()
	if err != nil {
		return err
	}

	g.Resources = storageAccounts

	return nil
}
