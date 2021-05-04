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

type PagerDutyService struct { //nolint
	terraformutils.Service
}

func (s *PagerDutyService) Client() (*pagerduty.Client, error) {
	client, err := pagerduty.NewClient(&pagerduty.Config{Token: s.GetArgs()["token"].(string)})
	if err != nil {
		return nil, err
	}
	return client, nil
}
