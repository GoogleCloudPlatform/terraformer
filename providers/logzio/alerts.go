// Copyright 2018 The Terraformer Authors.
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

package logzio

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/jonboydell/logzio_client/alerts"
)

type AlertsGenerator struct {
	LogzioService
}

// Generate Terraform Resources from Logzio API,
func (g *AlertsGenerator) InitResources() error {
	var client *alerts.AlertsClient
	client, _ = alerts.New(g.Args["token"].(string), g.Args["baseURL"].(string))

	alerts, err := client.ListAlerts()
	if err != nil {
		return err
	}
	allowedEmptyValues := []string{"alert_notification_endpoints.#", "notification_emails.#"}
	for _, alert := range alerts {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			strconv.FormatInt(alert.AlertId, 10),
			createSlug(alert.Title+"-"+strconv.FormatInt(alert.AlertId, 10)),
			"logzio_alert",
			"logzio",
			allowedEmptyValues,
		))
	}
	return nil
}
