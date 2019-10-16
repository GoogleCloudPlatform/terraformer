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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
)

// PvtzGenerator Struct for generating AliCloud private zone
type PvtzGenerator struct {
	AliCloudService
}

func resourceFromZoneResponse(zone pvtz.Zone) terraform_utils.Resource {
	return terraform_utils.NewResource(
		zone.ZoneId,                    // id
		zone.ZoneId+"__"+zone.ZoneName, // name
		"alicloud_pvtz_zone",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all pvtz Zone ids and generates resources
func (g *PvtzGenerator) InitResources() error {
	client, err := LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allZones := make([]pvtz.Zone, 1)

	for remaining > 0 {
		raw, err := client.WithPvtzClient(func(pvtzClient *pvtz.Client) (interface{}, error) {
			request := pvtz.CreateDescribeZonesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return pvtzClient.DescribeZones(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*pvtz.DescribeZonesResponse)
		for _, Zone := range response.Zones.Zone {
			allZones = append(allZones, Zone)

		}
		remaining = response.TotalItems - pageNumber*pageSize
		pageNumber++
	}

	for _, Zone := range allZones {
		resource := resourceFromZoneResponse(Zone)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
