package awsTerraforming

import (
	"log"
	"os"
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

func Generate(service, region string) {
	rootPath, _ := os.Getwd()
	currentPath := rootPath + pathForGenerateFiles + region + "/" + service
	os.MkdirAll(currentPath, os.ModePerm)
	os.Chdir(currentPath)
	defer os.Chdir(rootPath)
	var err error
	switch service {
	case "vpc":
		err = vpc.Generate(region)
	case "sg":
		err = sg.Generate(region)
	case "subnet":
		err = subnet.Generate(region)
	case "igw":
		err = igw.Generate(region)
	case "vpn_gateway":
		err = vpn_gateway.Generate(region)
	case "nacl":
		err = nacl.Generate(region)
	case "vpn_connections":
		err = vpn_connection.Generate(region)
	case "s3":
		err = s3.Generate(region)
	case "elb":
		err = elb.Generate(region)
	case "iam":
		err = iam.Generate(region)
	}

	if err != nil {
		log.Println(err)
		return
	}

}
