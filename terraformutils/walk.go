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
	"fmt"
	"reflect"
	"strings"
)

func WalkAndGet(path string, data interface{}) []interface{} {
	_, values := walkAndGet(path, data)
	return values
}
func WalkAndCheckField(path string, data interface{}) bool {
	hasField, _ := walkAndGet(path, data)
	return hasField
}

func WalkAndOverride(path, oldValue, newValue string, data interface{}) {
	pathSegments := strings.Split(path, ".")
	walkAndOverride(pathSegments, oldValue, newValue, data)
}

func walkAndGet(path string, data interface{}) (bool, []interface{}) {
	val := reflect.ValueOf(data)

	if data == nil {
		if path == "" {
			return true, []interface{}{}
		}
		return false, []interface{}{}
	}

	if isArray(val.Interface()) {
		var arrayValues []interface{}
		for i := 0; i < val.Len(); i++ {
			foundField, fieldValue := walkAndGet(path, val.Index(i).Interface())
			if foundField {
				arrayValues = append(arrayValues, fieldValue...)
			}
		}
		return len(arrayValues) > 0, arrayValues
	}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			pathFirstElement := strings.SplitN(path, ".", 2)
			if e.String() == pathFirstElement[0] {
				var pathReminder = ""
				if len(pathFirstElement) > 1 {
					pathReminder = pathFirstElement[1]
				}
				hasField, value := walkAndGet(pathReminder, v.Interface())
				if !hasField {
					hasField, value = walkAndGet(path, v.Interface())
				}
				return hasField, value
			} else if e.String() == path {
				return walkAndGet("", v.Interface())
			}
		}
	}

	if val.Kind() == reflect.String && path == "" {
		return true, []interface{}{val.Interface()}
	}

	return false, []interface{}{}
}

func walkAndOverride(pathSegments []string, oldValue, newValue string, data interface{}) {
	val := reflect.ValueOf(data)
	switch {
	case isArray(val.Interface()):
		for i := 0; i < val.Len(); i++ {
			arrayValue := val.Index(i).Interface()
			walkAndOverride(pathSegments, oldValue, newValue, arrayValue)
		}
	case len(pathSegments) == 1:
		if val.Kind() == reflect.Map {
			for _, e := range val.MapKeys() {
				v := val.MapIndex(e)
				if e.String() == pathSegments[0] {
					switch {
					case isArray(v.Interface()):
						valss := v.Interface().([]interface{})
						for idx, currentValue := range valss {
							curValString, ok := currentValue.(string)
							if ok && oldValue == curValString {
								valss[idx] = newValue
							}
							if !ok {
								fmt.Printf("Warning: expected string at path: %s, but found: %+v\n", e.String(), currentValue)
							}
						}
					case isStringArray(v.Interface()):
						valss := v.Interface().([]string)
						for idx, currentValue := range valss {
							if oldValue == currentValue {
								valss[idx] = newValue
							}
						}
					case oldValue == fmt.Sprint(v.Interface()):
						val.Interface().(map[string]interface{})[pathSegments[0]] = newValue
					}
				}
			}
		}
	case val.Kind() == reflect.Map:
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			if e.String() == pathSegments[0] {
				walkAndOverride(pathSegments[1:], oldValue, newValue, v.Interface())
			}
		}
	}
}

func isArray(val interface{}) bool { // Go reflect lib can't sometimes detect given value is array
	switch val.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func isStringArray(val interface{}) bool { // to support locally established arrays
	switch val.(type) {
	case []string:
		return true
	default:
		return false
	}
}
