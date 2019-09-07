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
	"strconv"

	datadog "github.com/zorkian/go-datadog-api"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

var (
	// DowntimeAllowEmptyValues ...
	DowntimeAllowEmptyValues = []string{}
	// DowntimeAttributes ...
	DowntimeAttributes = map[string]string{}
	// DowntimeAdditionalFields ...
	DowntimeAdditionalFields = map[string]string{}
)

// DowntimeGenerator ...
type DowntimeGenerator struct {
	DatadogService
}

func (DowntimeGenerator) createResources(downtimes []datadog.Downtime) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, downtime := range downtimes {
		resourceName := strconv.Itoa(downtime.GetId())
		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			fmt.Sprintf("downtime_%s", resourceName),
			"datadog_downtime",
			"datadog",
			DowntimeAttributes,
			DowntimeAllowEmptyValues,
			DowntimeAdditionalFields,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each downtime create 1 TerraformResource.
// Need Downtime ID as ID for terraform resource
func (g *DowntimeGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	downtimes, err := client.GetDowntimes()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(downtimes)
	g.PopulateIgnoreKeys()
	return nil
}
