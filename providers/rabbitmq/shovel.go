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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ShovelGenerator struct {
	RBTService
}

type Shovel struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

type ShovelDefinition struct {
	AckMode                          string   `json:"ack-mode,omitempty"`
	AddForwardHeaders                bool     `json:"add-forward-headers,omitempty"`
	DestinationAddForwardHeaders     bool     `json:"dest-add-forward-headers,omitempty"`
	DestinationAddTimestampHeader    bool     `json:"dest-add-timestamp-header,omitempty"`
	DestinationAddress               string   `json:"dest-address,omitempty"`
	DestinationApplicationProperties string   `json:"dest-application-properties,omitempty"`
	DestinationExchange              string   `json:"dest-exchange,omitempty"`
	DestinationExchangeKey           string   `json:"dest-exchange-key,omitempty"`
	DestinationProperties            string   `json:"dest-properties,omitempty"`
	DestinationProtocol              string   `json:"dest-protocol,omitempty"`
	DestinationPublishProperties     string   `json:"dest-publish-properties,omitempty"`
	DestinationQueue                 string   `json:"dest-queue,omitempty"`
	DestinationURI                   []string `json:"dest-uri"`
	PrefetchCount                    int      `json:"prefetch-count,omitempty"`
	ReconnectDelay                   int      `json:"reconnect-delay,omitempty"`
	SourceAddress                    string   `json:"src-address,omitempty"`
	SourceDeleteAfter                string   `json:"src-delete-after,omitempty"`
	SourceExchange                   string   `json:"src-exchange,omitempty"`
	SourceExchangeKey                string   `json:"src-exchange-key,omitempty"`
	SourcePrefetchCount              int      `json:"src-prefetch-count,omitempty"`
	SourceProtocol                   string   `json:"src-protocol,omitempty"`
	SourceQueue                      string   `json:"src-queue,omitempty"`
	SourceURI                        []string `json:"src-uri"`
}

type ShovelInfo struct {
	// Shovel name
	Name string `json:"name"`
	// Virtual host this shovel belongs to
	Vhost string `json:"vhost"`
	// Component shovels belong to
	Component string `json:"component"`
	// Details the configuration values of the shovel
	Definition ShovelDefinition `json:"value"`
}

type Shovels []ShovelInfo

var ShovelAllowEmptyValues = []string{}
var ShovelAdditionalFields = map[string]interface{}{}

func (g ShovelGenerator) createResources(shovels Shovels) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, shovel := range shovels {
		if len(shovel.Name) == 0 {
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", shovel.Name, shovel.Vhost),
			fmt.Sprintf("shovel_%s_%s", normalizeResourceName(shovel.Vhost), normalizeResourceName(shovel.Name)),
			"rabbitmq_shovel",
			"rabbitmq",
			map[string]string{
				"name":  shovel.Name,
				"vhost": shovel.Vhost,
			},
			ShovelAllowEmptyValues,
			ShovelAdditionalFields,
		))
	}
	log.Println(resources)
	return resources
}

func (g *ShovelGenerator) InitResources() error {
	body, err := g.generateRequest("/api/shovels?columns=name,vhost")
	if err != nil {
		return err
	}
	var shovels Shovels
	err = json.Unmarshal(body, &shovels)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(shovels)
	return nil
}
