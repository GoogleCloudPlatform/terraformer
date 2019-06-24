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
	"fmt"

	datadog "github.com/zorkian/go-datadog-api"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

var (
	// SyntheticsAllowEmptyValues ...
	SyntheticsAllowEmptyValues = []string{"tags."}
	// SyntheticsAttributes ...
	SyntheticsAttributes = map[string]string{}
	// SyntheticsAdditionalFields ...
	SyntheticsAdditionalFields = map[string]string{}
)

// SyntheticsGenerator ...
type SyntheticsGenerator struct {
	DatadogService
}

func (SyntheticsGenerator) createResources(syntheticsList []datadog.SyntheticsTest) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, synthetics := range syntheticsList {
		resourceName := synthetics.GetPublicId()
		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			fmt.Sprintf("synthetics_%s", resourceName),
			"datadog_synthetics_test",
			"datadog",
			SyntheticsAttributes,
			SyntheticsAllowEmptyValues,
			SyntheticsAdditionalFields,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each synthetics create 1 TerraformResource.
// Need Synthetics ID as ID for terraform resource
func (g *SyntheticsGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	syntheticsList, err := client.GetSyntheticsTests()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(syntheticsList)
	g.PopulateIgnoreKeys()
	return nil
}
