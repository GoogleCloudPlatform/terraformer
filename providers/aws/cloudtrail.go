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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

var cloudtrailAllowEmptyValues = []string{"tags."}

type CloudTrailGenerator struct {
	AWSService
}

func (g *CloudTrailGenerator) createResources(trailList []cloudtrail.Trail) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, trail := range trailList {
		resourceName := aws.StringValue(trail.Name)
		resources = append(resources, terraformutils.NewSimpleResource(
			resourceName,
			resourceName,
			"aws_cloudtrail",
			"aws",
			cloudtrailAllowEmptyValues))
	}
	return resources
}

func (g *CloudTrailGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := cloudtrail.New(config)
	output, err := svc.DescribeTrailsRequest(&cloudtrail.DescribeTrailsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output.TrailList)
	return nil
}
