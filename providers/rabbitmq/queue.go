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

type QueueGenerator struct {
	RBTService
}

type Queue struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

type Queues []Queue

var QueueAllowEmptyValues = []string{}
var QueueAdditionalFields = map[string]interface{}{}

func (g QueueGenerator) createResources(queues Queues) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, queue := range queues {
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s@%s", queue.Name, queue.Vhost),
			fmt.Sprintf("queue_%s_%s", normalizeResourceName(queue.Vhost), normalizeResourceName(queue.Name)),
			"rabbitmq_queue",
			"rabbitmq",
			map[string]string{
				"name":  queue.Name,
				"vhost": queue.Vhost,
			},
			QueueAllowEmptyValues,
			QueueAdditionalFields,
		))
	}
	return resources
}

func (g *QueueGenerator) InitResources() error {
	body, err := g.generateRequest("/api/queues?columns=name,vhost")
	if err != nil {
		return err
	}
	var queues Queues
	err = json.Unmarshal(body, &queues)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(queues)
	return nil
}
