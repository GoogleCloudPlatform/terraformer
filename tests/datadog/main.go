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
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/cmd"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"

	datadog_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/datadog"
)

func main() {
	var terraformerServices []string
	var terraformerFilters []string

	provider := &datadog_terraforming.DatadogProvider{}
	cfg, _ := getConfig()

	// CD into 'tests/datadog/resources'
	err := os.Chdir(datadogResourcesPath)
	if err != nil {
		handleFatalErr(cfg, err, "Error changing directory: ")
	}
	// Run the terraform v13 upgrade command if applicable
	if strings.Contains(cfg.tfVersion, "0.13.") {
		if err := cmdRun(cfg, []string{commandTerraformV13Upgrade}); err != nil {
			handleFatalErr(cfg, err, "Error running command 'terraform 0.13upgrade'")
		}
	}
	// Initialize the Datadog provider for resource creation
	err = initializeDatadogProvider(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error initializing the Datadog provider: ")
	}

	// Create datadog resources
	resourcesMap, err := createDatadogResource(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error creating resources")
	}

	for resource, resourceId := range *resourcesMap {
		terraformerServices = append(terraformerServices, resource)
		terraformerFilters = append(terraformerFilters, fmt.Sprintf("%s=%s", resource, strings.Join(resourceId, ":")))
	}
	if len(terraformerServices) == 0 {
		terraformerServices = getAllServices(provider)
	}

	// Import created resources with Terraformer
	err = cmd.Import(provider, cmd.ImportOptions{
		Resources:   terraformerServices,
		PathPattern: "{output}/",
		PathOutput:  cmd.DefaultPathOutput,
		State:       "local",
		Connect:     true,
		Output:      "hcl",
		Filter:      terraformerFilters,
		Verbose:     true,
	}, []string{cfg.Datadog.apiKey, cfg.Datadog.appKey, cfg.Datadog.apiURL})
	if err != nil {
		handleFatalErr(cfg, err, "Error while importing resources")
	}

	// Run tests to ensure created and imported resources match
	err = terraformerResourcesTest(cfg, resourcesMap)
	if err != nil {
		handleFatalErr(cfg, err, "Terraform resource test step failed")
	}

	// Destroy created resources
	err = destroyDatadogResources(cfg)
	if err != nil {
		log.Fatal("Error while destroying resources ", err)
	}

	log.Print("Successfully created and imported resources with Terraformer")
}

func terraformerResourcesTest(cfg *Config, resourcesMap *map[string][]string) error {
	if err := os.Chdir("generated/"); err != nil {
		return err
	}

	// Check if terraform version is 0.13.x. If so, remove the generated provider file and upgrade
	if strings.Contains(cfg.tfVersion, "0.13.") {
		e := os.Remove("provider.tf")
		if e != nil {
			log.Printf("provider.tf file does not exist in the directory")
		}
		if err := cmdRun(cfg, []string{commandTerraformV13Upgrade}); err != nil {
			return err
		}
	}

	//Initialize Datadog provider in the 'generated/' directory
	err := initializeDatadogProvider(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error initializing the Datadog provider: ")
	}

	// Collect outputs from generated resources
	terraformerResourcesOutput, err := terraformOutput()
	if err != nil {
		log.Println(err)
		return err
	}
	terraformResourcesMap := parseTerraformOutput(string(terraformerResourcesOutput))

	// Sort the map values for easier comparison
	for _, v := range *terraformResourcesMap {
		sort.Strings(v)
	}
	for _, v := range *resourcesMap {
		sort.Strings(v)
	}

	log.Printf("Comparing resource names and resources ids. \n Created resources: %s \n Imported Resources: %s", resourcesMap, terraformResourcesMap)
	match := reflect.DeepEqual(resourcesMap, terraformResourcesMap)
	if match {
		// Run plan against the generated resources to make sure there is no diff
		log.Println("Running terraform plan on generated resources. This should produce no diff")
		err := terraformPlan(cfg)
		if err != nil {
			return err
		}

		if err := os.Chdir(cfg.rootPath); err != nil {
			return err
		}
	} else {
		log.Printf("Imported resources did not match the created. \n Created resources: %s \n Imported Resources: %s", resourcesMap, terraformResourcesMap)
		return errors.New("imported resource names and/or ids did not match the created")
	}

	if err := os.Chdir(cfg.rootPath); err != nil {
		return err
	}

	return nil
}

func handleFatalErr(cfg *Config, err error, msg string) {
	// Destroy any lingering resources before exiting
	log.Print("Destroying resources before exiting")
	if err := destroyDatadogResources(cfg); err != nil {
		log.Printf("Error while destroying resources: %s", err)
	}

	log.Fatalf("Message: %s. Error: %s",msg,  err)
}
