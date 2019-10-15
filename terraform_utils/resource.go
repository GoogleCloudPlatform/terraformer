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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
	"github.com/hashicorp/terraform/terraform"
	"github.com/zclconf/go-cty/cty"
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

type ApplicableFilter interface {
	IsApplicable(resourceName string) bool
}

type ResourceFilter struct {
	ApplicableFilter
	ResourceName     string
	FieldPath        string
	AcceptableValues []string
}

func (rf *ResourceFilter) Filter(resource Resource) bool {
	if !rf.IsApplicable(resource.InstanceInfo.Type) {
		return true
	}
	var vals []interface{}
	if rf.FieldPath == "id" {
		vals = []interface{}{resource.InstanceState.ID}
	} else {
		vals = WalkAndGet(rf.FieldPath, resource.InstanceState.Attributes)
		if len(vals) == 0 {
			vals = WalkAndGet(rf.FieldPath, resource.Item)
		}
	}
	for _, val := range vals {
		for _, acceptableValue := range rf.AcceptableValues {
			if val == acceptableValue {
				return true
			}
		}
	}
	return false
}

func (rf *ResourceFilter) IsApplicable(resourceName string) bool {
	return rf.ResourceName == "" || rf.ResourceName == resourceName
}

func (rf *ResourceFilter) isInitial() bool {
	return rf.FieldPath == "id"
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

func (r *Resource) ParseTFstate(parser Flatmapper, impliedType cty.Type) error {
	attributes, err := parser.Parse(impliedType)
	if err != nil {
		return err
	}

	// add Additional Fields to resource
	for key, value := range r.AdditionalFields {
		attributes[key] = value
	}

	r.Item = attributes
	return nil
}

func (r *Resource) ConvertTFstate(provider *provider_wrapper.ProviderWrapper) error {
	ignoreKeys := []*regexp.Regexp{}
	for _, pattern := range r.IgnoreKeys {
		ignoreKeys = append(ignoreKeys, regexp.MustCompile(pattern))
	}
	allowEmptyValues := []*regexp.Regexp{}
	for _, pattern := range r.AllowEmptyValues {
		if pattern != "" {
			allowEmptyValues = append(allowEmptyValues, regexp.MustCompile(pattern))
		}
	}
	parser := NewFlatmapParser(r.InstanceState.Attributes, ignoreKeys, allowEmptyValues)
	schema := provider.Provider.GetSchema()
	impliedType := schema.ResourceTypes[r.InstanceInfo.Type].Block.ImpliedType()
	return r.ParseTFstate(parser, impliedType)
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
