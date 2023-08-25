// Copyright 2022 The Terraformer Authors.
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
	sumologic_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/sumologic"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdSumoLogicImporter(options ImportOptions) *cobra.Command {
	var accessId, accessKey, environment string
	cmd := &cobra.Command{
		Use:   "sumologic",
		Short: "Import current state to Terraform configuration from Sumo Logic",
		Long:  "Import current state to Terraform configuration from Sumo Logic",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newSumoLogicProvider()
			err := Import(provider, options, []string{accessId, accessKey, environment})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSumoLogicProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "users", "users=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&accessId, "access-id", "", "", "Sumo Logic Access ID or env param SUMOLOGIC_ACCESS_ID")
	cmd.PersistentFlags().StringVarP(&accessKey, "access-key", "", "", "Sumo Logic Access Key or env param SUMOLOGIC_ACCESS_KEY")
	cmd.PersistentFlags().StringVarP(&environment, "environment", "", "", "Sumo Logic environment or env param SUMOLOGIC_ENVIRONMENT")
	return cmd
}

func newSumoLogicProvider() terraformutils.ProviderGenerator {
	return &sumologic_terraforming.SumoLogicProvider{}
}
