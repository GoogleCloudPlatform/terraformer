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
	"waze/terraformer/aws_terraforming"
	"waze/terraformer/cmd"
	"waze/terraformer/terraform_utils"
)

const command = "terraform init && terraform plan"

func main() {
	region := "eu-west-1"
	services := []string{}
	provider := &aws_terraforming.AWSProvider{}
	for service := range provider.GetAWSSupportService() {
		if service == "route53" {
			continue
		}
		if service == "iam" {
			continue
		}
		if service == "sg" {
			continue
		}
		services = append(services, service)

	}
	sort.Strings(services)
	for _, service := range services {
		if _, err := os.Stat("/Users/sergeylanz/go/src/waze/terraformer/generated/aws/eu-west-1/" + service); !os.IsNotExist(err) {
			continue
		}
		err := cmd.Exec("aws", service, []string{region})
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		provider = &aws_terraforming.AWSProvider{
			Provider: terraform_utils.Provider{},
		}
		provider.Init([]string{region})
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
