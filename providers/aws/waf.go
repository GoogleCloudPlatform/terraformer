// Copyright 2020 The Terraformer Authors.
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

package aws

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
)

var wafAllowEmptyValues = []string{"tags."}

type WafGenerator struct {
	AWSService
}

func (g *WafGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := wafv2.New(config)

	if g.Args["region"] == "us-east-1" || g.Args["region"] == "" { // import CF ACLs as well
		output, err := svc.ListWebACLsRequest(&wafv2.ListWebACLsInput{
			Scope: wafv2.ScopeCloudfront,
		}).Send(context.Background())
		if err != nil {
			return err
		}
		g.Resources = g.createWebAclResources(output.WebACLs, g.Resources)
	}

	output, err := svc.ListWebACLsRequest(&wafv2.ListWebACLsInput{
		Scope: wafv2.ScopeRegional,
	}).Send(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createWebAclResources(output.WebACLs, g.Resources)

	return nil
}

func (g WafGenerator) createWebAclResources(acls []wafv2.WebACLSummary, resources []terraform_utils.Resource) []terraform_utils.Resource {
	newResources := resources
	for _, acl := range acls {
		newResources = append(newResources, terraform_utils.NewSimpleResource(
			*acl.Id,
			*acl.Name,
			"aws_waf_web_acl",
			"aws",
			wafAllowEmptyValues))
	}
	return newResources
}
