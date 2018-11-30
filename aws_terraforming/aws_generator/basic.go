package aws_generator

import "waze/terraform/terraform_utils"

type BasicGenerator struct{}

type Generator interface {
	Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error)
}
