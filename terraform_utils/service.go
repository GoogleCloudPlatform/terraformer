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
	"log"
	"strings"

	"github.com/zclconf/go-cty/cty"
)

type ServiceGenerator interface {
	InitResources() error
	GetResources() []Resource
	SetResources(resources []Resource)
	ParseFilter(rawFilter string) []ResourceFilter
	ParseFilters(rawFilters []string)
	PostConvertHook() error
	GetArgs() map[string]interface{}
	SetArgs(args map[string]interface{})
	SetName(name string)
	SetProviderName(name string)
	GetName() string
	InitialCleanup()
	PopulateIgnoreKeys(cty.Value)
	PostRefreshCleanup()
}

type Service struct {
	Name         string
	Resources    []Resource
	ProviderName string
	Args         map[string]interface{}
	Filter       []ResourceFilter
}

func (s *Service) SetProviderName(providerName string) {
	s.ProviderName = providerName
}

func (s *Service) ParseFilters(rawFilters []string) {
	s.Filter = []ResourceFilter{}
	for _, rawFilter := range rawFilters {
		filters := s.ParseFilter(rawFilter)
		for _, resourceFilter := range filters {
			s.Filter = append(s.Filter, resourceFilter)
		}
	}
}

func (s *Service) ParseFilter(rawFilter string) []ResourceFilter {
	var filters []ResourceFilter
	if len(strings.Split(rawFilter, "=")) == 2 {
		parts := strings.Split(rawFilter, "=")
		resourceName, resourcesID := parts[0], parts[1]
		filters = append(filters, ResourceFilter{
			ResourceName:     resourceName,
			FieldPath:        "id",
			AcceptableValues: ParseFilterValues(resourcesID),
		})
	} else {
		parts := strings.Split(rawFilter, ";")
		if len(parts) != 2 && len(parts) != 3 {
			log.Print("Invalid filter: " + rawFilter)
			return filters
		}
		var ResourceNamePart string
		var FieldPathPart string
		var AcceptableValuesPart string
		if len(parts) == 2 {
			ResourceNamePart = ""
			FieldPathPart = parts[0]
			AcceptableValuesPart = parts[1]
		} else {
			ResourceNamePart = strings.TrimPrefix(parts[0], "Type=")
			FieldPathPart = parts[1]
			AcceptableValuesPart = parts[2]
		}

		filters = append(filters, ResourceFilter{
			ResourceName:     ResourceNamePart,
			FieldPath:        strings.TrimPrefix(FieldPathPart, "Name="),
			AcceptableValues: ParseFilterValues(strings.TrimPrefix(AcceptableValuesPart, "Value=")),
		})
	}
	return filters
}

func (s *Service) SetName(name string) {
	s.Name = name
}
func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) InitialCleanup() {
	FilterCleanup(s, true)
}

func (s *Service) PostRefreshCleanup() {
	if len(s.Filter) != 0 {
		FilterCleanup(s, false)
	}
}

func (s *Service) GetArgs() map[string]interface{} {
	return s.Args
}
func (s *Service) SetArgs(args map[string]interface{}) {
	s.Args = args
}

func (s *Service) GetResources() []Resource {
	return s.Resources
}
func (s *Service) SetResources(resources []Resource) {
	s.Resources = resources
}

func (s *Service) InitResources() error {
	panic("implement me")
}

func (s *Service) PostConvertHook() error {
	return nil
}

func (s *Service) PopulateIgnoreKeys(providerConfig cty.Value) {
	resourcesTypes := []string{}
	for _, r := range s.Resources {
		resourcesTypes = append(resourcesTypes, r.InstanceInfo.Type)
	}
	keys := IgnoreKeys(resourcesTypes, s.ProviderName, providerConfig)
	for k, v := range keys {
		for i := range s.Resources {
			if s.Resources[i].InstanceInfo.Type == k {
				s.Resources[i].IgnoreKeys = append(s.Resources[i].IgnoreKeys, v...)
			}
		}
	}
}
