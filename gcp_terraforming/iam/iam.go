package iam

import (
	"context"
	"log"
	"os"

	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"

	"cloud.google.com/go/iam/admin/apiv1"
	"google.golang.org/api/iterator"
)

var ignoreKey = map[string]bool{
	"url":       true,
	"id":        true,
	"self_link": true,
	"unique_id": true,
	"email":     true,
	"name":      true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

var additionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type IamGenerator struct {
	gcp_generator.BasicGenerator
}

func (IamGenerator) createResources(serviceAccountsIterator *admin.ServiceAccountIterator) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for {
		serviceAccount, err := serviceAccountsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error with service account:", err)
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			serviceAccount.Name,
			serviceAccount.UniqueId,
			"google_service_account",
			"google",
			nil,
			map[string]string{},
		))
	}
	return resources
}

func (g IamGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := admin.NewIamClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	serviceAccountsIterator := client.ListServiceAccounts(ctx, &adminpb.ListServiceAccountsRequest{Name: "projects/" + projectID})

	resources := g.createResources(serviceAccountsIterator)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, additionalFields)
	return resources, metadata, nil

}
