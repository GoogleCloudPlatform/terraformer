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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/terraform-provider-okta/sdk"
)

type FactorGenerator struct {
	OktaService
}

func (g FactorGenerator) createResources(ctx context.Context, factorList []*okta.UserFactor, client *sdk.APISupplement) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, factor := range factorList {
		if factor.Status == "ACTIVE" {
			resources = append(resources, terraformutils.NewResource(
				factor.Id,
				"factor_"+normalizeResourceNameWithRandom(factor.Id, true),
				"okta_factor",
				"okta",
				map[string]string{
					"provider_id": factor.Id,
				},
				[]string{},
				map[string]interface{}{},
			))

			if factor.FactorType == "token:hotp" {
				hotpFactorProfiles, _, _ := getHotpFactorProfiles(ctx, client)

				for _, factorProfile := range hotpFactorProfiles {
					if factorProfile != nil {
						resources = append(resources, terraformutils.NewResource(
							factorProfile.ID,
							"factor_totp_"+normalizeResourceNameWithRandom(factorProfile.Name, true),
							"okta_factor_totp",
							"okta",
							map[string]string{},
							[]string{},
							map[string]interface{}{
								"name":                   factorProfile.Name,
								"otp_length":             factorProfile.Settings.OtpLength,
								"time_step":              factorProfile.Settings.TimeStep,
								"clock_drift_interval":   factorProfile.Settings.AcceptableAdjacentIntervals,
								"shared_secret_encoding": factorProfile.Settings.Encoding,
								"hmac_algorithm":         factorProfile.Settings.TimeStep,
							},
						))
					}
				}
			}
		}
	}
	return resources
}

func (g *FactorGenerator) InitResources() error {
	var factors []*okta.UserFactor

	ctx, client, err := g.APISupplementClient()
	if err != nil {
		return err
	}

	output, _, err := getListFactors(ctx, client)
	if err != nil {
		return err
	}

	factors = append(factors, output...)

	g.Resources = g.createResources(ctx, factors, client)
	return nil
}

func getListFactors(ctx context.Context, m *sdk.APISupplement) ([]*okta.UserFactor, *okta.Response, error) {
	//NOTE: Okta SDK does not support general ListFactors method so we got to manually implement the REST calls.
	url := "/api/v1/org/factors"
	req, err := m.RequestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	var factors []*okta.UserFactor
	resp, err := m.RequestExecutor.Do(ctx, req, &factors)
	if err != nil {
		return nil, resp, err
	}
	return factors, resp, nil
}

func getHotpFactorProfiles(ctx context.Context, m *sdk.APISupplement) ([]*sdk.HotpFactorProfile, *okta.Response, error) {
	url := "/api/v1/org/factors/hotp/profiles"
	req, err := m.RequestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	var factors []*sdk.HotpFactorProfile
	resp, err := m.RequestExecutor.Do(ctx, req, &factors)
	if err != nil {
		return nil, resp, err
	}
	return factors, resp, nil
}
