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

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type DiskGenerator struct {
	AzureService
}

func (g DiskGenerator) createResources(diskListIterator compute.DiskListIterator) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for diskListIterator.NotDone() {
		disk := diskListIterator.Value()
		resources = append(resources, terraform_utils.NewSimpleResource(
			*disk.ID,
			*disk.Name,
			"azurerm_managed_disk",
			"azurerm",
			[]string{}))
		if err := diskListIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *DiskGenerator) InitResources() error {
	ctx := context.Background()
	disksClient := compute.NewDisksClient(g.Args["subscription"].(string))
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return err
	}
	disksClient.Authorizer = authorizer
	output, err := disksClient.ListComplete(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
