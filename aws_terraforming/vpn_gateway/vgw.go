package vpn_gateway

import (
	"waze/terraform/aws_terraforming/awsGenerator"
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"id": true,
}
var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type VpnGatewayGenerator struct {
	awsGenerator.BasicGenerator
}

func (VpnGatewayGenerator) createResources(vpnGws *ec2.DescribeVpnGatewaysOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, vpnGw := range vpnGws.VpnGateways {
		resourceName := ""
		if len(vpnGw.Tags) > 0 {
			for _, tag := range vpnGw.Tags {
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
			ID:           aws.StringValue(vpnGw.VpnGatewayId),
			Provider:     "aws",
		})
	}
	return resoures
}

func (g VpnGatewayGenerator) Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	vpnGws, err := svc.DescribeVpnGateways(&ec2.DescribeVpnGatewaysInput{})
	if err != nil {
		return err
	}
	resources := g.createResources(vpnGws)
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
	err = terraform_utils.GenerateTf(resources, "vpn_gateway", region)
	if err != nil {
		return err
	}
	return nil

}
