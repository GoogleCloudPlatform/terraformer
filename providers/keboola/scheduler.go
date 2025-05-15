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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

// SchedulerGenerator generates Terraform resources for keboola_scheduler
type SchedulerGenerator struct {
	terraformutils.Service
}

// InitResources fetches and initializes all scheduler configurations
func (g *SchedulerGenerator) InitResources() error {
	// These variables will be used when implementing the actual API calls
	_ = g.Args["token"].(string)
	_ = g.Args["host"].(string)

	// TODO: Initialize the Keboola client with apiKey and apiURL
	// client := keboola.NewClient(apiKey, apiURL)

	// TODO: Call the Keboola API to list all schedulers
	// For now, we'll just create a placeholder for one scheduler

	// Create a terraform resource for each scheduler
	g.Resources = append(g.Resources, terraformutils.NewResource(
		"sample-scheduler-id",   // ID of the scheduler
		"sample-scheduler-name", // Name of the scheduler (a readable name for reference)
		"keboola_scheduler",
		"keboola",
		map[string]string{
			"configuration_id": "123456", // Example configuration ID
			// Additional attributes will need to be filled in based on the actual API response
		},
		[]string{}, // Don't skip any attributes during import
		map[string]interface{}{},
	))

	return nil
}
