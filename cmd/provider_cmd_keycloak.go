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
	"log"
	"os"
	"strconv"
	"strings"

	keycloak_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/keycloak"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

const (
	defaultKeycloakEndpoint              = "https://localhost:8443"
	defaultKeycloakRealm                 = "master"
	defaultKeycloakClientTimeout         = int64(30)
	defaultKeycloakTLSInsecureSkipVerify = false
)

func newCmdKeycloakImporter(options ImportOptions) *cobra.Command {
	targets := []string{}
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
			clientTimeout, err := strconv.ParseInt(os.Getenv("KEYCLOAK_CLIENT_TIMEOUT"), 10, 64)
			if err != nil {
				clientTimeout = defaultKeycloakClientTimeout
			}
			tlsInsecureSkipVerify, err := strconv.ParseBool(os.Getenv("KEYCLOAK_TLS_INSECURE_SKIP_VERIFY"))
			if err != nil {
				tlsInsecureSkipVerify = defaultKeycloakTLSInsecureSkipVerify
			}
			caCert := os.Getenv("KEYCLOAK_CACERT")
			if len(caCert) == 0 {
				caCert = "-"
			}
			if len(targets) > 0 {
				originalPathPattern := options.PathPattern
				for _, target := range targets {
					provider := newKeycloakProvider()
					log.Println(provider.GetName() + " importing realm " + target)
					options.PathPattern = originalPathPattern
					options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}", "{provider}/"+target)
					err := Import(provider, options, []string{url, clientID, clientSecret, realm, strconv.FormatInt(clientTimeout, 10), caCert, strconv.FormatBool(tlsInsecureSkipVerify), target})
					if err != nil {
						return err
					}
				}
			} else {
				provider := newKeycloakProvider()
				log.Println(provider.GetName() + " importing all realms")
				err := Import(provider, options, []string{url, clientID, clientSecret, realm, strconv.FormatInt(clientTimeout, 10), caCert, strconv.FormatBool(tlsInsecureSkipVerify), "-"})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newKeycloakProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "realms", "type=id1:id2:id4")
	cmd.PersistentFlags().StringSliceVarP(&targets, "targets", "", []string{}, "")
	return cmd
}

func newKeycloakProvider() terraformutils.ProviderGenerator {
	return &keycloak_terraforming.KeycloakProvider{}
}
