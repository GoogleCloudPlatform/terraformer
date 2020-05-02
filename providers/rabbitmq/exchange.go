// Copyright 2018 The Terraformer Authors.
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

package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ExchangeGenerator struct {
	RBTService
}

type Exchange struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

type Exchanges []Exchange

var ExchangeAllowEmptyValues = []string{}
var ExchangeAdditionalFields = map[string]interface{}{}

func (g ExchangeGenerator) createResources(exchanges Exchanges) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, exchange := range exchanges {
		if len(exchange.Name) == 0 {
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", exchange.Name, exchange.Vhost),
			fmt.Sprintf("exchange_%s_%s", normalizeResourceName(exchange.Vhost), normalizeResourceName(exchange.Name)),
			"rabbitmq_exchange",
			"rabbitmq",
			map[string]string{
				"name":  exchange.Name,
				"vhost": exchange.Vhost,
			},
			ExchangeAllowEmptyValues,
			ExchangeAdditionalFields,
		))
	}
	return resources
}

func (g *ExchangeGenerator) InitResources() error {
	body, err := g.generateRequest("/api/exchanges?columns=name,vhost")
	if err != nil {
		return err
	}
	var exchanges Exchanges
	err = json.Unmarshal(body, &exchanges)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(exchanges)
	return nil
}
