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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type AddOnAttachmentGenerator struct {
	HerokuService
}

func (g AddOnAttachmentGenerator) createResources(addOnAttachmentList []heroku.AddOnAttachment) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, addOnAttachment := range addOnAttachmentList {
		resources = append(resources, terraformutils.NewSimpleResource(
			addOnAttachment.ID,
			fmt.Sprintf("%s-%s", addOnAttachment.App.Name, addOnAttachment.Name),
			"heroku_addon_attachment",
			"heroku",
			[]string{}))
	}
	return resources
}

func (g *AddOnAttachmentGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()

	var output []heroku.AddOnAttachment
	var hasRequiredFilter bool

	if len(g.Filter) > 0 {
		for _, filter := range g.Filter {
			if filter.IsApplicable("app") {
				hasRequiredFilter = true
				for _, appID := range filter.AcceptableValues {
					appAddons, err := svc.AddOnListByApp(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
					if err != nil {
						return fmt.Errorf("Error filtering addons by app '%s': %w", appID, err)
					}
					for _, addOn := range appAddons {
						addonAttachments, err := svc.AddOnAttachmentListByAddOn(ctx, addOn.ID, &heroku.ListRange{Field: "id", Max: 1000})
						if err != nil {
							return fmt.Errorf("Error filtering addon attachments by app '%s': %w", appID, err)
						}
						for _, attachment := range addonAttachments {
							output = append(output, attachment)
						}
					}
				}
			}
		}
	}
	if !hasRequiredFilter {
		return fmt.Errorf("Heroku Addons Attachments must be filtered by app: --filter=app=<ID>")
	}

	g.Resources = g.createResources(output)
	return nil
}
