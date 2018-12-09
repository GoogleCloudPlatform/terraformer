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

var targetSslProxiesIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"proxy_id": true,
}

var targetSslProxiesAllowEmptyValues = map[string]bool{}

var targetSslProxiesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type TargetSslProxiesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on targetSslProxiesList and create for each TerraformResource
func (TargetSslProxiesGenerator) createResources(targetSslProxiesList *compute.TargetSslProxiesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := targetSslProxiesList.Pages(ctx, func(page *compute.TargetSslProxyList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_target_ssl_proxy",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
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

// Generate TerraformResources from GCP API,
// from each targetSslProxies create 1 TerraformResource
// Need targetSslProxies name as ID for terraform resource
func (g TargetSslProxiesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	targetSslProxiesList := computeService.TargetSslProxies.List(project)

	resources := g.createResources(targetSslProxiesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, targetSslProxiesIgnoreKey, targetSslProxiesAllowEmptyValues, targetSslProxiesAdditionalFields)
	return resources, metadata, nil

}
