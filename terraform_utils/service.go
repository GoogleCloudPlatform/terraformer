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

type ServiceGenerator interface {
	InitResources() error
	GetResources() []Resource
	SetResources(resources []Resource)
	PostConvertHook() error
	GetArgs() map[string]string
	SetArgs(args map[string]string)
	SetName(name string)
	SetProviderName(name string)
	GetName() string
}

type Service struct {
	Name         string
	Resources    []Resource
	ProviderName string
	Args         map[string]string
}

func (s *Service) SetProviderName(providerName string) {
	s.ProviderName = providerName
}

func (s *Service) SetName(name string) {
	s.Name = name
}
func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) GetArgs() map[string]string {
	return s.Args
}
func (s *Service) SetArgs(args map[string]string) {
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
