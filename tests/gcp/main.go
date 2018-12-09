package main

import (
	"log"
	"os"
	"os/exec"

	"waze/terraformer/gcp_terraforming"
	"waze/terraformer/gcp_terraforming/compute_resources"
)

const command = "terraform init && terraform plan"

func main() {
	zone := "europe-west1-c"
	services := []string{}
	for service := range computeTerrforming.ComputeService {
		if service == "images" {
			continue
		}
		if service == "instanceTemplates" {
			continue
		}
		if service == "instances" {
			continue
		}
		if service == "targetHttpProxies" {
			continue
		}
		services = append(services, service)

	}
	services = append(services, "gcs", "alerts")
	for _, service := range services {
		err := gcp_terraforming.Generate(service, []string{zone})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		rootPath, _ := os.Getwd()
		currentPath := rootPath + gcp_terraforming.PathForGenerateFiles + zone + "/" + service
		if err := os.Chdir(currentPath); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Chdir(rootPath)
	}
}
