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

package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

// AWSIntegrationGenerator ...
type AWSIntegrationGenerator struct {
	MackerelService
}

func (g *AWSIntegrationGenerator) createResources(awsIntegrations []*mackerel.AWSIntegration) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, awsIntegration := range awsIntegrations {
		resources = append(resources, g.createResource(awsIntegration.ID))
	}
	return resources
}

func (g *AWSIntegrationGenerator) createResource(awsIntegrationID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		awsIntegrationID,
		fmt.Sprintf("aws_integration_%s", awsIntegrationID),
		"mackerel_aws_integration",
		"mackerel",
		[]string{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each aws integration create 1 TerraformResource.
// Need AWS Integration ID as ID for terraform resource
func (g *AWSIntegrationGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)
	awsIntegrations, err := client.FindAWSIntegrations()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(awsIntegrations)...)
	return nil
}
