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
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type MysqlGenerator struct {
	TencentCloudService
}

func (g *MysqlGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := cdb.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err := g.loadMysqlMaster(client); err != nil {
		return err
	}
	if err := g.loadMysqlReadOnly(client); err != nil {
		return err
	}

	return nil
}

func (g *MysqlGenerator) loadMysqlMaster(client *cdb.Client) error {
	request := cdb.NewDescribeDBInstancesRequest()
	var instanceTypeMaster uint64 = 1
	request.InstanceTypes = []*uint64{&instanceTypeMaster}

	var offset uint64 = 0
	var pageSize uint64 = 50
	allInstances := make([]*cdb.InstanceInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeDBInstances(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.InstanceId,
			*instance.InstanceName+"_"+*instance.InstanceId,
			"tencentcloud_mysql_instance",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *MysqlGenerator) loadMysqlReadOnly(client *cdb.Client) error {
	request := cdb.NewDescribeDBInstancesRequest()
	var instanceTypeMaster uint64 = 3
	request.InstanceTypes = []*uint64{&instanceTypeMaster}

	var offset uint64 = 0
	var pageSize uint64 = 50
	allInstances := make([]*cdb.InstanceInfo, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeDBInstances(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.Items...)
		if len(response.Response.Items) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.InstanceId,
			*instance.InstanceName+"_"+*instance.InstanceId,
			"tencentcloud_mysql_readonly_instance",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *MysqlGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_mysql_instance" {
			delete(resource.Item, "pay_type")
			delete(resource.Item, "period")
		}

		if resource.InstanceInfo.Type != "tencentcloud_mysql_readonly_instance" {
			if masterID, exist := resource.InstanceState.Attributes["master_instance_id"]; exist {
				for _, r := range g.Resources {
					if r.InstanceInfo.Type != "tencentcloud_mysql_instance" {
						continue
					}
					if masterID == r.InstanceState.Attributes["id"] {
						g.Resources[i].Item["master_instance_id"] = "${tencentcloud_mysql_instance." + r.ResourceName + ".id}"
					}
				}
			}
		}
	}

	return nil
}
