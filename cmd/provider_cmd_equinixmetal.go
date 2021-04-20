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
	equinixmetal_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/equinixmetal"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdEquinixMetalImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metal",
		Short: "Import current state to Terraform configuration from Equinix Metal",
		Long:  "Import current state to Terraform configuration from Equinix Metal",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newEquinixMetalProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newEquinixMetalProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "project,device", "project=name1:name2:name3")

	return cmd
}

func newEquinixMetalProvider() terraformutils.ProviderGenerator {
	return &equinixmetal_terraforming.EquinixMetalProvider{}
}
