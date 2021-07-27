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
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type GroupGenerator struct {
	OktaService
}

func (g GroupGenerator) createResources(groupList []*okta.Group) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, group := range groupList {

		resources = append(resources, terraformutils.NewSimpleResource(
			group.Id,
			"group_"+group.Profile.Name,
			"okta_group",
			"okta",
			[]string{}))
	}
	return resources
}

func (g *GroupGenerator) InitResources() error {
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	filter := query.NewQueryParams(query.WithFilter("type eq \"OKTA_GROUP\""))
	output, resp, err := client.Group.ListGroups(ctx, filter)
	if err != nil {
		return e
	}

	for resp.HasNextPage() {
		var nextGroupSet []*okta.Group
		resp, _ = resp.Next(ctx, &nextGroupSet)
		output = append(output, nextGroupSet...)
	}

	g.Resources = g.createResources(output)
	return nil
}
