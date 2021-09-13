package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/purview/mgmt/2021-07-01/purview"
)

type PurviewGenerator struct {
	AzureService
}

func (az *PurviewGenerator) listAccounts() ([]purview.Account, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := purview.NewAccountsClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator purview.AccountListIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup, "")
	} else {
		iterator, err = client.ListBySubscriptionComplete(ctx, "")
	}
	if err != nil {
		return nil, err
	}
	var resources []purview.Account
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

func (az *PurviewGenerator) AppendAccount(account *purview.Account) {
	az.AppendSimpleResource(*account.ID, *account.Name, "azurerm_purview_account")
}

func (az *PurviewGenerator) InitResources() error {

	accounts, err := az.listAccounts()
	if err != nil {
		return err
	}
	for _, account := range accounts {
		az.AppendAccount(&account)
	}
	return nil
}
