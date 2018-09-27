package vpc

import (
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"arn":                       true,
	"main_route_table_id":       true,
	"id":                        true,
	"dhcp_options_id":           true,
	"default_security_group_id": true,
	"default_route_table_id":    true,
	"default_network_acl_id":    true,
}

func createResources(vpcs *ec2.DescribeVpcsOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, vpc := range vpcs.Vpcs {
		resourceName := ""
		if len(vpc.Tags) > 0 {
			for _, tag := range vpc.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.TerraformResource{
			ResourceType: "aws_vpc",
			ResourceName: resourceName,
			Item:         nil,
			ID:           aws.StringValue(vpc.VpcId),
			Provider:     "aws",
		})
	}
	return resoures
}

func Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	vpcs, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return err
	}
	resources := createResources(vpcs)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	resources, err = terraform_utils.TfstateToTfConverter("terraform.tfstate", "aws", ignoreKey)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "vpc", region)
	if err != nil {
		return err
	}
	return nil

}
