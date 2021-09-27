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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
)

type NetworkWatcherGenerator struct {
	AzureService
}

func (az *NetworkWatcherGenerator) listResources() ([]network.Watcher, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewWatchersClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		resources network.WatcherListResult
		err       error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		resources, err = client.List(ctx, resourceGroup)
	} else {
		resources, err = client.ListAll(ctx)
	}
	if err != nil {
		return nil, err
	}
	return *resources.Value, nil
}

func (az *NetworkWatcherGenerator) appendResource(resource *network.Watcher) {
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_network_watcher")
}

func (az *NetworkWatcherGenerator) appendFlowLogs(parent *network.Watcher, resourceGroupID *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := network.NewFlowLogsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, resourceGroupID.ResourceGroup, *parent.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_network_watcher_flow_log")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *NetworkWatcherGenerator) appendPacketCaptures(parent *network.Watcher, resourceGroupID *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := network.NewPacketCapturesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	resources, err := client.List(ctx, resourceGroupID.ResourceGroup, *parent.Name)
	if err != nil {
		return err
	}
	for _, item := range *resources.Value {
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_network_packet_capture")
	}
	return nil
}

func (az *NetworkWatcherGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
		resourceGroupID, err := ParseAzureResourceID(*resource.ID)
		if err != nil {
			return err
		}
		err = az.appendFlowLogs(&resource, resourceGroupID)
		if err != nil {
			return err
		}
		err = az.appendPacketCaptures(&resource, resourceGroupID)
		if err != nil {
			return err
		}
	}
	return nil
}
