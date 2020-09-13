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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

var kmsAllowEmptyValues = []string{"tags."}

type KmsGenerator struct {
	AWSService
}

func (g *KmsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	client := kms.New(config)

	err, aliases := g.addAliases(client)
	if err != nil {
		return err
	}
	err = g.addKeys(client, aliases)
	return err
}

func (g *KmsGenerator) addKeys(client *kms.Client, aliases map[string]string) error {
	p := kms.NewListKeysPaginator(client.ListKeysRequest(&kms.ListKeysInput{}))
	for p.Next(context.Background()) {
		for _, key := range p.CurrentPage().Keys {
			if strings.HasPrefix(aliases[*key.KeyId], "alias/aws/") {
				continue
			}
			resource := terraformutils.NewResource(
				*key.KeyId,
				*key.KeyId,
				"aws_kms_key",
				"aws",
				map[string]string{
					"key_id": *key.KeyId,
				},
				kmsAllowEmptyValues,
				map[string]interface{}{},
			)
			resource.SlowQueryRequired = true
			g.Resources = append(g.Resources, resource)
		}
	}
	return p.Err()
}

func (g *KmsGenerator) addAliases(client *kms.Client) (error, map[string]string) {
	p := kms.NewListAliasesPaginator(client.ListAliasesRequest(&kms.ListAliasesInput{}))
	aliases := make(map[string]string)
	for p.Next(context.Background()) {
		for _, alias := range p.CurrentPage().Aliases {
			if strings.HasPrefix(*alias.AliasName, "alias/aws/") {
				if alias.TargetKeyId != nil {
					aliases[*alias.TargetKeyId] = *alias.AliasName
				}
				continue
			}
			resource := terraformutils.NewSimpleResource(
				*alias.AliasName,
				*alias.AliasName,
				"aws_kms_alias",
				"aws",
				kmsAllowEmptyValues,
			)
			resource.SlowQueryRequired = true
			g.Resources = append(g.Resources, resource)
		}
	}
	return p.Err(), aliases
}
