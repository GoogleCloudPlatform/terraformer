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

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type ScaleSetGenerator struct {
	AzureService
}

func (g ScaleSetGenerator) createResourcesByResourceGroup(ctx context.Context, client compute.VirtualMachineScaleSetsClient, rg string) ([]terraformutils.Resource, error) {
	scaleSetIterator, err := client.ListComplete(ctx, rg)
	if err != nil {
		return nil, err
	}
	var resources []terraformutils.Resource
	for scaleSetIterator.NotDone() {
		scaleSet := scaleSetIterator.Value()
		newResource := terraformutils.NewSimpleResource(
			*scaleSet.ID,
			*scaleSet.Name,
			"azurerm_virtual_machine_scale_set",
			"azurerm",
			[]string{})
		resources = append(resources, newResource)
		if err := scaleSetIterator.Next(); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (g ScaleSetGenerator) createResources(ctx context.Context, client compute.VirtualMachineScaleSetsClient) ([]terraformutils.Resource, error) {
	scaleSetIterator, err := client.ListAllComplete(ctx)
	if err != nil {
		return nil, err
	}
	var resources []terraformutils.Resource
	for scaleSetIterator.NotDone() {
		scaleSet := scaleSetIterator.Value()
		newResource := terraformutils.NewSimpleResource(
			*scaleSet.ID,
			*scaleSet.Name,
			"azurerm_virtual_machine_scale_set",
			"azurerm",
			[]string{})
		resources = append(resources, newResource)
		if err := scaleSetIterator.Next(); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (g *ScaleSetGenerator) InitResources() error {
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	ScaleSetClient := compute.NewVirtualMachineScaleSetsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)

	ScaleSetClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	if rg := g.Args["resource_group"].(string); rg != "" {
		var err error
		g.Resources, err = g.createResourcesByResourceGroup(ctx, ScaleSetClient, rg)
		return err
	}
	var err error
	g.Resources, err = g.createResources(ctx, ScaleSetClient)
	return err
}
