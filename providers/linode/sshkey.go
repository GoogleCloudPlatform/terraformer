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

package linode

import (
	"context"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/linode/linodego"
)

type SSHKeyGenerator struct {
	LinodeService
}

func (g SSHKeyGenerator) createResources(keyList []linodego.SSHKey) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, key := range keyList {
		resources = append(resources, terraformutils.NewSimpleResource(
			strconv.Itoa(key.ID),
			strconv.Itoa(key.ID),
			"linode_sshkey",
			"linode",
			[]string{}))
	}
	return resources
}

func (g *SSHKeyGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.ListSSHKeys(context.Background(), nil)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
