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
	"reflect"
	"strings"
)

func WalkAndGet(path string, data interface{}) []interface{} {
	pathSegments := strings.Split(path, ".")
	return walkAndGet(pathSegments, data)
}

func WalkAndOverride(path, oldValue, newValue string, data interface{}) {
	pathSegments := strings.Split(path, ".")
	walkAndOverride(pathSegments, oldValue, newValue, data)
}

func walkAndGet(pathSegments []string, data interface{}) []interface{} {
	val := reflect.ValueOf(data)
	switch {
	case isArray(val.Interface()):
		var arrayValues []interface{}
		for i := 0; i < val.Len(); i++ {
			arrayValues = append(arrayValues, walkAndGet(pathSegments, val.Index(i).Interface())...)
		}
		return arrayValues
	case len(pathSegments) == 1:
		if val.Kind() == reflect.Map {
			for _, e := range val.MapKeys() {
				v := val.MapIndex(e)
				if e.String() == pathSegments[0] {
					switch {
					case isArray(v.Interface()):
						return v.Interface().([]interface{})
					case isStringArray(v.Interface()):
						return v.Interface().([]interface{})
					default:
						return []interface{}{v.Interface()}
					}
				}
			}
		}
		return []interface{}{}
	default:
		if val.Kind() == reflect.Map {
			for _, e := range val.MapKeys() {
				v := val.MapIndex(e)
				if e.String() == pathSegments[0] {
					return walkAndGet(pathSegments[1:], v.Interface())
				}
			}
			return []interface{}{}
		}
		return []interface{}{}
	}
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
							if oldValue == currentValue.(string) {
								valss[idx] = newValue
							}
						}
					case isStringArray(v.Interface()):
						valss := v.Interface().([]string)
						for idx, currentValue := range valss {
							if oldValue == currentValue {
								valss[idx] = newValue
							}
						}
					case oldValue == v.Interface().(string):
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
