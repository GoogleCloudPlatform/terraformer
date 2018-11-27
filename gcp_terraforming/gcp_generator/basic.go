package gcp_generator

import "waze/terraform/terraform_utils"

type BasicGenerator struct{}

type Generator interface {
	Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error)
}
