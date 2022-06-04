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
	snowflake_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/snowflake"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdSnowflakeImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snowflake",
		Short: "Import current State to terraform configuration from Snowflake",
		Long:  "Import current State to terraform configuration from Snowflake",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newSnowflakeProvider()
			return Import(provider, options, []string{})
		},
	}
	cmd.AddCommand(listCmd(newSnowflakeProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "database")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "snowflake_database=name1:name2:name3")

	return cmd
}

func newSnowflakeProvider() terraformutils.ProviderGenerator {
	return &snowflake_terraforming.SnowflakeProvider{}
}
