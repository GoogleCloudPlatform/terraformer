package gcs

import (
	"context"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var ignoreKey = map[string]bool{
	"url":       true,
	"id":        true,
	"self_link": true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

var additionalFields = map[string]string{
	"force_destroy": "false",
	"project":       "waze-development",
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

func (g GcsGenerator) Generate(region string) error {
	ctx := context.Background()

	projectID := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := storage.NewClient(ctx)
	bucketIterator := client.Buckets(ctx, projectID)

	resources := g.createResources(bucketIterator)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "buckets", region, "google")
	if err != nil {
		return err
	}
	return nil

}
