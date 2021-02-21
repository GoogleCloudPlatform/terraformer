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
	"fmt"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var sqsAllowEmptyValues = []string{"tags."}

type SqsGenerator struct {
	AWSService
}

func (g *SqsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := sqs.NewFromConfig(config)

	listQueuesInput := sqs.ListQueuesInput{}

	sqsPrefix, hasPrefix := os.LookupEnv("SQS_PREFIX")
	if hasPrefix {
		listQueuesInput.QueueNamePrefix = aws.String(sqsPrefix)
	}

	queuesList, err := svc.ListQueues(context.TODO(), &listQueuesInput)

	if err != nil {
		return err
	}

	for _, queueURL := range queuesList.QueueUrls {
		urlParts := strings.Split(queueURL, "/")
		queueName := urlParts[len(urlParts)-1]

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			queueURL,
			queueName,
			"aws_sqs_queue",
			"aws",
			sqsAllowEmptyValues,
		))
	}

	return nil
}

// PostConvertHook for add policy json as heredoc
func (g *SqsGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_sqs_queue" {
			if val, ok := g.Resources[i].Item["policy"]; ok {
				policy := g.escapeAwsInterpolation(val.(string))
				g.Resources[i].Item["policy"] = fmt.Sprintf(`<<POLICY
%s
POLICY`, policy)
			}
		}
	}
	return nil
}
