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

type ShovelGenerator struct {
	RBTService
}

type Shovel struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

type Shovels []Shovel

var ShovelAllowEmptyValues = []string{}
var ShovelAdditionalFields = map[string]interface{}{}

func (g ShovelGenerator) createResources(shovels Shovels) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, shovel := range shovels {
		if len(shovel.Name) == 0 {
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", shovel.Name, shovel.Vhost),
			fmt.Sprintf("shovel_%s_%s", normalizeResourceName(shovel.Vhost), normalizeResourceName(shovel.Name)),
			"rabbitmq_shovel",
			"rabbitmq",
			map[string]string{
				"name":  shovel.Name,
				"vhost": shovel.Vhost,
			},
			ShovelAllowEmptyValues,
			ShovelAdditionalFields,
		))
	}
	return resources
}

func (g *ShovelGenerator) InitResources() error {
	body, err := g.generateRequest("/api/shovels?columns=name,vhost")
	if err != nil {
		return err
	}
	var shovels Shovels
	err = json.Unmarshal(body, &shovels)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(shovels)
	return nil
}
