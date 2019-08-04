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

	"github.com/jonboydell/logzio_client/endpoints"
)

type AlertNotificationEndpointsGenerator struct {
	LogzioService
}

// Generate Terraform Resources from Logzio API,
func (g *AlertNotificationEndpointsGenerator) InitResources() error {
	var client *endpoints.EndpointsClient
	client, _ = endpoints.New(g.Args["token"].(string), g.Args["baseURL"].(string))

	endpoints, err := client.ListEndpoints()
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			strconv.FormatInt(endpoint.Id, 10),
			createSlug(endpoint.Title+"-"+string(endpoint.EndpointType)+"-"+strconv.FormatInt(endpoint.Id, 10)),
			"logzio_endpoint",
			"logzio",
			map[string]string{},
			[]string{},
			map[string]string{},
		))
	}
	g.PopulateIgnoreKeys()
	return nil
}
