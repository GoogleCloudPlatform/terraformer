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
	mongodb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20180408"
)

type MongodbGenerator struct {
	TencentCloudService
}

func (g *MongodbGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := mongodb.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := mongodb.NewDescribeDBInstancesRequest()

	var offset uint64
	var pageSize uint64 = 50
	allInstances := make([]*mongodb.MongoDBInstanceDetail, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeDBInstances(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.InstanceDetails...)
		if len(response.Response.InstanceDetails) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.InstanceId,
			*instance.InstanceName+"_"+*instance.InstanceId,
			"tencentcloud_mongodb_instance",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *MongodbGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_mongodb_instance" {
			g.Resources[i].Item["password"] = "test1234;"
		}
	}
	return nil
}
