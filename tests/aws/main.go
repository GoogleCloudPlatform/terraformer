package main

import (
	"log"
	"os"
	"os/exec"

	"waze/terraformer/aws_terraforming"
)

const command = "terraform init && terraform plan"

func main() {
	region := "eu-west-1"
	for service := range aws_terraforming.GetAWSSupportService() {
		if service == "iam" {
			continue
		}
		err := aws_terraforming.Generate(service, []string{region})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		rootPath, _ := os.Getwd()
		currentPath := rootPath + aws_terraforming.PathForGenerateFiles + region + "/" + service
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
