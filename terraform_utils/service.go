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
)

type ServiceGenerator interface {
	InitResources() error
	GetResources() []Resource
	SetResources(resources []Resource)
	ParseFilter(rawFilter []string)
	PostConvertHook() error
	GetArgs() map[string]interface{}
	SetArgs(args map[string]interface{})
	SetName(name string)
	SetProviderName(name string)
	GetName() string
	CleanupWithFilter()
}

type Service struct {
	Name         string
	Resources    []Resource
	ProviderName string
	Args         map[string]interface{}
	Filter       map[string][]string
}

func (s *Service) SetProviderName(providerName string) {
	s.ProviderName = providerName
}

func (s *Service) ParseFilter(rawFilter []string) {
	s.Filter = map[string][]string{}
	for _, resource := range rawFilter {
		t := strings.Split(resource, "=")
		if len(t) != 2 {
			log.Println("Pattern for filter must be resource_type=id1:id2:id4")
			continue
		}
		resourceName, resourcesID := t[0], t[1]
		s.Filter[resourceName] = strings.Split(resourcesID, ":")
	}
}

func (s *Service) SetName(name string) {
	s.Name = name
}
func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) CleanupWithFilter() {
	if len(s.Filter) == 0 {
		return
	}
	newListOfResources := []Resource{}
	for _, v := range s.Resources {
		if _, exist := s.Filter[v.InstanceInfo.Type]; exist {
			for _, r := range s.Filter[v.InstanceInfo.Type] {
				if v.InstanceState.ID == r {
					newListOfResources = append(newListOfResources, v)
				}
			}
		} else {
			newListOfResources = append(newListOfResources, v)
		}
	}
	s.Resources = newListOfResources
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

func (s *Service) PopulateIgnoreKeys() {
	resourcesTypes := []string{}
	for _, r := range s.Resources {
		resourcesTypes = append(resourcesTypes, r.InstanceInfo.Type)
	}
	keys := IgnoreKeys(resourcesTypes, s.ProviderName)
	for k, v := range keys {
		for i := range s.Resources {
			if s.Resources[i].InstanceInfo.Type == k {
				s.Resources[i].IgnoreKeys = append(s.Resources[i].IgnoreKeys, v...)
			}
		}
	}
}
