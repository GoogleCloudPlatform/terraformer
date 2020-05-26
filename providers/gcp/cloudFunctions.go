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

package gcp

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/compute/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var cloudFunctionsAllowEmptyValues = []string{""}

type CloudFunctionsGenerator struct {
	GCPService
}

// Run on CloudFunctionsList and create for each TerraformResource
func (g CloudFunctionsGenerator) createResources(ctx context.Context, functionsList *cloudfunctions.ProjectsLocationsFunctionsListCall) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := functionsList.Pages(ctx, func(page *cloudfunctions.ListFunctionsResponse) error {
		for _, functions := range page.Functions {
			t := strings.Split(functions.Name, "/")
			name := t[len(t)-1]
			resources = append(resources, terraformutils.NewSimpleResource(
				g.GetArgs()["project"].(string)+"/"+g.GetArgs()["region"].(compute.Region).Name+"/"+name,
				g.GetArgs()["region"].(compute.Region).Name+"_"+name,
				"google_cloudfunctions_function",
				"google",
				cloudFunctionsAllowEmptyValues,
			))
		}
		return nil
	}); err != nil {
		if strings.Contains(err.Error(), "403: Permission denied") {
			log.Print(fmt.Sprint("Got PermissionDenied for region %", g.GetArgs()["region"].(compute.Region).Name))
		} else {
			log.Fatal(err)
		}
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each CloudFunctions create 1 TerraformResource
// Need CloudFunctions name as ID for terraform resource
func (g *CloudFunctionsGenerator) InitResources() error {
	ctx := context.Background()
	cloudfunctionsService, err := cloudfunctions.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	functionsList := cloudfunctionsService.Projects.Locations.Functions.List("projects/" + g.GetArgs()["project"].(string) + "/locations/" + g.GetArgs()["region"].(compute.Region).Name)

	g.Resources = g.createResources(ctx, functionsList)
	return nil
}
