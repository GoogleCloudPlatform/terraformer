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

func (g NetworkSecurityGroupGenerator) createResources(securityGroupListResultPage network.SecurityGroupListResultPage) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for securityGroupListResultPage.NotDone() {
		nsgs := securityGroupListResultPage.Values()
		for _, nsg := range nsgs {
			resources = append(resources, terraformutils.NewSimpleResource(
				*nsg.ID,
				*nsg.Name+"-"+*nsg.ID,
				"azurerm_network_security_group",
				"azurerm",
				[]string{}))
		}
		if err := securityGroupListResultPage.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *NetworkSecurityGroupGenerator) InitResources() error {
	ctx := context.Background()
	securityGroupsClient := network.NewSecurityGroupsClient(g.Args["config"].(authentication.Config).SubscriptionID)
	securityGroupsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := securityGroupsClient.ListAll(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
