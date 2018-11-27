// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"
)

var instanceTemplatesIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,

	"tags_fingerprint": true,
}

var instanceTemplatesAllowEmptyValues = map[string]bool{}

var instanceTemplatesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type InstanceTemplatesGenerator struct {
	gcp_generator.BasicGenerator
}

func (InstanceTemplatesGenerator) createResources(instanceTemplatesList *compute.InstanceTemplatesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := instanceTemplatesList.Pages(ctx, func(page *compute.InstanceTemplateList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_instance_template",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": "waze-development",
					"region":  region,
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g InstanceTemplatesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	instanceTemplatesList := computeService.InstanceTemplates.List(project)

	resources := g.createResources(instanceTemplatesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, instanceTemplatesIgnoreKey, instanceTemplatesAllowEmptyValues, instanceTemplatesAdditionalFields)
	return resources, metadata, nil

}
