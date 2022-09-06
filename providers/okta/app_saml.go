// Copyright 2021 The Terraformer Authors.
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
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
	"strings"
)

type AppSamlGenerator struct {
	OktaService
}

func (g AppSamlGenerator) createResourcesApp(ctx context.Context, client *okta.Client, appList []*okta.Application) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		r := terraformutils.NewResource(
			app.Id,
			normalizeResourceName(app.Id+"_"+app.Name),
			"okta_app_saml",
			"okta",
			map[string]string{},
			[]string{},
			map[string]interface{}{})
		r.IgnoreKeys = append(r.IgnoreKeys, "^groups", "^users")
		r.SlowQueryRequired = true
		groups := g.initAppGroups(ctx, client, app)
		resources = append(resources, r)
		resources = append(resources, groups...)
	}
	return resources
}

func (g AppSamlGenerator) initAppGroups(ctx context.Context, client *okta.Client, app *okta.Application) []terraformutils.Resource {
	groupsIDs, err := listApplicationGroupsIDs(ctx, client, app.Id)
	if err != nil {
		log.Println(err)
	}
	var resources []terraformutils.Resource
	for _, groupID := range groupsIDs {
		r := terraformutils.NewResource(
			app.Id,
			normalizeResourceName(app.Id+"_"+groupID),
			"okta_app_group_assignment",
			"okta",
			map[string]string{
				"group_id": groupID,
				"app_id":   app.Id,
			},
			[]string{},
			map[string]interface{}{})
		r.SlowQueryRequired = true
		resources = append(resources, r)
	}
	return resources
}

func (g *AppSamlGenerator) InitResources() error {
	signOnMode := []string{"SAML_1_1", "SAML_2_0"}
	allSamlApps := []*okta.Application{}
	ctx, client, err := g.Client()
	if err != nil {
		return err
	}
	for _, signOnMode := range signOnMode {
		apps, err := getApplications(ctx, client, signOnMode)
		if err != nil {
			return err
		}
		allSamlApps = append(allSamlApps, apps...)
	}
	g.Resources = g.createResourcesApp(ctx, client, allSamlApps)
	return nil
}

func (g *AppSamlGenerator) PostConvertHook() error {
	for i := range g.Resources {
		g.Resources[i].Item = replaceParams(g.Resources[i].Item)
	}
	return nil
}

func replaceParams(item map[string]interface{}) map[string]interface{} {
	for k, f := range item {
		switch v := f.(type) {
		case string:
			item[k] = strings.ReplaceAll(v, "${", "$${")
		case map[string]interface{}:
			item[k] = replaceParams(v)
		case []string:
			t := []string{}
			for _, s := range v {
				t = append(t, strings.ReplaceAll(s, "${", "$${"))
			}
			item[k] = t
		default:
		}
	}
	return item
}
