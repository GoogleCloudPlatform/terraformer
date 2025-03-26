// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	oktaV2 "github.com/okta/okta-sdk-golang/v2/okta"
	oktaV5 "github.com/okta/okta-sdk-golang/v5/okta"
	"github.com/okta/terraform-provider-okta/sdk"
)

type OktaService struct { //nolint
	terraformutils.Service
}

func (s *OktaService) Client() (context.Context, *oktaV2.Client, error) {
	orgName := s.Args["org_name"].(string)
	baseURL := s.Args["base_url"].(string)
	apiToken := s.Args["api_token"].(string)

	orgURL := fmt.Sprintf("https://%v.%v", orgName, baseURL)

	ctx, client, err := oktaV2.NewClient(
		context.Background(),
		oktaV2.WithOrgUrl(orgURL),
		oktaV2.WithToken(apiToken),
	)
	if err != nil {
		return ctx, nil, err
	}

	return ctx, client, nil
}

func (s *OktaService) ClientV5() (context.Context, *oktaV5.APIClient, error) {
	orgName := s.Args["org_name"].(string)
	baseURL := s.Args["base_url"].(string)
	apiToken := s.Args["api_token"].(string)

	orgURL := fmt.Sprintf("https://%v.%v", orgName, baseURL)

	config, err := oktaV5.NewConfiguration(
		oktaV5.WithOrgUrl(orgURL),
		oktaV5.WithToken(apiToken),
	)
	if err != nil {
		return nil, nil, err
	}
	client := oktaV5.NewAPIClient(config)

	return context.Background(), client, nil
}

func (s *OktaService) APISupplementClient() (context.Context, *sdk.APISupplement, error) {
	baseURL := s.Args["base_url"].(string)
	orgName := s.Args["org_name"].(string)
	apiToken := s.Args["api_token"].(string)

	orgURL := fmt.Sprintf("https://%v.%v", orgName, baseURL)

	ctx, client, err := oktaV2.NewClient(
		context.Background(),
		oktaV2.WithOrgUrl(orgURL),
		oktaV2.WithToken(apiToken),
	)
	if err != nil {
		return ctx, nil, err
	}

	apiSupplementClient := &sdk.APISupplement{
		RequestExecutor: client.CloneRequestExecutor(),
	}

	return ctx, apiSupplementClient, nil
}
