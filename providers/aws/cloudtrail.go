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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

var cloudtrailAllowEmptyValues = []string{"tags."}

type CloudTrailGenerator struct {
	AWSService
}

func (g CloudTrailGenerator) createResources(sess *session.Session, trailList []*cloudtrail.Trail) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, trail := range trailList {
		resourceName := aws.StringValue(trail.Name)
		resources = append(resources, terraform_utils.NewSimpleResource(
			resourceName,
			resourceName,
			"aws_cloudtrail",
			"aws",
			cloudtrailAllowEmptyValues))
	}
	return resources
}

func (g *CloudTrailGenerator) InitResources() error {
	sess := g.generateSession()
	svc := cloudtrail.New(sess)
	output, err := svc.DescribeTrails(&cloudtrail.DescribeTrailsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(sess, output.TrailList)
	return nil
}
