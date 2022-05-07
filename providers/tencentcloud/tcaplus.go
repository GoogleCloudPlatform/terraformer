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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tcaplus "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb/v20190823"
)

type TcaplusGenerator struct {
	TencentCloudService
}

func (g *TcaplusGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := tcaplus.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := tcaplus.NewDescribeClustersRequest()

	var offset int64
	var pageSize int64 = 50
	allInstances := make([]*tcaplus.ClusterInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeClusters(request)
		if err != nil {
			sdkErr, ok := err.(*errors.TencentCloudSDKError)
			if ok && sdkErr.Code == "UnsupportedRegion" {
				return nil
			}
			return err
		}

		allInstances = append(allInstances, response.Response.Clusters...)
		if len(response.Response.Clusters) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ClusterId,
			*instance.ClusterName+"_"+*instance.ClusterId,
			"tencentcloud_tcaplus_cluster",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
