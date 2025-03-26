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
	tat "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat/v20201028"
)

type TatGenerator struct {
	TencentCloudService
}

func (g *TatGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := tat.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	return g.DescribeCommands(client)
}
func (g *TatGenerator) DescribeCommands(client *tat.Client) error {
	request := tat.NewDescribeCommandsRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_tat_command") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		request.CommandIds = append(request.CommandIds, &filters[i])
	}

	var offset uint64
	var limit uint64 = 50
	allInstances := make([]*tat.Command, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeCommands(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.CommandSet...)
		if len(response.Response.CommandSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.CommandId,
			*instance.CommandId+"_"+*instance.CommandId,
			"tencentcloud_tat_command",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
		if err := g.DescribeInvokers(client, *instance.CommandId, resource.ResourceName); err != nil {
			return err
		}
	}

	return nil
}
func (g *TatGenerator) DescribeInvokers(client *tat.Client, commandID, resourceName string) error {
	request := tat.NewDescribeInvokersRequest()
	request.Filters = []*tat.Filter{
		{
			Name:   String("command-id"),
			Values: []*string{&commandID},
		},
	}
	var offset uint64
	var limit uint64 = 50
	allInstances := make([]*tat.Invoker, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeInvokers(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.InvokerSet...)
		if len(response.Response.InvokerSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.InvokerId,
			*instance.InvokerId+"_"+*instance.InvokerId,
			"tencentcloud_tat_invoker",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["command_id"] = "${tencentcloud_tat_command." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
