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
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
)

type TeoGenerator struct {
	TencentCloudService
}

func (g *TeoGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := teo.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := teo.NewDescribeZonesRequest()
	response, err := client.DescribeZones(request)
	if err != nil {
		return err
	}

	for _, instance := range response.Response.Zones {
		resource := terraformutils.NewResource(
			*instance.ZoneId,
			*instance.ZoneName,
			"tencentcloud_teo_zone",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		if err := g.loadRuleEngines(client, *instance.ZoneId); err != nil {
			return err
		}
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *TeoGenerator) loadRuleEngines(client *teo.Client, zoneID string) error {
	request := teo.NewDescribeRulesRequest()
	request.ZoneId = &zoneID
	response, err := client.DescribeRules(request)
	if err != nil {
		return err
	}

	for _, rule := range response.Response.RuleItems {
		resource := terraformutils.NewResource(
			fmt.Sprintf("%s#%s", zoneID, *rule.RuleId),
			*rule.RuleName,
			"tencentcloud_teo_rule_engine",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		// resource.AdditionalFields["zone_id"] = "${local.zone_id}"
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
