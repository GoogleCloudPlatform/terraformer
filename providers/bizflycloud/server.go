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

package bizflycloud

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type ServerGenerator struct {
	BizflyCloudService
}

func (g ServerGenerator) listServers(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.Server, error) {

	opt := &gobizfly.ServerListOptions{}
	servers, err := client.Server.List(ctx, opt)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

func (g ServerGenerator) createResources(serverList []*gobizfly.Server) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, server := range serverList {
		resources = append(resources, terraformutils.NewSimpleResource(
			server.ID,
			server.Name,
			"bizflycloud_server",
			"bizflycloud",
			[]string{}))
	}
	return resources
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listServers(context.TODO(), client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
