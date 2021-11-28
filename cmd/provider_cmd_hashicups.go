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
	"os"

	hashicups_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/hashicups"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdHashicupsImporter(options ImportOptions) *cobra.Command {
	username := ""
	password := ""
	cmd := &cobra.Command{
		Use:   "hashicups",
		Short: "Import current state to Terraform configuration from Hashicups",
		Long:  "Import current state to Terraform configuration from Hashicups",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newHashicupsProvider()
			err := Import(provider, options, []string{username, password})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newHashicupsProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "order", "order=1:2:4")
	cmd.PersistentFlags().StringVarP(&username, "username", "", os.Getenv("HASHICUPS_USERNAME"), "Hashicups username or env param HASHICUPS_USERNAME")
	cmd.PersistentFlags().StringVarP(&password, "password", "", os.Getenv("HASHICUPS_PASSWORD"), "Hashicups password or env param HASHICUPS_PASSWORD")
	return cmd
}

func newHashicupsProvider() terraformutils.ProviderGenerator {
	return &hashicups_terraforming.HashicupsProvider{}
}
