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

type SubnetGenerator struct {
	TencentCloudService
}

func (g *SubnetGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vpc.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := vpc.NewDescribeSubnetsRequest()
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
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
