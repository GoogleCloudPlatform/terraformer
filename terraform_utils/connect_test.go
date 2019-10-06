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

package terraform_utils

import (
	"reflect"
	"testing"
)

func TestSimpleReference(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref": "ID2",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}

	resourceConnections :=  map[string]map[string][]string{
		"type1": {
			"type2": {"type2_ref", "id"},
		},
	}
	resources := ConnectServices(importResources, resourceConnections)

	if !reflect.DeepEqual(resources, map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}) {
		t.Errorf("failed to connect %v", resources)
	}
}

func TestManyReferences(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}

	resourceConnections := map[string]map[string][]string{
		"type1": {
			"type2": {
				"type2_ref1", "id",
				"type2_ref2", "id",
			},
		},
	}
	resources := ConnectServices(importResources, resourceConnections)

	if !reflect.DeepEqual(resources, map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
			"type2_ref2": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}) {
		t.Errorf("failed to connect %v", resources)
	}
}

func TestResourceGroups(t *testing.T) {
	importResources := map[string][]Resource{
		"group1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "ID2",
			"type2_ref2": "ID2",
		}),
			prepare("ID3", "type3", map[string]string{})},
		"group2": {
			prepare("ID2", "type2", map[string]string{
				"uid": "ID2",
			}),
			prepare("ID4", "type4", map[string]string{})},
	}

	resourceConnections := map[string]map[string][]string{
		"group1": {
			"group2": {
				"type2_ref1", "uid",
				"type2_ref2", "uid",
			},
		},
	}
	resources := ConnectServices(importResources, resourceConnections)

	if !reflect.DeepEqual(resources, map[string][]Resource{
		"group1": {prepare("ID1", "type1", map[string]string{
			"type2_ref1": "${data.terraform_remote_state.group2.outputs.type2_name-type2_uid}",
			"type2_ref2": "${data.terraform_remote_state.group2.outputs.type2_name-type2_uid}",
		}),
			prepare("ID3", "type3", map[string]string{}),
		},
		"group2": {
			prepare("ID2", "type2", map[string]string{
				"uid": "ID2",
			}),
			prepare("ID4", "type4", map[string]string{})},
	}) {
		t.Errorf("failed to connect %v", resources)
	}
}

func TestNestedReference(t *testing.T) {
	importResources := map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"nested.type2_ref": "ID2",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}

	resourceConnections := map[string]map[string][]string{
		"type1": {
			"type2": {"nested.type2_ref", "id"},
		},
	}
	resources := ConnectServices(importResources, resourceConnections)

	if !reflect.DeepEqual(resources, map[string][]Resource{
		"type1": {prepare("ID1", "type1", map[string]string{
			"nested.type2_ref": "${data.terraform_remote_state.type2.outputs.type2_name-type2_id}",
		})},
		"type2": {prepare("ID2", "type2", map[string]string{})},
	}) {
		t.Errorf("failed to connect %v", resources)
	}
}

func prepare(ID, resourceType string, attributes map[string]string) Resource {
	r := NewResource(ID, "name-"+resourceType, resourceType, "provider", attributes, []string{}, map[string]interface{}{})
	r.InstanceState.Attributes["id"] = r.InstanceState.ID
	r.ConvertTFstate()
	return r
}
