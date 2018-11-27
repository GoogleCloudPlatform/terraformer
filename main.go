package main

import (
	"os"

	"waze/terraform/aws_terraforming"
	"waze/terraform/gcp_terraforming"
)

func main() {
	provider := os.Args[1]
	service := os.Args[2]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[3:]
	}
	switch provider {
	case "aws":
		awsTerraforming.Generate(service, args)
		//for _, service := range []string{"vpc", "sg", "subnet", "igw", "vpn_gateway", "vpn_connections", "s3"} {
		//awsTerraforming.Generate(service, region)
		//}
	case "google":
		gcp_terraforming.Generate(service, args)
	}
}
