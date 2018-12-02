package subnet

import (
	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"^id$":  true,
	"^arn": true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type SubnetGenerator struct {
	aws_generator.BasicGenerator
}

func (SubnetGenerator) createResources(subnets *ec2.DescribeSubnetsOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, subnet := range subnets.Subnets {
		resourceName := ""
		if len(subnet.Tags) > 0 {
			for _, tag := range subnet.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.TerraformResource{
			ResourceType: "aws_subnet",
			ResourceName: resourceName,
			Item:         nil,
			ID:           aws.StringValue(subnet.SubnetId),
			Provider:     "aws",
		})
	}
	return resoures
}

func (g SubnetGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	subnets, err := svc.DescribeSubnets(&ec2.DescribeSubnetsInput{})
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	resources := g.createResources(subnets)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, map[string]string{})
	return resources, metadata, nil

}
