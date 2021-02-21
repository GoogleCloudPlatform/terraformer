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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamodbAllowEmptyValues = []string{"tags."}

type DynamoDbGenerator struct {
	AWSService
}

func (g *DynamoDbGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := dynamodb.NewFromConfig(config)
	p := dynamodb.NewListTablesPaginator(svc, &dynamodb.ListTablesInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, tableName := range page.TableNames {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				tableName,
				tableName,
				"aws_dynamodb_table",
				"aws",
				dynamodbAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *DynamoDbGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_dynamodb_table" {
			continue
		}
		if val, ok := r.InstanceState.Attributes["ttl.0.enabled"]; ok && val == "false" {
			delete(r.Item, "ttl")
		}
	}
	return nil
}
