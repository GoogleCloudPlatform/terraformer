// Copyright 2021 The Terraformer Authors.
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
	offset := 0
	pageSize := 50
	allRouteTables := make([]*vpc.RouteTable, 0)

	for {
		offsetString := strconv.Itoa(offset)
		limitString := strconv.Itoa(pageSize)
		request.Offset = &offsetString
		request.Limit = &limitString
		response, err := client.DescribeRouteTables(request)
		if err != nil {
			return err
		}

		allRouteTables = append(allRouteTables, response.Response.RouteTableSet...)
		if len(response.Response.RouteTableSet) < pageSize {
			break
		}
		offset += pageSize
	}

	for _, routeTable := range allRouteTables {
		resource := terraformutils.NewResource(
			*routeTable.RouteTableId,
			*routeTable.RouteTableName+"_"+*routeTable.RouteTableId,
			"tencentcloud_route_table",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
