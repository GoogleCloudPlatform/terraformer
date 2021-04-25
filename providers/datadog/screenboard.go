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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// ScreenboardAllowEmptyValues ...
	ScreenboardAllowEmptyValues = []string{"tags."}
)

// ScreenboardGenerator ...
type ScreenboardGenerator struct {
	DatadogService
}

func (ScreenboardGenerator) createResources(screenboards []*datadog.ScreenboardLite) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, screenboard := range screenboards {
		resourceName := strconv.Itoa(screenboard.GetId())
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			fmt.Sprintf("screenboard_%s", resourceName),
			"datadog_screenboard",
			"datadog",
			ScreenboardAllowEmptyValues,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each screenboard create 1 TerraformResource.
// Need Screenboard ID as ID for terraform resource
func (g *ScreenboardGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	screenboards, err := client.GetScreenboards()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(screenboards)
	return nil
}
