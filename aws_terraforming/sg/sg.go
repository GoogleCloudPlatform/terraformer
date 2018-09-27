package sg

import (
	"strings"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxResults = 1000

var ignoreKey = map[string]bool{
	"arn":      true,
	"owner_id": true,
	"id":       true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

func createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.TerraformResource{
			ResourceType: "aws_security_group",
			ResourceName: strings.Trim(aws.StringValue(sg.GroupName), " "),
			Item:         nil,
			ID:           aws.StringValue(sg.GroupId),
			Provider:     "aws",
		})
	}
	return resources
}

func Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	var securityGroups []*ec2.SecurityGroup
	var err error
	firstRun := true
	securityGroupsOutput := &ec2.DescribeSecurityGroupsOutput{}
	for {
		if firstRun || securityGroupsOutput.NextToken != nil {
			firstRun = false
			securityGroupsOutput, err = svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
				MaxResults: aws.Int64(maxResults),
				NextToken:  securityGroupsOutput.NextToken,
			})
			securityGroups = append(securityGroups, securityGroupsOutput.SecurityGroups...)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}
	resources := createResources(securityGroups)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{
		Provider:        "aws",
		IgnoreKeys:      ignoreKey,
		AllowEmptyValue: allowEmptyValues,
	}
	resources, err = converter.Convert("terraform.tfstate")
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "security_group", region)
	if err != nil {
		return err
	}
	return nil
}
