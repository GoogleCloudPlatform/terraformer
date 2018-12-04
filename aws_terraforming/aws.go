package aws_terraforming

import (
	"log"
	"os"

	"waze/terraform/aws_terraforming/aws_generator"
	"waze/terraform/aws_terraforming/elb"
	"waze/terraform/aws_terraforming/iam"
	"waze/terraform/aws_terraforming/igw"
	"waze/terraform/aws_terraforming/nacl"
	"waze/terraform/aws_terraforming/route53"
	"waze/terraform/aws_terraforming/s3"
	"waze/terraform/aws_terraforming/sg"
	"waze/terraform/aws_terraforming/subnet"
	"waze/terraform/aws_terraforming/vpc"
	"waze/terraform/aws_terraforming/vpn_connection"
	"waze/terraform/aws_terraforming/vpn_gateway"
	"waze/terraform/terraform_utils"

	"github.com/pkg/errors"
)

const PathForGenerateFiles = "/generated/aws/"

// GetAWSSupportService return map of support service for AWS
func GetAWSSupportService() map[string]aws_generator.Generator {
	return map[string]aws_generator.Generator{
		"vpc":            vpc.VpcGenerator{},
		"sg":             sg.SecurityGenerator{},
		"subnet":         subnet.SubnetGenerator{},
		"igw":            igw.IgwGenerator{},
		"vpn_gateway":    vpn_gateway.VpnGatewayGenerator{},
		"nacl":           nacl.NaclGenerator{},
		"vpn_connection": vpn_connection.VpnConnectionGenerator{},
		"s3":             s3.S3Generator{},
		"elb":            elb.ElbGenerator{},
		"iam":            iam.IamGenerator{},
		"route53":        route53.Route53Generator{},
	}
}

// Main function for generate tf and tfstate file by AWS service and region
func Generate(service string, args []string) error {
	region := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + PathForGenerateFiles + region + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		log.Print(err)
		return err
	}
	// change current dir for terraform refresh
	if err := os.Chdir(currentPath); err != nil {
		log.Print(err)
		return err
	}
	// terraform work with env param AWS_DEFAULT_REGION
	// save old AWS_DEFAULT_REGION
	oldRegion := os.Getenv("AWS_DEFAULT_REGION")
	// return AWS_DEFAULT_REGION
	defer os.Setenv("AWS_DEFAULT_REGION", oldRegion)
	// set env param AWS_DEFAULT_REGION for terraform
	os.Setenv("AWS_DEFAULT_REGION", region)
	// return current dir after terraform refresh run
	defer os.Chdir(rootPath)
	var generator aws_generator.Generator
	var isSupported bool
	if generator, isSupported = GetAWSSupportService()[service]; !isSupported {
		return errors.New("aws: not supported service")
	}
	// generate TerraformResources with type and ids + metadata
	resources, metadata, err := generator.Generate(region)
	if err != nil {
		return err
	}
	// generate empty(resource and ids) tfstate,
	// and run terraform refresh with empty tfstate for populate data
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	// convert tfstate to go struct for hcl print
	converter := terraform_utils.TfstateConverter{}
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	// change structs with additional data for each resource
	resources, err = generator.PostGenerateHook(resources)
	// print HCL file
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
