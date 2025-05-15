// Copyright 2025 The Terraformer Authors.
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

package keboola

import (
	"context"
	"time"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

// ComponentConfigurationGenerator generates Terraform resources for keboola_component_configuration
type ComponentConfigurationGenerator struct {
	terraformutils.Service
}

// InitResources fetches and initializes all component configurations
func (g *ComponentConfigurationGenerator) InitResources() error {
	// These variables will be used when implementing the actual API calls
	_ = g.Args["host"].(string)
	_ = g.Args["token"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	client, err := keboola.NewAuthorizedAPI(ctx, g.Args["token"].(string), g.Args["host"].(string))
	if err != nil {
		return err
	}

	branch, err := client.GetDefaultBranchRequest().Send(ctx)
	if err != nil {
		return err
	}

	components, err := client.ListConfigsAndRowsFrom(
		branch.BranchKey,
	).Send(ctx)
	if err != nil {
		return err
	}

	for _, component := range *components {
		for _, config := range component.Configs {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				config.ID.String(),
				config.Name,
				"keboola_component_configuration",
				"keboola",
				map[string]string{
					"component_id": component.ID.String(),
				},
				[]string{}, // Don't skip any attributes during import
				map[string]interface{}{},
			))
		}
	}
	return nil
}
