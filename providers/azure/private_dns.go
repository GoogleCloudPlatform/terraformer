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

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type PrivateDNSGenerator struct {
	AzureService
}

func (g *PrivateDNSGenerator) listRecordSets(resourceGroupName string, privateZoneName string, top *int32) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	RecordSetsClient := privatedns.NewRecordSetsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	RecordSetsClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	recordSetIterator, err := RecordSetsClient.ListComplete(ctx, resourceGroupName, privateZoneName, top, "")
	if err != nil {
		return nil, err
	}
	for recordSetIterator.NotDone() {
		recordSet := recordSetIterator.Value()
		// NOTE:
		// Format example: "Microsoft.Network/privateDnsZones/CNAME"
		recordTypeSplitted := strings.Split(*recordSet.Type, "/")
		recordType := recordTypeSplitted[len(recordTypeSplitted)-1]
		typeResourceNameMap := map[string]string{
			"A":     "azurerm_private_dns_a_record",
			"AAAA":  "azurerm_private_dns_aaaa_record",
			"CNAME": "azurerm_private_dns_cname_record",
			"MX":    "azurerm_private_dns_mx_record",
			"PTR":   "azurerm_private_dns_ptr_record",
			"SRV":   "azurerm_private_dns_srv_record",
			"TXT":   "azurerm_private_dns_txt_record",
		}
		if resName, exist := typeResourceNameMap[recordType]; exist {
			resources = append(resources, terraformutils.NewSimpleResource(
				*recordSet.ID,
				*recordSet.Name,
				resName,
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

func (g *PrivateDNSGenerator) listVirtualNetworkLinks(resourceGroupName string, privateZoneName string, pageSize *int32) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	VirtualNetworkLinksClient := privatedns.NewVirtualNetworkLinksClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	VirtualNetworkLinksClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	virtualNetworkLinkIterator, err := VirtualNetworkLinksClient.ListComplete(ctx, resourceGroupName, privateZoneName, pageSize)
	if err != nil {
		return nil, err
	}
	for virtualNetworkLinkIterator.NotDone() {
		virtualNetworkLink := virtualNetworkLinkIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*virtualNetworkLink.ID,
			*virtualNetworkLink.Name,
			"azurerm_private_dns_zone_virtual_network_link",
			g.ProviderName,
			[]string{}))

		if err := virtualNetworkLinkIterator.Next(); err != nil {
			log.Println(err)
			break
		}

	}

	return resources, nil
}

func (g *PrivateDNSGenerator) listAndAddForPrivateDNSZone() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	PrivateDNSZonesClient := privatedns.NewPrivateZonesClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	PrivateDNSZonesClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	var pageSize int32 = 50

	var (
		dnsZoneIterator privatedns.PrivateZoneListResultIterator
		err             error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		dnsZoneIterator, err = PrivateDNSZonesClient.ListByResourceGroupComplete(ctx, rg, &pageSize)
	} else {
		dnsZoneIterator, err = PrivateDNSZonesClient.ListComplete(ctx, &pageSize)
	}
	if err != nil {
		return nil, err
	}
	for dnsZoneIterator.NotDone() {
		zone := dnsZoneIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*zone.ID,
			*zone.Name,
			"azurerm_private_dns_zone",
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

		networkLinks, err := g.listVirtualNetworkLinks(id.ResourceGroup, *zone.Name, &pageSize)
		if err != nil {
			return nil, err
		}
		resources = append(resources, networkLinks...)

		if err := dnsZoneIterator.Next(); err != nil {
			log.Println(err)
			return resources, err
		}
	}

	return resources, nil
}

func (g *PrivateDNSGenerator) InitResources() error {
	functions := []func() ([]terraformutils.Resource, error){
		g.listAndAddForPrivateDNSZone,
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
