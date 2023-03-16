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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ssm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
)

type SsmGenerator struct {
	TencentCloudService
}

func (g *SsmGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := ssm.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err := g.ListSecrets(client); err != nil {
		return err
	}

	return nil
}
func (g *SsmGenerator) ListSecrets(client *ssm.Client) error {
	request := ssm.NewListSecretsRequest()

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*ssm.SecretMetadata, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.ListSecrets(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.SecretMetadatas...)
		if len(response.Response.SecretMetadatas) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.SecretName,
			*instance.SecretName+"_"+*instance.SecretName,
			"tencentcloud_ssm_secret",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
		if err := g.ListSecretVersionIds(client, *instance.SecretName, resource.ResourceName); err != nil {
			return err
		}
	}

	return nil
}

//func (g *SsmGenerator) ListSecretVersionIds(allInstances []*ssm.SecretMetadata, resourceName string) error {
//	for _, instance := range allInstances {
//		resource := terraformutils.NewResource(
//			*instance.SecretName+"#"+"_",
//			*instance.SecretName,
//			"tencentcloud_ssm_secret_version",
//			"tencentcloud",
//			map[string]string{},
//			[]string{},
//			map[string]interface{}{},
//		)
//		resource.AdditionalFields["secret_name"] = "${tencentcloud_ssm_secret." + resourceName + ".id}"
//		g.Resources = append(g.Resources, resource)
//	}
//
//	return nil
//}

func (g *SsmGenerator) ListSecretVersionIds(client *ssm.Client, secretName, resourceName string) error {
	request := ssm.NewListSecretVersionIdsRequest()

	request.SecretName = &secretName
	var allInstances []*ssm.VersionInfo
	response, err := client.ListSecretVersionIds(request)
	if err != nil {
		return err
	}

	allInstances = response.Response.Versions
	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			secretName+"#"+*instance.VersionId,
			secretName+"_"+*instance.VersionId,
			"tencentcloud_ssm_secret_version",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		resource.AdditionalFields["secret_name"] = "${tencentcloud_ssm_secret." + resourceName + ".id}"
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
