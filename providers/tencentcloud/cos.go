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
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosGenerator struct {
	TencentCloudService
}

func (g *CosGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	requestURL := fmt.Sprintf("https://cos.%s.myqcloud.com", region)
	u, _ := url.Parse(requestURL)
	uri := &cos.BaseURL{ServiceURL: u}
	client := cos.NewClient(uri, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:     credential.SecretId,
			SecretKey:    credential.SecretKey,
			SessionToken: credential.Token,
		},
	})

	result, _, err := client.Service.Get(context.Background())
	if err != nil {
		return err
	}

	for _, bucket := range result.Buckets {
		resource := terraformutils.NewResource(
			bucket.Name,
			bucket.Name,
			"tencentcloud_cos_bucket",
			"tencentcloud",
			map[string]string{
				"acl": "private",
			},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *CosGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "tencentcloud_cos_bucket" {
			if _, ok := resource.Item["lifecycle_rules"]; ok {
				lifecycleRules := resource.Item["lifecycle_rules"].([]interface{})
				for i := range lifecycleRules {
					rule := lifecycleRules[i].(map[string]interface{})
					if _, ok := rule["filter_prefix"]; !ok {
						rule["filter_prefix"] = ""
						lifecycleRules[i] = rule
					}
				}
			}
		}
	}
	return nil
}
