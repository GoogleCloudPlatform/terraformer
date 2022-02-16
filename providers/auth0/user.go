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
	UserAllowEmptyValues = []string{}
)

type UserGenerator struct {
	Auth0Service
}

func (g UserGenerator) createResources(users []*management.User) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, user := range users {
		resourceName := user.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			*resourceName,
			*resourceName,
			"auth0_user",
			"auth0",
			UserAllowEmptyValues,
		))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.User{}

	var page int
	for {
		l, err := m.User.List(management.Page(page))
		if err != nil {
			return err
		}
		list = append(list, l.Users...)
		if !l.HasNext() {
			break
		}
		page++
	}

	g.Resources = g.createResources(list)
	return nil
}
