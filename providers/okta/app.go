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

	"github.com/okta/okta-sdk-golang/v2/okta"
)

//NOTE: Okta SDK v2.6.1 ListApplications() method does not support applications by type at this time. So
//		we have to create the application filter by our self.

func getApplications(ctx context.Context, client *okta.Client, signOnMode string) ([]*okta.Application, error) {
	apps, resp, err := client.Application.ListApplications(ctx, nil)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextAppSet []okta.App
		resp, err = resp.Next(ctx, &nextAppSet)
		if err != nil {
			return nil, err
		}
		apps = append(apps, nextAppSet...)
	}

	resultingApps := make([]*okta.Application, len(apps))
	for i := range apps {
		resultingApps[i] = apps[i].(*okta.Application)
	}

	var supportedApps []*okta.Application
	for _, app := range resultingApps {
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

	oktaSupportApplications := map[*okta.Application]string{}
	for _, app := range supportedApps {
		oktaSupportApplications[app] = app.SignOnMode
	}

	var filterApps []*okta.Application
	for app, appSignOnNode := range oktaSupportApplications {
		if appSignOnNode == signOnMode {
			filterApps = append(filterApps, app)
		}
	}
	return filterApps, nil
}
