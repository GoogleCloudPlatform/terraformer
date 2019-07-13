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
	// MonitorAllowEmptyValues ...
	MonitorAllowEmptyValues = []string{"tags."}
	// MonitorAttributes ...
	MonitorAttributes = map[string]string{}
	// MonitorAdditionalFields ...
	MonitorAdditionalFields = map[string]string{}
)

// MonitorGenerator ...
type MonitorGenerator struct {
	DatadogService
}

func (MonitorGenerator) createResources(monitors []datadog.Monitor) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, monitor := range monitors {
		resourceName := strconv.Itoa(monitor.GetId())
		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			fmt.Sprintf("monitor_%s", resourceName),
			"datadog_monitor",
			"datadog",
			MonitorAttributes,
			MonitorAllowEmptyValues,
			MonitorAdditionalFields,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each monitor create 1 TerraformResource.
// Need Monitor ID as ID for terraform resource
func (g *MonitorGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	monitors, err := client.GetMonitors()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(monitors)
	g.PopulateIgnoreKeys()
	return nil
}
