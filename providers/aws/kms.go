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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
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
	client := kms.NewFromConfig(config)

	err := g.addKeys(client)
	if err != nil {
		return err
	}
	err = g.addAliases(client)
	return err
}

func (g *KmsGenerator) addKeys(client *kms.Client) error {
	p := kms.NewListKeysPaginator(client, &kms.ListKeysInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, key := range page.Keys {
			keyDescription, err := client.DescribeKey(context.TODO(), &kms.DescribeKeyInput{
				KeyId: key.KeyId,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			if keyDescription.KeyMetadata.KeyManager == types.KeyManagerTypeCustomer {
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

				g.addGrants(key.KeyId, client)
			}
		}
	}
	return nil
}

func (g *KmsGenerator) addAliases(client *kms.Client) error {
	p := kms.NewListAliasesPaginator(client, &kms.ListAliasesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, alias := range page.Aliases {
			if alias.TargetKeyId == nil {
				continue
			}
			keyDescription, err := client.DescribeKey(context.TODO(), &kms.DescribeKeyInput{
				KeyId: alias.TargetKeyId,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			if keyDescription.KeyMetadata.KeyManager == types.KeyManagerTypeCustomer {
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
	}
	return nil
}

func (g *KmsGenerator) addGrants(keyID *string, client *kms.Client) {
	p := kms.NewListGrantsPaginator(client, &kms.ListGrantsInput{
		KeyId: keyID,
	})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return
		}
		for _, grant := range page.Grants {
			grantID := *grant.KeyId + ":" + *grant.GrantId
			resource := terraformutils.NewSimpleResource(
				grantID,
				grantID,
				"aws_kms_grant",
				"aws",
				kmsAllowEmptyValues,
			)
			resource.SlowQueryRequired = true
			g.Resources = append(g.Resources, resource)
		}
	}
}
