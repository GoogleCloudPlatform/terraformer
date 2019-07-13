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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdGoogleImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "google",
		Short: "Import current State to terraform configuration from google cloud",
		Long:  "Import current State to terraform configuration from google cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, project := range options.Projects {
				for _, region := range options.Regions {
					provider := newGCPProvider()
					options.PathPattern = originalPathPattern
					options.PathPattern = strings.Replace(options.PathPattern, "{provider}/{service}", "{provider}/"+project+"/{service}/"+region, -1)
					log.Println(provider.GetName() + " importing project " + project + " region " + region)
					err := Import(provider, options, []string{region, project})
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGCPProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "firewalls,networks")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringSliceVarP(&options.Regions, "regions", "z", []string{"global"}, "europe-west1,")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "google_compute_firewall=id1:id2:id4")
	cmd.PersistentFlags().StringSliceVarP(&options.Projects, "projects", "", []string{}, "")
	return cmd
}

func newGCPProvider() terraform_utils.ProviderGenerator {
	return &gcp_terraforming.GCPProvider{}
}
