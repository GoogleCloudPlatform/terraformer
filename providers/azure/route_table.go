// Copyright 2021 Josh Garverick
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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type RouteTableGenerator struct {
	AzureService
}

func (g *RouteTableGenerator) getRouteTables(routeTables network.RouteTableListResultIterator) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	for routeTables.NotDone() {
		var routeTable = routeTables.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*routeTable.ID,
			*routeTable.Name,
			"azurerm_route_table",
			"azurerm",
			[]string{}))
		if err := routeTables.Next(); err != nil {
			return nil, err
		}
	}
	return resources, nil

}

func (g *RouteTableGenerator) InitResources() error {

	ctx := context.Background()
	SubscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	Authorizer := g.Args["authorizer"].(autorest.Authorizer)

	Client := network.NewRouteTablesClient(SubscriptionID)
	Client.Authorizer = Authorizer
	var (
		output network.RouteTableListResultIterator
		err    error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		output, err = Client.ListComplete(ctx, rg)
	} else {
		output, err = Client.ListAllComplete(ctx)
	}
	resources, err := g.getRouteTables(output)
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, resources...)

	return err
}
