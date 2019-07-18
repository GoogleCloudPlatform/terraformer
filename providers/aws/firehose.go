package aws


import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

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


type FirehoseGenerator struct {
	AWSService
}

func (g FirehoseGenerator) createResources(sess *session.Session, streamNames []*string, region string) []terraform_utils.Resource {

	var resources []terraform_utils.Resource
	for _, streamName := range streamNames {
		resourceName := aws.StringValue(streamName)

		resources = append(resources, terraform_utils.NewResource(
			resourceName,
			resourceName,
			"aws_kinesis_firehose_delivery_stream",
			"aws",
			map[string]string{"name": resourceName},
			[]string{".tags"},
			map[string]string{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// Need deliver stream name for terraform resource
func (g *FirehoseGenerator) InitResources() error {
	sess := g.generateSession()
	svc := firehose.New(sess)
	var streamNames []*string
	for {
		output, err := svc.ListDeliveryStreams(&firehose.ListDeliveryStreamsInput{})
		if err != nil {
			return err
		}
		streamNames = append(streamNames, output.DeliveryStreamNames...)
		if *output.HasMoreDeliveryStreams == false {
			break
		}
	}

	g.Resources = g.createResources(sess, streamNames, g.GetArgs()["region"].(string))

	g.PopulateIgnoreKeys()
	return nil
}
