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
	"github.com/GoogleCloudPlatform/terraformer/cmd"
	aws_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/aws"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"log"
	"os"
	"os/exec"
	"sort"
)

const command = "terraform init && terraform plan"

func main() {
	region := "us-east-1"
	profile := "personal"
	var services []string
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
	services = []string{"waf"}
	sort.Strings(services)
	provider = &aws_terraforming.AWSProvider{
		Provider: terraform_utils.Provider{},
	}
	err := cmd.Import(provider, cmd.ImportOptions{
		Resources:   services,
		PathPattern: "{output}/{provider}/",
		PathOutput:  cmd.DefaultPathOutput,
		State:       "local",
		Connect:     true,
		Compact:     true,
		Verbose:     true,
		Output:	     "hcl",
	}, []string{region, profile})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rootPath, _ := os.Getwd()
	currentPath := cmd.Path(cmd.DefaultPathPattern, provider.GetName(), "", cmd.DefaultPathOutput)
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
