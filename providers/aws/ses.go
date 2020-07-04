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
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go/aws"
)

var sesAllowEmptyValues = []string{"tags."}

type SesGenerator struct {
	AWSService
}

func (g *SesGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ses.New(config)

	if err := g.loadDomainIdentities(svc); err != nil {
		return err
	}
	if err := g.loadMailIdentities(svc); err != nil {
		return err
	}
	if err := g.loadTemplates(svc); err != nil {
		return err
	}
	if err := g.loadConfigurationSets(svc); err != nil {
		return err
	}
	if err := g.loadRuleSets(svc); err != nil {
		return err
	}

	return nil
}

func (g *SesGenerator) loadDomainIdentities(svc *ses.Client) error {
	p := ses.NewListIdentitiesPaginator(svc.ListIdentitiesRequest(&ses.ListIdentitiesInput{
		IdentityType: "Domain",
	}))
	for p.Next(context.Background()) {
		for _, identity := range p.CurrentPage().Identities {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				identity,
				identity,
				"aws_ses_domain_identity",
				"aws",
				sesAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *SesGenerator) loadMailIdentities(svc *ses.Client) error {
	p := ses.NewListIdentitiesPaginator(svc.ListIdentitiesRequest(&ses.ListIdentitiesInput{
		IdentityType: "EmailAddress",
	}))
	for p.Next(context.Background()) {
		for _, identity := range p.CurrentPage().Identities {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				identity,
				identity,
				"aws_ses_email_identity",
				"aws",
				sesAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *SesGenerator) loadTemplates(svc *ses.Client) error {
	templates, err := svc.ListTemplatesRequest(&ses.ListTemplatesInput{}).Send(context.Background())
	if err != nil {
		return err
	}

	for _, templateMetadata := range templates.TemplatesMetadata {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			aws.StringValue(templateMetadata.Name),
			aws.StringValue(templateMetadata.Name),
			"aws_ses_template",
			"aws",
			sesAllowEmptyValues))
	}
	return nil
}

func (g *SesGenerator) loadConfigurationSets(svc *ses.Client) error {
	configurationSets, err := svc.ListConfigurationSetsRequest(&ses.ListConfigurationSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}

	for _, configurationSet := range configurationSets.ConfigurationSets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			aws.StringValue(configurationSet.Name),
			aws.StringValue(configurationSet.Name),
			"aws_ses_configuration_set",
			"aws",
			sesAllowEmptyValues))
	}
	return nil
}

func (g *SesGenerator) loadRuleSets(svc *ses.Client) error {
	ruleSets, err := svc.ListReceiptRuleSetsRequest(&ses.ListReceiptRuleSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}

	for _, ruleSet := range ruleSets.RuleSets {
		ruleSetName := aws.StringValue(ruleSet.Name)
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			ruleSetName,
			ruleSetName,
			"aws_ses_receipt_rule_set",
			"aws",
			sesAllowEmptyValues))
		rules, err := svc.DescribeReceiptRuleSetRequest(&ses.DescribeReceiptRuleSetInput{
			RuleSetName: ruleSet.Name,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, rule := range rules.Rules {
			ruleID := ruleSetName + ":" + *rule.Name
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*rule.Name,
				ruleID,
				"aws_ses_receipt_rule",
				"aws",
				map[string]string{
					"name":          *rule.Name,
					"rule_set_name": ruleSetName,
				},
				sesAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return nil
}
