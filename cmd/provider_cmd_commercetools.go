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
	"errors"
	"os"

	commercetools_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/commercetools"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

const (
	defaultCommercetoolsBaseURL  = "https://api.sphere.io"
	defaultCommercetoolsTokenURL = "https://auth.sphere.io"
)

func newCmdCommercetoolsImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commercetools",
		Short: "Import current state to Terraform configuration from Commercetools",
		Long:  "Import current state to Terraform configuration from Commercetools",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientID := os.Getenv("CTP_CLIENT_ID")
			if len(clientID) == 0 {
				return errors.New("API client ID for commercetools must be set through `CTP_CLIENT_ID` env var")
			}
			clientScope := os.Getenv("CTP_CLIENT_SCOPE")
			if len(clientScope) == 0 {
				return errors.New("API client scope for comercetools must be set through `CTP_CLIENT_SCOPE` env var")
			}
			clientSecret := os.Getenv("CTP_CLIENT_SECRET")
			if len(clientSecret) == 0 {
				return errors.New("API client secret for comercetools must be set through `CTP_CLIENT_SECRET` env var")
			}
			projectKey := os.Getenv("CTP_PROJECT_KEY")
			if len(projectKey) == 0 {
				return errors.New("API project key for comercetools must be set through `CTP_PROJECT_KEY` env var")
			}
			baseURL := os.Getenv("CTP_BASE_URL")
			if len(baseURL) == 0 {
				baseURL = defaultCommercetoolsBaseURL
			}
			tokenURL := os.Getenv("CTP_TOKEN_URL")
			if len(tokenURL) == 0 {
				tokenURL = defaultCommercetoolsTokenURL
			}
			provider := newCommercetoolsProvider()
			err := Import(provider, options, []string{clientID, clientScope, clientSecret, projectKey, baseURL, tokenURL})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newCommercetoolsProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "types", "type=id1:id2:id4")
	return cmd
}

func newCommercetoolsProvider() terraformutils.ProviderGenerator {
	return &commercetools_terraforming.CommercetoolsProvider{}
}
