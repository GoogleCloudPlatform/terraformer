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

package okta

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

type UserGenerator struct {
	OktaService
}

func (g UserGenerator) createResources(userList []okta.User) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range userList {
		resources = append(resources, terraformutils.NewSimpleResource(
			user.GetId(),
			"user_"+user.GetId(),
			"okta_user",
			"okta",
			[]string{}))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	ctx, client, err := g.ClientV5()
	if err != nil {
		return err
	}

	output, resp, err := client.UserAPI.ListUsers(ctx).Execute()
	if err != nil {
		return err
	}

	for resp.HasNextPage() {
		var nextUserSet []okta.User
		resp, _ = resp.Next(&nextUserSet)
		output = append(output, nextUserSet...)
	}

	g.Resources = g.createResources(output)
	return nil
}
