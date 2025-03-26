// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/mq"
)

var mqAllowEmptyValues = []string{"tags."}

type MQGenerator struct {
	AWSService
}

func (g *MQGenerator) loadBrokers(svc *mq.Client) error {
	p := mq.NewListBrokersPaginator(svc, &mq.ListBrokersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, broker := range page.BrokerSummaries {
			resourceName := StringValue(broker.BrokerName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_mq_broker",
				"aws",
				mqAllowEmptyValues))
		}
	}
	return nil
}

func (g *MQGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := mq.NewFromConfig(config)

	err := g.loadBrokers(svc)
	if err != nil {
		return err
	}

	return nil
}
