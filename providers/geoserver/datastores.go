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
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gs "github.com/camptocamp/go-geoserver/client"
)

type DatastoresGenerator struct {
	GeoServerService
}

// InitResources generates TerraformResources from Github API,
func (g *DatastoresGenerator) InitResources() error {
	log.Println("InitResources for Datastores")
	ctx := context.Background()
	client := g.GeoserverClient()
	targetWorkspace := g.GetArgs()["targetWorkspace"].(string)

	if targetWorkspace == "" {
		log.Println("No target workspace is defined - Cannot work with datastores")
		return nil
	}

	g.Resources = append(g.Resources, createDatastoreResources(ctx, client, targetWorkspace)...)

	return nil
}

func createDatastoreResources(ctx context.Context, client *gs.Client, workspaceName string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	datastores, err := client.GetDatastores(workspaceName)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, datastore := range datastores {
		resource := terraformutils.NewSimpleResource(
			fmt.Sprintf("%s/%s", workspaceName, datastore.Name),
			datastore.Name,
			"geoserver_datastore",
			"geoserver",
			[]string{})
		resources = append(resources, resource)

		// init feature types
		resources = append(resources, createFeatureTypeResources(ctx, client, workspaceName, datastore.Name)...)

	}

	return resources
}
