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
	"strings"

	gcp_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/gcp"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGoogleImporter(options ImportOptions) *cobra.Command {
	providerType := ""
	cmd := &cobra.Command{
		Use:   "google",
		Short: "Import current state to Terraform configuration from Google Cloud",
		Long:  "Import current state to Terraform configuration from Google Cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, project := range options.Projects {
				for _, region := range options.Regions {
					provider := newGoogleProvider()
					options.PathPattern = originalPathPattern
					options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}/{service}", "{provider}/"+project+"/{service}/"+region)
					log.Println(provider.GetName() + " importing project " + project + " region " + region)
					err := Import(provider, options, []string{region, project, providerType})
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGoogleProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "firewalls,networks", "compute_firewall=id1:id2:id4")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "z", []string{"global"}, "europe-west1,")
	cmd.PersistentFlags().StringSliceVarP(&options.Projects, "projects", "", []string{}, "")
	cmd.PersistentFlags().StringVarP(&providerType, "provider-type", "", "", "beta")
	_ = cmd.MarkPersistentFlagRequired("projects")
	return cmd
}

func newGoogleProvider() terraformutils.ProviderGenerator {
	return &gcp_terraforming.GCPProvider{}
}
