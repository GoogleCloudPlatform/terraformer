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
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

var devicefarmAllowEmptyValues = []string{"tags."}

type DeviceFarmGenerator struct {
	AWSService
}

func (g *DeviceFarmGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := devicefarm.NewFromConfig(config)
	p := devicefarm.NewListProjectsPaginator(svc, &devicefarm.ListProjectsInput{})
	var resources []terraformutils.Resource
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, project := range page.Projects {
			projectArn := StringValue(project.Arn)
			projectName := StringValue(project.Name)
			resources = append(resources, terraformutils.NewSimpleResource(
				projectArn,
				projectName,
				"aws_devicefarm_project",
				"aws",
				devicefarmAllowEmptyValues))
		}
	}
	g.Resources = resources
	return nil
}
