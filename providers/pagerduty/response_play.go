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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"

	"fmt"
	"strings"
)

type ResponsePlayGenerator struct {
	PagerDutyService
}

func (g *ResponsePlayGenerator) createResponsePlayResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListUsersOptions{}
	responsePlayOptions := pagerduty.ListResponsePlayOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Users.List(&options)
		if err != nil {
			return err
		}
		for _, user := range resp.Users {
			responsePlayOptions.From = user.Email
			resp, _, err := client.ResponsePlays.List(&responsePlayOptions)
			if err != nil {
				return err
			}
			for _, responsePlay := range resp.ResponsePlays {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					responsePlay.ID,
					fmt.Sprintf("%s_%s", strings.Replace(responsePlay.Name, " ", "_", -1), strings.ToLower(strings.TrimSuffix(strings.Replace(user.Name, " ", "_", -1), "_("))),
					"pagerduty_response_play",
					g.ProviderName,
					map[string]string{
						"from": user.Email,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
		if !resp.More {
			break
		}

		offset += resp.Limit
	}

	return nil
}

func (g *ResponsePlayGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createResponsePlayResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
