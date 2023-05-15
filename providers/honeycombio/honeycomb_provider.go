// Copyright 2022 The Terraformer Authors.
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

package honeycombio

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

const honeycombDefaultURL = "https://api.honeycomb.io"
const honeycombTerraformerProviderVersion = "0.0.2"

type HoneycombProvider struct { //nolint
	terraformutils.Provider
	apiKey   string
	apiURL   string
	datasets []string
}

func (p HoneycombProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"honeycomb": map[string]interface{}{
				"api_url": p.apiURL,
			},
		},
	}
}

func (p *HoneycombProvider) GetName() string {
	return "honeycombio"
}

// This mapping will stop working if queries/query annotations are generated as
// sub-resources of boards or triggers
func (p HoneycombProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"board": {
			"dataset":          {"query.dataset", "name"},
			"query":            {"query.query_id", "id"},
			"query_annotation": {"query.query_annotation_id", "id"},
		},
		"column": {
			"dataset": {"dataset", "name"},
		},
		"derived_column": {
			"dataset": {"dataset", "name"},
		},
		"query": {
			"dataset": {"dataset", "name"},
		},
		"query_annotation": {
			"query":   {"query_id", "id"},
			"dataset": {"dataset", "name"},
		},
		"slo": {
			"dataset": {"dataset", "name"},
		},
		"burn_alert": {
			"slo":     {"slo_id", "id"},
			"dataset": {"dataset", "name"},
		},
		"trigger": {
			"query":   {"query_id", "id"},
			"dataset": {"dataset", "name"},
		},
	}
}
func (p *HoneycombProvider) Init(args []string) error {
	p.apiKey = os.Getenv("HONEYCOMB_API_KEY")
	if p.apiKey == "" {
		return errors.New("the Honeycomb API key must be set via `HONEYCOMB_API_KEY` env var")
	}
	p.apiURL = os.Getenv("HONEYCOMB_API_URL")
	if p.apiURL == "" {
		p.apiURL = honeycombDefaultURL
	}
	// datasets are the only argument
	p.datasets = args

	return nil
}

func (p *HoneycombProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(p.apiKey),
		"api_url": cty.StringVal(p.apiURL),
	})
}

func (p *HoneycombProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("honeycombio: " + serviceName + " is not a supported resource type")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key":  p.apiKey,
		"api_url":  p.apiURL,
		"datasets": p.datasets,
	})
	return nil
}

func (p *HoneycombProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"board":            &BoardGenerator{},
		"derived_column":   &DerivedColumnGenerator{},
		"trigger":          &TriggerGenerator{},
		"dataset":          &DatasetGenerator{},
		"column":           &ColumnGenerator{},
		"query":            &QueryGenerator{},
		"query_annotation": &QueryAnnotationGenerator{},
		"slo":              &SLOGenerator{},
		"burn_alert":       &BurnAlertGenerator{},
	}
}
