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
	datadog_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/datadog"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdDatadogImporter(options ImportOptions) *cobra.Command {
	var apiKey, appKey, apiURL, validate string
	cmd := &cobra.Command{
		Use:   "datadog",
		Short: "Import current state to Terraform configuration from Datadog",
		Long:  "Import current state to Terraform configuration from Datadog",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newDataDogProvider()
			err := Import(provider, options, []string{apiKey, appKey, apiURL, validate})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newDataDogProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "monitors,users", "monitor=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "YOUR_DATADOG_API_KEY or env param DATADOG_API_KEY")
	cmd.PersistentFlags().StringVarP(&appKey, "app-key", "", "", "YOUR_DATADOG_APP_KEY or env param DATADOG_APP_KEY")
	cmd.PersistentFlags().StringVarP(&apiURL, "api-url", "", "", "YOUR_DATADOG_API_URL or env param DATADOG_HOST")
	cmd.PersistentFlags().StringVar(&validate, "validate", "", "bool-parsable values only or env param DATADOG_VALIDATE. Enables validation of the provided API and APP keys during provider initialization. Default is true. When false, api_key and app_key won't be checked")
	return cmd
}

func newDataDogProvider() terraformutils.ProviderGenerator {
	return &datadog_terraforming.DatadogProvider{}
}
