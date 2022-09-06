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

type ACLGenerator struct {
	TencentCloudService
}

func (g *ACLGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeNetworkAclsRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vpc_acl") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.NetworkAclIds = append(request.NetworkAclIds, &filters[i])
	}

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*vpc.NetworkAcl, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeNetworkAcls(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.NetworkAclSet...)
		if len(response.Response.NetworkAclSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.NetworkAclId,
			*instance.NetworkAclName+"_"+*instance.NetworkAclId,
			"tencentcloud_vpc_acl",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		for _, subnet := range instance.SubnetSet {
			attachment := terraformutils.NewResource(
				*instance.NetworkAclId+"#"+*subnet.SubnetId,
				*instance.NetworkAclId+"_"+*subnet.SubnetId,
				"tencentcloud_vpc_acl_attachment",
				"tencentcloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			attachment.AdditionalFields["acl_id"] = "${tencentcloud_vpc_acl." + resource.ResourceName + ".id}"
			g.Resources = append(g.Resources, attachment)
		}
	}

	return nil
}
