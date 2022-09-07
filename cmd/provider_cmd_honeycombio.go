// Copyright 2022 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	honeycombio_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/honeycombio"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdHoneycombioImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "honeycombio",
		Short: "Import current state to Terraform configuration from Honeycomb.io",
		Long:  "Import current state to Terraform configuration from Honeycomb.io",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newHoneycombProvider()
			err := Import(provider, options, options.Projects)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newHoneycombProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "derived_column,board", "board=id1,id2")
	cmd.PersistentFlags().StringSliceVarP(&options.Projects, "datasets", "", []string{}, "hello-service,goodbye-service")

	return cmd
}

func newHoneycombProvider() terraformutils.ProviderGenerator {
	return &honeycombio_terraforming.HoneycombProvider{}
}
