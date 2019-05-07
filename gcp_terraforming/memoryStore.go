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

package gcp_terraforming

import (
	"context"
	"log"
	"strings"

	"google.golang.org/api/redis/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"golang.org/x/oauth2/google"
)

var redisAllowEmptyValues = []string{""}

var redisAdditionalFields = map[string]string{}

type MemoryStoreGenerator struct {
	GCPService
}

// Run on redisInstancesList and create for each TerraformResource
func (g MemoryStoreGenerator) createResources(redisInstancesList *redis.ProjectsLocationsInstancesListCall, ctx context.Context) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := redisInstancesList.Pages(ctx, func(page *redis.ListInstancesResponse) error {
		for _, obj := range page.Instances {
			t := strings.Split(obj.Name, "/")
			name := t[len(t)-1]
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				name,
				"google_redis_instance",
				"google",
				map[string]string{
					"name":    name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
				},
				redisAllowEmptyValues,
				redisAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each redis create 1 TerraformResource
// Need Redis name as ID for terraform resource
func (g *MemoryStoreGenerator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, redis.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	redisService, err := redis.New(c)
	if err != nil {
		log.Fatal(err)
	}

	redisInstancesList := redisService.Projects.Locations.Instances.List("projects/" + g.GetArgs()["project"] + "/locations/" + g.GetArgs()["region"])

	g.Resources = g.createResources(redisInstancesList, ctx)
	g.PopulateIgnoreKeys()
	return nil
}
