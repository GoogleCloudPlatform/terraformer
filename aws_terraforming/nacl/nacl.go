package nacl

import (
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"id": true,
}

func createResources(nacls *ec2.DescribeNetworkAclsOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, nacl := range nacls.NetworkAcls {
		resourceName := ""
		if len(nacl.Tags) > 0 {
			for _, tag := range nacl.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.TerraformResource{
			ResourceType: "aws_network_acl",
			ResourceName: resourceName,
			Item:         nil,
			ID:           aws.StringValue(nacl.NetworkAclId),
			Provider:     "aws",
		})
	}
	return resoures
}

func Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	nacls, err := svc.DescribeNetworkAcls(&ec2.DescribeNetworkAclsInput{})
	if err != nil {
		return err
	}
	resources := createResources(nacls)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	resources, err = terraform_utils.TfstateToTfConverter("terraform.tfstate", "aws", ignoreKey)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "nacl", region)
	if err != nil {
		return err
	}
	return nil

}
