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

package azuredevpos

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AzureDevOpsService struct { //nolint
	terraformutils.Service
	connection *azuredevops.Connection
}

func (p *AzureDevOpsService) getCoreClient() (core.Client, error) {
	ctx := context.Background()
	client, err := core.NewClient(ctx, p.connection)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func (p *AzureDevOpsService) AppendSimpleResource(id string, resourceName string, resourceType string) {
	newResource := terraformutils.NewSimpleResource(id, resourceName, resourceType, p.ProviderName, []string{})
	p.Resources = append(p.Resources, newResource)
}
