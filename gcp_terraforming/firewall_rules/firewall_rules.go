package firewall_rules

import (
	"context"
	"log"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/compute/v1"
)

var ignoreKey = map[string]bool{
	//"url":                true,
	"id":                 true,
	"self_link":          true,
	"creation_timestamp": true,
}

var allowEmptyValues = map[string]bool{}

var additionalFields = map[string]string{
	"project": "waze-development",
}

type FirewallRulesGenerator struct {
	gcp_generator.BasicGenerator
}

func (FirewallRulesGenerator) createResources(firewallsList *compute.FirewallsListCall, ctx context.Context) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := firewallsList.Pages(ctx, func(page *compute.FirewallList) error {
		for _, rule := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				rule.Name,
				rule.Name,
				"google_compute_firewall",
				"google",
				nil,
				map[string]string{"name": rule.Name, "project": "waze-development"},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g FirewallRulesGenerator) Generate(region string) error {
	projectID := "waze-development" //os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	firewallsList := computeService.Firewalls.List(projectID)

	resources := g.createResources(firewallsList, ctx)
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
	err = terraform_utils.GenerateTf(resources, "firewall_rules", region, "google")
	if err != nil {
		return err
	}
	return nil

}
