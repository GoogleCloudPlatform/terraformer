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
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

var datapipelineAllowEmptyValues = []string{"tags."}

type DataPipelineGenerator struct {
	AWSService
}

func (g *DataPipelineGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := datapipeline.NewFromConfig(config)
	p := datapipeline.NewListPipelinesPaginator(svc, &datapipeline.ListPipelinesInput{})
	var resources []terraformutils.Resource
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, pipeline := range page.PipelineIdList {
			pipelineID := StringValue(pipeline.Id)
			pipelineName := StringValue(pipeline.Name)
			resources = append(resources, terraformutils.NewSimpleResource(
				pipelineID,
				pipelineName,
				"aws_datapipeline_pipeline",
				"aws",
				datapipelineAllowEmptyValues))
		}
	}
	g.Resources = resources
	return nil
}
