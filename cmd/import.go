// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils/terraformerstring"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"

	"github.com/spf13/pflag"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/terraformoutput"

	"github.com/spf13/cobra"
)

type ImportOptions struct {
	Resources     []string
	Excludes      []string
	PathPattern   string
	PathOutput    string
	State         string
	Bucket        string
	Profile       string
	Verbose       bool
	Zone          string
	Regions       []string
	Projects      []string
	ResourceGroup string
	Connect       bool
	Compact       bool
	Filter        []string
	Plan          bool `json:"-"`
	Output        string
	NoSort        bool
	RetryCount    int
	RetrySleepMs  int
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
	cmd.AddCommand(&cobra.Command{
		Use:   "no-sort",
		Short: "Don't sort resources",
		Long:  "Don't sort resources",
	})
	for _, subcommand := range providerImporterSubcommands() {
		providerCommand := subcommand(options)
		_ = providerCommand.MarkPersistentFlagRequired("resources")
		cmd.AddCommand(providerCommand)
	}
	return cmd
}

func Import(provider terraformutils.ProviderGenerator, options ImportOptions, args []string) error {

	providerWrapper, options, err := initOptionsAndWrapper(provider, options, args)
	if err != nil {
		return err
	}
	defer providerWrapper.Kill()
	providerMapping := terraformutils.NewProvidersMapping(provider)

	err = initAllServicesResources(providerMapping, options, args, providerWrapper)
	if err != nil {
		return err
	}

	err = terraformutils.RefreshResourcesByProvider(providerMapping, providerWrapper)
	if err != nil {
		return err
	}

	providerMapping.ConvertTFStates(providerWrapper)
	// change structs with additional data for each resource
	providerMapping.CleanupProviders()

	err = importFromPlan(providerMapping, options, args)

	return err
}

func initOptionsAndWrapper(provider terraformutils.ProviderGenerator, options ImportOptions, args []string) (*providerwrapper.ProviderWrapper, ImportOptions, error) {
	err := provider.Init(args)
	if err != nil {
		return nil, options, err
	}

	if terraformerstring.ContainsString(options.Resources, "*") {
		log.Println("Attempting an import of ALL resources in " + provider.GetName())
		options.Resources = providerServices(provider)
	}

	if options.Excludes != nil {
		localSlice := []string{}
		for _, r := range options.Resources {
			remove := false
			for _, e := range options.Excludes {
				if r == e {
					remove = true
					log.Println("Excluding resource " + e)
				}
			}
			if !remove {
				localSlice = append(localSlice, r)
			}
		}
		options.Resources = localSlice
	}

	providerWrapper, err := providerwrapper.NewProviderWrapper(provider.GetName(), provider.GetConfig(), options.Verbose, map[string]int{"retryCount": options.RetryCount, "retrySleepMs": options.RetrySleepMs})
	if err != nil {
		return nil, options, err
	}

	return providerWrapper, options, nil
}

func initAllServicesResources(providersMapping *terraformutils.ProvidersMapping, options ImportOptions, args []string, providerWrapper *providerwrapper.ProviderWrapper) error {
	numOfResources := len(options.Resources)
	var wg sync.WaitGroup
	wg.Add(numOfResources)

	var failedServices []string

	for _, service := range options.Resources {
		serviceProvider := providersMapping.AddServiceToProvider(service)
		err := serviceProvider.Init(args)
		if err != nil {
			return err
		}
		err = initServiceResources(service, serviceProvider, options, providerWrapper)
		if err != nil {
			failedServices = append(failedServices, service)
		}
	}

	// remove providers that failed to init their service
	providersMapping.RemoveServices(failedServices)
	providersMapping.ProcessResources(false)

	return nil
}

func importFromPlan(providerMapping *terraformutils.ProvidersMapping, options ImportOptions, args []string) error {
	plan := &ImportPlan{
		Provider:         providerMapping.GetBaseProvider().GetName(),
		Options:          options,
		Args:             args,
		ImportedResource: map[string][]terraformutils.Resource{},
	}

	resourcesByService := providerMapping.GetResourcesByService()
	for service := range resourcesByService {
		plan.ImportedResource[service] = append(plan.ImportedResource[service], resourcesByService[service]...)
	}

	if options.Plan {
		path := Path(options.PathPattern, providerMapping.GetBaseProvider().GetName(), "terraformer", options.PathOutput)
		return ExportPlanFile(plan, path, "plan.json")
	}

	return ImportFromPlan(providerMapping.GetBaseProvider(), plan)
}

func initServiceResources(service string, provider terraformutils.ProviderGenerator,
	options ImportOptions, providerWrapper *providerwrapper.ProviderWrapper) error {
	log.Println(provider.GetName() + " importing... " + service)
	err := provider.InitService(service, options.Verbose)
	if err != nil {
		log.Printf("%s error importing %s, err: %s\n", provider.GetName(), service, err)
		return err
	}
	provider.GetService().ParseFilters(options.Filter)
	err = provider.GetService().InitResources()
	if err != nil {
		log.Printf("%s error initializing resources in service %s, err: %s\n", provider.GetName(), service, err)
		return err
	}

	provider.GetService().PopulateIgnoreKeys(providerWrapper)
	provider.GetService().InitialCleanup()
	log.Println(provider.GetName() + " done importing " + service)

	return nil
}

func ImportFromPlan(provider terraformutils.ProviderGenerator, plan *ImportPlan) error {
	options := plan.Options
	importedResource := plan.ImportedResource
	isServicePath := strings.Contains(options.PathPattern, "{service}")

	if options.Connect {
		log.Println(provider.GetName() + " Connecting.... ")
		importedResource = terraformutils.ConnectServices(importedResource, isServicePath, provider.GetResourceConnections())
	}

	if !isServicePath {
		var compactedResources []terraformutils.Resource
		for _, resources := range importedResource {
			compactedResources = append(compactedResources, resources...)
		}
		e := printService(provider, "", options, compactedResources, importedResource)
		if e != nil {
			return e
		}
	} else {
		for serviceName, resources := range importedResource {
			e := printService(provider, serviceName, options, resources, importedResource)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

func printService(provider terraformutils.ProviderGenerator, serviceName string, options ImportOptions, resources []terraformutils.Resource, importedResource map[string][]terraformutils.Resource) error {
	log.Println(provider.GetName() + " save " + serviceName)
	// Print HCL files for Resources
	path := Path(options.PathPattern, provider.GetName(), serviceName, options.PathOutput)
	err := terraformoutput.OutputHclFiles(resources, provider, path, serviceName, options.Compact, options.Output, !options.NoSort)
	if err != nil {
		return err
	}
	tfStateFile, err := terraformutils.PrintTfState(resources)
	if err != nil {
		return err
	}
	// print or upload State file
	if options.State == "bucket" {
		log.Println(provider.GetName() + " upload tfstate to  bucket " + options.Bucket)
		bucket := terraformoutput.BucketState{
			Name: options.Bucket,
		}
		if err := bucket.BucketUpload(path, tfStateFile); err != nil {
			return err
		}
		// create Bucket file
		if bucketStateDataFile, err := terraformutils.Print(bucket.BucketGetTfData(path), map[string]struct{}{}, options.Output, !options.NoSort); err == nil {
			terraformoutput.PrintFile(path+"/bucket.tf", bucketStateDataFile)
		}
	} else {
		if serviceName == "" {
			log.Println(provider.GetName() + " save tfstate")
		} else {
			log.Println(provider.GetName() + " save tfstate for " + serviceName)
		}
		if err := os.WriteFile(path+"/terraform.tfstate", tfStateFile, os.ModePerm); err != nil {
			return err
		}
	}
	// Print hcl variables.tf
	if serviceName != "" {
		if options.Connect && len(provider.GetResourceConnections()[serviceName]) > 0 {
			variables := map[string]map[string]map[string]interface{}{}
			variables["data"] = map[string]map[string]interface{}{}
			variables["data"]["terraform_remote_state"] = map[string]interface{}{}
			if options.State == "bucket" {
				bucket := terraformoutput.BucketState{
					Name: options.Bucket,
				}
				for k := range provider.GetResourceConnections()[serviceName] {
					if _, exist := importedResource[k]; !exist {
						continue
					}
					variables["data"]["terraform_remote_state"][k] = map[string]interface{}{
						"backend": "gcs",
						"config":  bucket.BucketGetTfData(strings.ReplaceAll(path, serviceName, k)),
					}
				}
			} else {
				for k := range provider.GetResourceConnections()[serviceName] {
					if _, exist := importedResource[k]; !exist {
						continue
					}
					variables["data"]["terraform_remote_state"][k] = map[string]interface{}{
						"backend": "local",
						"config": map[string]interface{}{
							"path": strings.Repeat("../", strings.Count(path, "/")) + strings.ReplaceAll(path, serviceName, k) + "terraform.tfstate",
						},
					}
				}
			}
			// create variables file
			if len(provider.GetResourceConnections()[serviceName]) > 0 && options.Connect && len(variables["data"]["terraform_remote_state"]) > 0 {
				variablesFile, err := terraformutils.Print(variables, map[string]struct{}{"config": {}}, options.Output, !options.NoSort)
				if err != nil {
					return err
				}
				terraformoutput.PrintFile(path+"/variables."+terraformoutput.GetFileExtension(options.Output), variablesFile)
			}
		}
	} else {
		if options.Connect {
			variables := map[string]map[string]map[string]interface{}{}
			variables["data"] = map[string]map[string]interface{}{}
			variables["data"]["terraform_remote_state"] = map[string]interface{}{}
			if options.State == "bucket" {
				bucket := terraformoutput.BucketState{
					Name: options.Bucket,
				}
				variables["data"]["terraform_remote_state"]["local"] = map[string]interface{}{
					"backend": "gcs",
					"config":  bucket.BucketGetTfData(path),
				}
			} else {
				variables["data"]["terraform_remote_state"]["local"] = map[string]interface{}{
					"backend": "local",
					"config": map[string]interface{}{
						"path": "terraform.tfstate",
					},
				}
			}
			// create variables file
			if options.Connect {
				variablesFile, err := terraformutils.Print(variables, map[string]struct{}{"config": {}}, options.Output, !options.NoSort)
				if err != nil {
					return err
				}
				terraformoutput.PrintFile(path+"/variables."+terraformoutput.GetFileExtension(options.Output), variablesFile)
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

func listCmd(provider terraformutils.ProviderGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List supported resources for " + provider.GetName() + " provider",
		Long:  "List supported resources for " + provider.GetName() + " provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			services := providerServices(provider)
			for _, k := range services {
				fmt.Println(k)
			}
			return nil
		},
	}
	cmd.Flags().AddFlag(&pflag.Flag{Name: "resources"})
	return cmd
}

func providerServices(provider terraformutils.ProviderGenerator) []string {
	var services []string
	for k := range provider.GetSupportedService() {
		services = append(services, k)
	}
	sort.Strings(services)
	return services
}

func baseProviderFlags(flag *pflag.FlagSet, options *ImportOptions, sampleRes, sampleFilters string) {
	flag.BoolVarP(&options.Connect, "connect", "c", true, "")
	flag.BoolVarP(&options.Compact, "compact", "C", false, "")
	flag.StringSliceVarP(&options.Resources, "resources", "r", []string{}, sampleRes)
	flag.StringSliceVarP(&options.Excludes, "excludes", "x", []string{}, sampleRes)
	flag.StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/")
	flag.StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	flag.StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	flag.StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	flag.StringSliceVarP(&options.Filter, "filter", "f", []string{}, sampleFilters)
	flag.BoolVarP(&options.Verbose, "verbose", "v", false, "")
	flag.BoolVarP(&options.NoSort, "no-sort", "S", false, "set to disable sorting of HCL")
	flag.StringVarP(&options.Output, "output", "O", "hcl", "output format hcl or json")
	flag.IntVarP(&options.RetryCount, "retry-number", "n", 5, "number of retries to perform when refresh fails")
	flag.IntVarP(&options.RetrySleepMs, "retry-sleep-ms", "m", 300, "time in ms to sleep between retries")
}
