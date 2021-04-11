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

package equinixmetal

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/packethost/packngo"
)

type VolumeGenerator struct {
	EquinixMetalService
}

func (g VolumeGenerator) listVolumes(client *packngo.Client) ([]packngo.Volume, error) {
	volumes, _, err := client.Volumes.List(g.GetArgs()["project_id"].(string), nil)
	if err != nil {
		return nil, err
	}

	return volumes, nil
}

func (g VolumeGenerator) createResources(volumeList []packngo.Volume) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, volume := range volumeList {
		resources = append(resources, terraformutils.NewSimpleResource(
			volume.ID,
			volume.Name,
			"metal_volume",
			"equinixmetal",
			[]string{}))
	}
	return resources
}

func (g *VolumeGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listVolumes(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
