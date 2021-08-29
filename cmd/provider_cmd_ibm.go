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
	ibm_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/ibm"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdIbmImporter(options ImportOptions) *cobra.Command {
	var resourceGroup string
	var region string
	var cis string
	cmd := &cobra.Command{
		Use:   "ibm",
		Short: "Import current state to Terraform configuration from ibm",
		Long:  "Import current state to Terraform configuration from ibm",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newIbmProvider()
			err := Import(provider, options, []string{resourceGroup, region, cis})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newIbmProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "server", "ibm_server=name1:name2:name3")
	cmd.PersistentFlags().StringVarP(&resourceGroup, "resource_group", "", "", "resource_group=default")
	cmd.PersistentFlags().StringVarP(&region, "region", "R", "", "region=us-south")
	cmd.PersistentFlags().StringVarP(&cis, "cis", "", "", "cis=TestCIS")
	return cmd
}

func newIbmProvider() terraformutils.ProviderGenerator {
	return &ibm_terraforming.IBMProvider{}
}
