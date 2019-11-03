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
	"os"

	rabbitmq_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/rabbitmq"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

const (
	defaultRabbitMQEndpoint = "http://localhost:15672"
)

func newCmdRabbitMQImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rabbitmq",
		Short: "Import current state to Terraform configuration from RabbitMQ",
		Long:  "Import current state to Terraform configuration from RabbitMQ",
		RunE: func(cmd *cobra.Command, args []string) error {
			endpoint := os.Getenv("RABBITMQ_SERVER_URL")
			if len(endpoint) == 0 {
				endpoint = defaultRabbitMQEndpoint
			}
			username := os.Getenv("RABBITMQ_USERNAME")
			password := os.Getenv("RABBITMQ_PASSWORD")
			provider := newRabbitMQProvider()
			err := Import(provider, options, []string{endpoint, username, password})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newRabbitMQProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", true, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "vhosts")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	cmd.PersistentFlags().StringSliceVarP(&options.Filter, "filter", "f", []string{}, "rabbitmq_type=id1:id2:id4")
	return cmd
}

func newRabbitMQProvider() terraform_utils.ProviderGenerator {
	return &rabbitmq_terraforming.RBTProvider{}
}
