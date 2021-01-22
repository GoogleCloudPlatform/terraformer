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
	"log"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SyntheticsPrivateLocationAllowEmptyValues ...
	SyntheticsPrivateLocationAllowEmptyValues = []string{"tags."}
)

// SyntheticsPrivateLocationGenerator ...
type SyntheticsPrivateLocationGenerator struct {
	DatadogService
}

func (g *SyntheticsPrivateLocationGenerator) createResources(privateLocations []datadogV1.SyntheticsPrivateLocation) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, privateLocation := range privateLocations {
		resourceName := privateLocation.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *SyntheticsPrivateLocationGenerator) createResource(plID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		plID,
		fmt.Sprintf("synthetics_private_location_%s", plID),
		"datadog_synthetics_private_location",
		"datadog",
		SyntheticsPrivateLocationAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each SyntheticsPrivateLocation create 1 TerraformResource.
// Need SyntheticsPrivateLocation ID as ID for terraform resource
func (g *SyntheticsPrivateLocationGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	var privateLocations []datadogV1.SyntheticsPrivateLocation
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("synthetics_private_location") {
			for _, v := range filter.AcceptableValues {
				resp, _, err := datadogClientV1.SyntheticsApi.GetPrivateLocation(authV1, v).Execute()
				if err != nil {
					log.Printf("error retrieving synthetics private location with id:%s - %s", v, err)
					continue
				}
				privateLocations = append(privateLocations, resp)
			}
		}
	}

	if len(privateLocations) == 0 {
		log.Print("Filter(Synthetics Private Location IDs) is required for importing datadog_synthetics_private_location resource")
		return nil
	}

	g.Resources = g.createResources(privateLocations)
	return nil
}
