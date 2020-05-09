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
package mikrotik

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
)

type DhcpLeaseGenerator struct {
	MikrotikService
}

func (g DhcpLeaseGenerator) createResources(leases []client.DhcpLease) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, lease := range leases {
		resourceName := lease.Id
		if lease.Hostname != "" {
			resourceName = fmt.Sprintf("%s-%s", lease.Hostname, lease.Id)
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			lease.Id,
			resourceName,
			"mikrotik_dhcp_lease",
			"mikrotik",
			[]string{}))
	}
	return resources
}

func (g *DhcpLeaseGenerator) InitResources() error {
	client := g.generateClient()
	leases, err := client.ListDhcpLeases()

	if err != nil {
		return err
	}
	g.Resources = g.createResources(leases)
	return nil
}
