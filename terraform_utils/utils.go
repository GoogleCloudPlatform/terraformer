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

func RefreshResources(resources []Resource, providerName string, providerConfig map[string]interface{}) ([]Resource, error) {
	refreshedResources := []Resource{}
	input := make(chan *Resource, 100)
	provider, err := provider_wrapper.NewProviderWrapper(providerName, providerConfig)
	if err != nil {
		return refreshedResources, err
	}
	defer provider.Kill()
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

func IgnoreKeys(resourcesTypes []string, providerName string) map[string][]string {
	p, err := provider_wrapper.NewProviderWrapper(providerName, map[string]interface{}{})
	if err != nil {
		log.Println("plugin error:", err)
		return map[string][]string{}
	}
	defer p.Kill()
	readOnlyAttributes, err := p.GetReadOnlyAttributes(resourcesTypes)
	if err != nil {
		log.Println("plugin error:", err)
		return map[string][]string{}
	}
	return readOnlyAttributes
}
