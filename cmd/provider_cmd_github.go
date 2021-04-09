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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGithubImporter(options ImportOptions) *cobra.Command {
	token := ""
	baseURL := ""
	owner := []string{}
	cmd := &cobra.Command{
		Use:   "github",
		Short: "Import current state to Terraform configuration from GitHub",
		Long:  "Import current state to Terraform configuration from GitHub",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, organization := range owner {
				provider := newGitHubProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}", "{provider}/"+organization)
				log.Println(provider.GetName() + " importing organization " + organization)
				err := Import(provider, options, []string{organization, token, baseURL})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGitHubProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "repository", "repository=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "YOUR_GITHUB_TOKEN or env param GITHUB_TOKEN")
	cmd.PersistentFlags().StringSliceVarP(&owner, "owner", "", []string{}, "")
	cmd.PersistentFlags().StringVarP(&baseURL, "base-url", "", "", "")
	return cmd
}

func newGitHubProvider() terraformutils.ProviderGenerator {
	return &github_terraforming.GithubProvider{}
}
