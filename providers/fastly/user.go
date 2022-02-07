// Copyright 2019 The Terraformer Authors.
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

package fastly

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/fastly/go-fastly/v6/fastly"
)

type UserGenerator struct {
	FastlyService
}

func (g *UserGenerator) loadUsers(client *fastly.Client, customerID string) error {
	users, err := client.ListCustomerUsers(&fastly.ListCustomerUsersInput{CustomerID: customerID})
	if err != nil {
		return err
	}
	for _, user := range users {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			user.ID,
			user.ID,
			"fastly_user_v1",
			"fastly",
			[]string{}))
	}
	return nil
}

func (g *UserGenerator) InitResources() error {
	client, err := fastly.NewClient(g.Args["api_key"].(string))
	if err != nil {
		return err
	}

	if err := g.loadUsers(client, g.Args["customer_id"].(string)); err != nil {
		return err
	}

	return nil
}
