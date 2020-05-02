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
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	synthetics "github.com/dollarshaveclub/new-relic-synthetics-go"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

type NewRelicService struct { //nolint
	terraformutils.Service
}

func (s *NewRelicService) Client() (*newrelic.Client, error) {
	apiKey := os.Getenv("NEWRELIC_API_KEY")

	if apiKey == "" {
		err := errors.New("No NEWRELIC_API_KEY environment set")
		return nil, err
	}

	client := newrelic.New(newrelic.Config{APIKey: apiKey})
	return &client, nil
}

func (s *NewRelicService) InfraClient() (*newrelic.InfraClient, error) {
	apiKey := os.Getenv("NEWRELIC_API_KEY")

	if apiKey == "" {
		err := errors.New("No NEWRELIC_API_KEY environment set")
		return nil, err
	}

	client := newrelic.NewInfraClient(newrelic.Config{APIKey: apiKey, Debug: s.Verbose})

	return &client, nil
}

func (s *NewRelicService) SyntheticsClient() (*synthetics.Client, error) {
	apiKey := os.Getenv("NEWRELIC_API_KEY")

	if apiKey == "" {
		err := errors.New("No NEWRELIC_API_KEY environment set")
		return nil, err
	}

	conf := func(c *synthetics.Client) {
		c.APIKey = apiKey
	}

	return synthetics.NewClient(conf)
}
