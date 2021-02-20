// Copyright 2019 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

var cloudFormationAllowEmptyValues = []string{"tags."}

type CloudFormationGenerator struct {
	AWSService
}

func (g *CloudFormationGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := cloudformation.NewFromConfig(config)
	p := cloudformation.NewListStacksPaginator(svc, &cloudformation.ListStacksInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, stackSummary := range page.StackSummaries {
			if stackSummary.StackStatus == types.StackStatusDeleteComplete {
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*stackSummary.StackId,
				*stackSummary.StackName,
				"aws_cloudformation_stack",
				"aws",
				cloudFormationAllowEmptyValues,
			))
		}
	}
	stackSets, err := svc.ListStackSets(context.TODO(), &cloudformation.ListStackSetsInput{})
	if err != nil {
		return err
	}
	for _, stackSetSummary := range stackSets.Summaries {
		if stackSetSummary.Status == types.StackSetStatusDeleted {
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*stackSetSummary.StackSetId,
			*stackSetSummary.StackSetName,
			"aws_cloudformation_stack_set",
			"aws",
			cloudFormationAllowEmptyValues,
		))

		stackSetInstances, err := svc.ListStackInstances(context.TODO(), &cloudformation.ListStackInstancesInput{
			StackSetName: stackSetSummary.StackSetName,
		})
		if err != nil {
			return err
		}
		for _, stackSetI := range stackSetInstances.Summaries {
			id := StringValue(stackSetI.StackSetId) + "," + StringValue(stackSetI.Account) + "," + StringValue(stackSetI.Region)

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				id,
				"aws_cloudformation_stack_set_instance",
				"aws",
				cloudFormationAllowEmptyValues,
			))
		}
	}

	return nil
}

func (g *CloudFormationGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_cloudformation_stack" {
			delete(resource.Item, "outputs")
			if templateBody, ok := resource.InstanceState.Attributes["template_body"]; ok {
				resource.Item["template_body"] = g.escapeAwsInterpolation(templateBody)
			}
		}
	}
	return nil
}
