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
	"context"
	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	es "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

var esAllowEmptyValues = []string{"tags."}

type EsGenerator struct {
	AWSService
}

func (g *EsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := es.NewFromConfig(config)

	domainNames, err := svc.ListDomainNames(context.TODO(), &es.ListDomainNamesInput{})
	if err != nil {
		return err
	}

	for _, domainName := range domainNames.DomainNames {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			StringValue(domainName.DomainName),
			StringValue(domainName.DomainName),
			"aws_elasticsearch_domain",
			"aws",
			map[string]string{
				"domain_name": StringValue(domainName.DomainName),
			},
			esAllowEmptyValues,
			map[string]interface{}{},
		))
	}

	return nil
}

func (g *EsGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.Address.Type != "aws_elasticsearch_domain" {
			continue
		}
		instanceStateMap := r.InstanceState.Value.AsValueMap()

		if r.InstanceState.Value.GetAttr("cognito_options").AsValueSlice()[0].GetAttr("enabled").AsString() == "false" {
			delete(instanceStateMap, "cognito_options")
		}
		if r.InstanceState.Value.GetAttr("cluster_config").AsValueSlice()[0].GetAttr("warm_count").AsString() == "0" {
			clusterConfigMap := instanceStateMap["cluster_config"].AsValueSlice()[0].AsValueMap()
			delete(clusterConfigMap, "warm_count")
			instanceStateMap["cluster_config"] = terraformutils.ListToValue([]cty.Value{cty.ObjectVal(clusterConfigMap)})
		}
		r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
	}
	return nil
}
