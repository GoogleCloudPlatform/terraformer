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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/sts"
)

type BudgetsGenerator struct {
	AWSService
}

func (g BudgetsGenerator) createResources(budgets []*budgets.Budget, account *string) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, budget := range budgets {
		resourceName := aws.StringValue(budget.BudgetName)
		resources = append(resources, terraform_utils.NewSimpleResource(
			fmt.Sprintf("%s:%s", *account, resourceName),
			resourceName,
			"aws_budgets_budget",
			"aws",
			[]string{}))
	}
	return resources
}

func (g *BudgetsGenerator) InitResources() error {
	sess := g.generateSession()
	stsSvc := sts.New(sess)
	budgetsSvc := budgets.New(sess)

	identity, err := stsSvc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}
	account := identity.Account

	output, err := budgetsSvc.DescribeBudgets(&budgets.DescribeBudgetsInput{AccountId: account})
	if err != nil {
		return err
	}

	g.Resources = g.createResources(output.Budgets, account)
	return nil
}
