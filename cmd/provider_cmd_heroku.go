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
	heroku_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/heroku"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdHerokuImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "heroku",
		Short: "Import current state to Terraform configuration from Heroku",
		Long:  "Import current state to Terraform configuration from Heroku",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newHerokuProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newHerokuProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "app,addon", "app=name1:name2:name3")
	return cmd
}

func newHerokuProvider() terraformutils.ProviderGenerator {
	return &heroku_terraforming.HerokuProvider{}
}
