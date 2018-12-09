// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"waze/terraformer/gcp_terraforming/gcp_generator"
	"waze/terraformer/terraform_utils"
)

var globalForwardingRulesIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"region": true,
}

var globalForwardingRulesAllowEmptyValues = map[string]bool{}

var globalForwardingRulesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type GlobalForwardingRulesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on globalForwardingRulesList and create for each TerraformResource
func (GlobalForwardingRulesGenerator) createResources(globalForwardingRulesList *compute.GlobalForwardingRulesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := globalForwardingRulesList.Pages(ctx, func(page *compute.ForwardingRuleList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_global_forwarding_rule",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each globalForwardingRules create 1 TerraformResource
// Need globalForwardingRules name as ID for terraform resource
func (g GlobalForwardingRulesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	globalForwardingRulesList := computeService.GlobalForwardingRules.List(project)

	resources := g.createResources(globalForwardingRulesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, globalForwardingRulesIgnoreKey, globalForwardingRulesAllowEmptyValues, globalForwardingRulesAdditionalFields)
	return resources, metadata, nil

}
