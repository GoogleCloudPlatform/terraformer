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
	HookAllowEmptyValues = []string{}
)

type HookGenerator struct {
	Auth0Service
}

func (g HookGenerator) createResources(hooks []*management.Hook) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, hook := range hooks {
		resourceName := *hook.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName+"_"+*hook.Name,
			"auth0_hook",
			"auth0",
			HookAllowEmptyValues,
		))
	}
	return resources
}

func (g *HookGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.Hook{}

	var page int
	for {
		l, err := m.Hook.List(management.Page(page))
		if err != nil {
			return err
		}
		list = append(list, l.Hooks...)
		if !l.HasNext() {
			break
		}
		page++
	}

	g.Resources = g.createResources(list)
	return nil
}
