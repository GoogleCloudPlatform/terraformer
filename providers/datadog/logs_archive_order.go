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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// LogsArchiveOrderAllowEmptyValues ...
	LogsArchiveOrderAllowEmptyValues = []string{}
)

// LogsArchiveOrderGenerator ...
type LogsArchiveOrderGenerator struct {
	DatadogService
}

// InitResources Generate TerraformResources
func (g *LogsArchiveOrderGenerator) InitResources() error {
	g.Resources = append(g.Resources, terraformutils.NewResource(
		"archiveOrderID",
		"archiveOrderID",
		"datadog_logs_archive_order",
		"datadog",
		map[string]string{},
		LogsArchiveOrderAllowEmptyValues,
		map[string]interface{}{},
	))
	return nil
}
