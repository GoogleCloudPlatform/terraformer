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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"

	"github.com/spf13/pflag"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/terraform_output"

	"github.com/spf13/cobra"
)

type ImportOptions struct {
	Resources   []string
	PathPattern string
	PathOutput  string
	State       string
	Bucket      string
	Profile     string
	Zone        string
	Regions     []string
	Projects    []string
	Connect     bool
	Filter      []string
	Plan        bool `json:"-"`
}

const DefaultPathPattern = "{output}/{provider}/{service}/"
const DefaultPathOutput = "generated"
const DefaultState = "local"

func newImportCmd() *cobra.Command {
	options := ImportOptions{}
	cmd := &cobra.Command{
		Use:           "import",
		Short:         "Import current state to Terraform configuration",
		Long:          "Import current state to Terraform configuration",
		SilenceUsage:  true,
		SilenceErrors: false,
		//Version:       version.String(),
	}

	cmd.AddCommand(newCmdPlanImporter(options))
	for _, subcommand := range providerImporterSubcommands() {
		providerCommand := subcommand(options)
		_ = providerCommand.MarkPersistentFlagRequired("resources")
		cmd.AddCommand(providerCommand)
	}
	return cmd
}

func Import(provider terraform_utils.ProviderGenerator, options ImportOptions, args []string) error {
	err := provider.Init(args)
	if err != nil {
		return err
	}
	plan := &ImportPlan{
		Provider:         provider.GetName(),
		Options:          options,
		Args:             args,
		ImportedResource: map[string][]terraform_utils.Resource{},
	}

	for _, service := range options.Resources {
		log.Println(provider.GetName() + " importing... " + service)
		err = provider.InitService(service)
		if err != nil {
			return err
		}
		provider.GetService().ParseFilters(options.Filter)
		err = provider.GetService().InitResources()
		provider.GetService().PopulateIgnoreKeys(provider.GetBasicConfig())
		if err != nil {
			return err
		}
		provider.GetService().InitialCleanup()

		providerWrapper, err := provider_wrapper.NewProviderWrapper(provider.GetName(), provider.GetConfig())
		if err != nil {
			return err
		}

		refreshedResources, err := terraform_utils.RefreshResources(provider.GetService().GetResources(), providerWrapper)
		if err != nil {
			return err
		}
		provider.GetService().SetResources(refreshedResources)

		for i := range provider.GetService().GetResources() {
			err = provider.GetService().GetResources()[i].ConvertTFstate(providerWrapper)
			if err != nil {
				return err
			}
		}

		providerWrapper.Kill()

		provider.GetService().PostRefreshCleanup()

		// change structs with additional data for each resource
		err = provider.GetService().PostConvertHook()
		if err != nil {
			return err
		}
		plan.ImportedResource[service] = append(plan.ImportedResource[service], provider.GetService().GetResources()...)
	}
	if options.Plan {
		path := Path(options.PathPattern, provider.GetName(), "terraformer", options.PathOutput)
		return ExportPlanFile(plan, path, "plan.json")
	} else {
		return ImportFromPlan(provider, plan)
	}
}

func ImportFromPlan(provider terraform_utils.ProviderGenerator, plan *ImportPlan) error {
	options := plan.Options
	importedResource := plan.ImportedResource

	if options.Connect {
		log.Println(provider.GetName() + " Connecting.... ")
		importedResource = terraform_utils.ConnectServices(importedResource, provider.GetResourceConnections())
	}

	for serviceName, resources := range importedResource {
		log.Println(provider.GetName() + " save " + serviceName)
		// Print HCL files for Resources
		path := Path(options.PathPattern, provider.GetName(), serviceName, options.PathOutput)
		err := terraform_output.OutputHclFiles(resources, provider, path, serviceName)
		if err != nil {
			return err
		}
		tfStateFile, err := terraform_utils.PrintTfState(resources)
		if err != nil {
			return err
		}
		// print or upload State file
		if options.State == "bucket" {
			log.Println(provider.GetName() + " upload tfstate to  bucket " + options.Bucket)
			bucket := terraform_output.BucketState{
				Name: options.Bucket,
			}
			if err := bucket.BucketUpload(path, tfStateFile); err != nil {
				return err
			}
			// create Bucket file
			if bucketStateDataFile, err := terraform_utils.HclPrint(bucket.BucketGetTfData(path), map[string]struct{}{}); err == nil {
				terraform_output.PrintFile(path+"/bucket.tf", bucketStateDataFile)
			}
		} else {
			log.Println(provider.GetName() + " save tfstate for " + serviceName)
			if err := ioutil.WriteFile(path+"/terraform.tfstate", tfStateFile, os.ModePerm); err != nil {
				return err
			}
		}

		// Print hcl variables.tf
		if options.Connect && len(provider.GetResourceConnections()[serviceName]) > 0 {
			variables := map[string]map[string]map[string]interface{}{}
			variables["data"] = map[string]map[string]interface{}{}
			variables["data"]["terraform_remote_state"] = map[string]interface{}{}
			if options.State == "bucket" {
				bucket := terraform_output.BucketState{
					Name: options.Bucket,
				}
				for k := range provider.GetResourceConnections()[serviceName] {
					if _, exist := importedResource[k]; !exist {
						continue
					}
					variables["data"]["terraform_remote_state"][k] = map[string]interface{}{
						"backend": "gcs",
						"config": map[string]interface{}{
							"bucket": bucket.Name,
							"prefix": bucket.BucketPrefix(strings.Replace(path, serviceName, k, -1)),
						},
					}
				}
			} else {
				for k := range provider.GetResourceConnections()[serviceName] {
					if _, exist := importedResource[k]; !exist {
						continue
					}
					variables["data"]["terraform_remote_state"][k] = map[string]interface{}{
						"backend": "local",
						"config": [1]interface{}{map[string]interface{}{
							"path": strings.Repeat("../", strings.Count(path, "/")) + strings.Replace(path, serviceName, k, -1) + "terraform.tfstate",
						}},
					}
				}
			}
			// create variables file
			if len(provider.GetResourceConnections()[serviceName]) > 0 && options.Connect && len(variables["data"]["terraform_remote_state"]) > 0 {
				variablesFile, err := terraform_utils.HclPrint(variables, map[string]struct{}{"config": {}})
				if err != nil {
					return err
				}
				terraform_output.PrintFile(path+"/variables.tf", variablesFile)
			}
		}
	}
	return nil
}

func Path(pathPattern, providerName, serviceName, output string) string {
	return strings.NewReplacer(
		"{provider}", providerName,
		"{service}", serviceName,
		"{output}", output,
	).Replace(pathPattern)
}

func listCmd(provider terraform_utils.ProviderGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List supported resources for " + provider.GetName() + " provider",
		Long:  "List supported resources for " + provider.GetName() + " provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			services := []string{}
			for k := range provider.GetSupportedService() {
				services = append(services, k)
			}
			sort.Strings(services)
			for _, k := range services {
				fmt.Println(k)
			}
			return nil
		},
	}
	cmd.Flags().AddFlag(&pflag.Flag{Name: "resources"})
	return cmd
}
