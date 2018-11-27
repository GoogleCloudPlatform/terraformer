package gcp_terraforming

import (
	"log"
	"os"
	"strings"

	"waze/terraform/gcp_terraforming/compute_code_gen"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/gcp_terraforming/gcs"
	"waze/terraform/gcp_terraforming/iam"
	"waze/terraform/terraform_utils"
)

const pathForGenerateFiles = "/generated/gcp/"

func Generate(service string, args []string) {
	zone := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + pathForGenerateFiles + zone + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		log.Print(err)
		return
	}
	if err := os.Chdir(currentPath); err != nil {
		log.Print(err)
		return
	}
	defer os.Chdir(rootPath)
	var generator gcp_generator.Generator
	switch service {
	case "gcs":
		generator = gcs.GcsGenerator{}
	case "iam":
		generator = iam.IamGenerator{}
	default:
		if service, exist := computeTerrforming.ComputeService[service]; exist {
			generator = service
		}
	}
	resources, metadata, err := generator.Generate(zone)
	if err != nil {
		log.Println(err)
		return
	}
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		log.Println(err)
		return
	}
	converter := terraform_utils.TfstateConverter{}
	resources, err = converter.Convert("terraform.tfstate", metadata)
	if err != nil {
		log.Println(err)
		return
	}
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	err = terraform_utils.GenerateTf(resources, service, region, "google")
	if err != nil {
		log.Println(err)
		return
	}
	return

}
