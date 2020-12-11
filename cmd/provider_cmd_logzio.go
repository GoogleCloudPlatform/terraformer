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

	logzio_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/logzio"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

const (
	defaultBaseURL = "https://api.logz.io"
)

func newCmdLogzioImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logzio",
		Short: "Import current state to Terraform configuration from Logz.io",
		Long:  "Import current state to Terraform configuration from Logz.io",
		RunE: func(cmd *cobra.Command, args []string) error {
			token := os.Getenv("LOGZIO_API_TOKEN")
			if len(token) == 0 {
				return errors.New("API Token for Logz.io must be set through `LOGZIO_API_TOKEN` env var")
			}
			baseURL := os.Getenv("LOGZIO_BASE_URL")
			if len(baseURL) == 0 {
				baseURL = defaultBaseURL
			}

			provider := newLogzioProvider()
			err := Import(provider, options, []string{token, baseURL})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newLogzioProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "repository", "alert=id1:id2:id4")
	return cmd
}

func newLogzioProvider() terraformutils.ProviderGenerator {
	return &logzio_terraforming.LogzioProvider{}
}
