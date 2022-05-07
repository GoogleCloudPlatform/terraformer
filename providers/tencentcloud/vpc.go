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
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type VpcGenerator struct {
	TencentCloudService
}

func (g *VpcGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeVpcsRequest()
	request.Filters = make([]*vpc.Filter, 0)
	vpcIds := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vpc") {
			vpcIds = append(vpcIds, filter.AcceptableValues...)
		}
	}
	if len(vpcIds) > 0 {
		request.VpcIds = make([]*string, 0, len(vpcIds))
		for i := range vpcIds {
			request.VpcIds = append(request.VpcIds, &vpcIds[i])
		}
	}

	offset := 0
	pageSize := 50
	allVpcs := make([]*vpc.Vpc, 0)

	for {
		offsetString := strconv.Itoa(offset)
		limitString := strconv.Itoa(pageSize)
		request.Offset = &offsetString
		request.Limit = &limitString
		response, err := client.DescribeVpcs(request)
		if err != nil {
			return err
		}

		allVpcs = append(allVpcs, response.Response.VpcSet...)
		if len(response.Response.VpcSet) < pageSize {
			break
		}
		offset += pageSize
	}

	for _, vpcInstance := range allVpcs {
		resource := terraformutils.NewResource(
			*vpcInstance.VpcId,
			*vpcInstance.VpcName+"_"+*vpcInstance.VpcId,
			"tencentcloud_vpc",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		// g.loadSubnets(client, *vpcInstance.VpcId, resource.ResourceName)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *VpcGenerator) loadSubnets(client *vpc.Client, vpcID, resourceName string) error {
	request := vpc.NewDescribeSubnetsRequest()
	request.Filters = make([]*vpc.Filter, 0, 1)
	idKey := "vpc-id"
	idFilter := vpc.Filter{
		Name:   &idKey,
		Values: []*string{&vpcID},
	}
	request.Filters = append(request.Filters, &idFilter)

	offset := 0
	pageSize := 50
	allSubnets := make([]*vpc.Subnet, 0)

	for {
		offsetString := strconv.Itoa(offset)
		limitString := strconv.Itoa(pageSize)
		request.Offset = &offsetString
		request.Limit = &limitString
		response, err := client.DescribeSubnets(request)
		if err != nil {
			return err
		}

		allSubnets = append(allSubnets, response.Response.SubnetSet...)
		if len(response.Response.SubnetSet) < pageSize {
			break
		}
		offset += pageSize
	}

	for _, subnet := range allSubnets {
		resource := terraformutils.NewResource(
			*subnet.SubnetId,
			*subnet.SubnetName+"_"+*subnet.SubnetId,
			"tencentcloud_subnet",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["vpc_id"] = "${tencentcloud_vpc." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
