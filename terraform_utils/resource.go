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
	"fmt"
	"log"
	"regexp"
	"strings"
	"waze/terraformer/terraform_utils/provider_wrapper"

	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/terraform"
)

type Resource struct {
	InstanceInfo     *terraform.InstanceInfo
	InstanceState    *terraform.InstanceState
	ResourceName     string
	Provider         string
	Item             map[string]interface{}
	IgnoreKeys       []string
	AllowEmptyValues []string
	AdditionalFields map[string]string
}

func NewResource(ID, resourceName, resourceType, provider string,
	attributes map[string]string,
	allowEmptyValues []string,
	additionalFields map[string]string) Resource {
	return Resource{
		ResourceName: TfSanitize(resourceName),
		Item:         nil,
		Provider:     provider,
		InstanceState: &terraform.InstanceState{
			ID:         ID,
			Attributes: attributes,
		},
		InstanceInfo: &terraform.InstanceInfo{
			Type: resourceType,
			Id:   fmt.Sprintf("%s.%s", resourceType, TfSanitize(resourceName)),
		},
		AdditionalFields: additionalFields,
		AllowEmptyValues: allowEmptyValues,
	}
}

func (r *Resource) Refresh(provider *provider_wrapper.ProviderWrapper) {
	var err error
	r.InstanceState, err = provider.Refresh(r.InstanceInfo, r.InstanceState)
	if err != nil {
		log.Println(err)
	}
}

func (r *Resource) ConvertTFstate() {
	r.Item = map[string]interface{}{}
	allAttributes := []string{}
	for key := range r.InstanceState.Attributes {
		allAttributes = append(allAttributes, key)
	}
	// delete empty array
	for _, key := range allAttributes {
		if strings.HasSuffix(key, ".#") && r.InstanceState.Attributes[key] == "0" {
			delete(r.InstanceState.Attributes, key)
		}
	}
	// delete ignored keys
	for keyAttribute := range r.InstanceState.Attributes {
		for _, patter := range r.IgnoreKeys {
			match, err := regexp.MatchString(patter, keyAttribute)
			if match && err == nil {
				delete(r.InstanceState.Attributes, keyAttribute)
			}
		}
	}
	// delete empty keys with empty value, but not from AllowEmptyValue list
	for keyAttribute, value := range r.InstanceState.Attributes {
		if value != "" {
			continue
		}
		allowEmptyValue := false
		for _, patter := range r.AllowEmptyValues {
			match, err := regexp.MatchString(patter, keyAttribute)
			if match && err == nil && patter != "" {
				allowEmptyValue = true
			}
		}
		if !allowEmptyValue {
			delete(r.InstanceState.Attributes, keyAttribute)
		}
	}
	// parse Attributes to go string with flatmap package
	for key := range r.InstanceState.Attributes {
		blockName := strings.Split(key, ".")[0]

		if _, exist := r.Item[blockName]; exist {
			continue
		}

		r.Item[blockName] = flatmap.Expand(r.InstanceState.Attributes, blockName)
	}
	// add Additional Fields to resource
	for key, value := range r.AdditionalFields {
		r.Item[key] = value
	}
}
