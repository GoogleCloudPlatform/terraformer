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
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

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

func (r *Resource) HasStateAttr(attr string) bool {
	return r.InstanceState.Value.HasIndex(cty.StringVal(attr)) == cty.True
}

func (r *Resource) GetStateAttr(attr string) string {
	if !r.HasStateAttr(attr) {
		return ""
	}
	return r.valueToString(r.InstanceState.Value.GetAttr(attr))
}

func (r *Resource) GetStateAttrSlice(attr string) []cty.Value {
	if !r.HasStateAttr(attr) {
		return []cty.Value{}
	}
	return r.InstanceState.Value.GetAttr(attr).AsValueSlice()
}

func (r *Resource) GetStateAttrMap(attr string) map[string]cty.Value {
	if !r.HasStateAttr(attr) {
		return map[string]cty.Value{}
	}
	return r.InstanceState.Value.GetAttr(attr).AsValueMap()
}

func (r *Resource) SetStateAttr(attr string, value cty.Value) {
	instanceStateMap := r.InstanceState.Value.AsValueMap()
	instanceStateMap[attr] = value
	r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
}

func (r *Resource) DeleteStateAttr(attr string) {
	instanceStateMap := r.InstanceState.Value.AsValueMap()
	delete(instanceStateMap, attr)
	r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
}

func (r *Resource) SortStateAttrStringSlice(attr string) {
	if r.HasStateAttr(attr) {
		var sortedStrings []string
		for _, v := range r.GetStateAttrSlice(attr) {
			sortedStrings = append(sortedStrings, v.AsString())
		}
		sort.Strings(sortedStrings)
		var sortedValues []cty.Value
		for _, v := range sortedStrings {
			sortedValues = append(sortedValues, cty.StringVal(v))
		}
		r.SetStateAttr(attr, cty.ListVal(sortedValues))
	}
}

func (r *Resource) HasStateAttrFirstAttr(firstAttr string, secondAttr string) bool {
	return r.HasStateAttr(firstAttr) &&
		r.InstanceState.Value.GetAttr(firstAttr).AsValueSlice()[0].HasIndex(cty.StringVal(secondAttr)) == cty.True
}

func (r *Resource) GetStateAttrFirstAttr(firstAttr string, secondAttr string) string {
	if !r.HasStateAttrFirstAttr(firstAttr, secondAttr) {
		return ""
	}
	return r.valueToString(r.InstanceState.Value.GetAttr(firstAttr).AsValueSlice()[0].GetAttr(secondAttr))
}

func (r *Resource) GetStateAttrFirstAttrMap(firstAttr string, secondAttr string) map[string]cty.Value {
	if !r.HasStateAttrFirstAttr(firstAttr, secondAttr) {
		return map[string]cty.Value{}
	}
	return r.InstanceState.Value.GetAttr(firstAttr).AsValueSlice()[0].GetAttr(secondAttr).AsValueMap()
}

func (r *Resource) DeleteStateAttrFirstAttr(firstAttr string, secondAttr string) {
	instanceStateMap := r.InstanceState.Value.AsValueMap()
	firstAttrMap := instanceStateMap[firstAttr].AsValueSlice()[0].AsValueMap()
	delete(firstAttrMap, secondAttr)
	instanceStateMap[firstAttr] = cty.ListVal([]cty.Value{cty.ObjectVal(firstAttrMap)})
	r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
}

func (r *Resource) SetStateAttrFirstAttr(firstAttr string, secondAttr string, val cty.Value) {
	instanceStateMap := r.InstanceState.Value.AsValueMap()
	firstAttrMap := instanceStateMap[firstAttr].AsValueSlice()[0].AsValueMap()
	firstAttrMap[secondAttr] = val
	instanceStateMap[firstAttr] = cty.ListVal([]cty.Value{cty.ObjectVal(firstAttrMap)})
	r.InstanceState.Value = cty.ObjectVal(instanceStateMap)
}

func (r *Resource) SortStateAttrEachAttrStringSlice(firstAttr string, secondAttr string) {
	if r.HasStateAttr(firstAttr) {
		firstAttrSlice := r.GetStateAttrSlice(firstAttr)
		for i, firstAttrSliceItem := range firstAttrSlice {
			if firstAttrSliceItem.HasIndex(cty.StringVal(secondAttr)) == cty.False {
				continue
			}
			secondAttrSlice := firstAttrSliceItem.GetAttr(secondAttr).AsValueSlice()
			var sortedSecondAttrSliceStrings []string
			for _, secondAttrSliceString := range secondAttrSlice {
				sortedSecondAttrSliceStrings = append(sortedSecondAttrSliceStrings, secondAttrSliceString.AsString())
			}
			sort.Strings(sortedSecondAttrSliceStrings)
			var sortedSecondAttrSliceValues []cty.Value
			for _, ssl := range sortedSecondAttrSliceStrings {
				sortedSecondAttrSliceValues = append(sortedSecondAttrSliceValues, cty.StringVal(ssl))
			}
			valueMap := firstAttrSliceItem.AsValueMap()
			valueMap[secondAttr] = cty.ListVal(sortedSecondAttrSliceValues)
			firstAttrSlice[i] = cty.ObjectVal(valueMap)
		}
		r.SetStateAttr(firstAttr, cty.ListVal(firstAttrSlice))
	}
}

func (r *Resource) valueToString(val cty.Value) string {
	switch val.Type() {
	case cty.String:
		return val.AsString()
	case cty.Number:
		fv := val.AsBigFloat()
		if fv.IsInt() {
			intVal, _ := fv.Int64()
			return strconv.FormatInt(intVal, 10)
		} else {
			return fmt.Sprintf("%f", fv)
		}
	default:
		return val.GoString()
	}
}
