// Copyright 2024 The Terraformer Authors.
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

package scaleway

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type KubernetesGenerator struct {
	ScalewayService
}

func (g KubernetesGenerator) ListKubernetes(client *scw.Client) ([]*k8s.Cluster, []*k8s.Pool, error) {
	listCluster := []*k8s.Cluster{}
	listPool := []*k8s.Pool{}
	clusterApi := k8s.NewAPI(client)

	page := int32(1)
	perpage := uint32(100)
	optCluster := &k8s.ListClustersRequest{
		Page:     &page,
		PageSize: &perpage,
	}

	for {
		resp, err := clusterApi.ListClusters(optCluster)
		if err != nil {
			return nil, nil, err
		}
		for _, cluster := range resp.Clusters {
			if cluster != nil {
				listCluster = append(listCluster, cluster)
			}
		}
		// Exit loop when we are on the last page.
		if resp.TotalCount < uint64(*optCluster.PageSize) {
			break
		}
		*optCluster.Page++
	}
	for _, cluster := range listCluster {
		optPool := &k8s.ListPoolsRequest{
			Page:      &page,
			PageSize:  &perpage,
			ClusterID: cluster.ID,
		}
		for {
			resp, err := clusterApi.ListPools(optPool)
			if err != nil {
				return nil, nil, err
			}
			for _, server := range resp.Pools {
				if server != nil {
					listPool = append(listPool, server)
				}
			}
			// Exit loop when we are on the last page.
			if resp.TotalCount < uint64(*optPool.PageSize) {
				break
			}
			*optPool.Page++
		}
	}
	return listCluster, listPool, nil
}

func (g KubernetesGenerator) createResources(clusterList []*k8s.Cluster, poolList []*k8s.Pool) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, cluster := range clusterList {
		resources = append(resources, terraformutils.NewSimpleResource(
			string(cluster.Region)+"/"+cluster.ID,
			cluster.Name,
			"scaleway_k8s_cluster",
			"scaleway",
			[]string{}))
	}
	for _, pool := range poolList {
		resources = append(resources, terraformutils.NewSimpleResource(
			string(pool.Region)+"/"+pool.ID,
			pool.Name,
			"scaleway_k8s_pool",
			"scaleway",
			[]string{}))
	}
	return resources
}

func (g *KubernetesGenerator) InitResources() error {
	client := g.generateClient()
	clusters, pools, err := g.ListKubernetes(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(clusters, pools)
	return nil
}
