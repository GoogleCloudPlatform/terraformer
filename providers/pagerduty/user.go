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

package pagerduty

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
)

type UserGenerator struct {
	PagerDutyService
}

func (g *UserGenerator) createUserResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListUsersOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Users.List(&options)
		if err != nil {
			return err
		}

		for _, user := range resp.Users {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				user.ID,
				fmt.Sprintf("user_%s", user.ID),
				"pagerduty_user",
				g.ProviderName,
				[]string{},
			))
		}

		if !resp.More {
			break
		}

		offset += resp.Limit
	}

	return nil
}

func (g *UserGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createUserResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
