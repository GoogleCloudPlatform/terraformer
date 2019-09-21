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
	"strings"

	"github.com/aws/aws-sdk-go/service/sqs"
)

var sqsAllowEmptyValues = []string{"tags."}

type SqsGenerator struct {
	AWSService
}

func (g *SqsGenerator) InitResources() error {
	sess := g.generateSession()
	svc := sqs.New(sess)

	queuesList, err := svc.ListQueues(&sqs.ListQueuesInput{})

	if err != nil {
		return err
	}

	for _, queueUrl := range queuesList.QueueUrls {
		urlParts := strings.Split(aws.StringValue(queueUrl), "/")
		queueName := urlParts[len(urlParts)-1]

		g.Resources = append(g.Resources, terraform_utils.NewResource(
			aws.StringValue(queueUrl),
			queueName,
			"aws_sqs_queue",
			"aws",
			map[string]string{},
			sqsAllowEmptyValues,
			map[string]string{},
		))
	}

	g.PopulateIgnoreKeys()
	return nil
}
