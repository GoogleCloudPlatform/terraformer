// Copyright 2020 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
)

var securityhubAllowEmptyValues = []string{"tags."}

type SecurityhubGenerator struct {
	AWSService
}

func (g *SecurityhubGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	client := securityhub.New(config)

	account, err := g.getAccountNumber(config)
	if err != nil {
		return err
	}

	err = g.addAccount(client, *account)
	if err != nil {
		return err
	}
	err = g.addMembers(client, *account)
	if err != nil {
		return err
	}
	err = g.addStandardsSubscription(client, *account)
	return err
}

func (g *SecurityhubGenerator) addAccount(client *securityhub.Client, accountNumber string) error {
	_, err := client.GetEnabledStandardsRequest(&securityhub.GetEnabledStandardsInput{}).Send(context.Background())

	if err != nil {
		if !strings.Contains(err.(awserr.Error).Message(), "not subscribed to AWS Security Hub") {
			return err
		}
	} else {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			accountNumber,
			accountNumber,
			"aws_securityhub_account",
			"aws",
			securityhubAllowEmptyValues,
		))
	}
	return nil
}

func (g *SecurityhubGenerator) addMembers(svc *securityhub.Client, accountNumber string) error {
	p := securityhub.NewListMembersPaginator(svc.ListMembersRequest(&securityhub.ListMembersInput{}))

	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, member := range page.Members {
			id := *member.AccountId
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				id,
				"securityhub_member_"+id,
				"aws_securityhub_member",
				"aws",
				map[string]string{
					"account_id": id,
					"email":      *member.Email,
				},
				securityhubAllowEmptyValues,
				map[string]interface{}{
					"depends_on": []string{"${aws_securityhub_account.tfer--" + accountNumber + "}"},
				},
			))
		}
	}
	return p.Err()
}

func (g *SecurityhubGenerator) addStandardsSubscription(svc *securityhub.Client, accountNumber string) error {
	p := securityhub.NewGetEnabledStandardsPaginator(
		svc.GetEnabledStandardsRequest(&securityhub.GetEnabledStandardsInput{}))

	for p.Next(context.Background()) {
		page := p.CurrentPage()
		for _, standardsSubscription := range page.StandardsSubscriptions {
			id := *standardsSubscription.StandardsSubscriptionArn
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				id,
				id,
				"aws_securityhub_standards_subscription",
				"aws",
				map[string]string{
					"standards_arn": id,
				},
				securityhubAllowEmptyValues,
				map[string]interface{}{
					"depends_on": []string{"aws_securityhub_account.tfer--" + accountNumber},
				},
			))
		}
	}
	return p.Err()
}
