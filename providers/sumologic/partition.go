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

type PartitionGenerator struct {
	SumoLogicService
}

func (g *PartitionGenerator) createResources(partitions []sumologic.Partition) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(partitions))

	for i, partition := range partitions {
		name := strcase.ToSnake(replaceSpaceAndDash(partition.Name))
		resource := terraformutils.NewSimpleResource(
			partition.Id,
			fmt.Sprintf("%s-%s", name, partition.Id),
			"sumologic_partition",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *PartitionGenerator) InitResources() error {
	client := g.Client()
	req := client.PartitionManagementApi.ListPartitions(g.AuthCtx())
	req = req.Limit(100)

	respBody, _, err := client.PartitionManagementApi.ListPartitionsExecute(req)
	if err != nil {
		return err
	}
	partitions := respBody.Data
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.PartitionManagementApi.ListPartitionsExecute(req)
		if err != nil {
			return err
		}
		partitions = append(partitions, respBody.Data...)
	}

	resources := g.createResources(partitions)
	g.Resources = resources
	return nil
}
