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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	cf "github.com/cloudflare/cloudflare-go"
)

type CloudflareService struct {
	terraform_utils.Service
}

func (s *CloudflareService) initializeAPI() (*cf.API, error) {
	apiKey := os.Getenv("CLOUDFLARE_TOKEN")
	apiEmail := os.Getenv("CLOUDFLARE_EMAIL")

	if apiEmail == "" || apiKey == "" {
		err := errors.New("No CLOUDFLARE_TOKEN/CLOUDFLARE_EMAIL environment set")
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return cf.New(apiKey, apiEmail)
}
