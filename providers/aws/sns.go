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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var snsAllowEmptyValues = []string{"tags."}

type SnsGenerator struct {
	AWSService
}

// TF currently doesn't support email subscriptions + subscriptions with pending confirmations
func (g *SnsGenerator) isSupportedSubscription(protocol, subscriptionId string) bool {
	return protocol != "email" && protocol != "email-json" && subscriptionId != "PendingConfirmation"
}

func (g *SnsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := sns.New(config)
	p := sns.NewListTopicsPaginator(svc.ListTopicsRequest(&sns.ListTopicsInput{}))
	for p.Next(context.Background()) {
		for _, topic := range p.CurrentPage().Topics {
			arnParts := strings.Split(aws.StringValue(topic.TopicArn), ":")
			topicName := arnParts[len(arnParts)-1]

			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				aws.StringValue(topic.TopicArn),
				topicName,
				"aws_sns_topic",
				"aws",
				snsAllowEmptyValues,
			))

			topicSubsPage := sns.NewListSubscriptionsByTopicPaginator(svc.ListSubscriptionsByTopicRequest(&sns.ListSubscriptionsByTopicInput{
				TopicArn: topic.TopicArn,
			}))
			for topicSubsPage.Next(context.Background()) {
				for _, subscription := range topicSubsPage.CurrentPage().Subscriptions {
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
			}
			if err := topicSubsPage.Err(); err != nil {
				log.Println(err)
			}
		}
	}
	return p.Err()
}
