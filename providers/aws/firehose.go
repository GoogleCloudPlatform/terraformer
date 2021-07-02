// Copyright 2018 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

type FirehoseGenerator struct {
	AWSService
}

func (g *FirehoseGenerator) createResources(streamNames []string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, resourceName := range streamNames {
		resources = append(resources, terraformutils.NewResource(
			resourceName,
			resourceName,
			"aws_kinesis_firehose_delivery_stream",
			"aws",
			map[string]string{"name": resourceName},
			[]string{".tags"},
			map[string]interface{}{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// Need deliver stream name for terraform resource
func (g *FirehoseGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := firehose.NewFromConfig(config)
	var streamNames []string
	var lastStreamName *string
	for {
		output, err := svc.ListDeliveryStreams(context.TODO(), &firehose.ListDeliveryStreamsInput{
			ExclusiveStartDeliveryStreamName: lastStreamName,
			Limit:                            aws.Int32(100),
		})
		if err != nil {
			return err
		}
		streamNames = append(streamNames, output.DeliveryStreamNames...)
		if !*output.HasMoreDeliveryStreams {
			break
		}

		lastStreamName = aws.String(streamNames[len(streamNames)-1])
	}

	g.Resources = g.createResources(streamNames)

	return nil
}

func (g *FirehoseGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		_, hasExtendedS3Configuration := resource.Item["extended_s3_configuration"]
		_, hasS3Configuration := resource.Item["s3_configuration"]
		if hasExtendedS3Configuration && hasS3Configuration {
			delete(resource.Item, "s3_configuration")
		}
	}
	return nil
}
