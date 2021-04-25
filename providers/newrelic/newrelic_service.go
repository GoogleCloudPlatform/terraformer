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

package newrelic

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	synthetics "github.com/dollarshaveclub/new-relic-synthetics-go"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

type NewRelicService struct { //nolint
	terraformutils.Service
}

func (s *NewRelicService) Client() (*newrelic.Client, error) {
	client := newrelic.New(newrelic.Config{APIKey: s.GetArgs()["apiKey"].(string)})
	return &client, nil
}

func (s *NewRelicService) InfraClient() (*newrelic.InfraClient, error) {
	client := newrelic.NewInfraClient(newrelic.Config{APIKey: s.GetArgs()["apiKey"].(string), Debug: s.Verbose})
	return &client, nil
}

func (s *NewRelicService) SyntheticsClient() (*synthetics.Client, error) {
	conf := func(c *synthetics.Client) {
		c.APIKey = s.GetArgs()["apiKey"].(string)
	}
	return synthetics.NewClient(conf)
}
