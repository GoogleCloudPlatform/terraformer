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
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

var kinesisAllowEmptyValues = []string{"tags."}

type KinesisGenerator struct {
	AWSService
}

func (g *KinesisGenerator) createResources(streamNames []string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, resourceName := range streamNames {
		resources = append(resources, terraformutils.NewResource(
			resourceName,
			resourceName,
			"aws_kinesis_stream",
			"aws",
			map[string]string{"name": resourceName},
			kinesisAllowEmptyValues,
			map[string]interface{}{}))
	}
	return resources
}

func (g *KinesisGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := kinesis.New(config)
	p := kinesis.NewListStreamsPaginator(svc.ListStreamsRequest(&kinesis.ListStreamsInput{}))
	for p.Next(context.Background()) {
		g.Resources = append(g.Resources, g.createResources(p.CurrentPage().StreamNames)...)
	}
	return p.Err()
}
