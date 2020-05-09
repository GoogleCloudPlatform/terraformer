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
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

var ecrAllowEmptyValues = []string{"tags."}

type EcrGenerator struct {
	AWSService
}

func (g *EcrGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ecr.New(config)

	p := ecr.NewDescribeRepositoriesPaginator(svc.DescribeRepositoriesRequest(&ecr.DescribeRepositoriesInput{}))
	for p.Next(context.Background()) {
		for _, repository := range p.CurrentPage().Repositories {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*repository.RepositoryName,
				*repository.RepositoryName,
				"aws_ecr_repository",
				"aws",
				ecrAllowEmptyValues))
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*repository.RepositoryName,
				*repository.RepositoryName,
				"aws_ecr_repository_policy",
				"aws",
				ecrAllowEmptyValues))
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*repository.RepositoryName,
				*repository.RepositoryName,
				"aws_ecr_lifecycle_policy",
				"aws",
				ecrAllowEmptyValues))
		}
	}
	return p.Err()
}
