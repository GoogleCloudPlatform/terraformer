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

package okta

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type AppUserSchemaPropertyGenerator struct {
	OktaService
}

func (g AppUserSchemaPropertyGenerator) createResources(appUserSchema *okta.UserSchema, appID string) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for index := range appUserSchema.Definitions.Custom.Properties {
		resources = append(resources, terraformutils.NewResource(
			index,
			normalizeResourceName(appID)+"_property_"+normalizeResourceName(index),
			"okta_app_user_schema_property",
			"okta",
			map[string]string{
				"app_id": appID,
				"index":  index,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	for index := range appUserSchema.Definitions.Base.Properties {
		resources = append(resources, terraformutils.NewResource(
			index,
			normalizeResourceName(appID)+"_property_"+normalizeResourceName(index),
			"okta_app_user_base_schema_property",
			"okta",
			map[string]string{
				"app_id": appID,
				"index":  index,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *AppUserSchemaPropertyGenerator) InitResources() error {
	var resources []terraformutils.Resource
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	apps, err := getAllApplications(ctx, client)
	if err != nil {
		return err
	}

	for _, app := range apps {
		appUserSchema, _, err := client.UserSchema.GetApplicationUserSchema(ctx, app.Id)
		if err != nil {
			return err
		}

		resources = append(resources, g.createResources(appUserSchema, app.Id)...)
	}
	g.Resources = resources
	return nil
}
