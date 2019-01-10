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


package cmd

import (
	"io/ioutil"
	"os"
	"waze/terraformer/aws_terraforming"
	"waze/terraformer/gcp_terraforming"
	"waze/terraformer/terraform_utils"
)

func Exec(providerName, service string, args []string) error {
	if len(os.Args) > 2 {
		args = os.Args[3:]
	}

	var err error
	var provider terraform_utils.ProviderGenerator
	switch providerName {
	case "aws":
		provider = &aws_terraforming.AWSProvider{}
	case "google":
		provider = &gcp_terraforming.GCPProvider{}
	}

	err = provider.Init(args)
	if err != nil {
		return err
	}

	err = provider.InitService(service)
	if err != nil {
		return err
	}

	err = provider.GenerateOutputPath()
	if err != nil {
		return err
	}

	err = provider.GetService().InitResources()
	if err != nil {
		return err
	}
	refreshedResources, err := terraform_utils.RefreshResources(provider.GetService().GetResources(), provider.GetName())
	if err != nil {
		return err
	}
	provider.GetService().SetResources(refreshedResources)

	// create tfstate
	tfStateFile, err := terraform_utils.PrintTfState(provider.GetService().GetResources())
	if err != nil {
		return err
	}
	// convert InstanceState to go struct for hcl print
	for i := range provider.GetService().GetResources() {
		provider.GetService().GetResources()[i].ConvertTFstate()
	}
	// change structs with additional data for each resource
	err = provider.GetService().PostConvertHook()
	if err != nil {
		return err
	}
	// create HCL
	tfFile := []byte{}
	tfFile, err = terraform_utils.HclPrintResource(provider.GetService().GetResources(), provider.RegionResource())
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(provider.CurrentPath()+"/"+service+".tf", tfFile, os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(provider.CurrentPath()+"/terraform.tfstate", tfStateFile, os.ModePerm)
}
