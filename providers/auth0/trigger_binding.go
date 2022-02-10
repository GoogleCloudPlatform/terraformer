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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"gopkg.in/auth0.v5/management"
)

var (
	TriggerBindingAllowEmptyValues = []string{}
)

type TriggerBindingGenerator struct {
	Auth0Service
}

func (g TriggerBindingGenerator) createResources(bindings []*management.ActionBinding) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, binding := range bindings {
		resourceName := *binding.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName,
			"auth0_trigger_binding",
			"auth0",
			TriggerBindingAllowEmptyValues,
		))
	}
	return resources
}

func (g *TriggerBindingGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.ActionBinding{}

	t, err := m.Action.Triggers()
	if err != nil {
		return err
	}

	for _, trigger := range t.Triggers {
		log.Println(trigger.ID)
		var page int
		for {
			l, err := m.Action.Bindings(*trigger.ID, management.Page(page))
			if err != nil {
				return err
			}
			list = append(list, l.Bindings...)
			if !l.HasNext() {
				break
			}
			page++
			log.Println(list)
		}
	}

	g.Resources = g.createResources(list)
	return nil
}
