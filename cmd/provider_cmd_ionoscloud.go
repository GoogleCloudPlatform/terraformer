// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	ionoscloud_terraformer "github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdIonosCloudImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ionoscloud",
		Short: "Import current state to Terraform configuration from IONOS Cloud",
		Long:  "Import current state to Terraform configuration from IONOS Cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newIonosCloudProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newIonosCloudProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "app,addon", "app=name1:name2:name3")
	return cmd
}

func newIonosCloudProvider() terraformutils.ProviderGenerator {
	return &ionoscloud_terraformer.IonosCloudProvider{}
}
