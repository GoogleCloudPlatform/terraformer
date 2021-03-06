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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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
	client := securityhub.NewFromConfig(config)

	account, err := g.getAccountNumber(config)
	if err != nil {
		return err
	}

	accountDisabled, err := g.addAccount(client, *account)
	if accountDisabled {
		return nil
	}
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

func (g *SecurityhubGenerator) addAccount(client *securityhub.Client, accountNumber string) (bool, error) {
	_, err := client.GetEnabledStandards(context.TODO(), &securityhub.GetEnabledStandardsInput{})

	if err != nil {
		errorMsg := err.Error()
		if !strings.Contains(errorMsg, "not subscribed to AWS Security Hub") {
			return false, err
		}
		return true, nil
	}
	g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
		accountNumber,
		accountNumber,
		"aws_securityhub_account",
		"aws",
		securityhubAllowEmptyValues,
	))
	return false, nil
}

func (g *SecurityhubGenerator) addMembers(svc *securityhub.Client, accountNumber string) error {
	p := securityhub.NewListMembersPaginator(svc, &securityhub.ListMembersInput{})

	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, member := range page.Members {
			id := *member.AccountId
			attributes := map[string]string{
				"account_id": id,
			}
			if member.Email != nil {
				attributes["email"] = *member.Email
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				id,
				"securityhub_member_"+id,
				"aws_securityhub_member",
				"aws",
				attributes,
				securityhubAllowEmptyValues,
				map[string]interface{}{
					"depends_on": []string{"${aws_securityhub_account.tfer--" + accountNumber + "}"},
				},
			))
		}
	}
	return nil
}

func (g *SecurityhubGenerator) addStandardsSubscription(svc *securityhub.Client, accountNumber string) error {
	p := securityhub.NewGetEnabledStandardsPaginator(svc, &securityhub.GetEnabledStandardsInput{})

	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, standardsSubscription := range page.StandardsSubscriptions {
			id := *standardsSubscription.StandardsSubscriptionArn
			g.Resources = append(g.Resources, terraformutils.NewResource(
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
	return nil
}
