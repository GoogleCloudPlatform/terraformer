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
	ProjectAllowEmptyValues = []string{}
)

type ProjectGenerator struct {
	MongoDBAtlasService
}

func (g ProjectGenerator) getProjects(ctx context.Context, client *mongodbatlas.Client) ([]*mongodbatlas.Project, error) {
	list := []*mongodbatlas.Project{}
	opt := &mongodbatlas.ListOptions{}
	for {
		projects, resp, err := client.Projects.GetAllProjects(context.TODO(), &mongodbatlas.ListOptions{})
		if err != nil {
			return nil, err
		}
		list = append(list, projects.Results...)
		if resp.Links == nil || resp.IsLastPage() {
			break
		}
		page, err := resp.CurrentPage()
		if err != nil {
			return nil, err
		}
		opt.PageNum = page + 1
	}
	return list, nil
}

func (g ProjectGenerator) createResources(projects []*mongodbatlas.Project) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, project := range projects {
		resourceName := project.ID
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName,
			"mongodbatlas_project",
			"mongodbatlas",
			ProjectAllowEmptyValues,
		))
	}
	return resources
}

func (g *ProjectGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.getProjects(context.TODO(), client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
