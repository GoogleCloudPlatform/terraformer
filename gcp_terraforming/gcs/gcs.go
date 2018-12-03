package gcs

import (
	"context"
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
}

var allowEmptyValues = map[string]bool{
	"tags.":          true,
	"storage_class":  true,
	"created_before": true,
}

var additionalFields = map[string]string{
	"force_destroy": "false",
	"project":       os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

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
