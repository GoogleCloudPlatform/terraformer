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

package azuredevops

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/microsoft/azure-devops-go-api/azuredevops/graph"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AzureDevOpsServiceGenerator interface {
	terraformutils.ServiceGenerator
	GetResourceConnections() map[string][]string
}

type AzureDevOpsService struct { //nolint
	terraformutils.Service
}

func (az *AzureDevOpsService) GetResourceConnections() map[string][]string {
	return nil
}

func (az *AzureDevOpsService) getConnection() *azuredevops.Connection {

	organizationURL := az.Args["organizationURL"].(string)
	personalAccessToken := az.Args["personalAccessToken"].(string)
	return azuredevops.NewPatConnection(organizationURL, personalAccessToken)
}

func (az *AzureDevOpsService) getCoreClient() (core.Client, error) {
	ctx := context.Background()
	client, err := core.NewClient(ctx, az.getConnection())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, nil
}

func (az *AzureDevOpsService) getGraphClient() (graph.Client, error) {
	ctx := context.Background()
	client, err := graph.NewClient(ctx, az.getConnection())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, nil
}

func (az *AzureDevOpsService) getGitClient() (git.Client, error) {
	ctx := context.Background()
	client, err := git.NewClient(ctx, az.getConnection())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, nil
}

func (az *AzureDevOpsService) appendSimpleResource(id string, resourceName string, resourceType string) {
	newResource := terraformutils.NewSimpleResource(id, resourceName, resourceType, az.ProviderName, []string{})
	az.Resources = append(az.Resources, newResource)
}
