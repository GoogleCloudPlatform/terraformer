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
	"log"
	"strings"

	cloudscheduler "google.golang.org/api/cloudscheduler/v1beta1"
	"google.golang.org/api/compute/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var schedulerJobsAllowEmptyValues = []string{""}

var schedulerJobsAdditionalFields = map[string]interface{}{}

type SchedulerJobsGenerator struct {
	GCPService
}

// Run on SchedulerJobsList and create for each TerraformResource
func (g SchedulerJobsGenerator) createResources(ctx context.Context, jobsList *cloudscheduler.ProjectsLocationsJobsListCall) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := jobsList.Pages(ctx, func(page *cloudscheduler.ListJobsResponse) error {
		for _, obj := range page.Jobs {
			t := strings.Split(obj.Name, "/")
			name := t[len(t)-1]
			resources = append(resources, terraformutils.NewResource(
				obj.Name,
				name,
				"google_cloud_scheduler_job",
				g.ProviderName,
				map[string]string{
					"name":    name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				schedulerJobsAllowEmptyValues,
				schedulerJobsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
func (g *SchedulerJobsGenerator) InitResources() error {
	ctx := context.Background()
	cloudSchedulerService, err := cloudscheduler.NewService(ctx)
	if err != nil {
		return err
	}

	jobsList := cloudSchedulerService.Projects.Locations.Jobs.List("projects/" + g.GetArgs()["project"].(string) + "/locations/" + g.GetArgs()["region"].(compute.Region).Name)

	g.Resources = g.createResources(ctx, jobsList)
	return nil
}
