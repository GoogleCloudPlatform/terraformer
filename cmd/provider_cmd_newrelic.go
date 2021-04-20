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
	newrelic_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/newrelic"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdNewRelicImporter(options ImportOptions) *cobra.Command {
	apiKey := ""
	accountID := ""
	region := ""
	cmd := &cobra.Command{
		Use:   "newrelic",
		Short: "Import current state to Terraform configuration from New Relic",
		Long:  "Import current state to Terraform configuration from New Relic",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newNewRelicProvider()
			err := Import(provider, options, []string{apiKey, accountID, region})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newNewRelicProvider()))
	cmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Your Personal API Key")
	cmd.PersistentFlags().StringVar(&accountID, "account-id", "", "Your Account ID")
	cmd.PersistentFlags().StringVar(&region, "region", "US", "")
	baseProviderFlags(cmd.PersistentFlags(), &options, "alert", "dashboard=id1:id2:id4")
	return cmd
}

func newNewRelicProvider() terraformutils.ProviderGenerator {
	return &newrelic_terraforming.NewRelicProvider{}
}
