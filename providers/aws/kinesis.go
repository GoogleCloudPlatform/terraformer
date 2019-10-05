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
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var kinesisAllowEmptyValues = []string{"tags."}

type KinesisGenerator struct {
	AWSService
}

func (g KinesisGenerator) createResources(streamNames []*string) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, streamName := range streamNames {
		resourceName := aws.StringValue(streamName)
		resources = append(resources, terraform_utils.NewResource(
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
	sess := g.generateSession()
	svc := kinesis.New(sess)
	output, err := svc.ListStreams(&kinesis.ListStreamsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output.StreamNames)
	return nil
}