package gcs

import (
	"context"
	"fmt"
	"log"
	"os"

	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var ignoreKey = map[string]bool{
	"^url$":       true,
	"^id$":        true,
	"^self_link$": true,
	"^etag$":      true,
}

var allowEmptyValues = map[string]bool{
	"tags.":          true,
	"storage_class":  true,
	"created_before": true,
}

var additionalFields = map[string]string{}

type GcsGenerator struct {
	gcp_generator.BasicGenerator
}

func (GcsGenerator) createResources(bucketIterator *storage.BucketIterator) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for {
		battrs, err := bucketIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with bucket:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket",
			"google",
			nil,
			map[string]string{
				"name": battrs.Name,
			},
		))
		resources = append(resources, terraform_utils.NewTerraformResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_acl",
			"google",
			nil,
			map[string]string{
				"bucket": battrs.Name,
			},
		))
		resources = append(resources, terraform_utils.NewTerraformResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_binding",
			"google",
			nil,
			map[string]string{
				"bucket": battrs.Name,
			},
		))
		resources = append(resources, terraform_utils.NewTerraformResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_member",
			"google",
			nil,
			map[string]string{
				"bucket": battrs.Name,
			},
		))
		resources = append(resources, terraform_utils.NewTerraformResource(
			battrs.Name,
			battrs.Name,
			"google_storage_bucket_iam_policy",
			"google",
			nil,
			map[string]string{
				"bucket": battrs.Name,
			},
		))
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each bucket  create 1 TerraformResource
// Need bucket name as ID for terraform resource
func (g GcsGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Print(err)
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	bucketIterator := client.Buckets(ctx, projectID)

	resources := g.createResources(bucketIterator)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	return resources, metadata, nil
}

// PostGenerateHook for add bucket policy json as heredoc
// support only bucket with policy
func (GcsGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for _, resource := range resources {
		if resource.ResourceType != "google_storage_bucket_iam_policy" {
			continue
		}
		policy := resource.Item.(interface{}).(map[string]interface{})["policy_data"].(string)
		resource.Item.(interface{}).(map[string]interface{})["policy_data"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
	}
	return resources, nil
}
