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
	"log"

	openstack_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/openstack"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdOpenStackImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openstack",
		Short: "Import current state to Terraform configuration from OpenStack",
		Long:  "Import current state to Terraform configuration from OpenStack",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, region := range options.Regions {
				provider := newOpenStackProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern += region + "/"
				log.Println(provider.GetName() + " importing region " + region)
				err := Import(provider, options, []string{region})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newOpenStackProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "compute,networking", "compute_instance_v2=id1:id2:id4")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "", []string{}, "RegionOne")
	return cmd
}

func newOpenStackProvider() terraformutils.ProviderGenerator {
	return &openstack_terraforming.OpenStackProvider{}
}
