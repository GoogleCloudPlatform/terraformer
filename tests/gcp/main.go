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
	"log"
	"os"
	"os/exec"
	"sort"
	"waze/terraformer/cmd"
	"waze/terraformer/terraform_utils"

	"waze/terraformer/gcp_terraforming"
)

const command = "terraform init && terraform plan"

func main() {
	zone := "europe-west1-c"
	services := []string{}
	provider := &gcp_terraforming.GCPProvider{}
	for service := range provider.GetGCPSupportService() {
		if service == "images" {
			continue
		}
		if service == "iam" {
			continue
		}
		if service == "instanceTemplates" {
			continue
		}
		if service == "targetHttpProxies" {
			continue
		}
		if service == "monitoring" {
			continue
		}
		if service == "cloudsql" {
			continue
		}
		services = append(services, service)

	}
	sort.Strings(services)
	for _, service := range services {
		err := cmd.Exec("google", service, []string{zone})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		provider = &gcp_terraforming.GCPProvider{
			Provider: terraform_utils.Provider{},
		}
		provider.Init([]string{zone})
		provider.InitService(service)
		rootPath, _ := os.Getwd()
		currentPath := provider.CurrentPath()
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
