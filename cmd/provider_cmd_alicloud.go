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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdAliCloudImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alicloud",
		Short: "Import current State to terraform configuration from alicloud",
		Long:  "Import current State to terraform configuration from alicloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, region := range options.Regions {
				provider := newAliCloudProvider()
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
	cmd.AddCommand(listCmd(newAliCloudProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "vpc,subnet,nacl", "slb=id1:id2:id4")
	cmd.PersistentFlags().StringVar(&options.Profile, "profile", "default", "prod")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "cn-hangzhou")
	return cmd
}

func newAliCloudProvider() terraformutils.ProviderGenerator {
	return &alicloud_terraforming.AliCloudProvider{}
}
