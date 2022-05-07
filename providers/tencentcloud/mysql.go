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
	"math/rand"

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

	request := cdb.NewDescribeDBInstancesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_mysql_instance") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.InstanceIds = append(request.InstanceIds, &filters[i])
	}

	var offset uint64
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
		if *instance.InstanceType == 1 {
			resource := terraformutils.NewResource(
				*instance.InstanceId,
				*instance.InstanceName+"_"+*instance.InstanceId,
				"tencentcloud_mysql_instance",
				"tencentcloud",
				map[string]string{
					"force_delete":   "false",
					"prepaid_period": "1",
				},
				[]string{},
				map[string]interface{}{},
			)
			g.Resources = append(g.Resources, resource)
		} else if *instance.InstanceType == 3 {
			resource := terraformutils.NewResource(
				*instance.InstanceId,
				*instance.InstanceName+"_"+*instance.InstanceId,
				"tencentcloud_mysql_readonly_instance",
				"tencentcloud",
				map[string]string{
					"force_delete":   "false",
					"prepaid_period": "1",
				},
				[]string{},
				map[string]interface{}{},
			)
			g.Resources = append(g.Resources, resource)
		}
	}

	return nil
}

func (g *MysqlGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_mysql_instance" {
			password := g.generatePassword(16)
			g.Resources[i].Item["root_password"] = password
			g.Resources[i].InstanceState.Attributes["root_password"] = password
		}
		delete(resource.Item, "pay_type")
		delete(resource.Item, "period")
	}

	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type != "tencentcloud_mysql_readonly_instance" {
			continue
		}
		delete(resource.Item, "pay_type")
		delete(resource.Item, "period")
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
	return nil
}

func (g *MysqlGenerator) generatePassword(length int) string {
	digits := "0123456789"
	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	specials := "_+-!@#$"
	all := digits + alphabets + specials

	password := make([]byte, length)
	password[0] = alphabets[rand.Intn(len(alphabets))]
	password[1] = digits[rand.Intn(len(digits))]
	for i := 2; i < length; i++ {
		password[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
	return string(password)
}
