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
	"fmt"
	"os"

	authlete_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/authlete"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/terraformerstring"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func newCmdAuthleteImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authlete",
		Short: "Import current state to Terraform configuration from Authlete",
		Long:  "Import current state to Terraform configuration from Authlete",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiServer := os.Getenv("AUTHLETE_API_SERVER")
			if len(apiServer) == 0 {
				apiServer = "https://api.authlete.com"
			}
			soKey := os.Getenv("AUTHLETE_SO_KEY")
			soSecret := os.Getenv("AUTHLETE_SO_SECRET")
			apiKey := os.Getenv("AUTHLETE_API_KEY")
			apiSecret := os.Getenv("AUTHLETE_API_SECRET")

			if terraformerstring.ContainsString(options.Resources, "authlete_service") ||
				terraformerstring.ContainsString(options.Resources, "*") {
				if len(soKey) == 0 {
					return errors.New("Service Owner Key for Authlete must be set through `AUTHLETE_SO_KEY` env var in order to import the services")
				}
				if len(soSecret) == 0 {
					return errors.New("Service Owner Secret for Authlete must be set through `AUTHLETE_SO_SECRET` env var in order to import the services")
				}
			}
			if terraformerstring.ContainsString(options.Resources, "authlete_client") ||
				terraformerstring.ContainsString(options.Resources, "*") {

				if len(apiKey) == 0 {
					return errors.New("API Key for Authlete must be set through `AUTHLETE_API_KEY` env var in order to import the clients")
				}
				if len(apiSecret) == 0 {
					return errors.New("API Secret for Authlete must be set through `AUTHLETE_API_SECRET` env var in order to import the clients")
				}
			}
			provider := newAuthleteProvider()
			err := Import(provider, options, []string{apiServer, soKey, soSecret, apiKey, apiSecret})
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.AddCommand(listAuthleteCmd())
	baseProviderFlags(cmd.PersistentFlags(), &options, "authlete_service", "authlete_service=apikey1:apikey2:apikey3")
	return cmd
}

func newAuthleteProvider() terraformutils.ProviderGenerator {
	return &authlete_terraforming.AuthleteProvider{}
}

func listAuthleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List supported resources for authlete provider",
		Long:  "List supported resources for authlete provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			services := []string{"authlete_service", "authlete_client"}
			for _, k := range services {
				fmt.Println(k)
			}
			return nil
		},
	}
	cmd.Flags().AddFlag(&pflag.Flag{Name: "resources"})
	return cmd
}
