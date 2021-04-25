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

package datadog

import (
	"context"
	"fmt"
	"strconv"

	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SecurityMonitoringRuleAllowEmptyValues ...
	SecurityMonitoringRuleAllowEmptyValues = []string{"tags."}
)

// SecurityMonitoringRuleGenerator ...
type SecurityMonitoringRuleGenerator struct {
	DatadogService
}

func (g *SecurityMonitoringRuleGenerator) createResources(rulesResponse []datadogV2.SecurityMonitoringRuleResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, rule := range rulesResponse {
		if !rule.GetIsDefault() {
			resourceName := rule.GetId()
			resources = append(resources, g.createResource(resourceName, rule.GetIsEnabled()))
		}
	}

	return resources
}

func (g *SecurityMonitoringRuleGenerator) createResource(ruleID string, ruleEnabled bool) terraformutils.Resource {
	return terraformutils.NewResource(
		ruleID,
		fmt.Sprintf("security_monitoring_rule_%s", ruleID),
		"datadog_security_monitoring_rule",
		"datadog",
		map[string]string{
			"enabled": strconv.FormatBool(ruleEnabled),
		},
		SecurityMonitoringRuleAllowEmptyValues,
		map[string]interface{}{},
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each SecurityMonitoringRule create 1 TerraformResource.
// Need SecurityMonitoringRule ID as ID for terraform resource
func (g *SecurityMonitoringRuleGenerator) InitResources() error {
	var securityMonitoringRuleResponses []datadogV2.SecurityMonitoringRuleResponse

	datadogClientV2 := g.Args["datadogClientV2"].(*datadogV2.APIClient)
	authV2 := g.Args["authV2"].(context.Context)

	pageSize := int64(1000)
	pageNumber := int64(0)
	remaining := int64(1)

	for remaining > int64(0) {
		resp, _, err := datadogClientV2.SecurityMonitoringApi.ListSecurityMonitoringRules(authV2).PageSize(pageSize).PageNumber(pageNumber).Execute()
		if err != nil {
			return err
		}
		securityMonitoringRuleResponses = append(securityMonitoringRuleResponses, resp.GetData()...)

		remaining = resp.Meta.Page.GetTotalCount() - pageSize*(pageNumber+1)
		pageNumber++
	}

	g.Resources = g.createResources(securityMonitoringRuleResponses)
	return nil
}
