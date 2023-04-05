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

package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type BindingGenerator struct {
	RBTService
}

type Binding struct {
	Source          string                 `json:"source"`
	Vhost           string                 `json:"vhost"`
	Destination     string                 `json:"destination"`
	DestinationType string                 `json:"destination_type"`
	PropertiesKey   string                 `json:"properties_key"`
	Arguments       map[string]interface{} `json:"arguments"`
}

type Bindings []Binding

var BindingAllowEmptyValues = []string{"source"}
var BindingAdditionalFields = map[string]interface{}{}

func (g BindingGenerator) createResources(bindings Bindings) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, binding := range bindings {
		argumentsJSON, errArgumentsJSON := json.Marshal(binding.Arguments)
		if errArgumentsJSON != nil {
			argumentsJSON = []byte("{}")
		}
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s/%s/%s/%s/%s", percentEncodeSlashes(binding.Vhost), binding.Source, binding.Destination, binding.DestinationType, binding.PropertiesKey),
			fmt.Sprintf("binding_%s_%s_%s_%s_%s", normalizeResourceName(binding.Source), normalizeResourceName(binding.Vhost), normalizeResourceName(binding.Destination), normalizeResourceName(binding.DestinationType), normalizeResourceName(binding.PropertiesKey)),
			"rabbitmq_binding",
			"rabbitmq",
			map[string]string{
				"source":           binding.Source,
				"vhost":            binding.Vhost,
				"destination":      binding.Destination,
				"destination_type": binding.DestinationType,
				"properties_key":   binding.PropertiesKey,
				"arguments_json":   string(argumentsJSON),
			},
			BindingAllowEmptyValues,
			BindingAdditionalFields,
		))
	}
	return resources
}

func (g *BindingGenerator) InitResources() error {
	body, err := g.generateRequest("/api/bindings?columns=source,vhost,destination,destination_type,properties_key,arguments")
	if err != nil {
		return err
	}
	var bindings Bindings
	err = json.Unmarshal(body, &bindings)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(bindings)
	return nil
}
