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
	// DowntimeAllowEmptyValues ...
	DowntimeAllowEmptyValues = []string{}
)

// DowntimeGenerator ...
type DowntimeGenerator struct {
	DatadogService
}

func (g *DowntimeGenerator) createResources(downtimes []datadogV2.DowntimeResponseData) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, downtime := range downtimes {
		resourceName := downtime.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *DowntimeGenerator) createResource(downtimeID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		downtimeID,
		fmt.Sprintf("downtime_schedule_%s", downtimeID),
		"datadog_downtime_schedule",
		"datadog",
		DowntimeAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each downtime create 1 TerraformResource.
// Need Downtime ID as ID for terraform resource
func (g *DowntimeGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV2.NewDowntimesApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("downtime") {
			for _, value := range filter.AcceptableValues {
				downtime, _, err := api.GetDowntime(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(downtime.Data.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	var downtimes []datadogV2.DowntimeResponseData
	optionalParameters := *datadogV2.NewListDowntimesOptionalParameters()
	downtimesChan, _ := api.ListDowntimesWithPagination(auth,
		*optionalParameters.WithPageLimit(1000))

	for {
		pageResult, more := <-downtimesChan
		if !more {
			break
		}
		if pageResult.Error != nil {
			return pageResult.Error
		}
		downtimes = append(downtimes, pageResult.Item)
	}

	g.Resources = g.createResources(downtimes)
	return nil
}
