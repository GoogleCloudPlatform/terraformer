// Copyright 2019 The Terraformer Authors.
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

package cloudflare

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	cf "github.com/cloudflare/cloudflare-go"
)

type DNSGenerator struct {
	CloudflareService
}

func (*DNSGenerator) createZonesResource(api *cf.API, zoneID string) ([]terraform_utils.Resource, error) {
	zoneDetails, err := api.ZoneDetails(zoneID)
	if err != nil {
		log.Println(err)
		return []terraform_utils.Resource{}, err
	}

	resource := terraform_utils.NewResource(
		zoneDetails.ID,
		zoneDetails.Name,
		"cloudflare_zone",
		"cloudflare",
		map[string]string{
			"id": zoneDetails.ID,
		},
		[]string{},
		map[string]string{},
	)

	return []terraform_utils.Resource{resource}, nil
}

func (*DNSGenerator) createRecordsResources(api *cf.API, zoneID string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}
	records, err := api.DNSRecords(zoneID, cf.DNSRecord{})
	if err != nil {
		log.Println(err)
		return resources, err
	}

	for _, record := range records {
		resources = append(resources, terraform_utils.NewResource(
			record.ID,
			fmt.Sprintf("%s_%s_%s", record.Type, record.ZoneName, record.ID),
			"cloudflare_record",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
				"domain":  record.ZoneName,
			},
			[]string{},
			map[string]string{},
		))
	}

	return resources, nil
}

func (g *DNSGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		log.Println(err)
		return err
	}

	zones, err := api.ListZones()
	if err != nil {
		log.Println(err)
		return err
	}

	funcs := []func(*cf.API, string) ([]terraform_utils.Resource, error){
		g.createZonesResource,
		g.createRecordsResources,
	}

	for _, zone := range zones {
		for _, f := range funcs {
			tmpRes, err := f(api, zone.ID)
			if err != nil {
				log.Println(err)
				return err
			}
			g.Resources = append(g.Resources, tmpRes...)
		}
	}

	g.PopulateIgnoreKeys()
	return nil

}

func (g *DNSGenerator) PostConvertHook() error {
	for i, resourceRecord := range g.Resources {
		if resourceRecord.InstanceInfo.Type == "cloudflare_zone" {
			continue
		}

		item := resourceRecord.Item
		zoneID := item["zone_id"].(string)
		for _, resourceZone := range g.Resources {
			if resourceZone.InstanceInfo.Type != "cloudflare_zone" {
				continue
			}
			if zoneID == resourceZone.InstanceState.ID {
				if resourceRecord.InstanceInfo.Type == "cloudflare_record" {
					g.Resources[i].Item["domain"] = "${cloudflare_zone." + resourceZone.ResourceName + ".zone}"
				}
			}
		}
	}
	return nil
}
