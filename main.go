package main

import "waze/terraform/aws_terraforming"

var (
	cloud   = "aws"
	service = "s3"
	region  = "eu-west-1"
)

func main() {
	switch cloud {
	case "aws":
		awsTerraforming.Generate(service, region)
	}
}
