package main

import (
	"os"
	"waze/terraform/aws_terraforming"
	"waze/terraform/gcp_terraforming"
)

var (
	//cloud   = "aws"
	//service = "iam"
	region = "eu-west-1"
)

func main() {
	cloud := os.Args[1]
	service := os.Args[2]
	switch cloud {
	case "aws":
		awsTerraforming.Generate(service, region)
		//for _, service := range []string{"vpc", "sg", "subnet", "igw", "vpn_gateway", "vpn_connections", "s3"} {
		//awsTerraforming.Generate(service, region)
		//}
	case "google":
		gcp_terraforming.Generate(service, region)
	}
}
