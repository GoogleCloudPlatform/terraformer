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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
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
	svc := cloudformation.New(config)
	p := cloudformation.NewListStacksPaginator(svc.ListStacksRequest(&cloudformation.ListStacksInput{}))
	for p.Next(context.Background()) {
		for _, stackSummary := range p.CurrentPage().StackSummaries {
			if stackSummary.StackStatus == cloudformation.StackStatusDeleteComplete {
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
	if err := p.Err(); err != nil {
		return err
	}
	stackSets, err := svc.ListStackSetsRequest(&cloudformation.ListStackSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, stackSetSummary := range stackSets.Summaries {
		if stackSetSummary.Status == cloudformation.StackSetStatusDeleted {
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*stackSetSummary.StackSetId,
			*stackSetSummary.StackSetName,
			"aws_cloudformation_stack_set",
			"aws",
			cloudFormationAllowEmptyValues,
		))

		stackSetInstances, err := svc.ListStackInstancesRequest(&cloudformation.ListStackInstancesInput{
			StackSetName: stackSetSummary.StackSetName,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, stackSetI := range stackSetInstances.Summaries {
			id := aws.StringValue(stackSetI.StackSetId) + "," + aws.StringValue(stackSetI.Account) + "," + aws.StringValue(stackSetI.Region)

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
