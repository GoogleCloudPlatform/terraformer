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

package mongodbatlas

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"go.mongodb.org/atlas/mongodbatlas"
)

var (
	ClusterAllowEmptyValues = []string{"tags."}
)

type ClusterGenerator struct {
	MongoDBAtlasService
}

func (g ClusterGenerator) createResources(clusters []mongodbatlas.Cluster) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, cluster := range clusters {
		resources = append(resources, terraformutils.NewSimpleResource(
			cluster.GroupID+"-"+cluster.Name,
			cluster.GroupID+"_"+cluster.Name,
			"mongodbatlas_cluster",
			"mongodbatlas",
			ProjectAllowEmptyValues,
		))
	}
	return resources
}

func (g *ClusterGenerator) InitResources() error {
	pg := ProjectGenerator{}
	client := g.generateClient()
	list := []mongodbatlas.Cluster{}
	opt := &mongodbatlas.ListOptions{}
	projects, err := pg.getProjects(context.TODO(), client)
	if err != nil {
		return err
	}
	//Enable filtering by project
	for _, project := range projects {
		projectID := project.ID
		for {
			clusters, resp, err := client.Clusters.List(context.TODO(), projectID, opt)
			if err != nil {
				return err
			}
			list = append(list, clusters...)
			if resp.Links == nil || resp.IsLastPage() {
				break
			}
			page, err := resp.CurrentPage()
			if err != nil {
				return err
			}
			opt.PageNum = page + 1
		}
	}
	g.Resources = g.createResources(list)
	return nil
}
