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

package terraformutils

import (
	"log"
	"reflect"
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestSimpleReference(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref": "ID2",
		}, map[string]interface{}{
			"type2_ref": "ID2",
		})},
		"type2": {prepareNoAttrs("ID2", "type2")},
	}

	resourceConnections := map[string]map[string][]string{
		"type1": {
			"type2": {"type2_ref", "id"},
		},
	}
	resources := ConnectServices(importResources, true, resourceConnections)

	if !reflect.DeepEqual(resources["type1"][0].Item, map[string]interface{}{
		"type2_ref": "${data.terraform_remote_state.type2.outputs.type2_tfer--name-type2_id}",
	}) {
		t.Errorf("failed to connect %v", resources["type1"][0].Item)
	}
}

func TestManyReferences(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		}, map[string]interface{}{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		})},
		"type2": {prepareNoAttrs("ID2", "type2")},
	}

	resourceConnections := map[string]map[string][]string{
		"type1": {
			"type2": {
				"type2_ref1", "id",
				"type2_ref2", "id",
			},
		},
	}
	resources := ConnectServices(importResources, true, resourceConnections)

	if !reflect.DeepEqual(resources["type1"][0].Item, map[string]interface{}{
		"type2_ref1": "${data.terraform_remote_state.type2.outputs.type2_tfer--name-type2_id}",
		"type2_ref2": "${data.terraform_remote_state.type2.outputs.type2_tfer--name-type2_id}",
	}) {
		t.Errorf("failed to connect %v", resources["type1"][0].Item)
	}
}

func TestResourceGroups(t *testing.T) {
	importResources := map[string][]Resource{
		"group1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		}, map[string]interface{}{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		}),
			prepareNoAttrs("ID3", "type3")},
		"group2": {
			prepare("ID2", "type2", map[string]string{
				"uid": "ID2",
			}, map[string]interface{}{
				"uid": "ID2",
			}),
			prepareNoAttrs("ID4", "type4")},
	}

	resourceConnections := map[string]map[string][]string{
		"group1": {
			"group2": {
				"type2_ref1", "uid",
				"type2_ref2", "uid",
			},
		},
	}
	resources := ConnectServices(importResources, true, resourceConnections)

	if !reflect.DeepEqual(resources["group1"][0].Item, map[string]interface{}{
		"type2_ref1": "${data.terraform_remote_state.group2.outputs.type2_tfer--name-type2_uid}",
		"type2_ref2": "${data.terraform_remote_state.group2.outputs.type2_tfer--name-type2_uid}",
	}) {
		t.Errorf("failed to connect %v", resources["group1"][0].Item)
	}
}

func TestNestedReference(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"nested.type2_ref": "ID2",
		}, mapI("nested", mapI("type2_ref", "ID2")))},
		"type2": {prepareNoAttrs("ID2", "type2")},
	}

	resourceConnections := map[string]map[string][]string{
		"type1": {
			"type2": {"nested.type2_ref", "id"},
		},
	}
	resources := ConnectServices(importResources, true, resourceConnections)

	if !reflect.DeepEqual(resources["type1"][0].Item, mapI("nested", mapI("type2_ref", "${data.terraform_remote_state.type2.outputs.type2_tfer--name-type2_id}"))) {
		t.Errorf("failed to connect %v", resources)
	}
}

func prepareNoAttrs(id, resourceType string) Resource {
	return prepare(id, resourceType, map[string]string{}, map[string]interface{}{})
}

func prepare(id, resourceType string, attributes map[string]string, attributesParsed map[string]interface{}) Resource {
	r := NewResource(id, "name-"+resourceType, resourceType, "provider", attributes, []string{}, map[string]interface{}{})
	r.InstanceState.Attributes["id"] = r.InstanceState.ID
	err := r.ParseTFstate(&MockedFlatmapParser{
		attributesParsed: attributesParsed,
	}, cty.NilType)
	if err != nil {
		log.Println(err)
	}
	return r
}

func mapI(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{key: value}
}

type MockedFlatmapParser struct {
	FlatmapParser
	attributesParsed map[string]interface{}
}

func (p *MockedFlatmapParser) Parse(ty cty.Type) (map[string]interface{}, error) {
	return p.attributesParsed, nil
}
