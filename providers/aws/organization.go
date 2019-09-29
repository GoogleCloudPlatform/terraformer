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
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/organizations"
)

var organizationAllowEmptyValues = []string{"tags."}

type OrganizationGenerator struct {
	AWSService
}

func (g *OrganizationGenerator) traverseNode(svc *organizations.Organizations, parentId string) {
	accountsForParent, err := svc.ListAccountsForParent(
		&organizations.ListAccountsForParentInput{ParentId: aws.String(parentId)})
	if err != nil {
		return
	}
	for _, account := range accountsForParent.Accounts {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			aws.StringValue(account.Id),
			aws.StringValue(account.Name),
			"aws_organizations_organization",
			"aws",
			map[string]string{
				"id":  aws.StringValue(account.Id),
				"arn": aws.StringValue(account.Arn),
			},
			organizationAllowEmptyValues,
			map[string]interface{}{},
		))
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			aws.StringValue(account.Id),
			aws.StringValue(account.Name),
			"aws_organizations_account",
			"aws",
			map[string]string{
				"id":  aws.StringValue(account.Id),
				"arn": aws.StringValue(account.Arn),
			},
			organizationAllowEmptyValues,
			map[string]interface{}{},
		))
	}

	unitsForParent, err := svc.ListOrganizationalUnitsForParent(
		&organizations.ListOrganizationalUnitsForParentInput{ParentId: aws.String(parentId)})
	if err != nil {
		return
	}
	for _, unit := range unitsForParent.OrganizationalUnits {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			aws.StringValue(unit.Id),
			aws.StringValue(unit.Name),
			"aws_organizations_organizational_unit",
			"aws",
			map[string]string{
				"id":  aws.StringValue(unit.Id),
				"arn": aws.StringValue(unit.Arn),
			},
			organizationAllowEmptyValues,
			map[string]interface{}{},
		))
		g.traverseNode(svc, aws.StringValue(unit.Id))
	}

}

func (g *OrganizationGenerator) InitResources() error {
	sess := g.generateSession()
	svc := organizations.New(sess)

	roots, err := svc.ListRoots(&organizations.ListRootsInput{})
	if err != nil {
		return err
	}

	for _, root := range roots.Roots {
		nodeId := aws.StringValue(root.Id)
		g.traverseNode(svc, nodeId)
	}

	err = svc.ListPoliciesPages(&organizations.ListPoliciesInput{
		Filter: aws.String(organizations.PolicyTypeServiceControlPolicy),
	},
		func(policies *organizations.ListPoliciesOutput, lastPage bool) bool {
			for _, policy := range policies.Policies {

				policyId := aws.StringValue(policy.Id)
				policyName := aws.StringValue(policy.Name)
				g.Resources = append(g.Resources, terraform_utils.NewResource(
					policyId,
					policyName,
					"aws_organizations_policy",
					"aws",
					map[string]string{
						"id":  policyId,
						"arn": aws.StringValue(policy.Arn),
					},
					organizationAllowEmptyValues,
					map[string]interface{}{},

				))

				targetsForPolicy, err := svc.ListTargetsForPolicy(
					&organizations.ListTargetsForPolicyInput{PolicyId: policy.Id})
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				for _, target := range targetsForPolicy.Targets {
					g.Resources = append(g.Resources, terraform_utils.NewResource(
						aws.StringValue(target.TargetId)+":"+policyId,
						"pa-"+aws.StringValue(target.TargetId)+":"+policyName,
						"aws_organizations_policy_attachment",
						"aws",
						map[string]string{
							"policy_id": policyId,
							"target_id": aws.StringValue(target.TargetId),
						},
						organizationAllowEmptyValues,
						map[string]interface{}{},

					))
				}
			}

			return !lastPage
		})
	if err != nil {
		return err
	}

	return nil
}
