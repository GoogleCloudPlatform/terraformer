// Copyright 2021 The Terraformer Authors.
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

	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var ecrPublicAllowEmptyValues = []string{"tags."}

type EcrPublicGenerator struct {
	AWSService
}

func (g *EcrPublicGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}

	ecrPublicConfig := config.Copy()
	ecrPublicConfig.Region = MainRegionPublicPartition
	svc := ecrpublic.NewFromConfig(ecrPublicConfig)

	p := ecrpublic.NewDescribeRepositoriesPaginator(svc, &ecrpublic.DescribeRepositoriesInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, repository := range page.Repositories {
			resource := terraformutils.NewSimpleResource(
				*repository.RepositoryName,
				*repository.RepositoryName,
				"aws_ecrpublic_repository",
				"aws",
				ecrPublicAllowEmptyValues)
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}
