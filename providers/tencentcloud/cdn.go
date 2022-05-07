// Copyright 2022 The Terraformer Authors.
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

package tencentcloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type CdnGenerator struct {
	TencentCloudService
}

func (g *CdnGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := cdn.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := cdn.NewDescribeDomainsConfigRequest()

	var offset int64
	var pageSize int64 = 50
	allInstances := make([]*cdn.DetailDomain, 0)

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		response, err := client.DescribeDomainsConfig(request)
		if err != nil {
			return err
		}

		allInstances = append(allInstances, response.Response.Domains...)
		if len(response.Response.Domains) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.Domain,
			*instance.Domain,
			"tencentcloud_cdn_domain",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *CdnGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_cdn_domain" {
			httpsConfigs := resource.Item["https_config"].([]interface{})
			if len(httpsConfigs) > 0 {
				config := httpsConfigs[0].(map[string]interface{})
				if config["https_switch"] == "on" &&
					resource.InstanceState.Attributes["https_config.0.server_certificate_config.#"] == "1" {
					serverCert := map[string]interface{}{
						"certificate_content": "",
						"private_key":         "",
					}
					serverCerts := make([]interface{}, 0, 1)
					serverCerts = append(serverCerts, serverCert)
					config["server_certificate_config"] = serverCerts
				}
				if config["verify_client"] == "on" {
					clientCert := map[string]interface{}{
						"certificate_content": "",
					}
					clientCerts := make([]interface{}, 0, 1)
					clientCerts = append(clientCerts, clientCert)
					config["client_certificate_config"] = clientCerts
				}
				httpsConfigs[0] = config
			}
			resource.Item["https_config"] = httpsConfigs
		}
	}
	return nil
}
