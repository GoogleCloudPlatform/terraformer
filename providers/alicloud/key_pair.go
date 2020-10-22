// Copyright 2018 The Terraformer Authors.
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

package alicloud

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// KeyPairGenerator Struct for generating AliCloud Key pair
type KeyPairGenerator struct {
	AliCloudService
}

func resourceFromKeyPair(keyPair ecs.KeyPair) terraformutils.Resource {
	return terraformutils.NewResource(
		keyPair.KeyPairName, // nolint
		keyPair.KeyPairName+"__"+keyPair.KeyPairName, // nolint
		"alicloud_key_pair",
		"alicloud",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)
}

// InitResources Gets the list of all key pair ids and generates resources
func (g *KeyPairGenerator) InitResources() error {
	client, err := g.LoadClientFromProfile()
	if err != nil {
		return err
	}
	remaining := 1
	pageNumber := 1
	pageSize := 10

	allKeyPairs := make([]ecs.KeyPair, 0)

	for remaining > 0 {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			request := ecs.CreateDescribeKeyPairsRequest()
			request.RegionId = client.RegionID
			request.PageSize = requests.NewInteger(pageSize)
			request.PageNumber = requests.NewInteger(pageNumber)
			return ecsClient.DescribeKeyPairs(request)
		})
		if err != nil {
			return err
		}

		response := raw.(*ecs.DescribeKeyPairsResponse)
		allKeyPairs = append(allKeyPairs, response.KeyPairs.KeyPair...)
		remaining = response.TotalCount - pageNumber*pageSize
		pageNumber++
	}

	for _, keypair := range allKeyPairs {
		resource := resourceFromKeyPair(keypair)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
