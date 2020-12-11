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
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var IamAllowEmptyValues = []string{"tags."}

var IamAdditionalFields = map[string]interface{}{}

type IamGenerator struct {
	AWSService
}

func (g *IamGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := iam.New(config)
	g.Resources = []terraformutils.Resource{}
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

	err = g.getInstanceProfiles(svc)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (g *IamGenerator) getRoles(svc *iam.Client) error {
	p := iam.NewListRolesPaginator(svc.ListRolesRequest(&iam.ListRolesInput{}))
	for p.Next(context.Background()) {
		for _, role := range p.CurrentPage().Roles {
			roleName := aws.StringValue(role.RoleName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				roleName,
				roleName,
				"aws_iam_role",
				"aws",
				IamAllowEmptyValues))
			rolePoliciesPage := iam.NewListRolePoliciesPaginator(svc.ListRolePoliciesRequest(&iam.ListRolePoliciesInput{RoleName: role.RoleName}))
			for rolePoliciesPage.Next(context.Background()) {
				for _, policyName := range rolePoliciesPage.CurrentPage().PolicyNames {
					g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
						roleName+":"+policyName,
						roleName+"_"+policyName,
						"aws_iam_role_policy",
						"aws",
						IamAllowEmptyValues))
				}
			}
			if err := rolePoliciesPage.Err(); err != nil {
				log.Println(err)
				continue
			}
			roleAttachedPoliciesPage := iam.NewListAttachedRolePoliciesPaginator(svc.ListAttachedRolePoliciesRequest(&iam.ListAttachedRolePoliciesInput{
				RoleName: &roleName,
			}))
			for roleAttachedPoliciesPage.Next(context.Background()) {
				for _, attachedPolicy := range roleAttachedPoliciesPage.CurrentPage().AttachedPolicies {
					g.Resources = append(g.Resources, terraformutils.NewResource(
						roleName+"/"+*attachedPolicy.PolicyArn,
						roleName+"_"+*attachedPolicy.PolicyName,
						"aws_iam_role_policy_attachment",
						"aws",
						map[string]string{
							"role":       roleName,
							"policy_arn": *attachedPolicy.PolicyArn,
						},
						IamAllowEmptyValues,
						map[string]interface{}{}))
				}
			}
			if err := roleAttachedPoliciesPage.Err(); err != nil {
				log.Println(err)
				continue
			}
		}
	}
	return p.Err()
}

func (g *IamGenerator) getUsers(svc *iam.Client) error {
	p := iam.NewListUsersPaginator(svc.ListUsersRequest(&iam.ListUsersInput{}))
	for p.Next(context.Background()) {
		for _, user := range p.CurrentPage().Users {
			resourceName := aws.StringValue(user.UserName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceName,
				aws.StringValue(user.UserId),
				"aws_iam_user",
				"aws",
				map[string]string{
					"force_destroy": "false",
				},
				IamAllowEmptyValues,
				map[string]interface{}{}))
			err := g.getUserPolices(svc, user.UserName)
			if err != nil {
				log.Println(err)
			}
			err = g.getUserPolicyAttachment(svc, user.UserName)
			if err != nil {
				log.Println(err)
			}
			err = g.getUserGroup(svc, user.UserName)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return p.Err()
}

func (g *IamGenerator) getUserGroup(svc *iam.Client, userName *string) error {
	p := iam.NewListGroupsForUserPaginator(svc.ListGroupsForUserRequest(&iam.ListGroupsForUserInput{UserName: userName}))
	for p.Next(context.Background()) {
		for _, group := range p.CurrentPage().Groups {
			userGroupMembership := *userName + "/" + *group.GroupName
			g.Resources = append(g.Resources, terraformutils.NewResource(
				userGroupMembership,
				userGroupMembership,
				"aws_iam_user_group_membership",
				"aws",
				map[string]string{
					"user":     *userName,
					"groups.#": "1",
					"groups.0": *group.GroupName,
				},
				IamAllowEmptyValues,
				IamAdditionalFields,
			))
		}
	}
	return p.Err()
}

func (g *IamGenerator) getUserPolices(svc *iam.Client, userName *string) error {
	p := iam.NewListUserPoliciesPaginator(svc.ListUserPoliciesRequest(&iam.ListUserPoliciesInput{UserName: userName}))
	for p.Next(context.Background()) {
		for _, policy := range p.CurrentPage().PolicyNames {
			resourceName := aws.StringValue(userName) + "_" + policy
			resourceName = strings.ReplaceAll(resourceName, "@", "")
			policyID := aws.StringValue(userName) + ":" + policy
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				policyID,
				resourceName,
				"aws_iam_user_policy",
				"aws",
				IamAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *IamGenerator) getUserPolicyAttachment(svc *iam.Client, userName *string) error {
	p := iam.NewListAttachedUserPoliciesPaginator(svc.ListAttachedUserPoliciesRequest(&iam.ListAttachedUserPoliciesInput{
		UserName: userName,
	}))
	for p.Next(context.Background()) {
		for _, attachedPolicy := range p.CurrentPage().AttachedPolicies {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*userName+"/"+*attachedPolicy.PolicyArn,
				*userName+"_"+*attachedPolicy.PolicyName,
				"aws_iam_user_policy_attachment",
				"aws",
				map[string]string{
					"user":       *userName,
					"policy_arn": *attachedPolicy.PolicyArn,
				},
				IamAllowEmptyValues,
				map[string]interface{}{}))
		}
	}
	return p.Err()
}

func (g *IamGenerator) getPolicies(svc *iam.Client) error {
	p := iam.NewListPoliciesPaginator(svc.ListPoliciesRequest(&iam.ListPoliciesInput{Scope: iam.PolicyScopeTypeLocal}))
	for p.Next(context.Background()) {
		for _, policy := range p.CurrentPage().Policies {
			resourceName := aws.StringValue(policy.PolicyName)
			policyARN := aws.StringValue(policy.Arn)

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				policyARN,
				resourceName,
				"aws_iam_policy",
				"aws",
				IamAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *IamGenerator) getGroups(svc *iam.Client) error {
	p := iam.NewListGroupsPaginator(svc.ListGroupsRequest(&iam.ListGroupsInput{}))
	for p.Next(context.Background()) {
		for _, group := range p.CurrentPage().Groups {
			resourceName := aws.StringValue(group.GroupName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_iam_group",
				"aws",
				IamAllowEmptyValues))
			g.getGroupPolicies(svc, group)
			g.getAttachedGroupPolicies(svc, group)
		}
	}
	return p.Err()
}

func (g *IamGenerator) getGroupPolicies(svc *iam.Client, group iam.Group) {
	groupPoliciesPage := iam.NewListGroupPoliciesPaginator(svc.ListGroupPoliciesRequest(&iam.ListGroupPoliciesInput{GroupName: group.GroupName}))
	for groupPoliciesPage.Next(context.Background()) {
		for _, policy := range groupPoliciesPage.CurrentPage().PolicyNames {
			id := *group.GroupName + ":" + policy
			groupPolicyName := *group.GroupName + "_" + policy
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				groupPolicyName,
				"aws_iam_group_policy",
				"aws",
				map[string]string{},
				IamAllowEmptyValues,
				IamAdditionalFields))
		}
	}
	if err := groupPoliciesPage.Err(); err != nil {
		log.Println(err)
	}
}

func (g *IamGenerator) getAttachedGroupPolicies(svc *iam.Client, group iam.Group) {
	groupAttachedPoliciesPage := iam.NewListAttachedGroupPoliciesPaginator(svc.ListAttachedGroupPoliciesRequest(
		&iam.ListAttachedGroupPoliciesInput{GroupName: group.GroupName}))
	for groupAttachedPoliciesPage.Next(context.Background()) {
		for _, attachedPolicy := range groupAttachedPoliciesPage.CurrentPage().AttachedPolicies {
			if !strings.Contains(*attachedPolicy.PolicyArn, "arn:aws:iam::aws") {
				continue // map only AWS managed policies since others should be managed by
			}
			id := *group.GroupName + "/" + *attachedPolicy.PolicyArn
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				*group.GroupName+"_"+*attachedPolicy.PolicyName,
				"aws_iam_group_policy_attachment",
				"aws",
				map[string]string{
					"group":      *group.GroupName,
					"policy_arn": *attachedPolicy.PolicyArn,
				},
				IamAllowEmptyValues,
				IamAdditionalFields))
		}
	}
	if err := groupAttachedPoliciesPage.Err(); err != nil {
		log.Println(err)
	}
}

func (g *IamGenerator) getInstanceProfiles(svc *iam.Client) error {
	p := iam.NewListInstanceProfilesPaginator(svc.ListInstanceProfilesRequest(&iam.ListInstanceProfilesInput{}))
	for p.Next(context.Background()) {
		for _, instanceProfile := range p.CurrentPage().InstanceProfiles {
			resourceName := *instanceProfile.InstanceProfileName

			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceName,
				resourceName,
				"aws_iam_instance_profile",
				"aws",
				map[string]string{
					"name": resourceName,
				},
				IamAllowEmptyValues,
				IamAdditionalFields))
		}
	}
	return p.Err()
}

// PostGenerateHook for add policy json as heredoc
func (g *IamGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		switch {
		case resource.InstanceInfo.Type == "aws_iam_policy" ||
			resource.InstanceInfo.Type == "aws_iam_user_policy" ||
			resource.InstanceInfo.Type == "aws_iam_group_policy" ||
			resource.InstanceInfo.Type == "aws_iam_role_policy":
			policy := g.escapeAwsInterpolation(resource.Item["policy"].(string))
			resource.Item["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
		case resource.InstanceInfo.Type == "aws_iam_role":
			policy := g.escapeAwsInterpolation(resource.Item["assume_role_policy"].(string))
			g.Resources[i].Item["assume_role_policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
		case resource.InstanceInfo.Type == "aws_iam_instance_profile":
			delete(resource.Item, "roles")
		}
	}
	return nil
}
