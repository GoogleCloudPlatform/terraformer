package clouddns

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/terraform_utils"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/dns/v1"
)

var cloudDNSIgnoreKey = map[string]bool{
	"^id$":          true,
	"^name_servers": true,
}

var cloudDNSAllowEmptyValues = map[string]bool{}

var cloudDNSAdditionalFields = map[string]string{}

type CloudDNSGenerator struct {
	gcp_generator.BasicGenerator
}

func (g CloudDNSGenerator) createZonesResources(ctx context.Context, svc *dns.Service, project string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	managedZonesListCall := svc.ManagedZones.List(project)

	err := managedZonesListCall.Pages(ctx, func(listDNS *dns.ManagedZonesListResponse) error {
		for _, zone := range listDNS.ManagedZones {
			resources = append(resources, terraform_utils.NewTerraformResource(
				zone.Name,
				zone.Name,
				"google_dns_managed_zone",
				"google",
				nil,
				map[string]string{
					"name": zone.Name,
				},
			))
			records := g.createRecordsResources(ctx, svc, project, zone.Name)
			resources = append(resources, records...)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return []terraform_utils.TerraformResource{}
	}
	return resources
}
func (CloudDNSGenerator) createRecordsResources(ctx context.Context, svc *dns.Service, project, zoneName string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	managedRecordsListCall := svc.ResourceRecordSets.List(project, zoneName)
	err := managedRecordsListCall.Pages(ctx, func(listDNS *dns.ResourceRecordSetsListResponse) error {
		for _, record := range listDNS.Rrsets {
			resources = append(resources, terraform_utils.NewTerraformResource(
				fmt.Sprintf("%s/%s/%s", zoneName, record.Name, record.Type),
				strings.TrimSuffix(record.Name, "."),
				"google_dns_record_set",
				"google",
				nil,
				map[string]string{
					"name":         record.Name,
					"managed_zone": zoneName,
					"type":         record.Type,
				},
			))
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return []terraform_utils.TerraformResource{}
	}
	return resources
}

// Generate TerraformResources from GCP API,
// create terraform resource for each zone + each record
func (g CloudDNSGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()
	var client *http.Client
	var err error
	client, err = google.DefaultClient(ctx, dns.NdevClouddnsReadwriteScope)
	svc, err := dns.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resources := g.createZonesResources(ctx, svc, project)
	metadata := terraform_utils.NewResourcesMetaData(resources, cloudDNSIgnoreKey, cloudDNSAllowEmptyValues, cloudDNSAdditionalFields)
	return resources, metadata, nil

}
