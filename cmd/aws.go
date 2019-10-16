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

	aws_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/aws"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

const defaultRegion = ""

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
				options.Resources = parseGlobalResources(originalResources)
				options.Regions = []string{defaultRegion}
				e := importGlobalResources(options)
				if e != nil {
					return e
				}

				options.Resources = parseRegionalResources(originalResources)
				options.Regions = originalRegions
				for _, region := range originalRegions {
					e := importRegionResources(options, originalPathPattern, region)
					if e != nil {
						return e
					}
				}
				return nil
			} else {
				err := importRegionResources(options, options.PathPattern, defaultRegion)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAWSProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "vpc,subnet,nacl")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringVar(&options.Profile, "profile", "default", "prod")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "eu-west-1,eu-west-2,us-east-1")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "aws_elb=id1:id2:id4")
	return cmd
}

func parseGlobalResources(allResources []string) []string {
	var globalResources []string
	for _, resourceName := range allResources {
		if contains(aws_terraforming.SupportedGlobalResources, resourceName) {
			globalResources = append(globalResources, resourceName)
		}
	}
	return globalResources
}

func importGlobalResources(options ImportOptions) error {
	if len(options.Resources) > 0 {
		return importRegionResources(options, options.PathPattern, defaultRegion)
	} else {
		return nil
	}
}

func parseRegionalResources(allResources []string) []string {
	var localResources []string
	for _, resourceName := range allResources {
		if !contains(aws_terraforming.SupportedGlobalResources, resourceName) {
			localResources = append(localResources, resourceName)
		}
	}
	return localResources
}

func importRegionResources(options ImportOptions, originalPathPattern string, region string) error {
	provider := newAWSProvider()
	options.PathPattern = originalPathPattern
	if region != "" {
		options.PathPattern += region + "/"
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

func newAWSProvider() terraform_utils.ProviderGenerator {
	return &aws_terraforming.AWSProvider{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
