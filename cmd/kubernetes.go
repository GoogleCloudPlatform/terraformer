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
	"github.com/GoogleCloudPlatform/terraformer/kubernetes_terraforming"

	"github.com/spf13/cobra"
)

func newCmdKubernetesImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubernetes",
		Short: "Import current State to terraform configuration from kubernetes",
		Long:  "Import current State to terraform configuration from kubernetes",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := &kubernetes_terraforming.KubernetesProvider{}
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(&kubernetes_terraforming.KubernetesProvider{}))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "configmaps,deployments,services")
	cmd.PersistentFlags().StringVarP(&options.PathPatter, "path-patter", "p", DefaultPathPatter, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	return cmd
}
