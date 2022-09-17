// Copyright 2021 The Terraformer Authors.
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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type UserGenerator struct {
	SumoLogicService
}

func (g *UserGenerator) createResources(users []sumologic.UserModel) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(users))

	for i, user := range users {
		name := strcase.ToSnake(replaceSpaceAndDash(user.FirstName + "_" + user.LastName))
		resource := terraformutils.NewSimpleResource(
			user.Id,
			fmt.Sprintf("%s-%s", name, user.Id),
			"sumologic_user",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *UserGenerator) InitResources() error {
	client := g.Client()
	req := client.UserManagementApi.ListUsers(g.AuthCtx())
	req = req.Limit(100)
	for _, filter := range g.Filter {
		if filter.IsApplicable("user") && filter.FieldPath == "email" {
			if len(filter.AcceptableValues) == 1 {
				req = req.Email(filter.AcceptableValues[0])
			}
		}
	}

	respBody, _, err := client.UserManagementApi.ListUsersExecute(req)
	if err != nil {
		return err
	}
	users := respBody.Data
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.UserManagementApi.ListUsersExecute(req)
		if err != nil {
			return err
		}
		users = append(users, respBody.Data...)
	}

	resources := g.createResources(users)
	g.Resources = resources
	return nil
}
