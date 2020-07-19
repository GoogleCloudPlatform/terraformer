// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"context"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type DnsGenerator struct {
	AzureService
}

func (g *DnsGenerator) listRecordSets(resourceGroupName string, zoneName string, top *int32) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	RecordSetsClient := dns.NewRecordSetsClient(subscriptionID)
	RecordSetsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	recordSetIterator, err := RecordSetsClient.ListAllByDNSZoneComplete(ctx, resourceGroupName, zoneName, top, "")
	if err != nil {
		return nil, err
	}
	for recordSetIterator.NotDone() {
		recordSet := recordSetIterator.Value()
		// NOTE:
		// Format example: "Microsoft.Network/dnszones/AAAA"
		recordType := *recordSet.Type
		if strings.HasSuffix(recordType, "/A") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_a_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/AAAA") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_aaaa_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/CAA") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_caa_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/CNAME") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_cname_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/NS") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_ns_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/MX") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_mx_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/PTR") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_ptr_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/SRV") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_srv_record",
				g.ProviderName,
				[]string{}))
		} else if strings.HasSuffix(recordType, "/TXT") {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				"azurerm_dns_txt_record",
				g.ProviderName,
				[]string{}))
		}
		if err := recordSetIterator.Next(); err != nil {
			log.Println(err)
			break
		}

	}
	return resources, nil
}

func (g *DnsGenerator) listAndAddForDnsZone() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	DnsZonesClient := dns.NewZonesClient(subscriptionID)
	DnsZonesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var pageSize int32 = 50
	dnsZoneIterator, err := DnsZonesClient.ListComplete(ctx, &pageSize)
	if err != nil {
		return nil, err
	}
	for dnsZoneIterator.NotDone() {
		zone := dnsZoneIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*zone.ID,
			*zone.Name,
			"azurerm_dns_zone",
			g.ProviderName,
			[]string{}))

		id, err := ParseAzureResourceID(*zone.ID)
		if err != nil {
			return nil, err
		}

		records, err := g.listRecordSets(id.ResourceGroup, *zone.Name, &pageSize)
		if err != nil {
			return nil, err
		}
		resources = append(resources, records...)

		if err := dnsZoneIterator.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return resources, nil
}

func (g *DnsGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listAndAddForDnsZone,
	}

	for _, f := range functions {
		resources, err := f()
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}
