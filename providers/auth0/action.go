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
	ActionAllowEmptyValues = []string{}
)

type ActionGenerator struct {
	Auth0Service
}

func (g ActionGenerator) createResources(actions []*management.Action) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, action := range actions {
		resourceName := *action.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName+"_"+*action.Name,
			"auth0_action",
			"auth0",
			ActionAllowEmptyValues,
		))
	}
	return resources
}

func (g *ActionGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.Action{}

	var page int
	for {
		l, err := m.Action.List(management.Page(page))
		if err != nil {
			return err
		}
		list = append(list, l.Actions...)
		if !l.HasNext() {
			break
		}
		page++
	}

	g.Resources = g.createResources(list)
	return nil
}
