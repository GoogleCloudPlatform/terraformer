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
	"log"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/cmd"
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
	// Initialize the Datadog provider in dir 'tests/datadog/resources'
	err = initializeDatadogProvider(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error initializing the Datadog provider: ")
	}

	// Create datadog resources
	resourcesMap, err := createDatadogResource(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error creating resources")
	}

	// Get list of terraformerServices and terraformerFilters from created resources
	for resource, resourceID := range *resourcesMap {
		terraformerServices = append(terraformerServices, resource)
		terraformerFilters = append(terraformerFilters, fmt.Sprintf("%s=%s", resource, strings.Join(resourceID, ":")))
	}
	if len(terraformerServices) == 0 {
		terraformerServices = getAllServices(provider)
	}

	// Delete the 'generated/' directory if it already exists
	_ = os.RemoveAll("generated/")

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

	// Run tests on created and imported resources
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

	// Run TF 0.13 upgrade command if applicable
	if strings.Contains(cfg.tfVersion, "0.13.") {
		if err := cmdRun(cfg, []string{commandTerraformV13Upgrade}); err != nil {
			return err
		}
	}

	// Initialize Datadog provider in the 'generated/' directory
	err := initializeDatadogProvider(cfg)
	if err != nil {
		handleFatalErr(cfg, err, "Error initializing the Datadog provider: ")
	}

	// Collect tf outputs from generated resources
	terraformerResourcesOutput, err := terraformOutput()
	if err != nil {
		log.Println(err)
		return err
	}
	terraformResourcesMap := parseTerraformOutput(string(terraformerResourcesOutput))

	// Sort the map values
	for _, v := range *terraformResourcesMap {
		sort.Strings(v)
	}
	for _, v := range *resourcesMap {
		sort.Strings(v)
	}

	log.Println("Comparing resource names and resources ids. \n Created resources:", resourcesMap, "\n Imported Resources:", terraformResourcesMap)
	match := reflect.DeepEqual(resourcesMap, terraformResourcesMap)
	if match {
		// Run 'terraform plan' against the generated resources
		// Command will exit with exit code 2 if diff is produced
		log.Println("Running terraform plan on generated resources. This should produce no diff")
		err := terraformPlan(cfg)
		if err != nil {
			return err
		}

		if err := os.Chdir(cfg.rootPath); err != nil {
			return err
		}
	} else {
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

	log.Fatalf("Message: %s. Error: %s", msg, err)
}
