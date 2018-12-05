package iam

import (
	"fmt"
	"strings"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/service/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var ignoreKey = map[string]bool{
	"^id$":       true,
	"^arn$":      true,
	"^unique_id": true,
}

var additionalFields = map[string]string{
	"force_destroy": "false",
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type IamGenerator struct {
	aws_generator.BasicGenerator
	metadata  map[string]terraform_utils.ResourceMetaData
	resources []terraform_utils.TerraformResource
}

// TODO All here
func (g IamGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := iam.New(sess)
	g.resources = []terraform_utils.TerraformResource{}
	g.metadata = map[string]terraform_utils.ResourceMetaData{}
	err := g.getUsers(svc)
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	//TODO ALL
	err = g.getGroups(svc)
	err = g.getPolicies(svc)
	/*
		svc.ListRolesPages()
		svc.ListAccessKeysPages()
		svc.ListGroupPoliciesPages()*/
	return g.resources, g.metadata, nil
}

func (g *IamGenerator) getUsers(svc *iam.IAM) error {
	err := svc.ListUsersPages(&iam.ListUsersInput{}, func(users *iam.ListUsersOutput, lastPage bool) bool {
		for _, user := range users.Users {
			resourceName := aws.StringValue(user.UserName)
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				resourceName,
				aws.StringValue(user.UserId),
				"aws_iam_user",
				"aws",
				nil,
				map[string]string{}))
			g.metadata[resourceName] = terraform_utils.ResourceMetaData{
				Provider:         "aws",
				IgnoreKeys:       ignoreKey,
				AllowEmptyValue:  allowEmptyValues,
				AdditionalFields: additionalFields,
			}

			g.getUserPolices(svc, user.UserName)
			//g.getUserGroup(svc, user.UserName) //not work maybe terraform-aws bug
			break
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
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				groupIDAttachment,
				resourceName,
				"aws_iam_user_group_membership",
				"aws",
				map[string]string{},
				map[string]string{"user": aws.StringValue(userName)},
			))
			g.metadata[groupIDAttachment] = terraform_utils.ResourceMetaData{
				Provider:         "aws",
				IgnoreKeys:       ignoreKey,
				AllowEmptyValue:  allowEmptyValues,
				AdditionalFields: map[string]string{},
			}
		}
		return !lastPage
	})
	return err
}

func (g *IamGenerator) getUserPolices(svc *iam.IAM, userName *string) error {
	err := svc.ListUserPoliciesPages(&iam.ListUserPoliciesInput{UserName: userName}, func(userPolices *iam.ListUserPoliciesOutput, lastPage bool) bool {
		for _, policy := range userPolices.PolicyNames {
			resourceName := strings.Replace(aws.StringValue(policy), "@", "", -1)
			policyID := aws.StringValue(userName) + ":" + aws.StringValue(policy)
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				policyID,
				resourceName,
				"aws_iam_user_policy",
				"aws",
				nil,
				map[string]string{}))
			g.metadata[policyID] = terraform_utils.ResourceMetaData{
				Provider:         "aws",
				IgnoreKeys:       ignoreKey,
				AllowEmptyValue:  allowEmptyValues,
				AdditionalFields: map[string]string{},
			}

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
			// not use AWS main policy
			//if strings.HasPrefix(policyARN, "arn:aws:iam::aws:policy") {
			//	continue
			//}
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				policyARN,
				resourceName,
				"aws_iam_policy",
				"aws",
				nil,
				map[string]string{}))
			g.metadata[resourceName] = terraform_utils.ResourceMetaData{
				Provider:         "aws",
				IgnoreKeys:       ignoreKey,
				AllowEmptyValue:  allowEmptyValues,
				AdditionalFields: map[string]string{},
			}
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				policyARN,
				resourceName,
				"aws_iam_policy_attachment",
				"aws",
				nil,
				map[string]string{
					"policy_arn": policyARN,
					"name":       resourceName,
				}))
			break

		}
		return !lastPage
	})
	return err
}

func (g *IamGenerator) getGroups(svc *iam.IAM) error {
	err := svc.ListGroupsPages(&iam.ListGroupsInput{}, func(groups *iam.ListGroupsOutput, lastPage bool) bool {
		for _, group := range groups.Groups {
			resourceName := aws.StringValue(group.GroupName)
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				resourceName,
				resourceName,
				"aws_iam_group",
				"aws",
				nil,
				map[string]string{}))
			g.metadata[resourceName] = terraform_utils.ResourceMetaData{
				Provider:         "aws",
				IgnoreKeys:       ignoreKey,
				AllowEmptyValue:  allowEmptyValues,
				AdditionalFields: map[string]string{},
			}
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				resourceName,
				resourceName,
				"aws_iam_group_membership",
				"aws",
				nil,
				map[string]string{
					"group": resourceName,
					"name":  resourceName,
				}))
		}
		return !lastPage
	})
	return err
}

// PostGenerateHook for add policy json as heredoc
func (IamGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for _, resource := range resources {
		if resource.ResourceType == "aws_iam_policy" || resource.ResourceType == "aws_iam_user_policy" {
			policy := resource.Item.(interface{}).(map[string]interface{})["policy"].(string)
			resource.Item.(interface{}).(map[string]interface{})["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
		}
	}
	return resources, nil
}
