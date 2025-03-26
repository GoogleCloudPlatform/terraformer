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
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type CbsGenerator struct {
	TencentCloudService
}

func (g *CbsGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := cbs.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := cbs.NewDescribeDisksRequest()

	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_cbs_storage") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.DiskIds = append(request.DiskIds, &filters[i])
	}

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*cbs.Disk, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeDisks(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.DiskSet...)
		if len(response.Response.DiskSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.DiskId,
			*instance.DiskId,
			"tencentcloud_cbs_storage",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)

		if *instance.Attached {
			attachment := terraformutils.NewResource(
				*instance.DiskId,
				*instance.DiskId,
				"tencentcloud_cbs_storage_attachment",
				"tencentcloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			)
			attachment.AdditionalFields["storage_id"] = "${tencentcloud_cbs_storage." + resource.ResourceName + ".id}"
			g.Resources = append(g.Resources, attachment)
		}
	}

	return nil
}
