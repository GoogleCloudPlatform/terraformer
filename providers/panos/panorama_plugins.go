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

package panos

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/PaloAltoNetworks/pango"
)

type PanoramaPluginsGenerator struct {
	PanosService
}

func (g *PanoramaPluginsGenerator) createGCPAccountResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Panorama.GcpAccount.GetList()
	if err != nil || len(l) == 0 {
		return resources
	}

	for _, r := range l {
		resources = append(resources, terraformutils.NewSimpleResource(
			r,
			normalizeResourceName(r),
			"panos_panorama_gcp_account",
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *PanoramaPluginsGenerator) createGKEClusterResources(group string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Panorama.GkeCluster.GetList(group)
	if err != nil || len(l) == 0 {
		return resources
	}

	for _, r := range l {
		id := group + ":" + r
		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(id),
			"panos_panorama_gke_cluster",
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *PanoramaPluginsGenerator) createGKEClusterGroupResources() (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Panorama.GkeClusterGroup.GetList()
	if err != nil || len(l) == 0 {
		return resources
	}

	for _, r := range l {
		resources = append(resources, terraformutils.NewSimpleResource(
			r,
			normalizeResourceName(r),
			"panos_panorama_gke_cluster_group",
			"panos",
			[]string{},
		))

		resources = append(resources, g.createGKEClusterResources(r)...)
	}

	return resources
}

func (g *PanoramaPluginsGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createGCPAccountResources()...)
	g.Resources = append(g.Resources, g.createGKEClusterGroupResources()...)

	return nil
}

func (g *PanoramaPluginsGenerator) PostConvertHook() error {
	mapGKEClusterGroupNames := map[string]string{}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_panorama_gke_cluster_group" {
			mapGKEClusterGroupNames[r.Item["name"].(string)] = "${" + r.InstanceInfo.Type + "." + r.ResourceName + ".name}"
		}
	}

	for _, r := range g.Resources {
		if r.InstanceInfo.Type == "panos_panorama_gke_cluster" {
			if mapExists(mapGKEClusterGroupNames, r.Item, "gke_cluster_group") {
				r.Item["gke_cluster_group"] = mapGKEClusterGroupNames[r.Item["gke_cluster_group"].(string)]
			}
		}
	}

	return nil
}
