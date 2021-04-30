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
	"fmt"
	"github.com/zclconf/go-cty/cty"

	"github.com/aws/aws-sdk-go-v2/service/ecr"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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

	svc := ecr.NewFromConfig(config)

	p := ecr.NewDescribeRepositoriesPaginator(svc, &ecr.DescribeRepositoriesInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, repository := range page.Repositories {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*repository.RepositoryName,
				*repository.RepositoryName,
				"aws_ecr_repository",
				"aws",
				ecrAllowEmptyValues))

			_, err := svc.GetRepositoryPolicy(context.TODO(), &ecr.GetRepositoryPolicyInput{
				RepositoryName: repository.RepositoryName,
				RegistryId:     repository.RegistryId,
			})
			if err == nil {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					*repository.RepositoryName,
					*repository.RepositoryName,
					"aws_ecr_repository_policy",
					"aws",
					ecrAllowEmptyValues))
			}

			_, err = svc.GetLifecyclePolicy(context.TODO(), &ecr.GetLifecyclePolicyInput{
				RepositoryName: repository.RepositoryName,
				RegistryId:     repository.RegistryId,
			})
			if err == nil {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					*repository.RepositoryName,
					*repository.RepositoryName,
					"aws_ecr_lifecycle_policy",
					"aws",
					ecrAllowEmptyValues))
			}
		}
	}
	return nil
}

func (g *EcrGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.Address.Type == "aws_ecr_repository_policy" || resource.Address.Type == "aws_ecr_lifecycle_policy" {
			if resource.InstanceState.Value.HasIndex(cty.StringVal("policy")) == cty.True {
				instanceStateMap := resource.InstanceState.Value.AsValueMap()
				policy := g.escapeAwsInterpolation(instanceStateMap["policy"].AsString())
				instanceStateMap["policy"] = cty.StringVal(fmt.Sprintf(`<<POLICY
%s
POLICY`, policy))
				resource.InstanceState.Value = cty.ObjectVal(instanceStateMap)
			}
		}
	}
	return nil
}
