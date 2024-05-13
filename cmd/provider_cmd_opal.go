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
	opal_terraformer "github.com/GoogleCloudPlatform/terraformer/providers/opal"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdOpalImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "opal",
		Short: "Import current state to Terraform configuration from opal.dev",
		Long:  "Import current state to Terraform configuration from opal.dev",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newOpalProvider()
			err := Import(provider, options, options.Projects)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newOpalProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")

	return cmd
}

func newOpalProvider() terraformutils.ProviderGenerator {
	return &opal_terraformer.OpalProvider{}
}
