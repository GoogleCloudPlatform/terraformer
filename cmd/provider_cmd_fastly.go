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
	fastly_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/fastly"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdFastlyImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fastly",
		Short: "Import current state to Terraform configuration from Fastly",
		Long:  "Import current state to Terraform configuration from Fastly",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newFastlyProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newFastlyProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "service_v1", "service_v1=id1:id2:id3")
	return cmd
}

func newFastlyProvider() terraformutils.ProviderGenerator {
	return &fastly_terraforming.FastlyProvider{}
}
