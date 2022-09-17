// Copyright 2021 The Terraformer Authors.
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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type ConnectionGenerator struct {
	SumoLogicService
}

func (g *ConnectionGenerator) createResources(connections []sumologic.Connection) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(connections))

	for i, connection := range connections {
		name := strcase.ToSnake(replaceSpaceAndDash(connection.Name))
		resource := terraformutils.NewSimpleResource(
			connection.Id,
			fmt.Sprintf("%s-%s", name, connection.Id),
			"sumologic_connection",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *ConnectionGenerator) InitResources() error {
	client := g.Client()
	req := client.ConnectionManagementApi.ListConnections(g.AuthCtx())
	req = req.Limit(100)

	respBody, _, err := client.ConnectionManagementApi.ListConnectionsExecute(req)
	if err != nil {
		return err
	}
	connections := respBody.Data
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.ConnectionManagementApi.ListConnectionsExecute(req)
		if err != nil {
			return err
		}
		connections = append(connections, respBody.Data...)
	}

	resources := g.createResources(connections)
	g.Resources = resources
	return nil
}
