// Copyright 2020 The Terraformer Authors.
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

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type KeyVaultGenerator struct {
	AzureService
}

func (g KeyVaultGenerator) createResources(resourceListResultIterator keyvault.ResourceListResultIterator) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for resourceListResultIterator.NotDone() {
		vault := resourceListResultIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*vault.ID,
			*vault.Name,
			"azurerm_key_vault",
			"azurerm",
			[]string{}))
		if err := resourceListResultIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *KeyVaultGenerator) InitResources() error {
	ctx := context.Background()
	vaultsClient := keyvault.NewVaultsClient(g.Args["config"].(authentication.Config).SubscriptionID)

	vaultsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := vaultsClient.ListComplete(ctx, nil)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
