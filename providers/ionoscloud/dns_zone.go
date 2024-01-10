package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DNSZoneGenerator struct {
	Service
}

func (g *DNSZoneGenerator) InitResources() error {
	client := g.generateClient()
	dnsAPIClient := client.DNSAPIClient
	resourceType := "ionoscloud_dns_zone"

	response, _, err := dnsAPIClient.ZonesApi.ZonesGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing DNS Zones, but received 'nil' instead")
		return nil
	}
	zones := *response.Items
	for _, zone := range zones {
		if zone.Properties == nil || zone.Properties.ZoneName == nil {
			log.Printf("[WARNING] 'nil' values in the response for the DNS Zone with ID: %v, skipping this resource", *zone.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*zone.Id,
			*zone.Properties.ZoneName+"-"+*zone.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
