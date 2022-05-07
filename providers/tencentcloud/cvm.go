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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type CvmGenerator struct {
	TencentCloudService
}

func (g *CvmGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := cvm.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := cvm.NewDescribeInstancesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_instance") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.InstanceIds = append(request.InstanceIds, &filters[i])
	}

	var offset int64
	var pageSize int64 = 50
	allInstances := make([]*cvm.Instance, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeInstances(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.InstanceSet...)
		if len(response.Response.InstanceSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.InstanceId,
			*instance.InstanceName+"_"+*instance.InstanceId,
			"tencentcloud_instance",
			"tencentcloud",
			map[string]string{
				"disable_monitor_service":  "false",
				"disable_security_service": "false",
				"force_delete":             "false",
			},
			[]string{},
			map[string]interface{}{},
		)
		if instance.LoginSettings != nil && len(instance.LoginSettings.KeyIds) > 0 {
			keyPairName, err := g.loadKeyPairs(client, instance.LoginSettings.KeyIds)
			if err == nil {
				resource.AdditionalFields["key_name"] = "${tencentcloud_key_pair." + keyPairName + ".id}"
			}
		}
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *CvmGenerator) loadKeyPairs(client *cvm.Client, keyIds []*string) (resourceName string, errRet error) {
	request := cvm.NewDescribeKeyPairsRequest()
	request.KeyIds = keyIds
	response, err := client.DescribeKeyPairs(request)
	if err != nil {
		errRet = err
		return
	}
	if len(response.Response.KeyPairSet) < 1 {
		errRet = fmt.Errorf("no key pair")
		return
	}

	instance := response.Response.KeyPairSet[0]
	resourceName = *instance.KeyName + "_" + *instance.KeyId
	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "tencentcloud_key_pair" && r.ResourceName == resourceName {
			return
		}
	}
	resource := terraformutils.NewResource(
		*instance.KeyId,
		resourceName,
		"tencentcloud_key_pair",
		"tencentcloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
	g.Resources = append(g.Resources, resource)
	return
}

/*
func (g *CvmGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_instance" {
			resource.InstanceState.Attributes["disable_monitor_service"] = "false"
			resource.InstanceState.Attributes["disable_security_service"] = "false"
			resource.InstanceState.Attributes["force_delete"] = "false"
		}
	}
	return nil
}
*/
