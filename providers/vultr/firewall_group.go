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

package vultr

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/vultr/govultr"
)

type FirewallGroupGenerator struct {
	VultrService
}

func (g FirewallGroupGenerator) createResources(firewallGroup []govultr.FirewallGroup) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, firewallGroup := range firewallGroup {
		resources = append(resources, terraform_utils.NewSimpleResource(
			firewallGroup.FirewallGroupID,
			firewallGroup.FirewallGroupID,
			"vultr_firewall_group",
			"vultr",
			[]string{}))
	}
	return resources
}

func (g *FirewallGroupGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.FirewallGroup.List(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
