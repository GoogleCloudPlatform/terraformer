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

package alerts

import (
	"context"
	"log"
	"os"

	"waze/terraformer/gcp_terraforming/gcp_generator"
	"waze/terraformer/terraform_utils"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/monitoring/apiv3"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

var alertsAllowEmptyValues = map[string]bool{}

var alertsAdditionalFields = map[string]string{}

type AlertsGenerator struct {
	gcp_generator.BasicGenerator
}

func (AlertsGenerator) createResources(alertIterator *monitoring.AlertPolicyIterator) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for {
		alert, err := alertIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with alert:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			alert.Name,
			alert.Name,
			"google_monitoring_alert_policy",
			"google",
			nil,
			map[string]string{
				"name": alert.Name,
			},
		))
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each alert  create 1 TerraformResource
// Need alert name as ID for terraform resource
func (g AlertsGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()
	req := &monitoringpb.ListAlertPoliciesRequest{
		Name: "projects/" + project,
	}

	client, err := monitoring.NewAlertPolicyClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	alertIterator := client.ListAlertPolicies(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	resources := g.createResources(alertIterator)
	metadata := terraform_utils.NewResourcesMetaData(resources, g.IgnoreKeys(resources, "google"), alertsAllowEmptyValues, alertsAdditionalFields)
	return resources, metadata, nil

}
