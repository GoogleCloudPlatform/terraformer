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
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		clustersID,
		clusterName,
		"ibm_container_cluster",
		"ibm",
		[]string{})
	return resources
}

func (g ContainerClusterGenerator) loadWorkerPools(clustersID, poolID, PoolName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", clustersID, poolID),
		PoolName,
		"ibm_container_worker_pool",
		"ibm",
		[]string{})
	return resources
}

func (g ContainerClusterGenerator) loadWorkerPoolZones(clustersID, poolID, zoneID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s/%s", clustersID, poolID, zoneID),
		zoneID,
		"ibm_container_worker_pool_zone_attachment",
		"ibm",
		[]string{})
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
		workerPools, err := client.WorkerPools().ListWorkerPools(cs.ID, containerv1.ClusterTargetHeader{})
		if err != nil {
			return err
		}
		for _, pool := range workerPools {
			if pool.Name != "default" {
				g.Resources = append(g.Resources, g.loadWorkerPools(cs.ID, pool.ID, pool.Name))

				zones := pool.Zones
				for _, zone := range zones {
					g.Resources = append(g.Resources, g.loadWorkerPoolZones(cs.ID, pool.ID, zone.ID))
				}
			}
		}
	}
	return nil
}
