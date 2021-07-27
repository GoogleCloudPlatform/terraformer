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
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type UserTypeGenerator struct {
	OktaService
}

func (g UserTypeGenerator) createResources(userTypeList []*okta.UserType) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, userType := range userTypeList {

		resources = append(resources, terraformutils.NewSimpleResource(
			userType.Id,
			"usertype_"+userType.Name,
			"okta_user_type",
			"okta",
			[]string{}))
	}
	return resources
}

func (g *UserTypeGenerator) InitResources() error {
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	output, resp, err := client.UserType.ListUserTypes(ctx)
	if err != nil {
		return e
	}

	for resp.HasNextPage() {
		var nextUserTypeSet []*okta.UserType
		resp, _ = resp.Next(ctx, &nextUserTypeSet)
		output = append(output, nextUserTypeSet...)
	}

	g.Resources = g.createResources(output)
	return nil
}
