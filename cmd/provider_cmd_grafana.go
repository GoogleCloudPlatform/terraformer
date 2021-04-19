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

	"github.com/GoogleCloudPlatform/terraformer/providers/grafana"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGrafanaImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grafana",
		Short: "Import current state to Terraform configuration from grafana",
		Long:  "Import current state to Terraform configuration from grafana",
		RunE: func(cmd *cobra.Command, args []string) error {
			auth := os.Getenv("GRAFANA_AUTH")
			if len(auth) == 0 {
				return errors.New("API Token must be set through `GRAFANA_AUTH` env var")
			}
			baseURL := os.Getenv("GRAFANA_URL")
			if len(auth) == 0 {
				return errors.New("grafana url must be set through `GRAFANA_URL` env var")
			}
			provider := newGrafanaProvider()
			err := Import(provider, options, []string{auth, baseURL})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGrafanaProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "grafana_folders", "alert=id1:id2:id4")
	return cmd
}

func newGrafanaProvider() terraformutils.ProviderGenerator {
	return &grafana.GrafanaProvider{}
}
