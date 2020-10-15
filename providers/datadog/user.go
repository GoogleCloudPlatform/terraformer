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

package datadog

import (
	"context"
	"fmt"
	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// UserAllowEmptyValues ...
	UserAllowEmptyValues = []string{}
)

// UserGenerator ...
type UserGenerator struct {
	DatadogService
}

func (g *UserGenerator) createResources(users []datadogV1.User) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, user := range users {
		resourceName := user.GetHandle()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *UserGenerator) createResource(userID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		userID,
		fmt.Sprintf("user_%s", userID),
		"datadog_user",
		"datadog",
		UserAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each user create 1 TerraformResource.
// Need User ID as ID for terraform resource
func (g *UserGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("user") {
			for _, value := range filter.AcceptableValues {
				user, _, err := datadogClientV1.UsersApi.GetUser(authV1, value).Execute()
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(user.User.GetHandle()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	users, _, err := datadogClientV1.UsersApi.ListUsers(authV1).Execute()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(users.GetUsers())
	return nil
}
