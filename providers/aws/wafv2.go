// Copyright 2020 The Terraformer Authors.
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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
)

var wafv2AllowEmptyValues = []string{"tags."}

type Wafv2Generator struct {
	AWSService
	scope types.Scope
}

func NewWafv2CloudfrontGenerator() *Wafv2Generator {
	return &Wafv2Generator{scope: types.ScopeCloudfront}
}

func NewWafv2RegionalGenerator() *Wafv2Generator {
	return &Wafv2Generator{scope: types.ScopeRegional}
}

func (g *Wafv2Generator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := wafv2.NewFromConfig(config)

	if err := g.loadWebACL(svc); err != nil {
		return err
	}
	if err := g.loadIPSet(svc); err != nil {
		return err
	}
	if err := g.loadRegexPatternSets(svc); err != nil {
		return err
	}
	if err := g.loadWafRuleGroups(svc); err != nil {
		return err
	}
	if err := g.loadWebACLLoggingConfiguration(svc); err != nil {
		return err
	}

	return nil
}

func (g *Wafv2Generator) loadWebACL(svc *wafv2.Client) error {
	output, err := svc.ListWebACLs(context.TODO(), &wafv2.ListWebACLsInput{Scope: g.scope})
	if err != nil {
		return err
	}
	for _, acl := range output.WebACLs {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*acl.Id,
			*acl.Name+"_"+(*acl.Id)[0:8],
			"aws_wafv2_web_acl",
			"aws",
			map[string]string{
				"name":  *acl.Name,
				"scope": string(g.scope),
			},
			wafv2AllowEmptyValues,
			map[string]interface{}{},
		))
		if g.scope == types.ScopeRegional {
			// cloudfront associations are not listed here since they should to defined in
			// aws_cloudfront_distribution resource instead
			err = g.loadWebACLAssociations(svc, acl.ARN)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Wafv2Generator) loadWebACLAssociations(svc *wafv2.Client, webACLArn *string) error {
	for _, resourceType := range types.ResourceTypeApplicationLoadBalancer.Values() {
		output, err := svc.ListResourcesForWebACL(context.TODO(),
			&wafv2.ListResourcesForWebACLInput{WebACLArn: webACLArn, ResourceType: resourceType})
		if err != nil {
			return err
		}
		for _, resource := range output.ResourceArns {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resource,
				resource,
				"aws_wafv2_web_acl_association",
				"aws",
				map[string]string{
					"resource_arn": resource,
					"web_acl_arn":  *webACLArn,
				},
				wafv2AllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}

func (g *Wafv2Generator) loadIPSet(svc *wafv2.Client) error {
	output, err := svc.ListIPSets(context.TODO(), &wafv2.ListIPSetsInput{Scope: g.scope})
	if err != nil {
		return err
	}
	for _, IPSet := range output.IPSets {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*IPSet.Id,
			*IPSet.Name+"_"+(*IPSet.Id)[0:8],
			"aws_wafv2_ip_set",
			"aws",
			map[string]string{
				"name":  *IPSet.Name,
				"scope": string(g.scope),
			},
			wafv2AllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *Wafv2Generator) loadRegexPatternSets(svc *wafv2.Client) error {
	output, err := svc.ListRegexPatternSets(context.TODO(), &wafv2.ListRegexPatternSetsInput{Scope: g.scope})
	if err != nil {
		return err
	}
	for _, regexPatternSet := range output.RegexPatternSets {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*regexPatternSet.Id,
			*regexPatternSet.Name+"_"+(*regexPatternSet.Id)[0:8],
			"aws_wafv2_regex_pattern_set",
			"aws",
			map[string]string{
				"name":  *regexPatternSet.Name,
				"scope": string(g.scope),
			},
			wafv2AllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *Wafv2Generator) loadWafRuleGroups(svc *wafv2.Client) error {
	output, err := svc.ListRuleGroups(context.TODO(), &wafv2.ListRuleGroupsInput{Scope: g.scope})
	if err != nil {
		return err
	}
	for _, ruleGroup := range output.RuleGroups {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*ruleGroup.Id,
			*ruleGroup.Name+"_"+(*ruleGroup.Id)[0:8],
			"aws_wafv2_rule_group",
			"aws",
			map[string]string{
				"arn":   *ruleGroup.ARN,
				"name":  *ruleGroup.Name,
				"scope": string(g.scope),
			},
			wafv2AllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *Wafv2Generator) loadWebACLLoggingConfiguration(svc *wafv2.Client) error {
	output, err := svc.ListLoggingConfigurations(context.TODO(), &wafv2.ListLoggingConfigurationsInput{Scope: g.scope})
	if err != nil {
		return err
	}
	for _, logConfig := range output.LoggingConfigurations {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*logConfig.ResourceArn,
			*logConfig.ResourceArn,
			"aws_wafv2_web_acl_logging_configuration",
			"aws",
			map[string]string{
				"resource_arn": *logConfig.ResourceArn,
			},
			wafv2AllowEmptyValues,
			map[string]interface{}{},
		))
	}
	return nil
}
