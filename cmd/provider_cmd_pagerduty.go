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
	pagerduty_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/pagerduty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdPagerDutyImporter(options ImportOptions) *cobra.Command {
	token := ""
	cmd := &cobra.Command{
		Use:   "pagerduty",
		Short: "Import current state to Terraform configuration from PagerDuty",
		Long:  "Import current state to Terraform configuration from PagerDuty",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newPagerDutyProvider()
			err := Import(provider, options, []string{token})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newPagerDutyProvider()))
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "env param PAGERDUTY_TOKEN")
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "user=id1:id2:id4")
	return cmd
}

func newPagerDutyProvider() terraformutils.ProviderGenerator {
	return &pagerduty_terraforming.PagerDutyProvider{}
}
