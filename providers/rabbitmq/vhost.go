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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type VhostGenerator struct {
	RBTService
}

type Vhost struct {
	Name string `json:"name"`
}

type Vhosts []Vhost

var VhostAllowEmptyValues = []string{}

func (g VhostGenerator) createResources(vhosts Vhosts) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, vhost := range vhosts {
		resources = append(resources, terraformutils.NewSimpleResource(
			vhost.Name,
			"vhost_"+normalizeResourceName(vhost.Name),
			"rabbitmq_vhost",
			"rabbitmq",
			VhostAllowEmptyValues,
		))
	}
	return resources
}

func (g *VhostGenerator) InitResources() error {
	body, err := g.generateRequest("/api/vhosts?columns=name")
	if err != nil {
		return err
	}
	var vhosts Vhosts
	err = json.Unmarshal(body, &vhosts)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(vhosts)
	return nil
}
