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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

var codecommitAllowEmptyValues = []string{"tags."}

type CodeCommitGenerator struct {
	AWSService
}

func (g *CodeCommitGenerator) loadRepository(svc *codecommit.Client) error {
	p := codecommit.NewListRepositoriesPaginator(svc, &codecommit.ListRepositoriesInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, repository := range page.Repositories {
			resourceName := StringValue(repository.RepositoryName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_codecommit_repository",
				"aws",
				codecommitAllowEmptyValues))
		}
	}
	return nil
}

func (g *CodeCommitGenerator) loadApprovalRuleTemplate(svc *codecommit.Client) error {
	p := codecommit.NewListApprovalRuleTemplatesPaginator(svc, &codecommit.ListApprovalRuleTemplatesInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, templateName := range page.ApprovalRuleTemplateNames {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				templateName,
				templateName,
				"aws_codecommit_approval_rule_template",
				"aws",
				codecommitAllowEmptyValues))
		}
	}
	return nil
}

func (g *CodeCommitGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := codecommit.NewFromConfig(config)
	err := g.loadRepository(svc)
	if err != nil {
		return err
	}
	err = g.loadApprovalRuleTemplate(svc)
	if err != nil {
		return err
	}

	return nil
}

func (g *CodeCommitGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_codecommit_approval_rule_template" {
			if content, ok := g.Resources[i].Item["content"]; ok {
				g.Resources[i].Item["content"] = fmt.Sprintf(`<<CONTENT
%s
CONTENT`, content)
			}
		}
	}
	return nil
}
