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

	keycloak_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/keycloak"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

const (
	defaultKeycloakEndpoint = "https://localhost:8443"
	defaultKeycloakRealm    = "master"
)

func newCmdKeycloakImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keycloak",
		Short: "Import current state to Terraform configuration from Keycloak",
		Long:  "Import current state to Terraform configuration from Keycloak",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := os.Getenv("KEYCLOAK_URL")
			if len(url) == 0 {
				url = defaultKeycloakEndpoint
			}
			clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
			clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
			realm := os.Getenv("KEYCLOAK_REALM")
			if len(realm) == 0 {
				realm = defaultKeycloakRealm
			}
			provider := newKeycloakProvider()
			err := Import(provider, options, []string{url, clientID, clientSecret, realm})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newKeycloakProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "realms", "keycloak_type=id1:id2:id4")
	return cmd
}

func newKeycloakProvider() terraform_utils.ProviderGenerator {
	return &keycloak_terraforming.KeycloakProvider{}
}
