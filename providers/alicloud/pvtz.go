// Copyright 2018 The Terraformer Authors.
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

package alicloud

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/providers/alicloud/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
)

// PvtzGenerator Struct for generating AliCloud private zone
type PvtzGenerator struct {
	AliCloudService
}

func resourceFromZoneResponse(zone pvtz.Zone) terraformutils.Resource {
	return terraformutils.NewResource(
		zone.ZoneId,                    // id
		zone.ZoneId+"__"+zone.ZoneName, // name
		"alicloud_pvtz_zone",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func resourceFromZoneAttachmentResponse(zone pvtz.Zone) terraformutils.Resource {
	return terraformutils.NewResource(
		zone.ZoneId, // id
		zone.ZoneId+"__"+zone.ZoneName+"_attachment", // name
		"alicloud_pvtz_zone_attachment",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func resourceFromZoneRecordResponse(record pvtz.Record, zoneID string) terraformutils.Resource {
	return terraformutils.NewResource(
		strconv.FormatInt(record.RecordId, 10)+":"+zoneID,     // id
		strconv.FormatInt(record.RecordId, 10)+"__"+record.Rr, // name
		"alicloud_pvtz_zone_record",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

func initZones(client *connectivity.AliyunClient) ([]pvtz.Zone, error) {
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allZones := make([]pvtz.Zone, 0)

	for remaining > 0 {
		raw, err := client.WithPvtzClient(func(pvtzClient *pvtz.Client) (interface{}, error) {
			request := pvtz.CreateDescribeZonesRequest()
			request.RegionId = client.RegionID
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return pvtzClient.DescribeZones(request)
		})
		if err != nil {
			return nil, err
		}

		response := raw.(*pvtz.DescribeZonesResponse)
		allZones = append(allZones, response.Zones.Zone...)
		remaining = response.TotalItems - pageNumber*pageSize
		pageNumber++
	}
	return allZones, nil
}

func initZoneRecords(client *connectivity.AliyunClient, allZones []pvtz.Zone) ([]pvtz.Record, []string, error) {
	allZoneRecords := make([]pvtz.Record, 0)
	zoneIds := make([]string, 0)

	for _, zone := range allZones {
		remaining := 1
		pageNumber := 1
		pageSize := 10
		if zone.ZoneId == "" {
			continue
		}
		for remaining > 0 {
			raw, err := client.WithPvtzClient(func(pvtzClient *pvtz.Client) (interface{}, error) {
				request := pvtz.CreateDescribeZoneRecordsRequest()
				request.RegionId = client.RegionID
				request.ZoneId = zone.ZoneId
				request.PageSize = requests.NewInteger(pageSize)
				request.PageNumber = requests.NewInteger(pageNumber)
				return pvtzClient.DescribeZoneRecords(request)
			})
			if err != nil {
				return nil, nil, err
			}

			response := raw.(*pvtz.DescribeZoneRecordsResponse)
			for _, zoneRecord := range response.Records.Record {
				allZoneRecords = append(allZoneRecords, zoneRecord)
				zoneIds = append(zoneIds, zone.ZoneId)
			}
			remaining = response.TotalItems - pageNumber*pageSize
			pageNumber++
		}
	}
	return allZoneRecords, zoneIds, nil
}

// InitResources Gets the list of all pvtz Zone ids and generates resources
func (g *PvtzGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}

	allZones, err := initZones(client)
	if err != nil {
		return err
	}

	allRecords, zoneIds, err := initZoneRecords(client, allZones)
	if err != nil {
		return err
	}

	for _, zone := range allZones {
		resource := resourceFromZoneResponse(zone)
		g.Resources = append(g.Resources, resource)
	}

	for _, zone := range allZones {
		resource := resourceFromZoneAttachmentResponse(zone)
		g.Resources = append(g.Resources, resource)
	}

	for i, record := range allRecords {
		resource := resourceFromZoneRecordResponse(record, zoneIds[i])
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

// PostConvertHook Runs before HCL files are generated
func (g *PvtzGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_pvtz_zone_record" {
			// https://www.terraform.io/docs/providers/alicloud/r/pvtz_zone_record.html#priority
			v, e := strconv.Atoi(r.Item["priority"].(string))
			if v < 1 || v > 50 || e != nil {
				delete(r.Item, "priority")
			}
		}
	}

	return nil
}
