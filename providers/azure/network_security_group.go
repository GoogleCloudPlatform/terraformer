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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type NetworkSecurityGroupGenerator struct {
	AzureService
}

func (g NetworkSecurityGroupGenerator) createResources(securityGroupListResult network.SecurityGroupListResultIterator) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for securityGroupListResult.NotDone() {
		nsg := securityGroupListResult.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*nsg.ID,
			*nsg.Name+"-"+*nsg.ID,
			"azurerm_network_security_group",
			"azurerm",
			[]string{}))
		if err := securityGroupListResult.Next(); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (g *NetworkSecurityGroupGenerator) InitResources() error {
	ctx := context.Background()
	securityGroupsClient := network.NewSecurityGroupsClient(g.Args["config"].(authentication.Config).SubscriptionID)
	securityGroupsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var (
		output network.SecurityGroupListResultIterator
		err    error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		output, err = securityGroupsClient.ListComplete(ctx, rg)
	} else {
		output, err = securityGroupsClient.ListAllComplete(ctx)
	}
	if err != nil {
		return err
	}
	g.Resources, err = g.createResources(output)
	return err
}
