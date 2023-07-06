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
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
)

type SesGenerator struct {
	TencentCloudService
}

func (g *SesGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := ses.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	if err = g.ListEmailIdentities(client); err != nil {
		return err
	}

	return g.ListEmailTemplates(client)
}
func (g *SesGenerator) ListEmailIdentities(client *ses.Client) error {
	request := ses.NewListEmailIdentitiesRequest()

	var allInstances []*ses.EmailIdentity
	response, err := client.ListEmailIdentities(request)
	if err != nil {
		return err
	}

	allInstances = response.Response.EmailIdentities
	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.IdentityName,
			*instance.IdentityName,
			"tencentcloud_ses_domain",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
		if err := g.ListEmailAddress(client); err != nil {
			return err
		}
	}

	return nil
}
func (g *SesGenerator) ListEmailAddress(client *ses.Client) error {
	request := ses.NewListEmailAddressRequest()

	var allInstances []*ses.EmailSender
	response, err := client.ListEmailAddress(request)
	if err != nil {
		return err
	}

	allInstances = response.Response.EmailSenders
	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.EmailAddress,
			*instance.EmailAddress,
			"tencentcloud_ses_email_address",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *SesGenerator) ListEmailTemplates(client *ses.Client) error {
	request := ses.NewListEmailTemplatesRequest()

	var offset uint64
	var limit uint64 = 50
	allInstances := make([]*ses.TemplatesMetadata, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.ListEmailTemplates(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.TemplatesMetadata...)
		if len(response.Response.TemplatesMetadata) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			strconv.FormatUint(*instance.TemplateID, 10),
			strconv.FormatUint(*instance.TemplateID, 10),
			"tencentcloud_ses_template",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
