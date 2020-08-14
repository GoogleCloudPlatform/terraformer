// Copyright 2020 The Terraformer Authors.
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
	gmailfilter_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/gmailfilter"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGmailfilterImporter(options ImportOptions) *cobra.Command {
	var creds, impersonatedUserEmail string
	cmd := &cobra.Command{
		Use:   "gmailfilter",
		Short: "Import current state to Terraform configuration from Gmail",
		Long:  "Import current state to Terraform configuration from Gmail",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newGmailfilterProvider()
			err := Import(provider, options, []string{
				creds,
				impersonatedUserEmail,
			})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newGmailfilterProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "label,filter", "label=name1:name2")
	cmd.PersistentFlags().StringVarP(&creds, "credentials", "", "", "/path/to/client_secret.json")
	cmd.PersistentFlags().StringVarP(&impersonatedUserEmail, "email", "", "", "foobar@example.com")
	return cmd
}

func newGmailfilterProvider() terraformutils.ProviderGenerator {
	return &gmailfilter_terraforming.GmailfilterProvider{}
}
