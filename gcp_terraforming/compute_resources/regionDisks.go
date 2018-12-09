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

var regionDisksIgnoreKey = map[string]bool{
	"^id$":                 true,
	"^self_link$":          true,
	"^fingerprint$":        true,
	"^label_fingerprint$":  true,
	"^creation_timestamp$": true,

	"last_attach_timestamp": true,
	"last_detach_timestamp": true,
	"users":                 true,
	"source_snapshot_id":    true,
}

var regionDisksAllowEmptyValues = map[string]bool{}

var regionDisksAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type RegionDisksGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on regionDisksList and create for each TerraformResource
func (RegionDisksGenerator) createResources(regionDisksList *compute.RegionDisksListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := regionDisksList.Pages(ctx, func(page *compute.DiskList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_region_disk",
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
// from each regionDisks create 1 TerraformResource
// Need regionDisks name as ID for terraform resource
func (g RegionDisksGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	regionDisksList := computeService.RegionDisks.List(project, region)

	resources := g.createResources(regionDisksList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, regionDisksIgnoreKey, regionDisksAllowEmptyValues, regionDisksAdditionalFields)
	return resources, metadata, nil

}
