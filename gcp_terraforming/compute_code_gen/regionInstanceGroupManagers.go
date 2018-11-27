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

var regionInstanceGroupManagersIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,

	"instance_group": true,
}

var regionInstanceGroupManagersAllowEmptyValues = map[string]bool{

	"name": true,

	"health_check": true,
}

var regionInstanceGroupManagersAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type RegionInstanceGroupManagersGenerator struct {
	gcp_generator.BasicGenerator
}

func (RegionInstanceGroupManagersGenerator) createResources(regionInstanceGroupManagersList *compute.RegionInstanceGroupManagersListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := regionInstanceGroupManagersList.Pages(ctx, func(page *compute.RegionInstanceGroupManagerList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_region_instance_group_manager",
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

func (g RegionInstanceGroupManagersGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	regionInstanceGroupManagersList := computeService.RegionInstanceGroupManagers.List(project, region)

	resources := g.createResources(regionInstanceGroupManagersList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, regionInstanceGroupManagersIgnoreKey, regionInstanceGroupManagersAllowEmptyValues, regionInstanceGroupManagersAdditionalFields)
	return resources, metadata, nil

}
