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

package ibm

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// InstanceTemplateGenerator ...
type InstanceTemplateGenerator struct {
	IBMService
}

func (g InstanceTemplateGenerator) createInstanceTemplateResources(templateID, templateName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		templateID,
		normalizeResourceName(templateName, false),
		"ibm_is_instance_template",
		"ibm",
		[]string{})
	return resources
}

// InitResources ...
func (g *InstanceTemplateGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
	}

	vpcurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region)
	vpcoptions := &vpcv1.VpcV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, vpcurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	options := &vpcv1.ListInstanceTemplatesOptions{}
	templates, response, err := vpcclient.ListInstanceTemplates(options)
	if err != nil {
		return fmt.Errorf("Error Fetching Instance Templates %s\n%s", err, response)
	}

	for _, template := range templates.Templates {
		instemp := template.(*vpcv1.InstanceTemplate)
		g.Resources = append(g.Resources, g.createInstanceTemplateResources(*instemp.ID, *instemp.Name))
	}
	return nil
}
