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
	"waze/terraformer/gcp_terraforming"

	"github.com/spf13/cobra"
)

func newCmdGoogleImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "google",
		Short: "Import current State to terraform configuration from google cloud",
		Long:  "Import current State to terraform configuration from google cloud",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPatter := options.PathPatter
			for _, project := range options.Projects {
				provider := &gcp_terraforming.GCPProvider{}
				options.PathPatter = originalPathPatter
				options.PathPatter = strings.Replace(options.PathPatter, "{provider}", "{provider}/"+project+"/", -1)
				log.Println(provider.GetName() + " importing project " + project)
				err := Import(provider, options, []string{options.Zone, project})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "firewalls,networks")
	cmd.PersistentFlags().StringVarP(&options.PathPatter, "path-patter", "p", DefaultPathPatter, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringVarP(&options.Zone, "zone", "z", "", "")
	cmd.PersistentFlags().StringSliceVarP(&options.Projects, "projects", "", []string{}, "")
	return cmd
}
