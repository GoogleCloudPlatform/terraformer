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
package hashicups

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type OrderGenerator struct {
	HashicupsService
}

func (a *OrderGenerator) InitResources() error {
	client, err := a.createClient()

	if err != nil {
		return err
	}

	cont, err := client.GetViaTokenAndPoint("/orders", client.Token)
	if err != nil {
		return err
	}

	c, err := cont.ArrayCount()
	if err != nil {
		return err
	}

	for i := 0; i < c; i++ {
		id:=RemoveQuotes(cont.Index(i).S("id").String())
		resource := terraformutils.NewResource(
			id,
			id,
			"hashicups_order",
			"hashicups",
			map[string]string{},
			[]string{
				"last_updated",
			},
			map[string]interface{}{},
		)
		resource.SlowQueryRequired = true
		a.Resources = append(a.Resources, resource)
	}
	return nil
}
