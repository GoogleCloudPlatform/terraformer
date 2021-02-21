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
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

var wafAllowEmptyValues = []string{"tags."}

type WafGenerator struct {
	AWSService
}

func (g *WafGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := waf.NewFromConfig(config)

	if err := g.loadWebACL(svc); err != nil {
		return err
	}
	if err := g.loadByteMatchSet(svc); err != nil {
		return err
	}
	if err := g.loadGeoMatchSet(svc); err != nil {
		return err
	}
	if err := g.loadIPSet(svc); err != nil {
		return err
	}
	if err := g.loadRateBasedRules(svc); err != nil {
		return err
	}
	if err := g.loadRegexMatchSets(svc); err != nil {
		return err
	}
	if err := g.loadRegexPatternSets(svc); err != nil {
		return err
	}
	if err := g.loadWafRules(svc); err != nil {
		return err
	}
	if err := g.loadWafRuleGroups(svc); err != nil {
		return err
	}
	if err := g.loadSizeConstraintSets(svc); err != nil {
		return err
	}
	if err := g.loadSQLInjectionMatchSets(svc); err != nil {
		return err
	}
	if err := g.loadXSSMatchSet(svc); err != nil {
		return err
	}

	return nil
}

func (g *WafGenerator) loadWebACL(svc *waf.Client) error {
	output, err := svc.ListWebACLs(context.TODO(), &waf.ListWebACLsInput{})
	if err != nil {
		return err
	}
	for _, acl := range output.WebACLs {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*acl.WebACLId,
			*acl.Name+"_"+(*acl.WebACLId)[0:8],
			"aws_waf_web_acl",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadByteMatchSet(svc *waf.Client) error {
	output, err := svc.ListByteMatchSets(context.TODO(), &waf.ListByteMatchSetsInput{})
	if err != nil {
		return err
	}
	for _, byteMatchSet := range output.ByteMatchSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*byteMatchSet.ByteMatchSetId,
			*byteMatchSet.Name+"_"+(*byteMatchSet.ByteMatchSetId)[0:8],
			"aws_waf_byte_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadGeoMatchSet(svc *waf.Client) error {
	output, err := svc.ListGeoMatchSets(context.TODO(), &waf.ListGeoMatchSetsInput{})
	if err != nil {
		return err
	}
	for _, matchSet := range output.GeoMatchSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*matchSet.GeoMatchSetId,
			*matchSet.Name+"_"+(*matchSet.GeoMatchSetId)[0:8],
			"aws_waf_geo_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadIPSet(svc *waf.Client) error {
	output, err := svc.ListIPSets(context.TODO(), &waf.ListIPSetsInput{})
	if err != nil {
		return err
	}
	for _, IPSet := range output.IPSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*IPSet.IPSetId,
			*IPSet.Name+"_"+(*IPSet.IPSetId)[0:8],
			"aws_waf_ipset",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadRateBasedRules(svc *waf.Client) error {
	output, err := svc.ListRateBasedRules(context.TODO(), &waf.ListRateBasedRulesInput{})
	if err != nil {
		return err
	}
	for _, rule := range output.Rules {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*rule.RuleId,
			*rule.Name+"_"+(*rule.RuleId)[0:8],
			"aws_waf_rate_based_rule",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadRegexMatchSets(svc *waf.Client) error {
	output, err := svc.ListRegexMatchSets(context.TODO(), &waf.ListRegexMatchSetsInput{})
	if err != nil {
		return err
	}
	for _, regexMatchSet := range output.RegexMatchSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*regexMatchSet.RegexMatchSetId,
			*regexMatchSet.Name+"_"+(*regexMatchSet.RegexMatchSetId)[0:8],
			"aws_waf_regex_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadRegexPatternSets(svc *waf.Client) error {
	output, err := svc.ListRegexPatternSets(context.TODO(), &waf.ListRegexPatternSetsInput{})
	if err != nil {
		return err
	}
	for _, regexPatternSet := range output.RegexPatternSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*regexPatternSet.RegexPatternSetId,
			*regexPatternSet.Name+"_"+(*regexPatternSet.RegexPatternSetId)[0:8],
			"aws_waf_regex_pattern_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadWafRules(svc *waf.Client) error {
	output, err := svc.ListRules(context.TODO(), &waf.ListRulesInput{})
	if err != nil {
		return err
	}
	for _, rule := range output.Rules {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*rule.RuleId,
			*rule.Name+"_"+(*rule.RuleId)[0:8],
			"aws_waf_rule",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadWafRuleGroups(svc *waf.Client) error {
	output, err := svc.ListRuleGroups(context.TODO(), &waf.ListRuleGroupsInput{})
	if err != nil {
		return err
	}
	for _, ruleGroup := range output.RuleGroups {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*ruleGroup.RuleGroupId,
			*ruleGroup.Name+"_"+(*ruleGroup.RuleGroupId)[0:8],
			"aws_waf_rule_group",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadSizeConstraintSets(svc *waf.Client) error {
	output, err := svc.ListSizeConstraintSets(context.TODO(), &waf.ListSizeConstraintSetsInput{})
	if err != nil {
		return err
	}
	for _, sizeConstraintSet := range output.SizeConstraintSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*sizeConstraintSet.SizeConstraintSetId,
			*sizeConstraintSet.Name+"_"+(*sizeConstraintSet.SizeConstraintSetId)[0:8],
			"aws_waf_size_constraint_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadSQLInjectionMatchSets(svc *waf.Client) error {
	output, err := svc.ListSqlInjectionMatchSets(context.TODO(), &waf.ListSqlInjectionMatchSetsInput{})
	if err != nil {
		return err
	}
	for _, sqlInjectionMatchSet := range output.SqlInjectionMatchSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*sqlInjectionMatchSet.SqlInjectionMatchSetId,
			*sqlInjectionMatchSet.Name+"_"+(*sqlInjectionMatchSet.SqlInjectionMatchSetId)[0:8],
			"aws_waf_sql_injection_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafGenerator) loadXSSMatchSet(svc *waf.Client) error {
	output, err := svc.ListXssMatchSets(context.TODO(), &waf.ListXssMatchSetsInput{})
	if err != nil {
		return err
	}
	for _, xssMatchSet := range output.XssMatchSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*xssMatchSet.XssMatchSetId,
			*xssMatchSet.Name+"_"+(*xssMatchSet.XssMatchSetId)[0:8],
			"aws_waf_xss_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}
