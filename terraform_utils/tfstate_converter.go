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

package terraform_utils

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/flatmap"
)

type InstanceStateConverter struct{}

func (c InstanceStateConverter) Convert(resources []TerraformResource, metadata map[string]ResourceMetaData) ([]TerraformResource, error) {
	newResources := []TerraformResource{}
	// read full tfstate file
	//data, err := ioutil.ReadFile(pathToTfstate)
	//if err != nil {
	//		return resources, err
	//	}
	//	tfState := terraform.State{}
	//	parse json to tfstate struct from terraform code
	//	err = json.Unmarshal(data, &tfState)
	//if err != nil {
	//return resources, err
	//}
	for _, resource := range resources {
		item := map[string]interface{}{}
		allAttributes := []string{}
		for key := range resource.InstanceState.Attributes {
			allAttributes = append(allAttributes, key)
		}
		// delete empty array
		for _, key := range allAttributes {
			if strings.HasSuffix(key, ".#") && resource.InstanceState.Attributes[key] == "0" {
				delete(resource.InstanceState.Attributes, key)
			}
		}
		// delete ignored keys
		for keyAttribute := range resource.InstanceState.Attributes {
			for _, patter := range metadata[resource.InstanceState.ID].IgnoreKeys[resource.ResourceType] {
				match, err := regexp.MatchString(patter, keyAttribute)
				if match && err == nil {
					delete(resource.InstanceState.Attributes, keyAttribute)
				}
			}
		}
		// delete empty keys with empty value, but not from AllowEmptyValue list
		for keyAttribute, value := range resource.InstanceState.Attributes {
			if value != "" {
				continue
			}
			allowEmptyValue := false
			for patter := range metadata[resource.InstanceState.ID].AllowEmptyValue {
				match, err := regexp.MatchString(patter, keyAttribute)
				if match && err == nil {
					allowEmptyValue = true
				}
			}
			if !allowEmptyValue {
				delete(resource.InstanceState.Attributes, keyAttribute)
			}
		}
		// parse Attributes to go string with flatmap package
		for key := range resource.InstanceState.Attributes {
			blockName := strings.Split(key, ".")[0]

			if _, exist := item[blockName]; exist {
				continue
			}

			item[blockName] = flatmap.Expand(resource.InstanceState.Attributes, blockName)
		}
		// add Additional Fields to resource
		for key, value := range metadata[resource.InstanceState.ID].AdditionalFields {
			item[key] = value
		}
		newResources = append(newResources, TerraformResource{
			ResourceType: resource.ResourceType,
			ResourceName: resource.ResourceName,
			Item:         item,
			ID:           resource.InstanceState.ID,
			Provider:     metadata[resource.InstanceState.ID].Provider,
		})
	}
	return newResources, nil
}
