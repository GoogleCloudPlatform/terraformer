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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
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
	SetVerbose(bool)
	SetProviderName(name string)
	GetProviderName() string
	GetName() string
	InitialCleanup()
	PopulateIgnoreKeys(*providerwrapper.ProviderWrapper)
	PostRefreshCleanup()
}

type Service struct {
	Name         string
	Resources    []Resource
	ProviderName string
	Args         map[string]interface{}
	Filter       []ResourceFilter
	Verbose      bool
}

func (s *Service) SetProviderName(providerName string) {
	s.ProviderName = providerName
}

func (s *Service) GetProviderName() string {
	return s.ProviderName
}

func (s *Service) SetVerbose(verbose bool) {
	s.Verbose = verbose
}

func (s *Service) ParseFilters(rawFilters []string) {
	s.Filter = []ResourceFilter{}
	for _, rawFilter := range rawFilters {
		filters := s.ParseFilter(rawFilter)
		s.Filter = append(s.Filter, filters...)
	}
}

func (s *Service) ParseFilter(rawFilter string) []ResourceFilter {
	var filters []ResourceFilter
	if !strings.HasPrefix(rawFilter, "Name=") && len(strings.Split(rawFilter, "=")) == 2 {
		parts := strings.Split(rawFilter, "=")
		serviceName, resourcesID := parts[0], parts[1]
		filters = append(filters, ResourceFilter{
			ServiceName:      serviceName,
			FieldPath:        "id",
			AcceptableValues: ParseFilterValues(resourcesID),
		})
	} else {
		parts := strings.Split(rawFilter, ";")
		if !((len(parts) == 1 && strings.HasPrefix(rawFilter, "Name=")) || len(parts) == 2 || len(parts) == 3) {
			log.Print("Invalid filter: " + rawFilter)
			return filters
		}
		var ServiceNamePart string
		var FieldPathPart string
		var AcceptableValuesPart string
		switch len(parts) {
		case 1:
			ServiceNamePart = ""
			FieldPathPart = parts[0]
			AcceptableValuesPart = ""
		case 2:
			ServiceNamePart = ""
			FieldPathPart = parts[0]
			AcceptableValuesPart = parts[1]
		default:
			ServiceNamePart = strings.TrimPrefix(parts[0], "Type=")
			FieldPathPart = parts[1]
			AcceptableValuesPart = parts[2]
		}

		filters = append(filters, ResourceFilter{
			ServiceName:      ServiceNamePart,
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

func (s *Service) PopulateIgnoreKeys(providerWrapper *providerwrapper.ProviderWrapper) {
	var resourcesTypes []string
	for _, r := range s.Resources {
		resourcesTypes = append(resourcesTypes, r.InstanceInfo.Type)
	}
	keys := IgnoreKeys(resourcesTypes, providerWrapper)
	for k, v := range keys {
		for i := range s.Resources {
			if s.Resources[i].InstanceInfo.Type == k {
				s.Resources[i].IgnoreKeys = append(s.Resources[i].IgnoreKeys, v...)
			}
		}
	}
}
