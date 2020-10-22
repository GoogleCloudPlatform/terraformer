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
	"log"

	awsterraformer "github.com/GoogleCloudPlatform/terraformer/providers/aws"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdAwsImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws",
		Short: "Import current state to Terraform configuration from AWS",
		Long:  "Import current state to Terraform configuration from AWS",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalResources := options.Resources
			originalRegions := options.Regions
			originalPathPattern := options.PathPattern

			if len(options.Regions) > 0 {
				shouldSpecifyPathRegion := len(options.Regions) > 1
				globalResources := parseGlobalResources(originalResources)
				options.Resources = globalResources
				options.Regions = []string{awsterraformer.GlobalRegion}
				e := importGlobalResources(options)
				if e != nil {
					return e
				}

				options.Resources = parseRegionalResources(originalResources)
				options.Regions = originalRegions
				if len(options.Resources) > 0 { // don't import anything and potentially override global resources
					if len(globalResources) > 0 {
						shouldSpecifyPathRegion = true // we should keep global resources away from regional
					}
					for _, region := range originalRegions {
						e := importRegionResources(options, originalPathPattern, region, shouldSpecifyPathRegion)
						if e != nil {
							return e
						}
					}
				}
				return nil
			}
			err := importRegionResources(options, options.PathPattern, awsterraformer.NoRegion, false)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAWSProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "vpc,subnet,nacl", "elb=id1:id2:id4")

	cmd.PersistentFlags().StringVarP(&options.Profile, "profile", "", "default", "prod")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "eu-west-1,eu-west-2,us-east-1")
	return cmd
}

func parseGlobalResources(allResources []string) []string {
	var globalResources []string
	for _, resourceName := range allResources {
		if contains(awsterraformer.SupportedGlobalResources, resourceName) {
			globalResources = append(globalResources, resourceName)
		}
	}
	return globalResources
}

func importGlobalResources(options ImportOptions) error {
	if len(options.Resources) > 0 {
		return importRegionResources(options, options.PathPattern, awsterraformer.GlobalRegion, false)
	}
	return nil
}

func parseRegionalResources(allResources []string) []string {
	var localResources []string
	for _, resourceName := range allResources {
		if !contains(awsterraformer.SupportedGlobalResources, resourceName) {
			localResources = append(localResources, resourceName)
		}
	}
	return localResources
}

func importRegionResources(options ImportOptions, originalPathPattern string, region string, shouldSpecifyPathRegion bool) error {
	provider := newAWSProvider()
	options.PathPattern = originalPathPattern
	if region != awsterraformer.GlobalRegion && region != awsterraformer.NoRegion {
		if shouldSpecifyPathRegion {
			options.PathPattern += region + "/"
		}
		log.Println(provider.GetName() + " importing region " + region)
	} else {
		log.Println(provider.GetName() + " importing default region")
	}
	err := Import(provider, options, []string{region, options.Profile})
	if err != nil {
		return err
	}
	return nil
}

func newAWSProvider() terraformutils.ProviderGenerator {
	return &awsterraformer.AWSProvider{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
