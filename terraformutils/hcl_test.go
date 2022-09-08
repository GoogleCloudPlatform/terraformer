// Copyright 2020 The Terraformer Authors.
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
	"strings"
	"testing"
)

func TestPrintResource(t *testing.T) {
	var resources []Resource
	var nested []map[string]interface{}
	nested = append(nested, mapI("field1", "egg"))
	importResource := prepare("ID1", "type1", map[string]string{
		"type1":                  "ID2",
		"map1.%":                 "1",
		"map1.foo":               "bar",
		"nested.#":               "1",
		"nested.0.map1.#":        "1",
		"nested.0.map1.0.field1": "egg",
		"nested2.#":              "1",
		"nested2.0.field1":       "spam",
		"nested2.0.map2.%":       "1",
		"nested2.0.map2.foo":     "bar",
	}, map[string]interface{}{
		"type1":   "ID2",
		"map1":    mapI("foo", "bar"),
		"nested":  mapI("map1", nested),
		"nested2": map[string]interface{}{"map2": mapI("bar", "foo"), "field1": "egg"},
	})
	resources = append(resources, importResource)
	providerData := map[string]interface{}{}
	output := "hcl"
	data, _ := HclPrintResource(resources, providerData, output, true)

	if strings.Count(string(data), "map1 = ") != 1 {
		t.Errorf("failed to parse data %s", string(data))
	}
	if strings.Count(string(data), "map2 = ") != 1 {
		t.Errorf("failed to parse data %s", string(data))
	}
}
