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

package heroku

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type AccountFeatureGenerator struct {
	HerokuService
}

func (g AccountFeatureGenerator) createResources(accountFeatureList []heroku.AccountFeature) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, accountFeature := range accountFeatureList {
		resources = append(resources, terraformutils.NewResource(
			accountFeature.ID,
			accountFeature.Name,
			"heroku_account_feature",
			"heroku",
			map[string]string{"name": accountFeature.Name},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *AccountFeatureGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()
	list := []heroku.AccountFeature{}

	accountFeatures, err := svc.AccountFeatureList(ctx, &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	for _, accountFeature := range accountFeatures {
		if accountFeature.Enabled {
			list = append(list, accountFeature)
		}
	}
	g.Resources = g.createResources(list)
	return nil
}
