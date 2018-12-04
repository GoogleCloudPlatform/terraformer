package gcp_terraforming

import (
	"errors"
	"os"
	"strings"

	"waze/terraform/gcp_terraforming/alerts"
	"waze/terraform/gcp_terraforming/clouddns"
	"waze/terraform/gcp_terraforming/compute_resources"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/gcp_terraforming/gcs"
	"waze/terraform/gcp_terraforming/iam"
	"waze/terraform/terraform_utils"
)

const PathForGenerateFiles = "/generated/gcp/"

// GetGCPSupportService return map of support service for GCP
func GetGCPSupportService() map[string]gcp_generator.Generator {
	services := computeTerrforming.ComputeService
	services["gcs"] = gcs.GcsGenerator{}
	services["alerts"] = alerts.AlertsGenerator{}
	services["iam"] = iam.IamGenerator{}
	services["dns"] = clouddns.CloudDNSGenerator{}
	return services
}

// Main function for generate tf and tfstate file by GCP service and region
func Generate(service string, args []string) error {
	zone := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + PathForGenerateFiles + zone + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		return err
	}
	// change current dir for terraform refresh
	if err := os.Chdir(currentPath); err != nil {
		return err
	}
	// return current dir after terraform refresh run
	defer os.Chdir(rootPath)
	var generator gcp_generator.Generator
	var isSupported bool
	if generator, isSupported = GetGCPSupportService()[service]; !isSupported {
		return errors.New("gcp: not supported service")
	}
	// generate TerraformResources with type and ids + metadata
	resources, metadata, err := generator.Generate(zone)
	if err != nil {
		return err
	}
	// generate empty(resource and ids) tfstate,
	// and run terraform refresh with empty tfstate for populate data
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	// convert tfstate to go struct for hcl print
	converter := terraform_utils.TfstateConverter{}
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		return err
	}
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	// change structs with additional data for each resource
	resources, err = generator.PostGenerateHook(resources)
	// print HCL file
	err = terraform_utils.GenerateTf(resources, service, NewGcpRegionResource(region))
	if err != nil {
		return err
	}
	return nil

}

func NewGcpRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"google": map[string]interface{}{
			"region":  region,
			"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
		},
	}
}
