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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
)

type GaapGenerator struct {
	TencentCloudService
}

func (g *GaapGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := gaap.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err := g.loadProxy(client); err != nil {
		return err
	}
	if err := g.loadRealServer(client); err != nil {
		return err
	}

	return nil
}

func (g *GaapGenerator) loadProxy(client *gaap.Client) error {
	request := gaap.NewDescribeProxiesRequest()

	var offset uint64 = 0
	var pageSize uint64 = 50
	allInstances := make([]*gaap.ProxyInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeProxies(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.ProxySet...)
		if len(response.Response.ProxySet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ProxyId,
			*instance.ProxyName+"_"+*instance.ProxyId,
			"tencentcloud_gaap_proxy",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *GaapGenerator) loadRealServer(client *gaap.Client) error {
	request := gaap.NewDescribeRealServersRequest()
	var projectID int64 = -1
	request.ProjectId = &projectID

	var offset uint64 = 0
	var pageSize uint64 = 50
	allInstances := make([]*gaap.BindRealServerInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeRealServers(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.RealServerSet...)
		if len(response.Response.RealServerSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.RealServerId,
			*instance.RealServerName+"_"+*instance.RealServerId,
			"tencentcloud_gaap_realserver",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
