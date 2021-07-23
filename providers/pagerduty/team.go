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

type TeamGenerator struct {
	PagerDutyService
}

func (g *TeamGenerator) createTeamResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListTeamsOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Teams.List(&options)
		if err != nil {
			return err
		}

		for _, team := range resp.Teams {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				team.ID,
				fmt.Sprintf("Team_%s", team.Name),
				"pagerduty_team",
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

func (g *TeamGenerator) createTeamMembershipResources(client *pagerduty.Client) error {
	var teamOffset = 0
	teamOptions := pagerduty.ListTeamsOptions{}

	for {
		teamOptions.Offset = teamOffset
		resp, _, err := client.Teams.List(&teamOptions)
		if err != nil {
			return err
		}

		memberOptions := pagerduty.GetMembersOptions{}
		for _, team := range resp.Teams {
			members, _, err := client.Teams.GetMembers(team.ID, &memberOptions)

			if err != nil {
				return err
			}

			for _, member := range members.Members {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					fmt.Sprintf("%s:%s", member.User.ID, team.ID),
					fmt.Sprintf("%s_%s", member.User.ID, team.Name),
					"pagerduty_team_membership",
					g.ProviderName,
					[]string{},
				))
			}
		}

		if !resp.More {
			break
		}

		teamOffset += resp.Limit
	}

	return nil
}

func (g *TeamGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createTeamResources,
		g.createTeamMembershipResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
