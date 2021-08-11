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

package okta

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"

	"github.com/okta/terraform-provider-okta/sdk"
)

type FactorGenerator struct {
	OktaService
}

func (g FactorGenerator) createResources(factorList []*sdk.Factor) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, factor := range factorList {
		if factor.Status == "ACTIVE" {
			resources = append(resources, terraformutils.NewResource(
				factor.Id,
				factor.Id,
				"okta_factor",
				"okta",
				map[string]string{
					"provider_id": factor.Id,
				},
				[]string{},
				map[string]interface{}{},
			))

		}
	}
	return resources
}

func (g *FactorGenerator) InitResources() error {
	var factors = []*sdk.Factor{}

	ctx, client, err := g.APISupplementClient()
	if err != nil {
		return err
	}

	output, _, err := getListFactors(ctx, client)
	if err != nil {
		return err
	}
	factors = append(factors, output...)

	g.Resources = g.createResources(factors)
	return nil
}

func getListFactors(ctx context.Context, m *sdk.ApiSupplement) ([]*sdk.Factor, *okta.Response, error) {
	//NOTE: Okta SDK does not support general ListFactors method so we got to manually implement the REST calls.
	url := fmt.Sprintf("/api/v1/org/factors")
	req, err := m.RequestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	var factors []*sdk.Factor
	resp, err := m.RequestExecutor.Do(ctx, req, &factors)
	if err != nil {
		return nil, resp, err
	}
	return factors, resp, nil
}
