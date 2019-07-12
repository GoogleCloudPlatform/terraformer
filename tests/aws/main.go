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

	"github.com/GoogleCloudPlatform/terraformer/cmd"
	aws_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/aws"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

const command = "terraform init && terraform plan"

func main() {
	region := "ap-southeast-1"
	profile := "dev_sin_admin"
	services := []string{}
	provider := &aws_terraforming.AWSProvider{}
	for service := range provider.GetSupportedService() {
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
	services = []string{"vpc", "subnet", "nacl"}
	sort.Strings(services)
	provider = &aws_terraforming.AWSProvider{
		Provider: terraform_utils.Provider{},
	}
	err := cmd.Import(provider, cmd.ImportOptions{
		Resources:   services,
		PathPattern: cmd.DefaultPathPattern,
		PathOutput:  cmd.DefaultPathOutput,
		State:       "local",
		Connect:     true,
	}, []string{region, profile})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rootPath, _ := os.Getwd()
	for _, serviceName := range services {
		currentPath := cmd.Path(cmd.DefaultPathPattern, provider.GetName(), serviceName, cmd.DefaultPathOutput)
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
