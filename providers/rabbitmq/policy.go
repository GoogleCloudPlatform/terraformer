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

type PolicyGenerator struct {
	RBTService
}

type Policy struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

type Policies []Policy

var PolicyAllowEmptyValues = []string{}
var PolicyAdditionalFields = map[string]interface{}{}

func (g PolicyGenerator) createResources(policies Policies) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, policy := range policies {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", policy.Name, policy.Vhost),
			fmt.Sprintf("policy_%s_%s", normalizeResourceName(policy.Vhost), normalizeResourceName(policy.Name)),
			"rabbitmq_policy",
			"rabbitmq",
			map[string]string{
				"name":  policy.Name,
				"vhost": policy.Vhost,
			},
			PolicyAllowEmptyValues,
			PolicyAdditionalFields,
		))
	}
	return resources
}

func (g *PolicyGenerator) InitResources() error {
	body, err := g.generateRequest("/api/policies?columns=name,vhost")
	if err != nil {
		return err
	}
	var policies Policies
	err = json.Unmarshal(body, &policies)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(policies)
	return nil
}
