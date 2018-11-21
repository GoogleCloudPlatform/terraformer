package elb

import (
	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

var ignoreKey = map[string]bool{
	"id":                       true,
	"arn":                      true,
	"dns_name":                 true,
	"source_security_group_id": true,
	"zone_id":                  true,
	"instances":                true, //dynamic value
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type ElbGenerator struct {
	aws_generator.BasicGenerator
}

func (ElbGenerator) Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := elb.New(sess)
	resources := []terraform_utils.TerraformResource{}
	err := svc.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, func(loadBalancers *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, loadBalancer := range loadBalancers.LoadBalancerDescriptions {
			resourceName := aws.StringValue(loadBalancer.LoadBalancerName)
			resources = append(resources, terraform_utils.NewTerraformResource(
				aws.StringValue(loadBalancer.LoadBalancerName),
				resourceName,
				"aws_elb",
				"aws",
				nil,
				map[string]string{}))
		}
		return !lastPage
	})
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, map[string]string{})
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "elb", region, "aws")
	if err != nil {
		return err
	}
	return nil

}
