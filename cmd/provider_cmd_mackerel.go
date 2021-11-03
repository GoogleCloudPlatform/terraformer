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
	mackerel_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/mackerel"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdMackerelImporter(options ImportOptions) *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "mackerel",
		Short: "Import current state to Terraform configuration from Mackerel",
		Long:  "Import current state to Terraform configuration from Mackerel",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newMackerelProvider()
			err := Import(provider, options, []string{apiKey})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newMackerelProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "service,role,aws_integration", "aws_integration=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "YOUR_MACKEREL_API_KEY or env param MACKEREL_API_KEY")
	return cmd
}

func newMackerelProvider() terraformutils.ProviderGenerator {
	return &mackerel_terraforming.MackerelProvider{}
}
