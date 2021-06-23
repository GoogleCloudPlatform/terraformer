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
	vault_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/vault"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdVaultImporter(options ImportOptions) *cobra.Command {
	var token, address string
	cmd := &cobra.Command{
		Use:   "vault",
		Short: "Import current state to Terraform configuration from Vault",
		Long:  "Import current state to Terraform configuration from Vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newVaultProvider()
			err := Import(provider, options, []string{address, token})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newVaultProvider()))
	cmd.PersistentFlags().StringVarP(&address, "address", "a", "", "env param VAULT_ADDR")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "env param VAULT_TOKEN")
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")
	return cmd
}

func newVaultProvider() terraformutils.ProviderGenerator {
	return &vault_terraforming.Provider{}
}
