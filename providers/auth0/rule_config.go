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

package auth0

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"gopkg.in/auth0.v5/management"
)

var (
	RuleConfigAllowEmptyValues = []string{}
)

type RuleConfigGenerator struct {
	Auth0Service
}

func (g RuleConfigGenerator) createResources(ruleConfigConfigs []*management.RuleConfig) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, ruleConfig := range ruleConfigConfigs {
		resourceName := *ruleConfig.Key
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName,
			"auth0_rule_config",
			"auth0",
			RuleConfigAllowEmptyValues,
		))
	}
	return resources
}

func (g *RuleConfigGenerator) InitResources() error {
	m := g.generateClient()

	list, err := m.RuleConfig.List()
	if err != nil {
		return err
	}

	g.Resources = g.createResources(list)
	return nil
}
