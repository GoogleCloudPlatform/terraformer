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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// SecurityMonitoringDefaultRuleAllowEmptyValues ...
	SecurityMonitoringDefaultRuleAllowEmptyValues = []string{"tags."}
)

// SecurityMonitoringDefaultRuleGenerator ...
type SecurityMonitoringDefaultRuleGenerator struct {
	DatadogService
}

func (g *SecurityMonitoringDefaultRuleGenerator) createResources(rulesResponse []datadogV2.SecurityMonitoringRuleResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, rule := range rulesResponse {
		if rule.SecurityMonitoringSignalRuleResponse != nil {
			if rule.SecurityMonitoringSignalRuleResponse.GetIsDefault() {
				resourceName := rule.SecurityMonitoringSignalRuleResponse.GetId()
				resources = append(resources, g.createResource(resourceName))
			}
		}

		if rule.SecurityMonitoringStandardRuleResponse != nil {
			if rule.SecurityMonitoringStandardRuleResponse.GetIsDefault() {
				resourceName := rule.SecurityMonitoringStandardRuleResponse.GetId()
				resources = append(resources, g.createResource(resourceName))
			}
		}
	}

	return resources
}

func (g *SecurityMonitoringDefaultRuleGenerator) createResource(ruleID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		ruleID,
		fmt.Sprintf("security_monitoring_default_rule_%s", ruleID),
		"datadog_security_monitoring_default_rule",
		"datadog",
		SecurityMonitoringDefaultRuleAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each SecurityMonitoringDefaultRule create 1 TerraformResource.
// Need SecurityMonitoringDefaultRule ID as ID for terraform resource
func (g *SecurityMonitoringDefaultRuleGenerator) InitResources() error {
	var securityMonitoringRuleResponses []datadogV2.SecurityMonitoringRuleResponse

	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV2.NewSecurityMonitoringApi(datadogClient)

	pageSize := int64(1000)
	pageNumber := int64(0)
	remaining := int64(1)

	for remaining > int64(0) {
		resp, _, err := api.ListSecurityMonitoringRules(auth,
			*datadogV2.NewListSecurityMonitoringRulesOptionalParameters().
				WithPageSize(pageSize).
				WithPageNumber(pageNumber))
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
