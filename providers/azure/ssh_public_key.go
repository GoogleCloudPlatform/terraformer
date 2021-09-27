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

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
)

type SSHPublicKeyGenerator struct {
	AzureService
}

func (az *SSHPublicKeyGenerator) listResources() ([]compute.SSHPublicKeyResource, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := compute.NewSSHPublicKeysClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator compute.SSHPublicKeysGroupListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListBySubscriptionComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []compute.SSHPublicKeyResource
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

func (az *SSHPublicKeyGenerator) appendResource(resource *compute.SSHPublicKeyResource) {
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_ssh_public_key")
}

func (az *SSHPublicKeyGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
	}
	return nil
}
