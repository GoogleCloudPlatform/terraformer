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
	"regexp"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SyntheticsPrivateLocationAllowEmptyValues ...
	SyntheticsPrivateLocationAllowEmptyValues = []string{"tags."}
	plIDRegex                                 = regexp.MustCompile("^pl:.*")
)

// SyntheticsPrivateLocationGenerator ...
type SyntheticsPrivateLocationGenerator struct {
	DatadogService
}

func (g *SyntheticsPrivateLocationGenerator) createResources(locations []datadogV1.SyntheticsLocation) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, location := range locations {
		locationID := location.GetId()
		if plIDRegex.MatchString(locationID) {
			resources = append(resources, g.createResource(locationID))
		}
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
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewSyntheticsApi(datadogClient)

	data, _, err := api.ListLocations(auth)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(data.GetLocations())
	return nil
}
