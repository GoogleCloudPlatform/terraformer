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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type EipGenerator struct {
	TencentCloudService
}

func (g *EipGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	var offset int64
	var pageSize int64 = 50
	allInstances := make([]*vpc.Address, 0)
	request := vpc.NewDescribeAddressesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_eip") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.AddressIds = append(request.AddressIds, &filters[i])
	}

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeAddresses(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.AddressSet...)
		if len(response.Response.AddressSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.AddressId,
			*instance.AddressId,
			"tencentcloud_eip",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		if instance.InstanceId != nil && *instance.InstanceId != "" {
			association := terraformutils.NewResource(
				*instance.AddressId+"::"+*instance.InstanceId,
				*instance.AddressId,
				"tencentcloud_eip_association",
				"tencentcloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			association.AdditionalFields["eip_id"] = "${tencentcloud_eip." + resource.ResourceName + ".id}"
			g.Resources = append(g.Resources, association)
		}
	}

	return nil
}
