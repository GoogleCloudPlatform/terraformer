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

package base_terraforming

import (
	"log"
	"waze/terraformer/terraform_utils"
	"waze/terraformer/terraform_utils/provider_wrapper"
)

type Generator interface {
	Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error)
	PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error)
}

type BasicGenerator struct{}

func (BasicGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	panic("implement me")
}

func (BasicGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	return resources, nil
}

func (BasicGenerator) IgnoreKeys(resources []terraform_utils.TerraformResource, providerName string) map[string][]string {
	p, err := provider_wrapper.NewProviderWrapper(providerName)
	if err != nil {
		log.Println(err)
		return map[string][]string{}
	}
	defer p.Kill()
	resourcesTypes := []string{}
	for _, k := range resources {
		resourcesTypes = append(resourcesTypes, k.InstanceInfo.Type)
	}
	readOnlyAttributes, err := p.GetReadOnlyAttributes(resourcesTypes)
	if err != nil {
		log.Println(err)
		return map[string][]string{}
	}
	return readOnlyAttributes
}
