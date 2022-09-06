// Copyright 2022 The Terraformer Authors.
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

package tencentcloud

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type RouteTableGenerator struct {
	TencentCloudService
}

func (g *RouteTableGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeRouteTablesRequest()

	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_route_table") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.RouteTableIds = append(request.RouteTableIds, &filters[i])
	}

	offset := 0
	pageSize := 50
	allInstances := make([]*vpc.RouteTable, 0)

	for {
		offsetString := strconv.Itoa(offset)
		limitString := strconv.Itoa(pageSize)
		request.Offset = &offsetString
		request.Limit = &limitString
		response, err := client.DescribeRouteTables(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.RouteTableSet...)
		if len(response.Response.RouteTableSet) < pageSize {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.RouteTableId,
			*instance.RouteTableName+"_"+*instance.RouteTableId,
			"tencentcloud_route_table",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		for _, entry := range instance.RouteSet {
			entryID := fmt.Sprintf("%d.%s", *entry.RouteId, *instance.RouteTableId)
			entryName := fmt.Sprintf("%s_%d", *instance.RouteTableId, *entry.RouteId)
			entryResource := terraformutils.NewResource(
				entryID,
				entryName,
				"tencentcloud_route_table_entry",
				"tencentcloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			entryResource.AdditionalFields["route_table_id"] = "${tencentcloud_route_table." + resource.ResourceName + ".id}"
			g.Resources = append(g.Resources, entryResource)
		}
	}

	return nil
}
