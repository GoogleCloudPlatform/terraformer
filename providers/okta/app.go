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
	"log"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

//NOTE: Okta SDK v2.6.1 ListApplications() method does not support applications by type at this time. So
//		we have to create the application filter by our self.
func getApplications(ctx context.Context, client *okta.Client, signOnMode string) ([]*okta.Application, error) {
	supportedApps, err := getAllApplications(ctx, client)
	if err != nil {
		return nil, err
	}

	var filterApps []*okta.Application
	for _, app := range supportedApps {
		if app.SignOnMode == signOnMode {
			filterApps = append(filterApps, app)
		}
	}
	return filterApps, nil
}

func getAllApplications(ctx context.Context, client *okta.Client) ([]*okta.Application, error) {
	var apps []*okta.Application
	data, resp, err := client.Application.ListApplications(ctx, nil)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextAppSet []*okta.Application
		resp, err = resp.Next(ctx, &nextAppSet)
		if err != nil {
			log.Println("fff")
			return nil, err
		}
		apps = append(apps, nextAppSet...)
	}
	for _, a := range data {
		apps = append(apps, a.(*okta.Application))
	}

	var supportedApps []*okta.Application
	for _, app := range apps {
		//NOTE: Okta provider does not support the following app type/name
		if app.Name == "template_wsfed" ||
			app.Name == "template_swa_two_page" ||
			app.Name == "okta_enduser" ||
			app.Name == "okta_browser_plugin" ||
			app.Name == "saasure" {
			continue
		}
		supportedApps = append(supportedApps, app)
	}

	return supportedApps, nil
}

func listApplicationGroupsIDs(ctx context.Context, client *okta.Client, id string) ([]string, error) {
	var groupIDs []string
	groups, resp, err := client.Application.ListApplicationGroupAssignments(ctx, id, &query.Params{})
	if err != nil {
		return nil, err
	}
	for {
		for _, groupID := range groups {
			groupIDs = append(groupIDs, groupID.Id)
		}
		if resp.HasNextPage() {
			resp, err = resp.Next(ctx, &groups)
			if err != nil {
				return nil, err
			}
			continue
		} else {
			break
		}
	}
	return groupIDs, nil
}
