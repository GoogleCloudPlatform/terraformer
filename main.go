package main

import "waze/terraform/aws_terraforming"

var (
	cloud = "aws"
	//service = "s3"
	region = "eu-west-1"
)

func main() {
	switch cloud {
	case "aws":
		for _, service := range []string{"vpc", "sg", "subnet", "igw", "vpn_gateway", "vpn_connections", "s3"} {
			awsTerraforming.Generate(service, region)
		}
	}
}
