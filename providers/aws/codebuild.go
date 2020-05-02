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
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
)

var codebuildAllowEmptyValues = []string{"tags."}

type CodeBuildGenerator struct {
	AWSService
}

func (g *CodeBuildGenerator) createResources(projectList []string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, project := range projectList {
		resources = append(resources, terraformutils.NewSimpleResource(
			project,
			project,
			"aws_codebuild_project",
			"aws",
			codebuildAllowEmptyValues))
	}
	return resources
}

func (g *CodeBuildGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := codebuild.New(config)
	output, err := svc.ListProjectsRequest(&codebuild.ListProjectsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output.Projects)
	return nil
}
