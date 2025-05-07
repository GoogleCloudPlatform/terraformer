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
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// PowerpackAllowEmptyValues ...
	PowerpackAllowEmptyValues = []string{}
)

// PowerpackGenerator ...
type PowerpackGenerator struct {
	DatadogService
}

func (g *PowerpackGenerator) createResources(powerpacks []datadogV2.PowerpackData) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, powerpack := range powerpacks {
		resourceName := powerpack.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *PowerpackGenerator) createResource(powerpackName string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		powerpackName,
		fmt.Sprintf("powerpack_%s", powerpackName),
		"datadog_powerpack",
		"datadog",
		PowerpackAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each powerpack create 1 TerraformResource.
// Need Powerpack Name as ID for terraform resource
func (g *PowerpackGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV2.NewPowerpackApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("powerpack") {
			for _, value := range filter.AcceptableValues {
				powerpack, _, err := api.GetPowerpack(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(powerpack.Data.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	var powerpacks []datadogV2.PowerpackData
	optionalParameters := &datadogV2.ListPowerpacksOptionalParameters{}
	paginationChan, _ := api.ListPowerpacksWithPagination(auth,
		*optionalParameters.WithPageLimit(1000))
	for {
		pageResult, more := <-paginationChan
		if !more {
			break
		}
		if pageResult.Error != nil {
			return pageResult.Error
		}
		powerpacks = append(powerpacks, pageResult.Item)
	}

	g.Resources = g.createResources(powerpacks)
	return nil
}
