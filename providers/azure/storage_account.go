// Copyright 2019 The Terraformer Authors.
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

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type StorageAccountGenerator struct {
	AzureService
}

func (g StorageAccountGenerator) createResources(accountListResultIterator storage.AccountListResultIterator) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for accountListResultIterator.NotDone() {
		account := accountListResultIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*account.ID,
			*account.Name,
			"azurerm_storage_account",
			"azurerm",
			[]string{}))
		if err := accountListResultIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *StorageAccountGenerator) InitResources() error {
	ctx := context.Background()
	accountsClient := storage.NewAccountsClient(g.Args["config"].(authentication.Config).SubscriptionID)

	accountsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := accountsClient.ListComplete(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
