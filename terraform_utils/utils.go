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
	"bytes"
	"log"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"

	"github.com/hashicorp/terraform/terraform"

	"github.com/zclconf/go-cty/cty"
)

type BaseResource struct {
	Tags map[string]string `json:"tags,omitempty"`
}

func NewTfState(resources []Resource) *terraform.State {
	tfstate := &terraform.State{
		Version:   terraform.StateVersion,
		TFVersion: terraform.VersionString(),
		Serial:    1,
	}
	outputs := map[string]*terraform.OutputState{}
	for _, r := range resources {
		for k, v := range r.Outputs {
			outputs[k] = v
		}
	}
	tfstate.Modules = []*terraform.ModuleState{
		{
			Path:      []string{"root"},
			Resources: map[string]*terraform.ResourceState{},
			Outputs:   outputs,
		},
	}
	for _, resource := range resources {
		resourceState := &terraform.ResourceState{
			Type:     resource.InstanceInfo.Type,
			Primary:  resource.InstanceState,
			Provider: "provider." + resource.Provider,
		}
		tfstate.Modules[0].Resources[resource.InstanceInfo.Type+"."+resource.ResourceName] = resourceState
	}
	return tfstate
}

func PrintTfState(resources []Resource) ([]byte, error) {
	state := NewTfState(resources)
	var buf bytes.Buffer
	err := terraform.WriteState(state, &buf)
	return buf.Bytes(), err
}

func RefreshResources(resources []Resource, provider *provider_wrapper.ProviderWrapper) ([]Resource, error) {
	refreshedResources := []Resource{}
	input := make(chan *Resource, 100)
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		go RefreshResourceWorker(input, &wg, provider)
	}
	for i := range resources {
		wg.Add(1)
		input <- &resources[i]
	}
	wg.Wait()
	close(input)
	for _, r := range resources {
		if r.InstanceState != nil && r.InstanceState.ID != "" {
			refreshedResources = append(refreshedResources, r)
		} else {
			log.Printf("ERROR: Unable to refresh resource %s", r.ResourceName)
		}
	}
	return refreshedResources, nil
}

func RefreshResourceWorker(input chan *Resource, wg *sync.WaitGroup, provider *provider_wrapper.ProviderWrapper) {
	for r := range input {
		log.Println("Refreshing state...", r.InstanceInfo.Id)
		r.Refresh(provider)
		wg.Done()
	}
}

func IgnoreKeys(resourcesTypes []string, providerName string, providerConfig cty.Value) map[string][]string {
	p, err := provider_wrapper.NewProviderWrapper(providerName, providerConfig)
	if err != nil {
		log.Println("plugin error 1:", err)
		return map[string][]string{}
	}
	defer p.Kill()
	readOnlyAttributes, err := p.GetReadOnlyAttributes(resourcesTypes)
	if err != nil {
		log.Println("plugin error 2:", err)
		return map[string][]string{}
	}
	return readOnlyAttributes
}

func ParseFilterValues(value string) []string {
	var values []string

	valueBuffering := true
	wrapped := false
	var valueBuffer []byte
	for i := 0; i < len(value); i++ {
		if value[i] == '\'' {
			wrapped = !wrapped
			continue
		} else if value[i] == ':' {
			if len(valueBuffer) == 0 {
				continue
			} else if valueBuffering && !wrapped {
				values = append(values, string(valueBuffer))
				valueBuffering = false
				valueBuffer = []byte{}
				continue
			}
		}
		valueBuffering = true
		valueBuffer = append(valueBuffer, value[i])
	}
	if len(valueBuffer) > 0 {
		values = append(values, string(valueBuffer))
	}

	return values
}

func FilterCleanup(s *Service, isInitial bool) {
	if len(s.Filter) == 0 {
		return
	}
	var newListOfResources []Resource
	for _, resource := range s.Resources {
		allPredicatesTrue := true
		for _, filter := range s.Filter {
			if filter.isInitial() == isInitial {
				allPredicatesTrue = allPredicatesTrue && filter.Filter(resource)
			}
		}
		if allPredicatesTrue && !ContainsResource(newListOfResources, resource) {
			newListOfResources = append(newListOfResources, resource)
		}
	}
	s.Resources = newListOfResources
}

func ContainsResource(s []Resource, e Resource) bool {
	for _, a := range s {
		if a.InstanceInfo.Id == e.InstanceInfo.Id {
			return true
		}
	}
	return false
}
