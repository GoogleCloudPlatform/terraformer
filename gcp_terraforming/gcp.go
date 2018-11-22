package gcp_terraforming

import (
	"log"
	"os"
	"strings"
	"waze/terraform/gcp_terraforming/compute_code_gen"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/gcp_terraforming/gcs"
)

const pathForGenerateFiles = "/generated/gcp/"

func Generate(service, zone string) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	rootPath, _ := os.Getwd()
	currentPath := rootPath + pathForGenerateFiles + region + "/" + service
	os.MkdirAll(currentPath, os.ModePerm)
	os.Chdir(currentPath)
	defer os.Chdir(rootPath)
	var generator gcp_generator.Generator
	switch service {
	case "gcs":
		generator = gcs.GcsGenerator{}
	default:
		if service, exist := compute_code_gen.ComputeService[service]; exist {
			generator = service
		}
	}
	err := generator.Generate(zone)
	if err != nil {
		log.Println(err)
		return
	}

}
