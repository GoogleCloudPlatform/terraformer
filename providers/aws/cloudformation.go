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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var cloudFormationAllowEmptyValues = []string{"tags."}

type CloudFormationGenerator struct {
	AWSService
}

func (g *CloudFormationGenerator) InitResources() error {
	sess := g.generateSession()
	svc := cloudformation.New(sess)

	err := svc.ListStacksPages(&cloudformation.ListStacksInput{}, func(stacks *cloudformation.ListStacksOutput, lastPage bool) bool {
		for _, stackSummary := range stacks.StackSummaries {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(stackSummary.StackId),
				aws.StringValue(stackSummary.StackName),
				"aws_cloudformation_stack",
				"aws",
				cloudFormationAllowEmptyValues,
			))
		}
		return !lastPage
	})
	if err != nil {
		return err
	}

	stackSets, err := svc.ListStackSets(&cloudformation.ListStackSetsInput{})
	if err != nil {
		return err
	}
	for _, stackSetSummary := range stackSets.Summaries {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			aws.StringValue(stackSetSummary.StackSetId),
			aws.StringValue(stackSetSummary.StackSetName),
			"aws_cloudformation_stack_set",
			"aws",
			cloudFormationAllowEmptyValues,
		))

		stackSetInstances, err := svc.ListStackInstances(&cloudformation.ListStackInstancesInput{
			StackSetName: stackSetSummary.StackSetName,
		})
		if err != nil {
			return err
		}
		for _, stackSetI := range stackSetInstances.Summaries {
			id := aws.StringValue(stackSetI.StackSetId) + "," + aws.StringValue(stackSetI.Account) + "," + aws.StringValue(stackSetI.Region)

			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
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
