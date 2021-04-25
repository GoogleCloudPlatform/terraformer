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
	"strconv"

	kubernetes_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/kubernetes"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdKubernetesImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubernetes",
		Short: "Import current state to Terraform configuration from Kubernetes",
		Long:  "Import current state to Terraform configuration from Kubernetes",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newKubernetesProvider()
			err := Import(provider, options, []string{strconv.FormatBool(options.Verbose)})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newKubernetesProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "configmaps,deployments,services", "deployment=name1:name2:name3")
	return cmd
}

func newKubernetesProvider() terraformutils.ProviderGenerator {
	return &kubernetes_terraforming.KubernetesProvider{}
}
