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

package terraformutils

import (
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

type Resource struct {
	Address           addrs.Resource
	InstanceState     *states.ResourceInstanceObject // the resource will always contain one instance as terraformer blocks don't use "count" or "for_each"
	Outputs           map[string]*states.OutputValue
	ImportID          string // identifier to be used by terraformer when importing a resource
	Provider          string
	PriorState        map[string]string      // used when refreshing a resource
	IgnoreKeys        []string               `json:",omitempty"`
	AllowEmptyValues  []string               `json:",omitempty"`
	AdditionalFields  map[string]interface{} `json:",omitempty"`
	SlowQueryRequired bool
	DataFiles         map[string][]byte
}

type ApplicableFilter interface {
	IsApplicable(resourceName string) bool
}

type ResourceFilter struct {
	ApplicableFilter
	ServiceName      string
	FieldPath        string
	AcceptableValues []string
}

func (rf *ResourceFilter) Filter(resource Resource) bool {
	if !rf.IsApplicable(strings.TrimPrefix(resource.Address.Type, resource.Provider+"_")) {
		return true
	}
	var vals []interface{}
	switch {
	case rf.FieldPath == "id":
		vals = []interface{}{resource.ImportID}
	case rf.AcceptableValues == nil:
		var dst interface{}
		err := gocty.FromCtyValue(resource.InstanceState.Value, &dst)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return WalkAndCheckField(rf.FieldPath, dst)
	default:
		var dst interface{}
		err := gocty.FromCtyValue(resource.InstanceState.Value, &dst)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		vals = WalkAndGet(rf.FieldPath, dst)
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

func (rf *ResourceFilter) IsApplicable(serviceName string) bool {
	return rf.ServiceName == "" || rf.ServiceName == serviceName
}

func (rf *ResourceFilter) isInitial() bool {
	return rf.FieldPath == "id"
}

func NewResource(id, resourceName, resourceType, provider string,
	attributes map[string]string,
	allowEmptyValues []string,
	additionalFields map[string]interface{}) Resource {
	attributes["id"] = id // to ensure resource refresh will work well
	return Resource{
		Address: addrs.Resource{
			Mode: addrs.ManagedResourceMode,
			Type: resourceType,
			Name: TfSanitize(resourceName),
		},
		ImportID:         id,
		Provider:         provider,
		PriorState:       attributes,
		AdditionalFields: additionalFields,
		AllowEmptyValues: allowEmptyValues,
	}
}

func NewSimpleResource(id, resourceName, resourceType, provider string, allowEmptyValues []string) Resource {
	return NewResource(
		id,
		resourceName,
		resourceType,
		provider,
		map[string]string{},
		allowEmptyValues,
		map[string]interface{}{},
	)
}

func (r *Resource) Refresh(provider *providerwrapper.ProviderWrapper) {
	var err error
	if r.SlowQueryRequired {
		time.Sleep(200 * time.Millisecond)
	}
	r.InstanceState, err = provider.Refresh(&r.Address, r.PriorState, r.ImportID)
	if err != nil {
		log.Println(err)
	}
}

func (r Resource) GetIDKey() string {
	if _, exist := r.PriorState["self_link"]; exist {
		return "self_link"
	}
	return "id"
}

func (r *Resource) ServiceName() string {
	return strings.TrimPrefix(r.Address.Type, r.Provider+"_")
}
