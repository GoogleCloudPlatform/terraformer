// Copyright 2020 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

var codepipelineAllowEmptyValues = []string{"tags."}

type CodePipelineGenerator struct {
	AWSService
}

func (g *CodePipelineGenerator) loadPipelines(svc *codepipeline.Client) error {
	p := codepipeline.NewListPipelinesPaginator(svc, &codepipeline.ListPipelinesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, pipeline := range page.Pipelines {
			resourceName := StringValue(pipeline.Name)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_codepipeline",
				"aws",
				codepipelineAllowEmptyValues))
		}
	}
	return nil
}

func (g *CodePipelineGenerator) loadWebhooks(svc *codepipeline.Client) error {
	p := codepipeline.NewListWebhooksPaginator(svc, &codepipeline.ListWebhooksInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, webhook := range page.Webhooks {
			resourceArn := StringValue(webhook.Arn)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceArn,
				resourceArn,
				"aws_codepipeline_webhook",
				"aws",
				codepipelineAllowEmptyValues))
		}
	}
	return nil
}

func (g *CodePipelineGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := codepipeline.NewFromConfig(config)

	if err := g.loadPipelines(svc); err != nil {
		return err
	}
	if err := g.loadWebhooks(svc); err != nil {
		return err
	}

	return nil
}
