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

package aws

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"

	es "github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

var esAllowEmptyValues = []string{"tags."}

type EsGenerator struct {
	AWSService
}

func (g *EsGenerator) InitResources() error {
	sess := g.generateSession()
	svc := es.New(sess)

	domainNames, err := svc.ListDomainNames(&es.ListDomainNamesInput{})
	if err != nil {
		return err
	}

	for _, domainName := range domainNames.DomainNames {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			aws.StringValue(domainName.DomainName),
			aws.StringValue(domainName.DomainName),
			"aws_elasticsearch_domain",
			"aws",
			map[string]string{
				"domain_name": aws.StringValue(domainName.DomainName),
			},
			esAllowEmptyValues,
			map[string]interface{}{},
		))
	}

	g.PopulateIgnoreKeys()
	return nil
}
