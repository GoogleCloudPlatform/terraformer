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

package vultr

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/vultr/govultr"
)

type SnapshotGenerator struct {
	VultrService
}

func (g SnapshotGenerator) createResources(snapshotList []govultr.Snapshot) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, snapshot := range snapshotList {
		resources = append(resources, terraformutils.NewSimpleResource(
			snapshot.SnapshotID,
			snapshot.SnapshotID,
			"vultr_snapshot",
			"vultr",
			[]string{}))
	}
	return resources
}

func (g *SnapshotGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.Snapshot.List(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
