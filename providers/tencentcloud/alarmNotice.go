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
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"log"
)

type AlarmNoticeGenerator struct {
	TencentCloudService
}

func (g *AlarmNoticeGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := monitor.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}
	println("++++++++++++++")

	//request := monitor.NewDescribeBasicAlarmListRequest()
	request := monitor.NewDescribeBasicAlarmListRequest()

	var s = "monitor"
	request.Module = &s

	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_monitor_alarm_notice") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}
	for i := range filters {
		request.ViewNames = append(request.MetricNames, &filters[i])
	}

	var offset int64
	var pageSize int64 = 50
	allInstances := make([]*monitor.DescribeBasicAlarmListAlarms, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeBasicAlarmList(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.Alarms...)
		if len(response.Response.Alarms) < int(pageSize) {
			break
		}
		offset += pageSize

		println("这里是日志")
		log.Printf("response body%s  , request body [%s]\n",
			response.ToJsonString(), request.ToJsonString())
	}

	println("_______________")
	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.ObjId,
			*instance.ObjId+"_"+*instance.ObjName,
			"tencentcloud_monitor_alarm_notice",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
		println("for 循环里面")
	}

	return nil
}
