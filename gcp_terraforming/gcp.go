package gcp_terraforming

import (
	"log"
	"os"
	"waze/terraform/gcp_terraforming/firewall_rules"
	"waze/terraform/gcp_terraforming/gcp_generator"
	"waze/terraform/gcp_terraforming/gcs"
)

const pathForGenerateFiles = "/generated/gcp/"

func Generate(service, region string) {
	rootPath, _ := os.Getwd()
	currentPath := rootPath + pathForGenerateFiles + region + "/" + service
	os.MkdirAll(currentPath, os.ModePerm)
	os.Chdir(currentPath)
	defer os.Chdir(rootPath)
	var generator gcp_generator.Generator
	switch service {
	case "gcs":
		generator = gcs.GcsGenerator{}
	case "firewall":
		generator = firewall_rules.FirewallRulesGenerator{}
	}
	err := generator.Generate(region)
	if err != nil {
		log.Println(err)
		return
	}

}
