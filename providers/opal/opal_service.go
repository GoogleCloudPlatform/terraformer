// Copyright 2022 The Terraformer Authors.
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

package opal

import (
	"fmt"
	"net/url"
	"path"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/opalsecurity/opal-go"
)

type OpalService struct { //nolint
	terraformutils.Service
}

func (s *OpalService) newClient() (*opal.APIClient, error) {
	conf := opal.NewConfiguration()

	conf.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", s.GetArgs()["token"].(string))
	u, err := url.Parse(s.GetArgs()["base_url"].(string))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/v1")
	conf.Servers = opal.ServerConfigurations{{
		URL: u.String(),
	}}

	return opal.NewAPIClient(conf), nil
}
