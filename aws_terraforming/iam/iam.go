package iam

import (
	"strings"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/service/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var ignoreKey = map[string]bool{
	"^id$":       true,
	"^arn":       true,
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
	/*svc.ListGroupsPages()
	svc.ListPoliciesPages()
	svc.ListRolesPages()
	svc.ListAccessKeysPages()
	svc.ListGroupPoliciesPages()
	svc.ListGroupsPages()*/
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
			g.getUserGroup(svc, user.UserName)
		}

		return !lastPage
	})
	return err
}

func (g *IamGenerator) getUserGroup(svc *iam.IAM, userName *string) {
	svc.ListGroupsForUserPages(&iam.ListGroupsForUserInput{UserName: userName}, func(userGroup *iam.ListGroupsForUserOutput, lastPage bool) bool {
		for _, group := range userGroup.Groups {
			resourceName := aws.StringValue(group.GroupName)
			groupIDAttachment := aws.StringValue(group.GroupName)
			g.resources = append(g.resources, terraform_utils.NewTerraformResource(
				groupIDAttachment,
				resourceName,
				"aws_iam_user_group_membership",
				"aws",
				map[string]interface{}{"user": userName},
				map[string]string{},
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
}

func (g *IamGenerator) getUserPolices(svc *iam.IAM, userName *string) {
	svc.ListUserPoliciesPages(&iam.ListUserPoliciesInput{UserName: userName}, func(userPolices *iam.ListUserPoliciesOutput, lastPage bool) bool {
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
}
