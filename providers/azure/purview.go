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
