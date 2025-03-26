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

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type ResourceGroupGenerator struct {
	AzureService
}

func (g ResourceGroupGenerator) createResources(groupListResultIterator resources.GroupListResultIterator) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for groupListResultIterator.NotDone() {
		group := groupListResultIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*group.ID,
			*group.Name,
			"azurerm_resource_group",
			"azurerm",
			[]string{}))
		if err := groupListResultIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *ResourceGroupGenerator) InitResources() error {
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	groupsClient := resources.NewGroupsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)

	groupsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	if rg := g.Args["resource_group"].(string); rg != "" {
		group, err := groupsClient.Get(ctx, rg)
		if err != nil {
			return err
		}
		g.Resources = []terraformutils.Resource{
			terraformutils.NewSimpleResource(
				*group.ID,
				*group.Name,
				"azurerm_resource_group",
				"azurerm",
				[]string{}),
		}
		return nil
	}
	output, err := groupsClient.ListComplete(ctx, "", nil)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
