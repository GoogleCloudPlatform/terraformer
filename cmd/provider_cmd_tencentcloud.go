// Copyright 2021 The Terraformer Authors.
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

	tencentcloud_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/tencentcloud"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdTencentCloudImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tencentcloud",
		Short: "Import current state to Terraform configuration from Tencent Cloud",
		Long:  "Import current state to Terraform configuration from Tencent Cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, region := range options.Regions {
				provider := newTencentCloudProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern += region + "/"
				log.Println(provider.GetName() + " importing region " + region)
				err := Import(provider, options, []string{region})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newTencentCloudProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "cvm,vpc,cdn", "tencentcloud_vpc=id1:id2:id3")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "ap-guangzhou")
	return cmd
}

func newTencentCloudProvider() terraformutils.ProviderGenerator {
	return &tencentcloud_terraforming.TencentCloudProvider{}
}
