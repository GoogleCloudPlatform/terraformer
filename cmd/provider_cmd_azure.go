// Copyright 2019 The Terraformer Authors.
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
	azure_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/azure"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdAzureImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure",
		Short: "Import current state to Terraform configuration from Azure",
		Long:  "Import current state to Terraform configuration from Azure",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newAzureProvider()
			err := Import(provider, options, []string{options.ResourceGroup})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newAzureProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "resource_group", "resource_group=name1:name2:name3")
	cmd.PersistentFlags().StringVarP(&options.ResourceGroup, "resource-group", "R", "", "")
	return cmd
}

func newAzureProvider() terraformutils.ProviderGenerator {
	return &azure_terraforming.AzureProvider{}
}
