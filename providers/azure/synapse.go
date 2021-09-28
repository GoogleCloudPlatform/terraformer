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

	"github.com/Azure/azure-sdk-for-go/services/preview/synapse/2019-06-01-preview/managedvirtualnetwork"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2020-12-01/synapse"
	// "github.com/Azure/azure-sdk-for-go/services/preview/synapse/2020-08-01-preview/accesscontrol"
)

type SynapseGenerator struct {
	AzureService
}

func (az *SynapseGenerator) listWorkspaces() ([]synapse.Workspace, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := synapse.NewWorkspacesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator synapse.WorkspaceInfoListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []synapse.Workspace
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

func (az *SynapseGenerator) appendWorkspace(workspace *synapse.Workspace) {
	az.AppendSimpleResource(*workspace.ID, *workspace.Name, "azurerm_synapse_workspace")
}

func (az *SynapseGenerator) appendSQLPools(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewSQLPoolsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_sql_pool")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendSparkPools(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewBigDataPoolsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_spark_pool")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendFirewallRule(workspace *synapse.Workspace, workspaceRg *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := synapse.NewIPFirewallRulesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListByWorkspaceComplete(ctx, workspaceRg.ResourceGroup, *workspace.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_firewall_rule")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) appendManagedPrivateEndpoint(workspace *synapse.Workspace) error {

	if workspace.WorkspaceProperties == nil || workspace.WorkspaceProperties.ManagedVirtualNetwork == nil {
		return nil
	}
	virtualNetworkName := *workspace.WorkspaceProperties.ManagedVirtualNetwork
	if virtualNetworkName == "" || virtualNetworkName == "default" {
		return nil
	}
	subscriptionID, _, authorizer := az.getClientArgs()
	client := managedvirtualnetwork.NewManagedPrivateEndpointsClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, virtualNetworkName)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_synapse_managed_private_endpoint")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SynapseGenerator) listPrivateLinkHubs() ([]synapse.PrivateLinkHub, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := synapse.NewPrivateLinkHubsClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator synapse.PrivateLinkHubInfoListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []synapse.PrivateLinkHub
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

func (az *SynapseGenerator) appendtPrivateLinkHubs(workspace *synapse.PrivateLinkHub) {
	az.AppendSimpleResource(*workspace.ID, *workspace.Name, "azurerm_synapse_private_link_hub")
}

func (az *SynapseGenerator) InitResources() error {

	workspaces, err := az.listWorkspaces()
	if err != nil {
		return err
	}
	for _, workspace := range workspaces {
		az.appendWorkspace(&workspace)
		workspaceRg, err := ParseAzureResourceID(*workspace.ID)
		if err != nil {
			return err
		}
		err = az.appendSQLPools(&workspace, workspaceRg)
		if err != nil {
			return err
		}
		err = az.appendSparkPools(&workspace, workspaceRg)
		if err != nil {
			return err
		}
		err = az.appendFirewallRule(&workspace, workspaceRg)
		if err != nil {
			return err
		}
		err = az.appendManagedPrivateEndpoint(&workspace)
		if err != nil {
			return err
		}
	}

	hubs, err := az.listPrivateLinkHubs()
	if err == nil {
		for _, hub := range hubs {
			az.appendtPrivateLinkHubs(&hub)
		}
	}
	return nil
}
