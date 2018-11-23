// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"strings"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var instanceGroupManagersIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,

	"instance_group": true,
}

var instanceGroupManagersAllowEmptyValues = map[string]bool{}

var instanceGroupManagersAdditionalFields = map[string]string{
	"project": "waze-development",
}

type InstanceGroupManagersGenerator struct {
	gcp_generator.BasicGenerator
}

func (InstanceGroupManagersGenerator) createResources(instanceGroupManagersList *compute.InstanceGroupManagersListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := instanceGroupManagersList.Pages(ctx, func(page *compute.InstanceGroupManagerList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				zone+"/"+obj.Name,
				obj.Name,
				"google_compute_instance_group_manager",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": "waze-development",
					"region":  region,
					"zone":    zone,
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g InstanceGroupManagersGenerator) Generate(zone string) error {
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

	instanceGroupManagersList := computeService.InstanceGroupManagers.List(project, zone)

	resources := g.createResources(instanceGroupManagersList, ctx, region, zone)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, instanceGroupManagersIgnoreKey, instanceGroupManagersAllowEmptyValues, instanceGroupManagersAdditionalFields)
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "instanceGroupManagers", region, "google")
	if err != nil {
		return err
	}
	return nil

}
