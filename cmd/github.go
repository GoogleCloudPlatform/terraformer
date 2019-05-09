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

	github_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/github"

	"github.com/spf13/cobra"
)

func newCmdGithubImporter(options ImportOptions) *cobra.Command {
	token := ""
	organizations := []string{}
	cmd := &cobra.Command{
		Use:   "github",
		Short: "Import current State to terraform configuration from github",
		Long:  "Import current State to terraform configuration from github",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPatter := options.PathPatter
			for _, organization := range organizations {
				provider := &github_terraforming.GithubProvider{}
				options.PathPatter = originalPathPatter
				options.PathPatter = strings.Replace(options.PathPatter, "{provider}", "{provider}/"+organization, -1)
				log.Println(provider.GetName() + " importing organization " + organization)
				err := Import(provider, options, []string{organization, token})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(&github_terraforming.GithubProvider{}))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "repository")
	cmd.PersistentFlags().StringVarP(&options.PathPatter, "path-patter", "p", DefaultPathPatter, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "YOUR_GITHUB_TOKEN or env param GITHUB_TOKEN")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "github_repository=id1:id2:id4")
	cmd.PersistentFlags().StringSliceVarP(&organizations, "organizations", "", []string{}, "")
	return cmd
}
