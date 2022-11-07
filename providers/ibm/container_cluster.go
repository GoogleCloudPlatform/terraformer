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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

const (
	defaultWorkerPool = "default"
	hardwareShared    = "shared"
	hardwareDedicated = "dedicated"
	isolationPublic   = "public"
	isolationPrivate  = "private"
)

type ContainerClusterGenerator struct {
	IBMService
}

func (g ContainerClusterGenerator) loadcluster(clustersID, clusterName, datacenter, hardware string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		clustersID,
		normalizeResourceName(clusterName, false),
		"ibm_container_cluster",
		"ibm",
		map[string]string{
			"force_delete_storage":   "true",
			"update_all_workers":     "false",
			"wait_for_worker_update": "true",
			"datacenter":             datacenter,
			"hardware":               hardware,
		},
		[]string{},
		map[string]interface{}{})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^worker_num$", "^region$",
	)

	return resource
}

func (g ContainerClusterGenerator) loadWorkerPools(clustersID, poolID, poolName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", clustersID, poolID),
		normalizeResourceName(poolName, true),
		"ibm_container_worker_pool",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	return resources
}

func (g ContainerClusterGenerator) loadWorkerPoolZones(clustersID, poolID, zoneID string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", clustersID, poolID, zoneID),
		normalizeResourceName("ibm_container_worker_pool_zone_attachment", true),
		"ibm_container_worker_pool_zone_attachment",
		"ibm",
		map[string]string{
			"wait_till_albs": "true",
		},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g ContainerClusterGenerator) loadNlbDNS(clusterID string, nlbIPs []interface{}) terraformutils.Resource {
	resources := terraformutils.NewResource(
		clusterID,
		normalizeResourceName(clusterID, true),
		"ibm_container_nlb_dns",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"nlb_ips": nlbIPs,
		})

	return resources
}

func (g *ContainerClusterGenerator) InitResources() error {
	region := g.Args["region"].(string)
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

	clientNlb, err := containerv2.New(sess)
	if err != nil {
		return err
	}

	for _, cs := range clusters {
		if region == cs.Region {
			hardware := hardwareShared

			workerPools, err := client.WorkerPools().ListWorkerPools(cs.ID, containerv1.ClusterTargetHeader{
				ResourceGroup: cs.ResourceGroupID,
			})
			if err != nil {
				return err
			}

			if len(workerPools) > 0 && workerPoolContains(workerPools, defaultWorkerPool) {
				hardware = workerPools[0].Isolation
				switch strings.ToLower(hardware) {
				case "":
					hardware = hardwareShared
				case isolationPrivate:
					hardware = hardwareDedicated
				case isolationPublic:
					hardware = hardwareShared
				}
			}

			g.Resources = append(g.Resources, g.loadcluster(cs.ID, cs.Name, cs.DataCenter, hardware))

			for _, pool := range workerPools {
				g.Resources = append(g.Resources, g.loadWorkerPools(cs.ID, pool.ID, pool.Name))

				zones := pool.Zones
				for _, zone := range zones {
					g.Resources = append(g.Resources, g.loadWorkerPoolZones(cs.ID, pool.ID, zone.ID))
				}
			}

			nlbData, err := clientNlb.NlbDns().GetNLBDNSList(cs.Name)
			if err != nil {
				return err
			}

			for _, data := range nlbData {
				g.Resources = append(g.Resources, g.loadNlbDNS(data.Nlb.Cluster, data.Nlb.NlbIPArray))
			}
		}
	}

	return nil
}

func workerPoolContains(workerPools []v1.WorkerPoolResponse, pool string) bool {
	for _, workerPool := range workerPools {
		if workerPool.Name == pool {
			return true
		}
	}
	return false
}

func (g *ContainerClusterGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_container_cluster" {
			continue
		}
		for i, wp := range g.Resources {
			if wp.InstanceInfo.Type != "ibm_container_worker_pool" {
				continue
			}

			if wp.InstanceState.Attributes["cluster"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["cluster"] = "${ibm_container_cluster." + r.ResourceName + ".id}"
			}
		}

		for i, wpZoneAttach := range g.Resources {
			if wpZoneAttach.InstanceInfo.Type != "ibm_container_worker_pool_zone_attachment" {
				continue
			}

			if wpZoneAttach.InstanceState.Attributes["cluster"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["cluster"] = "${ibm_container_cluster." + r.ResourceName + ".id}"
			}
		}

		for i, wp := range g.Resources {
			if wp.InstanceInfo.Type != "ibm_container_worker_pool" {
				continue
			}
			if wp.InstanceState.Attributes["cluster"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["cluster"] = "${ibm_container_cluster." + r.ResourceName + ".id}"
			}
		}

		for i, nlb := range g.Resources {
			if nlb.InstanceInfo.Type != "ibm_container_nlb_dns" {
				continue
			}

			if nlb.InstanceState.Attributes["cluster"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["cluster"] = "${ibm_container_cluster." + r.ResourceName + ".id}"
			}
		}
	}

	return nil
}
