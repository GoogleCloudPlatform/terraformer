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

package ibm

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/session"
)

type ContainerClusterGenerator struct {
	IBMService
}

func (g ContainerClusterGenerator) loadcluster(clustersID, clusterName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		clustersID,
		normalizeResourceName(clusterName, false),
		"ibm_container_cluster",
		"ibm",
		map[string]string{
			"force_delete_storage":   "true",
			"update_all_workers":     "false",
			"wait_for_worker_update": "true",
		},
		[]string{},
		map[string]interface{}{})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^worker_num$", "^region$",
	)
	return resource
}

func (g ContainerClusterGenerator) loadWorkerPools(clustersID, poolID, poolName string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", clustersID, poolID),
		normalizeResourceName(poolName, true),
		"ibm_container_worker_pool",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g ContainerClusterGenerator) loadWorkerPoolZones(clustersID, poolID, zoneID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", clustersID, poolID, zoneID),
		normalizeResourceName("ibm_container_worker_pool_zone_attachment", true),
		"ibm_container_worker_pool_zone_attachment",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g *ContainerClusterGenerator) InitResources() error {
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}
	client, err := containerv1.New(sess)
	if err != nil {
		return err
	}

	clusters, err := client.Clusters().List(containerv1.ClusterTargetHeader{})
	if err != nil {
		return err
	}

	for _, cs := range clusters {
		g.Resources = append(g.Resources, g.loadcluster(cs.ID, cs.Name))
		clusterResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
		workerPools, err := client.WorkerPools().ListWorkerPools(cs.ID, containerv1.ClusterTargetHeader{
			ResourceGroup: cs.ResourceGroupID,
		})
		if err != nil {
			return err
		}
		for _, pool := range workerPools {
			var dependsOn []string
			dependsOn = append(dependsOn,
				"ibm_container_cluster."+clusterResourceName)
			g.Resources = append(g.Resources, g.loadWorkerPools(cs.ID, pool.ID, pool.Name, dependsOn))
			poolResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName

			dependsOn = append(dependsOn,
				"ibm_container_worker_pool."+poolResourceName)
			zones := pool.Zones
			for _, zone := range zones {
				g.Resources = append(g.Resources, g.loadWorkerPoolZones(cs.ID, pool.ID, zone.ID, dependsOn))
			}
		}
	}
	return nil
}
