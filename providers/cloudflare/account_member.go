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

package cloudflare

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type AccountMemberGenerator struct {
	CloudflareService
}

func (g *AccountMemberGenerator) createAccountMemberResources(api *cf.API) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	pageOpt := cf.PaginationOptions{
		Page:    1,
		PerPage: 10}

	for {
		members, info, err := api.AccountMembers(api.AccountID, pageOpt)
		if err != nil {
			return resources, err
		}

		for _, member := range members {
			var roleIDs []string
			for _, role := range member.Roles {
				roleIDs = append(roleIDs, role.ID)
			}

			resources = append(resources, terraformutils.NewResource(
				member.ID,
				member.ID,
				"cloudflare_account_member",
				"cloudflare",
				map[string]string{
					"email_address": member.User.Email,
				},
				[]string{},
				map[string]interface{}{
					"role_ids": roleIDs,
				},
			))
		}

		if pageOpt.Page < info.TotalPages {
			pageOpt.Page++
		} else {
			break
		}
	}

	return resources, nil
}

func (g *AccountMemberGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}
	resources, err := g.createAccountMemberResources(api)
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, resources...)

	return nil
}
