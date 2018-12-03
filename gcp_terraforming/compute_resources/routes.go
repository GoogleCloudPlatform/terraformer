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

var routesIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"google_compute_route": true,
	"next_hop_network":     true,
}

var routesAllowEmptyValues = map[string]bool{}

var routesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type RoutesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on routesList and create for each TerraformResource
func (RoutesGenerator) createResources(routesList *compute.RoutesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := routesList.Pages(ctx, func(page *compute.RouteList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_route",
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
// from each routes create 1 TerraformResource
// Need routes name as ID for terraform resource
func (g RoutesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	routesList := computeService.Routes.List(project)

	resources := g.createResources(routesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, routesIgnoreKey, routesAllowEmptyValues, routesAdditionalFields)
	return resources, metadata, nil

}
