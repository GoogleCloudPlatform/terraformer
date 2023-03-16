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
	tem "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem/v20210701"
)

type TemGenerator struct {
	TencentCloudService
}

func (g *TemGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := tem.NewClient(&credential, region, profile)
	client.WithHttpTransport(&LogRoundTripper{
		Verbose: g.Verbose,
	})
	if err != nil {
		return err
	}

	if err := g.DescribeEnvironments(client); err != nil {
		return err
	}

	return nil
}
func (g *TemGenerator) DescribeEnvironments(client *tem.Client) error {
	request := tem.NewDescribeEnvironmentsRequest()

	var offset int64 = 0
	var limit int64 = 50
	allInstances := make([]*tem.TemNamespaceInfo, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeEnvironments(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.Result.Records...)
		if len(response.Response.Result.Records) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.EnvironmentId,
			*instance.EnvironmentId+"_"+*instance.EnvironmentId,
			"tencentcloud_tem_environment",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
		if err := g.DescribeIngresses(client, *instance.EnvironmentId, resource.ResourceName); err != nil {
			return err
		}
	}

	return nil
}
func (g *TemGenerator) DescribeIngresses(client *tem.Client, environmentId, resourceName string) error {
	request := tem.NewDescribeIngressesRequest()

	request.EnvironmentId = &environmentId
	request.ClusterNamespace = String("default")
	var allInstances []*tem.IngressInfo
	response, err := client.DescribeIngresses(request)
	if err != nil {
		return err
	}

	allInstances = response.Response.Result
	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.EnvironmentId+"#"+*instance.IngressName,
			*instance.EnvironmentId+"_"+*instance.IngressName,
			"tencentcloud_tem_gateway",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)

		//ingress["environment_id"] = "${tencentcloud_tem_environment." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

//func (g *TemGenerator) PostConvertHook() error {
//	for _, resource := range g.Resources {
//		if resource.InstanceInfo.Type == "tencentcloud_tem_gateway" {
//			resource.AdditionalFields["ingress.environment_id"] = "${tencentcloud_tem_environment." + resource.ResourceName + ".id}"
//		}
//	}
//
//	return nil
//}
