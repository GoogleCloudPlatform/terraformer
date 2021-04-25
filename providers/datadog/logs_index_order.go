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

package datadog

import (
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// LogsIndexOrderAllowEmptyValues ...
	LogsIndexOrderAllowEmptyValues = []string{}
)

// LogsIndexOrderGenerator ...
type LogsIndexOrderGenerator struct {
	DatadogService
}

// InitResources Generate TerraformResources
func (g *LogsIndexOrderGenerator) InitResources() error {
	currentDate := time.Now().Format("20060102150405")
	resourceName := fmt.Sprintf("logs_index_order_%s", currentDate)
	g.Resources = append(g.Resources, terraformutils.NewResource(
		resourceName,
		resourceName,
		"datadog_logs_index_order",
		"datadog",
		map[string]string{
			"name": resourceName,
		},
		LogsIndexOrderAllowEmptyValues,
		map[string]interface{}{},
	))
	return nil
}
