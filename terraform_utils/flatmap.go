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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/configs/hcl2shim"
	"github.com/zclconf/go-cty/cty"
)

type Flatmapper interface {
	Parse(ty cty.Type) (map[string]interface{}, error)
}

type FlatmapParser struct {
	Flatmapper
	attributes       map[string]string
	ignoreKeys       []*regexp.Regexp
	allowEmptyValues []*regexp.Regexp
}

func NewFlatmapParser(attributes map[string]string, ignoreKeys []*regexp.Regexp, allowEmptyValues []*regexp.Regexp) *FlatmapParser {
	return &FlatmapParser{
		attributes:       attributes,
		ignoreKeys:       ignoreKeys,
		allowEmptyValues: allowEmptyValues,
	}
}

// FromFlatmap converts a map compatible with what would be produced
// by the "flatmap" package to a map[string]interface{} object type.
//
// The intended result type must be provided in order to guide how the
// map contents are decoded. This must be an object type or this function
// will panic.
//
// Flatmap values can only represent maps when they are of primitive types,
// so the given type must not have any maps of complex types or the result
// is undefined.
//
// The result may contain null values if the given map does not contain keys
// for all of the different key paths implied by the given type.
func (p *FlatmapParser) Parse(ty cty.Type) (map[string]interface{}, error) {
	if p.attributes == nil {
		return nil, nil
	}
	if !ty.IsObjectType() {
		return nil, fmt.Errorf("FlatmapParser#Parse called on %#v", ty)
	}
	return p.fromFlatmapObject("", ty.AttributeTypes())
}

func (p *FlatmapParser) fromFlatmapValue(key string, ty cty.Type) (interface{}, error) {
	switch {
	case ty.IsPrimitiveType():
		return p.fromFlatmapPrimitive(key)
	case ty.IsObjectType():
		return p.fromFlatmapObject(key+".", ty.AttributeTypes())
	case ty.IsTupleType():
		return p.fromFlatmapTuple(key+".", ty.TupleElementTypes())
	case ty.IsMapType():
		return p.fromFlatmapMap(key+".", ty.ElementType())
	case ty.IsListType():
		return p.fromFlatmapList(key+".", ty.ElementType())
	case ty.IsSetType():
		return p.fromFlatmapSet(key+".", ty.ElementType())
	default:
		return nil, fmt.Errorf("cannot decode %s from flatmap", ty.FriendlyName())
	}
}

func (p *FlatmapParser) fromFlatmapPrimitive(key string) (interface{}, error) {
	value, ok := p.attributes[key]
	if !ok {
		return nil, nil
	}
	return value, nil
}

func (p *FlatmapParser) fromFlatmapObject(prefix string, tys map[string]cty.Type) (map[string]interface{}, error) {
	values := make(map[string]interface{})
	for name, ty := range tys {
		inAttributes := false
		attributeName := ""
		for k := range p.attributes {
			if k == prefix+name {
				attributeName = k
				inAttributes = true
				break
			}
			if k == name {
				attributeName = k
				inAttributes = true
				break
			}

			if strings.HasPrefix(k, prefix+name+".") {
				attributeName = k
				inAttributes = true
				break
			}
			lastAttribute := (prefix + name)[len(prefix):]
			if lastAttribute == k {
				attributeName = k
				inAttributes = true
				break
			}
		}

		if _, exist := p.attributes[prefix+name+".#"]; exist {
			attributeName = prefix + name + ".#"
			inAttributes = true
		}

		if _, exist := p.attributes[prefix+name+".%"]; exist {
			attributeName = prefix + name + ".%"
			inAttributes = true
		}

		if !inAttributes {
			continue
		}
		if p.isAttributeIgnored(attributeName) {
			continue
		}
		value, err := p.fromFlatmapValue(prefix+name, ty)
		if err != nil {
			return nil, err
		}
		if p.isValueAllowed(value, attributeName) {
			values[name] = value
		}
	}
	if len(values) == 0 {
		return nil, nil
	}
	return values, nil
}

func (p *FlatmapParser) fromFlatmapTuple(prefix string, tys []cty.Type) ([]interface{}, error) {
	// if the container is unknown, there is no count string
	listName := strings.TrimRight(prefix, ".")
	if p.attributes[listName] == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	countStr, exists := p.attributes[prefix+"#"]
	if !exists {
		return nil, nil
	}
	if countStr == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, fmt.Errorf("invalid count value for %q in state: %s", prefix, err)
	}
	if count != len(tys) {
		return nil, fmt.Errorf("wrong number of values for %q in state: got %d, but need %d", prefix, count, len(tys))
	}

	var values []interface{}
	for i, ty := range tys {
		key := prefix + strconv.Itoa(i)
		value, err := p.fromFlatmapValue(key, ty)
		if err != nil {
			return nil, err
		}
		if p.isValueAllowed(value, prefix) {
			values = append(values, value)
		}
	}
	if len(values) == 0 {
		return nil, nil
	}
	return values, nil
}

func (p *FlatmapParser) fromFlatmapMap(prefix string, ty cty.Type) (map[string]interface{}, error) {
	// if the container is unknown, there is no count string
	listName := strings.TrimRight(prefix, ".")
	if p.attributes[listName] == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	// We actually don't really care about the "count" of a map for our
	// purposes here, but we do need to check if it _exists_ in order to
	// recognize the difference between null (not set at all) and empty.
	strCount, exists := p.attributes[prefix+"%"]
	if !exists {
		return nil, nil
	}
	if strCount == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	values := make(map[string]interface{})
	for fullKey := range p.attributes {
		if !strings.HasPrefix(fullKey, prefix) {
			continue
		}

		// The flatmap format doesn't allow us to distinguish between keys
		// that contain periods and nested objects, so by convention a
		// map is only ever of primitive type in flatmap, and we just assume
		// that the remainder of the raw key (dots and all) is the key we
		// want in the result value.
		key := fullKey[len(prefix):]
		if key == "%" {
			// Ignore the "count" key
			continue
		}

		if p.isAttributeIgnored(fullKey) {
			continue
		}

		value, err := p.fromFlatmapValue(fullKey, ty)
		if err != nil {
			return nil, err
		}
		if p.isValueAllowed(value, prefix) {
			values[key] = value
		}
	}
	if len(values) == 0 {
		return nil, nil
	}
	return values, nil
}

func (p *FlatmapParser) fromFlatmapList(prefix string, ty cty.Type) ([]interface{}, error) {
	// if the container is unknown, there is no count string
	listName := strings.TrimRight(prefix, ".")
	if p.attributes[listName] == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	countStr, exists := p.attributes[prefix+"#"]
	if !exists {
		return nil, nil
	}
	if countStr == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, fmt.Errorf("invalid count value for %q in state: %s", prefix, err)
	}

	if count == 0 {
		return nil, nil
	}

	var values []interface{}
	for i := 0; i < count; i++ {
		key := prefix + strconv.Itoa(i)
		value, err := p.fromFlatmapValue(key, ty)
		if err != nil {
			return nil, err
		}
		if p.isValueAllowed(value, prefix) {
			values = append(values, value)
		}
	}
	return values, nil
}

func (p *FlatmapParser) fromFlatmapSet(prefix string, ty cty.Type) ([]interface{}, error) {
	// if the container is unknown, there is no count string
	listName := strings.TrimRight(prefix, ".")
	if p.attributes[listName] == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	strCount, exists := p.attributes[prefix+"#"]
	if !exists {
		return nil, nil
	}
	if strCount == hcl2shim.UnknownVariableValue {
		return nil, nil
	}

	// Keep track of keys we've seen, se we don't add the same set value
	// multiple times. The cty.Set will normally de-duplicate values, but we may
	// have unknown values that would not show as equivalent.
	seen := map[string]bool{}

	var values []interface{}
	for fullKey := range p.attributes {
		if !strings.HasPrefix(fullKey, prefix) {
			continue
		}

		subKey := fullKey[len(prefix):]
		if subKey == "#" {
			// Ignore the "count" key
			continue
		}

		key := fullKey
		if dot := strings.IndexByte(subKey, '.'); dot != -1 {
			key = fullKey[:dot+len(prefix)]
		}

		if seen[key] {
			continue
		}
		seen[key] = true

		// The flatmap format doesn't allow us to distinguish between keys
		// that contain periods and nested objects, so by convention a
		// map is only ever of primitive type in flatmap, and we just assume
		// that the remainder of the raw key (dots and all) is the key we
		// want in the result value.

		value, err := p.fromFlatmapValue(key, ty)
		if err != nil {
			return nil, err
		}
		if p.isValueAllowed(value, prefix) {
			values = append(values, value)
		}
	}
	if len(values) == 0 {
		return nil, nil
	}
	return values, nil
}

func (p *FlatmapParser) isAttributeIgnored(name string) bool {
	ignored := false
	for _, pattern := range p.ignoreKeys {
		if pattern.MatchString(name) {
			ignored = true
			break
		}
	}
	return ignored
}

func (p *FlatmapParser) isValueAllowed(value interface{}, prefix string) bool {
	if !reflect.ValueOf(value).IsValid() {
		return false
	}
	switch reflect.ValueOf(value).Kind() {
	case reflect.Slice:
		if reflect.ValueOf(value).Len() == 0 {
			return false
		}

		for i := 0; i < reflect.ValueOf(value).Len(); i++ {
			if !reflect.ValueOf(value).Index(i).IsZero() {
				return true
			}
		}
	case reflect.Map:
		if reflect.ValueOf(value).Len() == 0 {
			return false
		}
	}
	if !reflect.ValueOf(value).IsZero() {
		return true
	}

	allowed := false
	for _, pattern := range p.allowEmptyValues {
		if pattern.MatchString(prefix) {
			allowed = true
			break
		}
	}
	return allowed
}
