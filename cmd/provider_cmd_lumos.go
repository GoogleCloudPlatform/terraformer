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
	lumos_terraformer "github.com/GoogleCloudPlatform/terraformer/providers/lumos"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdLumosImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lumos",
		Short: "Import current state to Terraform configuration from lumos.com",
		Long:  "Import current state to Terraform configuration from lumos.com",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newLumosProvider()
			err := Import(provider, options, options.Projects)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newLumosProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")

	return cmd
}

func newLumosProvider() terraformutils.ProviderGenerator {
	return &lumos_terraformer.LumosProvider{}
}
