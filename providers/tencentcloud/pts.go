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
	pts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts/v20210728"
)

type PtsGenerator struct {
	TencentCloudService
}

func (g *PtsGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := pts.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	return g.DescribeProjects(client)
}
func (g *PtsGenerator) DescribeProjects(client *pts.Client) error {
	request := pts.NewDescribeProjectsRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_pts_project") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		request.ProjectIds = append(request.ProjectIds, &filters[i])
	}

	var offset int64
	var limit int64 = 50
	allInstances := make([]*pts.Project, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeProjects(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.ProjectSet...)
		if len(response.Response.ProjectSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ProjectId,
			*instance.ProjectId+"_"+*instance.ProjectId,
			"tencentcloud_pts_project",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
