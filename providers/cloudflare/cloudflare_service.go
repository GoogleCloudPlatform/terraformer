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

package cloudflare

import (
	"errors"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type CloudflareService struct { //nolint
	terraformutils.Service
}

func (s *CloudflareService) initializeAPI() (*cf.API, error) {
	apiKey := os.Getenv("CLOUDFLARE_API_KEY")
	apiEmail := os.Getenv("CLOUDFLARE_EMAIL")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	if apiToken == "" && (apiEmail == "" || apiKey == "") {
		err := errors.New("Either CLOUDFLARE_API_TOKEN or CLOUDFLARE_API_KEY/CLOUDFLARE_EMAIL environment variables must be set")
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	if apiToken != "" {
		return cf.NewWithAPIToken(apiToken, cf.UsingAccount(accountID))
	}

	return cf.New(apiKey, apiEmail, cf.UsingAccount(accountID))
}
