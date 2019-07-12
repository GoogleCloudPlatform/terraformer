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
	"github.com/GoogleCloudPlatform/terraformer/providers/gsuite"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdGSuiteImporter(options ImportOptions) *cobra.Command {
	var apiKey, appKey string
	cmd := &cobra.Command{
		Use:   "gsuite",
		Short: "Import current State to terraform configuration from gsuite",
		Long:  "Import current State to terraform configuration from gsuite",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newGSuiteProvider()
			err := Import(provider, options, []string{apiKey, appKey})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGSuiteProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "monitors,users")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "YOUR_DATADOG_API_KEY or env param DATADOG_API_KEY")
	cmd.PersistentFlags().StringVarP(&appKey, "app-key", "", "", "YOUR_DATADOG_APP_KEY or env param DATADOG_APP_KEY")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "datadog_monitor=id1:id2:id4")
	return cmd
}

func newGSuiteProvider() terraform_utils.ProviderGenerator {
	return &gsuite.GSuiteProvider{}
}
