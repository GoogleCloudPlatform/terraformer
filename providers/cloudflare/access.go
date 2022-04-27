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

package cloudflare

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type AccessGenerator struct {
	CloudflareService
}

func (g *AccessGenerator) createAccessIdentityProviders(ctx context.Context, api *cf.API, accountID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	idps, err := api.AccessIdentityProviders(ctx, accountID)
	if err != nil {
		return []terraformutils.Resource{}, err
	}

	for _, idp := range idps {
		resources = append(resources, terraformutils.NewResource(
			idp.ID,
			idp.Name,
			"cloudflare_access_identity_provider",
			"cloudflare",
			map[string]string{
				"account_id": accountID,
				"name":       idp.Name,
			},
			[]string{},
			map[string]interface{}{
				"config": idp.Config,
			},
		))
	}

	return resources, nil
}

func (g *AccessGenerator) createAccessGroups(ctx context.Context, api *cf.API, accountID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	groups, _, err := api.AccessGroups(ctx, accountID, cf.PaginationOptions{})
	if err != nil {
		return []terraformutils.Resource{}, err
	}

	for _, group := range groups {
		resources = append(resources, terraformutils.NewResource(
			group.ID,
			group.Name,
			"cloudflare_access_group",
			"cloudflare",
			map[string]string{
				"account_id": accountID,
				"name":       group.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources, nil
}

func (g *AccessGenerator) createAccessApplicationsAndPolicies(ctx context.Context, api *cf.API, zoneID, accountID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	accessApplications, _, err := api.ZoneLevelAccessApplications(ctx, zoneID, cf.PaginationOptions{})
	if err != nil {
		return []terraformutils.Resource{}, err
	}

	for _, app := range accessApplications {
		additionalFields := map[string]interface{}{
			"same_site_cookie_attribute": "none",
		}
		if len(app.AllowedIdps) > 0 {
			additionalFields["allowed_idps"] = app.AllowedIdps
		}
		if app.SameSiteCookieAttribute != "" {
			additionalFields["same_site_cookie_attribute"] = app.SameSiteCookieAttribute
		}
		if app.CorsHeaders != nil {
			additionalFields["cors_headers"] = app.CorsHeaders
		}
		resources = append(resources, terraformutils.NewResource(
			app.ID,
			app.Name,
			"cloudflare_access_application",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
				"name":    app.Name,
			},
			[]string{},
			additionalFields,
		))
		accessPolicies, _, err := api.ZoneLevelAccessPolicies(ctx, zoneID, app.ID, cf.PaginationOptions{})
		if err != nil {
			return []terraformutils.Resource{}, err
		}
		for _, policy := range accessPolicies {
			resources = append(resources, terraformutils.NewResource(
				policy.ID,
				fmt.Sprintf("%s_%s", app.Name, policy.Name),
				"cloudflare_access_policy",
				"cloudflare",
				map[string]string{
					"application_id": app.ID,
					"zone_id":        zoneID,
					"name":           policy.Name,
				},
				[]string{},
				map[string]interface{}{},
			))

		}
	}

	return resources, nil
}

func (g *AccessGenerator) InitResources() error {
	ctx := context.Background()
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	tmpRes, err := g.createAccessIdentityProviders(ctx, api, g.getAccountID())
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, tmpRes...)

	tmpRes, err = g.createAccessGroups(ctx, api, g.getAccountID())
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, tmpRes...)

	zones, err := api.ListZones(ctx)
	if err != nil {
		return err
	}

	for _, zone := range zones {
		tmpRes, err := g.createAccessApplicationsAndPolicies(ctx, api, zone.ID, g.getAccountID())
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, tmpRes...)
	}

	return nil
}

func (g *AccessGenerator) getAccessIdentityProviderReferenceById(idps map[string]terraformutils.Resource, id string) string {
	if idpResource, ok := idps[id]; ok {
		return "${cloudflare_access_identity_provider." + idpResource.ResourceName + ".id}"
	}
	return id
}

func (g *AccessGenerator) getAccessGroupReferenceById(groups map[string]terraformutils.Resource, id string) string {
	if groupResource, ok := groups[id]; ok {
		return "${cloudflare_access_group." + groupResource.ResourceName + ".id}"
	}
	return id
}

func (g *AccessGenerator) replaceAccessIdentityProviderReferenceForAccessGroupConditions(idps, groups map[string]terraformutils.Resource, conditions []interface{}) []interface{} {

	// fmt.Printf("conditions: %+v\n", conditions)
	for _, condition := range conditions {
		conditionConfig, _ := condition.(map[string]interface{})
		// fmt.Printf("conditionConfig: %+v\n", conditionConfig)
		for configKey, configValue := range conditionConfig {
			// fmt.Printf("configKey: %+v configValue: %+v\n", configKey, configValue)
			if configKey == "group" {
				if groupConfig, ok := configValue.([]interface{}); ok {
					newGroupConfig := []string{}
					for _, id := range groupConfig {
						if groupID, ok := id.(string); ok {
							newGroupConfig = append(newGroupConfig, g.getAccessGroupReferenceById(groups, groupID))
						}
					}
					conditionConfig["group"] = newGroupConfig
				}
			} else if authConfigs, ok := configValue.([]interface{}); ok {
				for _, authConfig := range authConfigs {
					// fmt.Printf("authConfig: %+v\n", authConfig)
					if authConfigMap, ok := authConfig.(map[string]interface{}); ok {

						for authKey, authValue := range authConfigMap {
							// fmt.Printf("authValue: %+v\n", authValue)
							if authKey == "identity_provider_id" {
								idpID, _ := authValue.(string)
								authConfigMap[authKey] = g.getAccessIdentityProviderReferenceById(idps, idpID)
							}
						}
					}
				}
			}
		}
	}
	return conditions
}

func (g *AccessGenerator) PostConvertHook() error {
	idps := map[string]terraformutils.Resource{}
	groups := map[string]terraformutils.Resource{}
	apps := map[string]terraformutils.Resource{}
	for _, resourceRecord := range g.Resources {
		switch resourceRecord.InstanceInfo.Type {
		case "cloudflare_access_identity_provider":
			idps[resourceRecord.InstanceState.ID] = resourceRecord
		case "cloudflare_access_group":
			groups[resourceRecord.InstanceState.ID] = resourceRecord
		case "cloudflare_access_application":
			apps[resourceRecord.InstanceState.ID] = resourceRecord
		}
	}

	for i, resourceRecord := range g.Resources {

		// Reference to 'cloudflare_access_identity_provider' resource in 'cloudflare_access_group'
		if resourceRecord.InstanceInfo.Type == "cloudflare_access_group" || resourceRecord.InstanceInfo.Type == "cloudflare_access_policy" {
			// fmt.Printf("Before: %+v\n", g.Resources[i].Item)
			if requireConditions, ok := resourceRecord.Item["require"].([]interface{}); ok {
				// fmt.Printf("requireConditions: %T %v %+v\n", resourceRecord.Item["require"], ok, requireConditions)
				g.Resources[i].Item["require"] = g.replaceAccessIdentityProviderReferenceForAccessGroupConditions(idps, groups, requireConditions)
			}
			if excludeConditions, ok := resourceRecord.Item["exclude"].([]interface{}); ok {
				// fmt.Printf("excludeConditions: %T %v %+v\n", resourceRecord.Item["exclude"], ok, excludeConditions)
				g.Resources[i].Item["exclude"] = g.replaceAccessIdentityProviderReferenceForAccessGroupConditions(idps, groups, excludeConditions)
			}
			if includeConditions, ok := resourceRecord.Item["include"].([]interface{}); ok {
				// fmt.Printf("includeConditions: %T %v %+v\n", resourceRecord.Item["include"], ok, includeConditions)
				g.Resources[i].Item["include"] = g.replaceAccessIdentityProviderReferenceForAccessGroupConditions(idps, groups, includeConditions)
			}

			// fmt.Printf("After: %+v\n", g.Resources[i].Item)

		}
		if resourceRecord.InstanceInfo.Type == "cloudflare_access_policy" {
			if appID, ok := resourceRecord.Item["application_id"].(string); ok {
				g.Resources[i].Item["application_id"] = "${cloudflare_access_application." + apps[appID].ResourceName + ".id}"
			}
		}
		if resourceRecord.InstanceInfo.Type == "cloudflare_access_application" {

			// fmt.Printf("Before: %+v\n", g.Resources[i].Item)
			if allowedIdps, ok := resourceRecord.Item["allowed_idps"].([]string); ok {
				newAllowedIdps := []string{}
				for _, allowedIdp := range allowedIdps {
					newAllowedIdps = append(newAllowedIdps, g.getAccessIdentityProviderReferenceById(idps, allowedIdp))
				}
				g.Resources[i].Item["allowed_idps"] = newAllowedIdps
			}
			// fmt.Printf("After: %+v\n", g.Resources[i].Item)
		}
	}

	return nil
}
