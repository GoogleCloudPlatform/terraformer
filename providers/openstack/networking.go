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

package openstack

import (
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/gophercloud/gophercloud/pagination"
)

type NetworkingGenerator struct {
	OpenStackService
}

// createResources iterate on all openstack_networking_secgroup_v2
func (g *NetworkingGenerator) createSecgroupResources(list *pagination.Pager) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	err := list.EachPage(func(page pagination.Page) (bool, error) {
		groups, err := groups.ExtractGroups(page)
		if err != nil {
			return false, err
		}

		for _, grp := range groups {
			resource := terraformutils.NewSimpleResource(
				grp.ID,
				grp.Name,
				"openstack_networking_secgroup_v2",
				"openstack",
				[]string{},
			)
			resources = append(resources, resource)
			resources = append(resources, g.createSecgroupRuleResources(grp.Rules)...)
		}

		return true, nil
	})
	if err != nil {
		log.Println(err)
	}
	return resources
}

// createResources iterate on all openstack_networking_secgroup_v2
func (g *NetworkingGenerator) createSecgroupRuleResources(rules []rules.SecGroupRule) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, r := range rules {
		resource := terraformutils.NewSimpleResource(
			r.ID,
			r.ID,
			"openstack_networking_secgroup_rule_v2",
			"openstack",
			[]string{},
		)
		resources = append(resources, resource)
	}
	return resources
}

// Generate TerraformResources from OpenStack API,
func (g *NetworkingGenerator) InitResources() error {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return err
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return err
	}

	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: g.GetArgs()["region"].(string),
	})
	if err != nil {
		return err
	}

	list := groups.List(client, groups.ListOpts{})

	g.Resources = g.createSecgroupResources(&list)

	return nil
}

func (g *NetworkingGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "openstack_networking_secgroup_rule_v2" {
			continue
		}
		for _, sg := range g.Resources {
			if sg.InstanceInfo.Type != "openstack_networking_secgroup_v2" {
				continue
			}
			if r.InstanceState.Attributes["security_group_id"] == sg.InstanceState.Attributes["id"] {
				g.Resources[i].Item["security_group_id"] = "${openstack_networking_secgroup_v2." + sg.ResourceName + ".id}"
			}
		}
	}

	return nil
}
