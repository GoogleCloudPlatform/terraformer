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

package gcp

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	container "google.golang.org/api/container/v1beta1"
)

var GkeAllowEmptyValues = []string{"labels."}

var GkeAdditionalFields = map[string]interface{}{}

type GkeGenerator struct {
	GCPService
}

func (g *GkeGenerator) initClusters(clusters *container.ListClustersResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, cluster := range clusters.Clusters {
		if _, exist := cluster.ResourceLabels["goog-composer-environment"]; exist { // don't manage composer clusters
			continue
		}
		resource := terraformutils.NewResource(
			cluster.Name,
			cluster.Name,
			"google_container_cluster",
			g.ProviderName,
			map[string]string{
				"name":     cluster.Name, // provider need cluster name as Required
				"project":  g.GetArgs()["project"].(string),
				"location": cluster.Location,
				"zone":     cluster.Zone,
			},
			GkeAllowEmptyValues,
			GkeAdditionalFields,
		)
		resource.IgnoreKeys = append(resource.IgnoreKeys,
			"^region$",
			"^additional_zones\\.(.*)",
			"^zone$",
			"^node_pool\\.(.*)",   // delete node_pool config from google_container_cluster
			"^node_config\\.(.*)", // delete node_config config from google_container_cluster
			"^ip_allocation_policy\\.[0-9]\\.cluster_secondary_range_name$",  // conflict with cluster_ipv4_cidr_block
			"^ip_allocation_policy\\.[0-9]\\.services_secondary_range_name$", // conflict with services_ipv4_cidr_block
			"^ip_allocation_policy\\.[0-9]\\.create_subnetwork")              // only for create new cluster conflict with others ip_allocation_policy fields
		resources = append(resources, resource)
		resources = append(resources, g.initNodePools(cluster.NodePools, cluster.Name, cluster.Location)...)
	}
	return resources
}

func (g *GkeGenerator) initNodePools(nodePools []*container.NodePool, clusterName, location string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, nodePool := range nodePools {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s/%s/%s", location, clusterName, nodePool.Name),
			clusterName+"_"+nodePool.Name,
			"google_container_node_pool",
			g.ProviderName,
			map[string]string{
				"location": location,
				"zone":     location,
				"project":  g.GetArgs()["project"].(string),
				"cluster":  clusterName, // provider need cluster name as Required
				"name":     nodePool.Name,
			},
			GkeAllowEmptyValues,
			GkeAdditionalFields,
		))
	}
	return resources
}

// Generate TerraformResources from GCP API,
func (g *GkeGenerator) InitResources() error {
	ctx := context.Background()
	service, err := container.NewService(ctx)
	if err != nil {
		log.Print(err)
		return err
	}
	// GKE support zone and regional cluster, api use location, it's can be region or zone, for all "-"
	location := fmt.Sprintf("projects/%s/locations/%s", g.GetArgs()["project"].(string), "-")
	clusters, err := service.Projects.Locations.Clusters.List(location).Do()
	if err != nil {
		log.Print(err)
		return err
	}

	g.Resources = g.initClusters(clusters)
	return nil
}

func (g *GkeGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "google_container_node_pool" {
			continue
		}
		if _, existNodeConfig := g.Resources[i].Item["node_config"]; existNodeConfig {
			if _, existMetadata := g.Resources[i].Item["node_config"].([]interface{})[0].(map[string]interface{})["metadata"]; existMetadata {
				for k, v := range g.Resources[i].Item["node_config"].([]interface{})[0].(map[string]interface{})["metadata"].(map[string]interface{}) {
					switch x := v.(type) {
					case bool:
						g.Resources[i].Item["node_config"].([]interface{})[0].(map[string]interface{})["metadata"].(map[string]interface{})[k] = strconv.FormatBool(x)
					default:
					}
				}
			}
		}
		for _, cluster := range g.Resources {
			if cluster.InstanceState.Attributes["name"] == r.InstanceState.Attributes["cluster"] {
				g.Resources[i].Item["cluster"] = "${google_container_cluster." + cluster.ResourceName + ".name}"
			}
		}
	}

	// hacks for fix GCP API<=>provider<=>parser inconsistency
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "google_container_cluster" {
			continue
		}
		if r.Item["master_authorized_networks_config"] != nil {
			if len(r.Item["master_authorized_networks_config"].([]interface{})) == 0 {
				g.Resources[i].Item["master_authorized_networks_config"] = map[string]interface{}{}
			}
		}
		if r.Item["ip_allocation_policy"] != nil {
			if len(r.Item["ip_allocation_policy"].([]interface{})) == 0 {
				g.Resources[i].Item["ip_allocation_policy"] = map[string]interface{}{}
			}
		}
	}
	return nil
}
