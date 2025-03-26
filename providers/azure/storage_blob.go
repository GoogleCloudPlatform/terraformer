package azure

import (
	"context"
	"fmt"

	"log"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

const (
	blobFormatString = `https://%s.blob.core.windows.net`
	blobIDFormat     = `https://%s.blob.core.windows.net/%s/%s`
)

type StorageBlobGenerator struct {
	AzureService
}

func (g StorageBlobGenerator) getAccountPrimaryKey(ctx context.Context, accountName, accountGroupName string) string {
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	storageAccountsClient := storage.NewAccountsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	storageAccountsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	response, err := storageAccountsClient.ListKeys(ctx, accountGroupName, accountName, "kerb")
	if err != nil {
		log.Fatalf("failed to list keys: %v", err)
	}
	return *(((*response.Keys)[0]).Value)
}

func (g StorageBlobGenerator) getContainerURL(ctx context.Context, accountName, accountGroupName, containerName string) (azblob.ContainerURL, error) {
	accountPrimaryKey := g.getAccountPrimaryKey(ctx, accountName, accountGroupName)
	sharedKeyCredential, err := azblob.NewSharedKeyCredential(accountName, accountPrimaryKey)
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	p := azblob.NewPipeline(sharedKeyCredential, azblob.PipelineOptions{})
	accountURL, err := url.Parse(fmt.Sprintf(blobFormatString, accountName))
	if err != nil {
		return azblob.ContainerURL{}, err
	}

	serviceURL := azblob.NewServiceURL(*accountURL, p)
	containerURL := serviceURL.NewContainerURL(containerName)

	return containerURL, nil
}

func (g StorageBlobGenerator) getBlobsFromContainer(ctx context.Context, accountName, accountGroupName, containerName string) ([]azblob.BlobItem, error) {
	containerURL, err := g.getContainerURL(ctx, accountName, accountGroupName, containerName)
	if err != nil {
		return nil, err
	}

	blobListResponse, err := containerURL.ListBlobsFlatSegment(
		ctx,
		azblob.Marker{},
		azblob.ListBlobsSegmentOptions{
			Details: azblob.BlobListingDetails{
				Snapshots: true,
			},
		})
	if err != nil {
		return nil, err
	}

	return blobListResponse.Segment.BlobItems, nil
}

func (g StorageBlobGenerator) listStorageBlobs() ([]terraformutils.Resource, error) {
	var storageBlobsResources []terraformutils.Resource
	ctx := context.Background()

	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	authorizer := g.Args["authorizer"].(autorest.Authorizer)
	resourceGroup := g.Args["resource_group"].(string)
	blobContainerGenerator := NewStorageContainerGenerator(resourceManagerEndpoint, subscriptionID, authorizer, resourceGroup)
	blobContainersResources, err := blobContainerGenerator.ListBlobContainers()
	if err != nil {
		return storageBlobsResources, err
	}

	for _, blobContainerResource := range blobContainersResources {
		containerID := blobContainerResource.InstanceState.ID
		parsedContainerID, err := ParseAzureResourceID(containerID)
		if err != nil {
			return storageBlobsResources, err
		}

		storageAccountName := blobContainerResource.InstanceState.Attributes["storage_account_name"]
		containerName := blobContainerResource.InstanceState.Attributes["name"]
		blobsList, err := g.getBlobsFromContainer(ctx, storageAccountName, parsedContainerID.ResourceGroup, containerName)
		if err != nil {
			return storageBlobsResources, err
		}

		for _, blobItem := range blobsList {
			storageBlobsResources = append(storageBlobsResources, terraformutils.NewSimpleResource(
				fmt.Sprintf(blobIDFormat, storageAccountName, containerName, blobItem.Name),
				blobItem.Name,
				"azurerm_storage_blob",
				"azurerm",
				[]string{}))
		}
	}

	return storageBlobsResources, err
}

func (g *StorageBlobGenerator) InitResources() error {
	resources, err := g.listStorageBlobs()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, resources...)

	return nil
}
