// Copyright 2018 The Terraformer Authors.
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

package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// RdsGenerator Struct for generating AliCloud Elastic Compute Service
type RdsGenerator struct {
	AliCloudService
}

func resourceFromrdsResponse(rds rds.DBInstance) terraform_utils.Resource {
	return terraform_utils.NewResource(
		rds.DBInstanceId, // id
		rds.DBInstanceId+"__"+rds.DBInstanceDescription, // name
		"alicloud_db_instance",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all rds ids and generates resources
func (g *RdsGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allrdss := make([]rds.DBInstance, 0)

	for remaining > 0 {
		raw, err := client.WithRdsClient(func(rdsClient *rds.Client) (interface{}, error) {
			request := rds.CreateDescribeDBInstancesRequest()
			request.RegionId = client.RegionId
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return rdsClient.DescribeDBInstances(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*rds.DescribeDBInstancesResponse)
		for _, rds := range response.Items.DBInstance {
			allrdss = append(allrdss, rds)

		}
		remaining = response.TotalRecordCount - pageNumber*pageSize
		pageNumber++
	}

	for _, rds := range allrdss {
		resource := resourceFromrdsResponse(rds)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

// PostConvertHook Runs before HCL files are generated
func (g *RdsGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "alicloud_db_instance" {
			// https://www.terraform.io/docs/providers/alicloud/r/db_instance.html#period
			if r.Item["instance_charge_type"] != "PrePaid" {
				delete(r.Item, "period")
			}
		}
	}

	return nil
}
