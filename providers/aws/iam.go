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

package aws

import (
	"fmt"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/service/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var IamAllowEmptyValues = []string{"tags."}

var IamAdditionalFields = map[string]string{}

type IamGenerator struct {
	AWSService
}

func (g *IamGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := iam.New(sess)
	g.Resources = []terraform_utils.Resource{}
	err := g.getUsers(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getGroups(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getPolicies(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getRoles(svc)
	if err != nil {
		log.Println(err)
	}

	g.PopulateIgnoreKeys()
	return nil
}

func (g *IamGenerator) getRoles(svc *iam.IAM) error {
	err := svc.ListRolesPages(&iam.ListRolesInput{}, func(roles *iam.ListRolesOutput, lastPages bool) bool {
		for _, role := range roles.Roles {
			roleID := aws.StringValue(role.RoleId)
			roleName := aws.StringValue(role.RoleName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				roleID,
				roleName,
				"aws_iam_role",
				"aws",
				map[string]string{},
				IamAllowEmptyValues,
				IamAdditionalFields))
			listRolePolicies, err := svc.ListRolePolicies(&iam.ListRolePoliciesInput{RoleName: role.RoleName})
			if err != nil {
				log.Println(err)
				continue
			}
			for _, policyName := range listRolePolicies.PolicyNames {
				g.Resources = append(g.Resources, terraform_utils.NewResource(
					roleName+":"+aws.StringValue(policyName),
					roleName+"_"+aws.StringValue(policyName),
					"aws_iam_role_policy",
					"aws",
					map[string]string{},
					IamAllowEmptyValues,
					IamAdditionalFields))
			}
		}
		return !lastPages
	})
	return err
}

func (g *IamGenerator) getUsers(svc *iam.IAM) error {
	err := svc.ListUsersPages(&iam.ListUsersInput{}, func(users *iam.ListUsersOutput, lastPage bool) bool {
		for _, user := range users.Users {
			resourceName := aws.StringValue(user.UserName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				aws.StringValue(user.UserId),
				"aws_iam_user",
				"aws",
				map[string]string{
					"force_destroy": "false",
				},
				IamAllowEmptyValues,
				IamAdditionalFields))
			g.getUserPolices(svc, user.UserName)
			//g.getUserGroup(svc, user.UserName) //not work maybe terraform-aws bug
		}

		return !lastPage
	})
	return err
}

func (g *IamGenerator) getUserGroup(svc *iam.IAM, userName *string) error {
	err := svc.ListGroupsForUserPages(&iam.ListGroupsForUserInput{UserName: userName}, func(userGroup *iam.ListGroupsForUserOutput, lastPage bool) bool {
		for _, group := range userGroup.Groups {
			resourceName := aws.StringValue(group.GroupName)
			groupIDAttachment := aws.StringValue(group.GroupName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				groupIDAttachment,
				resourceName,
				"aws_iam_user_group_membership",
				"aws",
				map[string]string{"user": aws.StringValue(userName)},
				IamAllowEmptyValues,
				IamAdditionalFields,
			))
		}
		return !lastPage
	})
	return err
}

func (g *IamGenerator) getUserPolices(svc *iam.IAM, userName *string) error {
	err := svc.ListUserPoliciesPages(&iam.ListUserPoliciesInput{UserName: userName}, func(userPolices *iam.ListUserPoliciesOutput, lastPage bool) bool {
		for _, policy := range userPolices.PolicyNames {
			resourceName := aws.StringValue(userName) + "_" + aws.StringValue(policy)
			resourceName = strings.Replace(resourceName, "@", "", -1)
			policyID := aws.StringValue(userName) + ":" + aws.StringValue(policy)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				policyID,
				resourceName,
				"aws_iam_user_policy",
				"aws",
				map[string]string{},
				IamAllowEmptyValues,
				IamAdditionalFields))
		}
		return !lastPage
	})
	return err
}

func (g *IamGenerator) getPolicies(svc *iam.IAM) error {
	err := svc.ListPoliciesPages(&iam.ListPoliciesInput{}, func(policies *iam.ListPoliciesOutput, lastPage bool) bool {
		for _, policy := range policies.Policies {
			resourceName := aws.StringValue(policy.PolicyName)
			policyARN := aws.StringValue(policy.Arn)

			g.Resources = append(g.Resources, terraform_utils.NewResource(
				policyARN,
				resourceName,
				"aws_iam_policy_attachment",
				"aws",
				map[string]string{
					"policy_arn": policyARN,
					"name":       resourceName,
				},
				IamAllowEmptyValues,
				IamAdditionalFields))
			// not use AWS main policy
			if strings.HasPrefix(policyARN, "arn:aws:iam::aws:policy") {
				continue
			}
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				policyARN,
				resourceName,
				"aws_iam_policy",
				"aws",
				map[string]string{},
				IamAllowEmptyValues,
				IamAdditionalFields))

		}
		return !lastPage
	})
	return err
}

func (g *IamGenerator) getGroups(svc *iam.IAM) error {
	err := svc.ListGroupsPages(&iam.ListGroupsInput{}, func(groups *iam.ListGroupsOutput, lastPage bool) bool {
		for _, group := range groups.Groups {
			resourceName := aws.StringValue(group.GroupName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_iam_group",
				"aws",
				map[string]string{},
				IamAllowEmptyValues,
				IamAdditionalFields))
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_iam_group_membership",
				"aws",
				map[string]string{
					"group": resourceName,
					"name":  resourceName,
				},
				IamAllowEmptyValues,
				IamAdditionalFields))
			_ = svc.ListGroupPoliciesPages(&iam.ListGroupPoliciesInput{GroupName: group.GroupName}, func(policyGroup *iam.ListGroupPoliciesOutput, lastPage bool) bool {
				for _, policy := range policyGroup.PolicyNames {
					id := resourceName + ":" + aws.StringValue(policy)
					groupPolicyName := resourceName + "_" + aws.StringValue(policy)
					g.Resources = append(g.Resources, terraform_utils.NewResource(
						id,
						groupPolicyName,
						"aws_iam_group_policy",
						"aws",
						map[string]string{},
						IamAllowEmptyValues,
						IamAdditionalFields))
				}
				return !lastPage
			})

		}
		return !lastPage
	})
	return err
}

// PostGenerateHook for add policy json as heredoc
func (g *IamGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_iam_policy" ||
			resource.InstanceInfo.Type == "aws_iam_user_policy" ||
			resource.InstanceInfo.Type == "aws_iam_group_policy" ||
			resource.InstanceInfo.Type == "aws_iam_role_policy" {
			policy := resource.Item["policy"].(string)
			resource.Item["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
		} else if resource.InstanceInfo.Type == "aws_iam_role" {
			policy := resource.Item["assume_role_policy"].(string)
			g.Resources[i].Item["assume_role_policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
		}
	}
	return nil
}
