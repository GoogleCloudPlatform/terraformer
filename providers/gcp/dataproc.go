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

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/dataproc/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

var dataprocAllowEmptyValues = []string{""}

var dataprocAdditionalFields = map[string]string{}

type DataprocGenerator struct {
	GCPService
}

// Run on DataprocClusterList and create for each TerraformResource
func (g DataprocGenerator) createClusterResources(clusterList *dataproc.ProjectsRegionsClustersListCall, ctx context.Context) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := clusterList.Pages(ctx, func(page *dataproc.ListClustersResponse) error {
		for _, cluster := range page.Clusters {
			resource := terraform_utils.NewResource(
				cluster.ClusterName,
				cluster.ClusterName,
				"google_dataproc_cluster",
				"google",
				map[string]string{
					"name":    cluster.ClusterName,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				dataprocAllowEmptyValues,
				dataprocAdditionalFields,
			)
			resource.IgnoreKeys = append(resource.IgnoreKeys, "^cluster_config.[0-9].delete_autogen_bucket$")
			resources = append(resources, resource)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

/*
// Run on DataprocJobList and create for each TerraformResource
func (g DataprocGenerator) createJobResources(jobList *dataproc.ProjectsRegionsJobsListCall, ctx context.Context) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := jobList.Pages(ctx, func(page *dataproc.ListJobsResponse) error {
		for _, job := range page.Jobs {
			resources = append(resources, terraform_utils.NewResource(
				job.Reference.JobId,
				job.Reference.JobId,
				"google_dataproc_job",
				"google",
				map[string]string{
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				dataprocAllowEmptyValues,
				dataprocAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}
*/

// Generate TerraformResources from GCP API,
// from each DataprocGenerator create 1 TerraformResource
// Need DataprocGenerator name as ID for terraform resource
func (g *DataprocGenerator) InitResources() error {
	ctx := context.Background()
	dataprocService, err := dataproc.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	clusterList := dataprocService.Projects.Regions.Clusters.List(g.GetArgs()["project"].(string), g.GetArgs()["region"].(compute.Region).Name)
	g.Resources = g.createClusterResources(clusterList, ctx)

	//jobList := dataprocService.Projects.Regions.Jobs.List(g.GetArgs()["project"].(string), g.GetArgs()["region"])
	//g.Resources = append(g.Resources, g.createJobResources(jobList, ctx)...)

	g.PopulateIgnoreKeys()
	return nil

}

func (g *DataprocGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		// hack for provider bug
		if v, exist := r.Item["cluster_config"].([]interface{})[0].(map[string]interface{})["preemptible_worker_config"]; exist {
			if v.([]interface{})[0].(map[string]interface{})["num_instances"].(string) == "0" && len(v.([]interface{})[0].(map[string]interface{})["disk_config"].([]interface{})) == 0 {
				delete(g.Resources[i].Item["cluster_config"].([]interface{})[0].(map[string]interface{}), "preemptible_worker_config")
			}
		}
	}
	return nil
}
