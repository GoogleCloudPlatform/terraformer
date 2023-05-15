// Copyright 2018 The Terraformer Authors.
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

package datadog

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SyntheticsAllowEmptyValues ...
	SyntheticsAllowEmptyValues = []string{"tags."}
)

// SyntheticsTestGenerator ...
type SyntheticsTestGenerator struct {
	DatadogService
}

func (g *SyntheticsTestGenerator) createResources(syntheticsList []datadogV1.SyntheticsTestDetails) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, synthetics := range syntheticsList {
		resourceName := synthetics.GetPublicId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *SyntheticsTestGenerator) createResource(syntheticsID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		syntheticsID,
		fmt.Sprintf("synthetics_%s", syntheticsID),
		"datadog_synthetics_test",
		"datadog",
		SyntheticsAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each synthetics create 1 TerraformResource.
// Need Synthetics ID as ID for terraform resource
func (g *SyntheticsTestGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewSyntheticsApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("synthetics_test") {
			for _, value := range filter.AcceptableValues {
				syntheticsTest, _, err := api.GetTest(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(syntheticsTest.GetPublicId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	syntheticsTests, _, err := api.ListTests(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(syntheticsTests.GetTests())
	return nil
}
