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
