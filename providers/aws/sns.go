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

	"github.com/aws/aws-sdk-go/service/sns"
)

var snsAllowEmptyValues = []string{"tags."}

type SnsGenerator struct {
	AWSService
}

// TF currently doesn't support email subscriptions + subscriptions with pending confirmations
func (g SnsGenerator) isSupportedSubscription(protocol, subscriptionId string) bool {
	return protocol != "email" && protocol != "email-json" && subscriptionId != "PendingConfirmation"
}

func (g *SnsGenerator) InitResources() error {
	sess := g.generateSession()
	svc := sns.New(sess)

	err := svc.ListTopicsPages(&sns.ListTopicsInput{}, func(topics *sns.ListTopicsOutput, lastPage bool) bool {
		for _, topic := range topics.Topics {
			arnParts := strings.Split(aws.StringValue(topic.TopicArn), ":")
			topicName := arnParts[len(arnParts)-1]

			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(topic.TopicArn),
				topicName,
				"aws_sns_topic",
				"aws",
				snsAllowEmptyValues,
			))

			_ = svc.ListSubscriptionsByTopicPages(&sns.ListSubscriptionsByTopicInput{
				TopicArn: topic.TopicArn,
			}, func(subscriptions *sns.ListSubscriptionsByTopicOutput, lastPage bool) bool {
				for _, subscription := range subscriptions.Subscriptions {
					subscriptionArnParts := strings.Split(aws.StringValue(subscription.SubscriptionArn), ":")
					subscriptionId := subscriptionArnParts[len(subscriptionArnParts)-1]

					if g.isSupportedSubscription(aws.StringValue(subscription.Protocol), subscriptionId) {
						g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
							aws.StringValue(subscription.SubscriptionArn),
							"subscription-"+subscriptionId,
							"aws_sns_topic_subscription",
							"aws",
							snsAllowEmptyValues,
						))
					}
				}

				return !lastPage
			})
		}

		return !lastPage
	})
	if err != nil {
		return err
	}
	return nil
}
