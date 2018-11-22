
/*
AUTO GENERATE CODE - NOT FOR EDIT

*/

package regionBackendServices

import (
	"context"
	"strings"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var ignoreKey = map[string]bool{
	//"url":                true,
	"id":                 	true,
	"self_link":          	true,
	"fingerprint": 			true,
	"label_fingerprint": 	true,
	"creation_timestamp": 	true,
	
}

var allowEmptyValues = map[string]bool{}

var additionalFields = map[string]string{
	"project": "waze-development",
}

type RegionBackendServicesGenerator struct {
	gcp_generator.BasicGenerator
}

func (RegionBackendServicesGenerator) createResources(RegionBackendServicesList *compute.RegionBackendServicesListCall, ctx context.Context, region string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := RegionBackendServicesList.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_region_backend_service",
				"google",
				nil,
				map[string]string{"name": obj.Name, "project": "waze-development", "region": region},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g RegionBackendServicesGenerator) Generate(zone string) error {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	RegionBackendServicesList := computeService.RegionBackendServices.List(project, region)

	resources := g.createResources(RegionBackendServicesList, ctx, region)
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
	err = terraform_utils.GenerateTf(resources, "RegionBackendServices", region, "google")
	if err != nil {
		return err
	}
	return nil

}

