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

package geoserver

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gs "github.com/camptocamp/go-geoserver/client"
)

type WorkspacesGenerator struct {
	GeoServerService
}

// InitResources generates TerraformResources from Github API,
func (g *WorkspacesGenerator) InitResources() error {
	log.Println("InitResources")
	ctx := context.Background()
	client := g.GeoserverClient()

	g.Resources = append(g.Resources, createWorkspaceResources(ctx, client)...)

	return nil
}

func createWorkspaceResources(ctx context.Context, client *gs.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	workspaces, err := client.GetWorkspaces()
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, workspace := range workspaces {
		resource := terraformutils.NewSimpleResource(
			workspace.Name,
			workspace.Name,
			"geoserver_workspace",
			"geoserver",
			[]string{})
		resources = append(resources, resource)
	}

	return resources
}
