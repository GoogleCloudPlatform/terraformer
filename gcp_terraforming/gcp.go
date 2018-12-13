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

package gcp_terraforming

import (
	"errors"

	"io/ioutil"
	"log"
	"os"
	"strings"
	"waze/terraformer/gcp_terraforming/alerts"
	"waze/terraformer/gcp_terraforming/clouddns"
	"waze/terraformer/gcp_terraforming/compute_resources"
	"waze/terraformer/gcp_terraforming/gcp_generator"
	"waze/terraformer/gcp_terraforming/gcs"
	"waze/terraformer/gcp_terraforming/iam"
	"waze/terraformer/terraform_utils"
)

const PathForGenerateFiles = "/generated/gcp/"

// GetGCPSupportService return map of support service for GCP
func GetGCPSupportService() map[string]gcp_generator.Generator {
	services := computeTerrforming.ComputeService
	services["gcs"] = gcs.GcsGenerator{}
	services["alerts"] = alerts.AlertsGenerator{}
	services["iam"] = iam.IamGenerator{}
	services["dns"] = clouddns.CloudDNSGenerator{}
	return services
}

// Main function for generate tf and tfstate file by GCP service and region
func Generate(service string, args []string) error {
	var generator gcp_generator.Generator
	var isSupported bool
	if generator, isSupported = GetGCPSupportService()[service]; !isSupported {
		return errors.New("gcp: not supported service")
	}
	// check projectName in env params
	projectName := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectName == "" {
		return errors.New("google cloud project name must be set")
	}
	// try connect to provider in $HOME/.terraform.d/....
	provider, err := terraform_utils.GetProvider("google")
	if err != nil {
		return err
	}

	zone := args[0]
	rootPath, _ := os.Getwd()
	currentPath := rootPath + PathForGenerateFiles + projectName + "/" + zone + "/" + service
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		log.Print(err)
		return err
	}
	// generate TerraformResources with type and ids + metadata
	cloudResources, metadata, err := generator.Generate(zone)
	if err != nil {
		return err
	}

	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	providerObject := NewGcpRegionResource(region)

	refreshedResources := terraform_utils.RefreshResources(provider, cloudResources)

	// create tfstate
	tfstateFile, err := terraform_utils.PrintTfState(refreshedResources)
	if err != nil {
		return err
	}
	// convert InstanceState to go struct for hcl print
	converter := terraform_utils.InstanceStateConverter{}
	refreshedResources, err = converter.Convert(refreshedResources, metadata)
	if err != nil {
		return err
	}
	// change structs with additional data for each resource
	refreshedResources, err = generator.PostGenerateHook(refreshedResources)
	// create HCL
	tfFile := []byte{}
	tfFile, err = terraform_utils.HclPrint(refreshedResources, providerObject)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(currentPath+"/"+service+".tf", tfFile, os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(currentPath+"/terraform.tfstate", tfstateFile, os.ModePerm)

}

func NewGcpRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"google": map[string]interface{}{
			"region":  region,
			"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
		},
	}
}
