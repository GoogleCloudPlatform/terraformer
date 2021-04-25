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

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/redis/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var redisAllowEmptyValues = []string{""}

var redisAdditionalFields = map[string]interface{}{}

type MemoryStoreGenerator struct {
	GCPService
}

// Run on redisInstancesList and create for each TerraformResource
func (g MemoryStoreGenerator) createResources(ctx context.Context, redisInstancesList *redis.ProjectsLocationsInstancesListCall) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := redisInstancesList.Pages(ctx, func(page *redis.ListInstancesResponse) error {
		for _, obj := range page.Instances {
			t := strings.Split(obj.Name, "/")
			name := t[len(t)-1]
			resources = append(resources, terraformutils.NewResource(
				obj.Name,
				name,
				"google_redis_instance",
				g.ProviderName,
				map[string]string{
					"name":    name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				redisAllowEmptyValues,
				redisAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each redis create 1 TerraformResource
// Need Redis name as ID for terraform resource
func (g *MemoryStoreGenerator) InitResources() error {
	ctx := context.Background()
	redisService, err := redis.NewService(ctx)
	if err != nil {
		return err
	}

	redisInstancesList := redisService.Projects.Locations.Instances.List("projects/" + g.GetArgs()["project"].(string) + "/locations/" + g.GetArgs()["region"].(compute.Region).Name)

	g.Resources = g.createResources(ctx, redisInstancesList)
	return nil
}
