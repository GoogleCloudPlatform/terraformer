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
	azuredevops "github.com/GoogleCloudPlatform/terraformer/providers/azuredevops"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdAzureDevOpsImporter(options ImportOptions) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "azuredevops",
		Short: "Import current state to Terraform configuration from Azure DevOps",
		Long:  "Import current state to Terraform configuration from Azure DevOps",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newAzureDevOpsProvider()
			err := Import(provider, options, []string{options.ResourceGroup})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAzureDevOpsProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "project,team,git", "project=name1:name2:name3")
	return cmd
}

func newAzureDevOpsProvider() terraformutils.ProviderGenerator {
	return &azuredevops.AzureDevOpsProvider{}
}
