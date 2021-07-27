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

type EventHookGenerator struct {
	OktaService
}

func (g EventHookGenerator) createResources(eventHookList []*okta.EventHook) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, eventHook := range eventHookList {

		resources = append(resources, terraformutils.NewSimpleResource(
			eventHook.Id,
			"event_hook_"+eventHook.Name,
			"okta_event_hook",
			"okta",
			[]string{}))
	}
	return resources
}

func (g *EventHookGenerator) InitResources() error {
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	output, resp, err := client.EventHook.ListEventHooks(ctx)
	if err != nil {
		return e
	}

	for resp.HasNextPage() {
		var nextEventHookSet []*okta.EventHook
		resp, _ = resp.Next(ctx, &nextEventHookSet)
		output = append(output, nextEventHookSet...)
	}

	g.Resources = g.createResources(output)
	return nil
}
