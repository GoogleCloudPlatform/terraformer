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
	"errors"
	"os"

	auth0_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/auth0"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdAuth0Importer(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth0",
		Short: "Import current state to Terraform configuration from Auth0",
		Long:  "Import current state to Terraform configuration from Auth0",
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := os.Getenv("AUTH0_DOMAIN")
			if len(domain) == 0 {
				return errors.New("Domain for Auth0 must be set through `AUTH0_DOMAIN` env var")
			}
			clientID := os.Getenv("AUTH0_CLIENT_ID")
			if len(clientID) == 0 {
				return errors.New("Client ID for Auht0 must be set through `AUTH0_CLIENT_ID` env var")
			}
			clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
			if len(clientSecret) == 0 {
				return errors.New("Clien Secret for Auth0 must be set through `AUTH0_CLIENT_SECRET` env var")
			}

			provider := newAuth0Provider()
			err := Import(provider, options, []string{domain, clientID, clientSecret})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newAuth0Provider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "action", "action=name1:name2:name3")
	return cmd
}

func newAuth0Provider() terraformutils.ProviderGenerator {
	return &auth0_terraforming.Auth0Provider{}
}
