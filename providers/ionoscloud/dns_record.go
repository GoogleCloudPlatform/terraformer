package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DNSRecordGenerator struct {
	Service
}

func (g *DNSRecordGenerator) InitResources() error {
	client := g.generateClient()
	dnsAPIClient := client.DNSAPIClient
	resourceType := "ionoscloud_dns_record"

	zonesResponse, _, err := dnsAPIClient.ZonesApi.ZonesGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if zonesResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing DNS Zones, but received 'nil' instead")
		return nil
	}
	zones := *zonesResponse.Items
	for _, zone := range zones {
		recordsResponse, _, err := dnsAPIClient.RecordsApi.ZonesRecordsGet(context.TODO(), *zone.Id).Execute()
		if err != nil {
			return err
		}
		if recordsResponse.Items == nil {
			log.Printf("[WARNING] expected a response containing DNS Records, but received 'nil' instead, skipping search for DNS Zone with ID: %v", *zone.Id)
			continue
		}
		records := *recordsResponse.Items
		for _, record := range records {
			if record.Properties == nil || record.Properties.Name == nil {
				log.Printf("[WARNING] 'nil' values in the response for DNS Record with ID: %v, Zone ID: %v, skipping this resource", *record.Id, *zone.Id)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*record.Id,
				*record.Properties.Name+"-"+*record.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.ZoneID: *zone.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
