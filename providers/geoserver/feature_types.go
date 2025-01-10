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

type FeatureTypesGenerator struct {
	GeoServerService
}

// InitResources generates TerraformResources from Github API,
func (g *FeatureTypesGenerator) InitResources() error {
	log.Println("InitResources for Datastores")
	ctx := context.Background()
	client := g.GeoserverClient()
	targetWorkspace := g.GetArgs()["targetWorkspace"].(string)
	targetDatastore := g.GetArgs()["targetDatastore"].(string)

	if targetWorkspace == "" || targetDatastore == "" {
		log.Println("No target workspace or datastore is defined - Cannot work with feature types")
		return nil
	}

	g.Resources = append(g.Resources, createFeatureTypeResources(ctx, client, targetWorkspace, targetDatastore)...)

	return nil
}

func createFeatureTypeResources(ctx context.Context, client *gs.Client, workspaceName string, datastoreName string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	featuretypes, err := client.GetFeatureTypes(workspaceName, datastoreName)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, featuretype := range featuretypes {
		resource := terraformutils.NewSimpleResource(
			fmt.Sprintf("%s/%s/%s", workspaceName, datastoreName, featuretype.Name),
			featuretype.Name,
			"geoserver_featuretype",
			"geoserver",
			[]string{})
		resources = append(resources, resource)
	}

	return resources
}
