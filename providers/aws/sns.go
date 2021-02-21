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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var snsAllowEmptyValues = []string{"tags."}

type SnsGenerator struct {
	AWSService
}

// TF currently doesn't support email subscriptions + subscriptions with pending confirmations
func (g *SnsGenerator) isSupportedSubscription(protocol, subscriptionID string) bool {
	return protocol != "email" && protocol != "email-json" && subscriptionID != "PendingConfirmation"
}

func (g *SnsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := sns.NewFromConfig(config)
	p := sns.NewListTopicsPaginator(svc, &sns.ListTopicsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, topic := range page.Topics {
			arnParts := strings.Split(StringValue(topic.TopicArn), ":")
			topicName := arnParts[len(arnParts)-1]

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(topic.TopicArn),
				topicName,
				"aws_sns_topic",
				"aws",
				snsAllowEmptyValues,
			))

			topicSubsPage := sns.NewListSubscriptionsByTopicPaginator(svc, &sns.ListSubscriptionsByTopicInput{
				TopicArn: topic.TopicArn,
			})
			for topicSubsPage.HasMorePages() {
				topicSubsNextPage, err := topicSubsPage.NextPage(context.TODO())
				if err != nil {
					log.Println(err)
					continue
				}
				for _, subscription := range topicSubsNextPage.Subscriptions {
					subscriptionArnParts := strings.Split(StringValue(subscription.SubscriptionArn), ":")
					subscriptionID := subscriptionArnParts[len(subscriptionArnParts)-1]

					if g.isSupportedSubscription(StringValue(subscription.Protocol), subscriptionID) {
						g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
							StringValue(subscription.SubscriptionArn),
							"subscription-"+subscriptionID,
							"aws_sns_topic_subscription",
							"aws",
							snsAllowEmptyValues,
						))
					}
				}
			}
		}
	}
	return nil
}

// PostConvertHook for add policy json as heredoc
func (g *SnsGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_sns_topic" {
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
