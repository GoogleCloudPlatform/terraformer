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
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"

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

func (g *UserGenerator) createResources(users []datadogV2.User) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, user := range users {
		relations := user.GetRelationships()
		roles := relations.GetRoles()
		// If no roles are present, we can assume user was created via the V1 API
		// Hence, import the user via their handle
		if len(roles.GetData()) == 0 {
			attr := user.GetAttributes()
			resources = append(resources, g.createResource(attr.GetHandle()))
			continue
		}

		resources = append(resources, g.createResource(user.GetId()))
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
	var users []datadogV2.User

	datadogClientV2 := g.Args["datadogClientV2"].(*datadogV2.APIClient)
	authV2 := g.Args["authV2"].(context.Context)

	pageSize := int64(1000)
	pageNumber := int64(0)
	remaining := int64(1)

	for remaining > int64(0) {
		resp, _, err := datadogClientV2.UsersApi.ListUsers(authV2).PageSize(pageSize).PageNumber(pageNumber).Execute()
		if err != nil {
			return err
		}
		users = append(users, resp.GetData()...)

		remaining = resp.Meta.Page.GetTotalCount() - pageSize*(pageNumber+1)
		pageNumber++
	}

	g.Resources = g.createResources(users)
	return nil
}
