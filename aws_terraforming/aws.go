package awsTerraforming

import (
	"log"
	"os"
	"strings"
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
)

const pathForGenerateFiles = "/generated/aws/"

func Generate(service, zone string) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	rootPath, _ := os.Getwd()
	currentPath := rootPath + pathForGenerateFiles + region + "/" + service
	os.MkdirAll(currentPath, os.ModePerm)
	os.Chdir(currentPath)
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
	case "vpn_connections":
		generator = vpn_connection.VpnConnectionGenerator{}
	case "s3":
		generator = s3.S3Generator{}
	case "elb":
		generator = elb.ElbGenerator{}
	case "iam":
		generator = iam.IamGenerator{}
	}
	err := generator.Generate(region)
	if err != nil {
		log.Println(err)
		return
	}

}
