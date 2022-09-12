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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SyntheticsGlobalVariableAllowEmptyValues ...
	SyntheticsGlobalVariableAllowEmptyValues = []string{"tags."}
)

// SyntheticsGlobalVariableGenerator ...
type SyntheticsGlobalVariableGenerator struct {
	DatadogService
}

func (g *SyntheticsGlobalVariableGenerator) createResources(globalVariables []datadogV1.SyntheticsGlobalVariable) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, globalVariable := range globalVariables {
		resourceID := globalVariable.GetId()
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *SyntheticsGlobalVariableGenerator) createResource(globalVariableID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		globalVariableID,
		fmt.Sprintf("synthetics_global_variable_%s", globalVariableID),
		"datadog_synthetics_global_variable",
		"datadog",
		SyntheticsGlobalVariableAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each SyntheticsGlobalVariable create 1 TerraformResource.
// Need SyntheticsGlobalVariable ID as ID for terraform resource
func (g *SyntheticsGlobalVariableGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewSyntheticsApi(datadogClient)

	var globalVariableIDs []datadogV1.SyntheticsGlobalVariable
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("synthetics_global_variable") {
			for _, v := range filter.AcceptableValues {
				resp, _, err := api.GetGlobalVariable(auth, v)
				if err != nil {
					log.Printf("error retrieving synthetics gloval variable with id:%s - %s", v, err)
					continue
				}
				globalVariableIDs = append(globalVariableIDs, resp)
			}
		}
	}

	if len(globalVariableIDs) == 0 {
		log.Print("Filter(Synthetics Global Variable IDs) is required for importing datadog_synthetics_global_variable resource")
		return nil
	}

	g.Resources = g.createResources(globalVariableIDs)
	return nil
}
