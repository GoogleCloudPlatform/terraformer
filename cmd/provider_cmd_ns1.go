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
	ns1_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/ns1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdNs1Importer(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ns1",
		Short: "Import current state to Terraform configuration from NS1",
		Long:  "Import current state to Terraform configuration from NS1",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newNs1Provider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newNs1Provider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "zone", "zone=id1:id2:id4")
	return cmd
}

func newNs1Provider() terraformutils.ProviderGenerator {
	return &ns1_terraforming.Ns1Provider{}
}
