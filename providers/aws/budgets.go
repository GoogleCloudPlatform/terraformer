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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
)

type BudgetsGenerator struct {
	AWSService
}

func (g *BudgetsGenerator) createResources(budgets []types.Budget, account *string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, budget := range budgets {
		resourceName := StringValue(budget.BudgetName)
		resources = append(resources, terraformutils.NewSimpleResource(
			fmt.Sprintf("%s:%s", *account, resourceName),
			resourceName,
			"aws_budgets_budget",
			"aws",
			[]string{}))
	}
	return resources
}

func (g *BudgetsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	budgetsSvc := budgets.NewFromConfig(config)

	account, err := g.getAccountNumber(config)
	if err != nil {
		return err
	}

	output, err := budgetsSvc.DescribeBudgets(context.TODO(), &budgets.DescribeBudgetsInput{AccountId: account})
	if err != nil {
		return err
	}

	g.Resources = g.createResources(output.Budgets, account)
	return nil
}
