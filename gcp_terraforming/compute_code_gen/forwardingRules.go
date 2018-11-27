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

var forwardingRulesIgnoreKey = map[string]bool{
	"id":                 true,
	"self_link":          true,
	"fingerprint":        true,
	"label_fingerprint":  true,
	"creation_timestamp": true,

	"service_name": true,
}

var forwardingRulesAllowEmptyValues = map[string]bool{}

var forwardingRulesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type ForwardingRulesGenerator struct {
	gcp_generator.BasicGenerator
}

func (ForwardingRulesGenerator) createResources(forwardingRulesList *compute.ForwardingRulesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := forwardingRulesList.Pages(ctx, func(page *compute.ForwardingRuleList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_forwarding_rule",
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

func (g ForwardingRulesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	forwardingRulesList := computeService.ForwardingRules.List(project, region)

	resources := g.createResources(forwardingRulesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, forwardingRulesIgnoreKey, forwardingRulesAllowEmptyValues, forwardingRulesAdditionalFields)
	return resources, metadata, nil

}
