package gcp_terraforming

import (
	"os"
	"strings"

	"waze/terraform/gcp_terraforming/alerts"
	"waze/terraform/gcp_terraforming/compute_resources"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/gcp_terraforming/gcs"
	"waze/terraform/gcp_terraforming/iam"
	"waze/terraform/terraform_utils"
)

const PathForGenerateFiles = "/generated/gcp/"

func NewGcpRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"google": map[string]interface{}{
			"region":  region,
			"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
		},
	}
}

func Generate(service string, args []string) error {
	zone := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + PathForGenerateFiles + zone + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		return err
	}
	if err := os.Chdir(currentPath); err != nil {
		return err
	}
	defer os.Chdir(rootPath)
	var generator gcp_generator.Generator
	switch service {
	case "gcs":
		generator = gcs.GcsGenerator{}
	case "iam":
		generator = iam.IamGenerator{}
	case "alerts":
		generator = alerts.AlertsGenerator{}
	default:
		if service, exist := computeTerrforming.ComputeService[service]; exist {
			generator = service
		}
	}
	resources, metadata, err := generator.Generate(zone)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	converter := terraform_utils.TfstateConverter{}
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	err = terraform_utils.GenerateTf(resources, service, NewGcpRegionResource(region))
	if err != nil {
		return err
	}
	return nil

}
