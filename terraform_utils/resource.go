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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"

	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/terraform"
)

type Resource struct {
	InstanceInfo     *terraform.InstanceInfo
	InstanceState    *terraform.InstanceState
	Outputs          map[string]*terraform.OutputState `json:",omitempty"`
	ResourceName     string
	Provider         string
	Item             map[string]interface{} `json:",omitempty"`
	IgnoreKeys       []string               `json:",omitempty"`
	AllowEmptyValues []string               `json:",omitempty"`
	AdditionalFields map[string]interface{} `json:",omitempty"`
}

func NewResource(ID, resourceName, resourceType, provider string,
	attributes map[string]string,
	allowEmptyValues []string,
	additionalFields map[string]interface{}) Resource {
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

func NewSimpleResource(ID, resourceName, resourceType, provider string, allowEmptyValues []string) Resource {
	return NewResource(
		ID,
		resourceName,
		resourceType,
		provider,
		map[string]string{},
		allowEmptyValues,
		map[string]interface{}{},
	)
}

func (r *Resource) Refresh(provider *provider_wrapper.ProviderWrapper) {
	var err error
	r.InstanceState, err = provider.Refresh(r.InstanceInfo, r.InstanceState)
	if err != nil {
		log.Println(err)
	}
}

func (r Resource) GetIDKey() string {
	if _, exist := r.InstanceState.Attributes["self_link"]; exist {
		return "self_link"
	}
	return "id"
}

func (r *Resource) ConvertTFstate() {
	r.Item = map[string]interface{}{}
	attributes := map[string]string{}
	for k, v := range r.InstanceState.Attributes {
		attributes[k] = v
	}

	// TODO: Delete optional numeric zero values

	// delete empty array
	for key := range r.InstanceState.Attributes {
		if strings.HasSuffix(key, ".#") && r.InstanceState.Attributes[key] == "0" {
			if !r.isAllowedEmptyValue(key) {
				delete(attributes, key)
			}
		}
	}
	// delete ignored keys
	for keyAttribute := range r.InstanceState.Attributes {
		for _, pattern := range r.IgnoreKeys {
			match, err := regexp.MatchString(pattern, keyAttribute)
			if match && err == nil {
				delete(attributes, keyAttribute)
			}
		}
	}
	// delete empty keys with empty value, but not from AllowEmptyValue list
	for keyAttribute, value := range r.InstanceState.Attributes {
		if value != "" {
			continue
		}

		if !r.isAllowedEmptyValue(keyAttribute) {
			delete(attributes, keyAttribute)
		}
	}
	// parse Attributes to go string with flatmap package
	for key := range attributes {
		blockName := strings.Split(key, ".")[0]

		if _, exist := r.Item[blockName]; exist {
			continue
		}

		r.Item[blockName] = flatmap.Expand(attributes, blockName)
	}
	// add Additional Fields to resource
	for key, value := range r.AdditionalFields {
		r.Item[key] = value
	}
}

// isAllowedEmptyValue checks if a key is an allowed empty value with regular expression
func (r *Resource) isAllowedEmptyValue(key string) bool {
	for _, pattern := range r.AllowEmptyValues {
		match, err := regexp.MatchString(pattern, key)
		if match && err == nil && pattern != "" {
			return true
		}
	}
	return false
}
