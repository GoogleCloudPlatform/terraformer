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

	"github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-08-01/analysisservices"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type AnalysisGenerator struct {
	AzureService
}

func (g *AnalysisGenerator) listServiceServers() ([]terraformutils.Resource, error) {
	log.Println("\tImporting Service Servers")
	var resources []terraformutils.Resource
	ctx := context.Background()
	AnalysisClient := analysisservices.NewServersClient(g.Args["config"].(authentication.Config).SubscriptionID)
	AnalysisClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var (
		servers analysisservices.Servers
		err     error
	)

	if rg := g.Args["resource_group"].(string); rg != "" {
		servers, err = AnalysisClient.ListByResourceGroup(ctx, rg)
	} else {
		servers, err = AnalysisClient.List(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, svr := range *servers.Value {
		resources = append(resources, terraformutils.NewSimpleResource(
			*svr.ID,
			*svr.Name,
			"azurerm_analysis_services_server",
			g.ProviderName,
			[]string{}))
	}

	return resources, nil
}

func (g *AnalysisGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listServiceServers,
	}

	for _, f := range functions {
		resources, err := f()
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}
