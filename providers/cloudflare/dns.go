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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type DNSGenerator struct {
	CloudflareService
}

func (*DNSGenerator) createZonesResource(api *cf.API, zoneID string) ([]terraformutils.Resource, error) {
	zoneDetails, err := api.ZoneDetails(zoneID)
	if err != nil {
		log.Println(err)
		return []terraformutils.Resource{}, err
	}

	resource := terraformutils.NewResource(
		zoneDetails.ID,
		zoneDetails.Name,
		"cloudflare_zone",
		"cloudflare",
		map[string]string{
			"id": zoneDetails.ID,
		},
		[]string{},
		map[string]interface{}{},
	)
	resource.IgnoreKeys = append(resource.IgnoreKeys, "^meta$")

	return []terraformutils.Resource{resource}, nil
}

func (*DNSGenerator) createRecordsResources(api *cf.API, zoneID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	records, err := api.DNSRecords(zoneID, cf.DNSRecord{})
	if err != nil {
		log.Println(err)
		return resources, err
	}

	for _, record := range records {
		r := terraformutils.NewResource(
			record.ID,
			fmt.Sprintf("%s_%s_%s", record.Type, record.ZoneName, record.ID),
			"cloudflare_record",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
				"domain":  record.ZoneName,
				"name":    record.Name,
			},
			[]string{},
			map[string]interface{}{},
		)

		r.IgnoreKeys = append(r.IgnoreKeys, "^metadata")
		resources = append(resources, r)
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

	funcs := []func(*cf.API, string) ([]terraformutils.Resource, error){
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
	return nil
}

func (g *DNSGenerator) PostConvertHook() error {
	// 'record' resource have 'data' and 'value' is mutual-exclude
	// delete which one have empty value
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "cloudflare_record" {
			if val, ok := resource.Item["data"]; ok && len(val.(map[string]interface{})) == 0 {
				delete(g.Resources[i].Item, "data")
			} else if val, ok := resource.Item["value"]; ok && len(val.(string)) == 0 {
				delete(g.Resources[i].Item, "value")
			}
		}
	}

	return nil
}
