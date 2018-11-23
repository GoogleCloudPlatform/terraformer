package iam

import (
	"context"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"

	admin "cloud.google.com/go/iam/admin/apiv1"
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

type IAMGenerator struct {
	gcp_generator.BasicGenerator
}

func (IAMGenerator) createResources(serviceAccountsIterator *admin.ServiceAccountIterator) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for {
		serviceAccount, err := serviceAccountsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with bucket:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			serviceAccount.Name,
			serviceAccount.Name,
			"google_storage_bucket",
			"google",
			nil,
			map[string]string{
				"name": serviceAccount.Name,
			},
		))
	}
	return resources
}

func (g IAMGenerator) Generate(region string) error {
	ctx := context.Background()

	//projectID := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := admin.NewIamClient(ctx)
	serviceAccountsIterator := client.ListServiceAccounts(ctx, &adminpb.ListServiceAccountsRequest{})

	resources := g.createResources(serviceAccountsIterator)
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
