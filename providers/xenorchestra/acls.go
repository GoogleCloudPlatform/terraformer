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
package xenorchestra

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ddelnano/terraform-provider-xenorchestra/client"
)

type AclGenerator struct { //nolint
	XenorchestraService
}

func (g AclGenerator) createResources(acls []client.Acl) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, acl := range acls {
		resourceName := acl.Id
		resources = append(resources, terraformutils.NewSimpleResource(
			acl.Id,
			resourceName,
			"xenorchestra_acl",
			"xenorchestra",
			[]string{}))
	}
	return resources
}

func (g *AclGenerator) InitResources() error {
	client := g.generateClient()
	acls, err := client.GetAcls()

	if err != nil {
		return err
	}
	g.Resources = g.createResources(acls)
	return nil
}
