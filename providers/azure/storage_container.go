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

func NewStorageContainerGenerator(resourceManagerEndpoint string, subscriptionID string, authorizer autorest.Authorizer, rg string) *StorageContainerGenerator {
	storageContainerGenerator := new(StorageContainerGenerator)
	storageContainerGenerator.Args = map[string]interface{}{}
	storageContainerGenerator.Args["config"] = authentication.Config{CustomResourceManagerEndpoint: resourceManagerEndpoint, SubscriptionID: subscriptionID}
	storageContainerGenerator.Args["authorizer"] = authorizer
	storageContainerGenerator.Args["resource_group"] = rg

	return storageContainerGenerator
}

func (g StorageContainerGenerator) ListBlobContainers() ([]terraformutils.Resource, error) {
	var containerResources []terraformutils.Resource
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	blobContainersClient := storage.NewBlobContainersClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	blobContainersClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	ctx := context.Background()

	accounts, err := g.getStorageAccounts()
	if err != nil {
		return containerResources, err
	}

	for _, storageAccount := range accounts {
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
	}

	return containerResources, nil
}

func (g *StorageContainerGenerator) getStorageAccounts() ([]storage.Account, error) {
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	accountsClient := storage.NewAccountsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)

	accountsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	var accounts []storage.Account
	if rg := g.Args["resource_group"].(string); rg != "" {
		accountsResult, err := accountsClient.ListByResourceGroup(ctx, rg)
		if err != nil {
			return nil, err
		}
		if paccounts := accountsResult.Value; paccounts != nil {
			accounts = append(accounts, *paccounts...)
		}
	} else {
		accountsIterator, err := accountsClient.ListComplete(ctx)
		if err != nil {
			return nil, err
		}
		for accountsIterator.NotDone() {
			account := accountsIterator.Value()
			accounts = append(accounts, account)
			if err := accountsIterator.NextWithContext(ctx); err != nil {
				return accounts, err
			}
		}
	}

	return accounts, nil
}

func (g *StorageContainerGenerator) InitResources() error {
	storageAccounts, err := g.ListBlobContainers()
	if err != nil {
		return err
	}

	g.Resources = storageAccounts

	return nil
}
