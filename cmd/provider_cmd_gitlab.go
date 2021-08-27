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

	gitLab_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/gitlab"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdGitLabImporter(options ImportOptions) *cobra.Command {
	token := ""
	baseURL := ""
	groups := []string{}
	cmd := &cobra.Command{
		Use:   "gitlab",
		Short: "Import current state to Terraform configuration from GitLab",
		Long:  "Import current state to Terraform configuration from GitLab",
		RunE: func(cmd *cobra.Command, args []string) error {
			originalPathPattern := options.PathPattern
			for _, group := range groups {
				provider := newGitLabProvider()
				options.PathPattern = originalPathPattern
				options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}", "{provider}/"+group)
				log.Println(provider.GetName() + " importing group " + group)
				err := Import(provider, options, []string{group, token, baseURL})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newGitLabProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "repository", "repository=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "YOUR_GITLAB_TOKEN or env param GITLAB_TOKEN")
	cmd.PersistentFlags().StringSliceVarP(&groups, "group", "", []string{}, "paths to groups")
	cmd.PersistentFlags().StringVarP(&baseURL, "base-url", "", "", "")
	return cmd
}

func newGitLabProvider() terraformutils.ProviderGenerator {
	return &gitLab_terraforming.GitLabProvider{}
}
