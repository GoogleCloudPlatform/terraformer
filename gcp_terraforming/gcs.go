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
	"fmt"
	"log"
	"waze/terraformer/terraform_utils"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var GcsAllowEmptyValues = []string{"labels.", "created_before"}

var GcsAdditionalFields = map[string]string{}

type GcsGenerator struct {
	GCPService
}

func (g *GcsGenerator) createResources(bucketIterator *storage.BucketIterator) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for {
		battrs, err := bucketIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with bucket:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket",
			"google",
			map[string]string{
				"name":          battrs.Name,
				"force_destroy": "false",
			},
			GcsAllowEmptyValues,
			GcsAdditionalFields,
		))
		resources = append(resources, terraform_utils.NewResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_acl",
			"google",
			map[string]string{
				"bucket": battrs.Name,
			},
			GcsAllowEmptyValues,
			GcsAdditionalFields,
		))
		resources = append(resources, terraform_utils.NewResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_binding",
			"google",
			map[string]string{
				"bucket": battrs.Name,
			},
			GcsAllowEmptyValues,
			GcsAdditionalFields,
		))
		resources = append(resources, terraform_utils.NewResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_member",
			"google",
			map[string]string{
				"bucket": battrs.Name,
			},
			GcsAllowEmptyValues,
			GcsAdditionalFields,
		))
		resources = append(resources, terraform_utils.NewResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_policy",
			"google",
			map[string]string{
				"bucket": battrs.Name,
			},
			GcsAllowEmptyValues,
			GcsAdditionalFields,
		))
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each bucket  create 1 TerraformResource
// Need bucket name as ID for terraform resource
func (g *GcsGenerator) InitResources() error {
	ctx := context.Background()

	projectID := g.GetArgs()["project"]
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Print(err)
		return err
	}
	bucketIterator := client.Buckets(ctx, projectID)

	g.Resources = g.createResources(bucketIterator)
	g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook for add bucket policy json as heredoc
// support only bucket with policy
func (g *GcsGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type != "google_storage_bucket_iam_policy" {
			continue
		}
		policy := resource.Item["policy_data"].(string)
		g.Resources[i].Item["policy_data"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
	}
	return nil
}
