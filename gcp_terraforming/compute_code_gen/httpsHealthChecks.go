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

var httpsHealthChecksIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,
}

var httpsHealthChecksAllowEmptyValues = map[string]bool{}

var httpsHealthChecksAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type HttpsHealthChecksGenerator struct {
	gcp_generator.BasicGenerator
}

func (HttpsHealthChecksGenerator) createResources(httpsHealthChecksList *compute.HttpsHealthChecksListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := httpsHealthChecksList.Pages(ctx, func(page *compute.HttpsHealthCheckList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_https_health_check",
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

func (g HttpsHealthChecksGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	httpsHealthChecksList := computeService.HttpsHealthChecks.List(project)

	resources := g.createResources(httpsHealthChecksList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, httpsHealthChecksIgnoreKey, httpsHealthChecksAllowEmptyValues, httpsHealthChecksAdditionalFields)
	return resources, metadata, nil

}
