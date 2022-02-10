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
	ClientGrantAllowEmptyValues = []string{}
)

type ClientGrantGenerator struct {
	Auth0Service
}

func (g ClientGrantGenerator) createResources(clientGrantGrants []*management.ClientGrant) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, clientGrant := range clientGrantGrants {
		resourceName := *clientGrant.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName+"_"+*clientGrant.ClientID,
			"auth0_client_grant",
			"auth0",
			ClientGrantAllowEmptyValues,
		))
	}
	return resources
}

func (g *ClientGrantGenerator) InitResources() error {
	m := g.generateClient()
	list := []*management.ClientGrant{}

	var page int
	for {
		l, err := m.ClientGrant.List(management.Page(page))
		if err != nil {
			return err
		}
		list = append(list, l.ClientGrants...)
		if !l.HasNext() {
			break
		}
		page++
	}

	g.Resources = g.createResources(list)
	return nil
}
