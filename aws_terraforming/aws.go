package awsTerraforming

import (
	"log"
	"os"
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
	switch service {
	case "vpc":
		err := vpc.Generate(region)
		if err != nil {
			log.Println(err)
			return
		}
	case "sg":
		err := sg.Generate(region)
		if err != nil {
			log.Println(err)
			return
		}
	case "subnet":
		err := subnet.Generate(region)
		if err != nil {
			log.Println(err)
			return
		}
	case "igw":
		err := igw.Generate(region)
		if err != nil {
			log.Println(err)
			return
		}
	case "vpn_gateway":
		err := vpn_gateway.Generate(region)
		if err != nil {
			log.Println(err)
			return
		}
	case "nacl":
		err := nacl.Generate(region)
		if err != nil {
			log.Println(err)
			return

		}
	case "vpn_connetions":
		err := vpn_connection.Generate(region)
		if err != nil {
			log.Println(err)
			return

		}
	case "s3":
		err := s3.Generate(region)
		if err != nil {
			log.Println(err)
			return

		}

	}

}
