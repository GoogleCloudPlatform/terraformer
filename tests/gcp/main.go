// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
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
		currentPath := fmt.Sprintf(gcp_terraforming.GenerateFilesFolderPath, rootPath, gcp_terraforming.PathForGenerateFiles, os.Getenv("GOOGLE_CLOUD_PROJECT"), zone, service)
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
