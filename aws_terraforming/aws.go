package awsTerraforming

import (
	"log"
	"os"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/aws_terraforming/elb"
	"waze/terraform/aws_terraforming/iam"
	"waze/terraform/aws_terraforming/igw"
	"waze/terraform/aws_terraforming/nacl"
	"waze/terraform/aws_terraforming/s3"
	"waze/terraform/aws_terraforming/sg"
	"waze/terraform/aws_terraforming/subnet"
	"waze/terraform/aws_terraforming/vpc"
	"waze/terraform/aws_terraforming/vpn_connection"
	"waze/terraform/aws_terraforming/vpn_gateway"
	"waze/terraform/terraform_utils"
)

const PathForGenerateFiles = "/generated/aws/"

func Generate(service string, args []string) error {
	region := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + PathForGenerateFiles + region + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		log.Print(err)
		return err
	}
	if err := os.Chdir(currentPath); err != nil {
		log.Print(err)
		return err
	}
	oldRegion := os.Getenv("AWS_DEFAULT_REGION")
	defer os.Setenv("AWS_DEFAULT_REGION", oldRegion)
	os.Setenv("AWS_DEFAULT_REGION", region)
	defer os.Chdir(rootPath)
	var generator aws_generator.Generator
	switch service {
	case "vpc":
		generator = vpc.VpcGenerator{}
	case "sg":
		generator = sg.SecurityGenerator{}
	case "subnet":
		generator = subnet.SubnetGenerator{}
	case "igw":
		generator = igw.IgwGenerator{}
	case "vpn_gateway":
		generator = vpn_gateway.VpnGatewayGenerator{}
	case "nacl":
		generator = nacl.NaclGenerator{}
	case "vpn_connection":
		generator = vpn_connection.VpnConnectionGenerator{}
	case "s3":
		generator = s3.S3Generator{}
	case "elb":
		generator = elb.ElbGenerator{}
	case "iam":
		generator = iam.IamGenerator{}
	}
	resources, metadata, err := generator.Generate(region)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	resources, err = generator.PostGenerateHook(resources)
	err = terraform_utils.GenerateTf(resources, service, NewAwsRegionResource(region))
	if err != nil {
		return err
	}
	return nil

}

func NewAwsRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"aws": map[string]interface{}{
			"region": region,
		},
	}
}
