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

type SSHKeyGenerator struct {
	EquinixMetalService
}

func (g SSHKeyGenerator) listSSHKeys(client *packngo.Client) ([]packngo.SSHKey, error) {
	sshKeys, _, err := client.SSHKeys.List()
	if err != nil {
		return nil, err
	}

	return sshKeys, nil
}

func (g SSHKeyGenerator) createResources(sshLeyList []packngo.SSHKey) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, sshKey := range sshLeyList {
		resources = append(resources, terraformutils.NewSimpleResource(
			sshKey.ID,
			sshKey.Label,
			"metal_ssh_key",
			"equinixmetal",
			[]string{}))
	}
	return resources
}

func (g *SSHKeyGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listSSHKeys(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
