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
	prismacloud_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/prismacloud"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdPrismaCloudImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prismacloud",
		Short: "Import current state to Terraform configuration from Prisma Cloud",
		Long:  "Import current state to Terraform configuration from Prisma Cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newPrismaCloudProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.AddCommand(listCmd(newPrismaCloudProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "policy", "resource_group=id1:id2:id3")
	return cmd
}

func newPrismaCloudProvider() terraformutils.ProviderGenerator {
	return &prismacloud_terraforming.PrismaCloudProvider{}
}
