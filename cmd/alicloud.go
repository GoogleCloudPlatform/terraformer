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

	alicloud_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/alicloud"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdAlicloudImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alicloud",
		Short: "Import current State to terraform configuration from alicloud",
		Long:  "Import current State to terraform configuration from alicloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, region := range options.Regions {
				provider := newAlicloudProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern += region + "/"
				log.Println(provider.GetName() + " importing region " + region)
				profile := options.Profile
				err := Import(provider, options, []string{region, profile})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAlicloudProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "vpc,subnet,nacl")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringVar(&options.Profile, "profile", "default", "prod")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "eu-west-1,eu-west-2,us-east-1")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "alicloud_slb=id1:id2:id4") // TODO: Not implemented
	return cmd
}

func newAlicloudProvider() terraform_utils.ProviderGenerator {
	return &alicloud_terraforming.AlicloudProvider{}
}
