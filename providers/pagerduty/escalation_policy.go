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
)

type EscalationPolicyGenerator struct {
	PagerDutyService
}

func (g *EscalationPolicyGenerator) createEscalationPolicyResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListEscalationPoliciesOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.EscalationPolicies.List(&options)
		if err != nil {
			return err
		}

		for _, policy := range resp.EscalationPolicies {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				policy.ID,
				policy.Name,
				"pagerduty_escalation_policy",
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

func (g *EscalationPolicyGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createEscalationPolicyResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
