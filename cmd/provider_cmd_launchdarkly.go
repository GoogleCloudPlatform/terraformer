// Copyright 2021 The Terraformer Authors.
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
	launchdarkly_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/launchdarkly"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdLaunchDarklyImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launchdarkly",
		Short: "Import current state to Terraform configuration from LaunchDarkly",
		Long:  "Import current state to Terraform configuration from LaunchDarkly",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newLaunchDarklyProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newLaunchDarklyProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "project", "launchdarkly_project=id1:id2:id3")
	return cmd
}

func newLaunchDarklyProvider() terraformutils.ProviderGenerator {
	return &launchdarkly_terraforming.LaunchDarklyProvider{}
}
