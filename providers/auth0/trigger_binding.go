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
	TriggerBindingAllowEmptyValues = []string{}
)

type TriggerBindingGenerator struct {
	Auth0Service
}

func (g TriggerBindingGenerator) createResources(bindings map[string]*management.ActionBinding) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	for _, binding := range bindings {
		resourceName := *binding.TriggerID
		resources = append(resources, terraformutils.NewResource(
			resourceName,
			*binding.ID,
			"auth0_trigger_binding",
			"auth0",
			map[string]string{},
			TriggerBindingAllowEmptyValues,
			map[string]interface{}{
				"trigger": *binding.TriggerID,
			},
		))
	}
	return resources
}

func (g *TriggerBindingGenerator) InitResources() error {
	m := g.generateClient()
	bindings := map[string]*management.ActionBinding{}

	t, err := m.Action.Triggers()
	if err != nil {
		return err
	}

	for _, trigger := range t.Triggers {
		var page int
		for {
			l, err := m.Action.Bindings(*trigger.ID, management.Page(page))
			if err != nil {
				return err
			}
			for _, binding := range l.Bindings {
				if _, ok := bindings[*binding.ID]; !ok {
					bindings[*binding.ID] = binding
				}
			}
			if !l.HasNext() {
				break
			}
			page++
		}
	}

	g.Resources = g.createResources(bindings)
	return nil
}
