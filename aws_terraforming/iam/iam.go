package iam

import (
	"strings"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/service/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var ignoreKey = map[string]bool{
	"id":        true,
	"arn":       true,
	"unique_id": true,
}

//var additionalFields = map[string]string{
//	"force_destroy": "false",
//}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

func Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := iam.New(sess)
	resources := []terraform_utils.TerraformResource{}
	err := svc.ListUsersPages(&iam.ListUsersInput{}, func(users *iam.ListUsersOutput, lastPage bool) bool {
		for _, user := range users.Users {
			resourceName := aws.StringValue(user.UserName)
			resources = append(resources, terraform_utils.NewTerraformResource(
				resourceName,
				aws.StringValue(user.UserId),
				"aws_iam_user",
				"aws",
				nil))
			svc.ListUserPoliciesPages(&iam.ListUserPoliciesInput{UserName: user.UserName}, func(userPolices *iam.ListUserPoliciesOutput, lastPage bool) bool {
				for _, policy := range userPolices.PolicyNames {
					resourceName := strings.Replace(aws.StringValue(policy), "@", "", -1)
					policyID := aws.StringValue(user.UserName) + ":" + aws.StringValue(policy)
					resources = append(resources, terraform_utils.NewTerraformResource(
						policyID,
						resourceName,
						"aws_iam_user_policy",
						"aws",
						nil))
				}
				return !lastPage
			})
		}

		return !lastPage
	})

	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{
		Provider:        "aws",
		IgnoreKeys:      ignoreKey,
		AllowEmptyValue: allowEmptyValues,
		//AdditionalFields: additionalFields,
	}
	resources, err = converter.Convert("terraform.tfstate")
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "iam", region)
	if err != nil {
		return err
	}
	return nil

}
