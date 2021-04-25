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
	"bytes"
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statefile"
	"log"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

type BaseResource struct {
	Tags map[string]string `json:"tags,omitempty"`
}

func NewTfState(resources []Resource, providerSource addrs.Provider, resourceTypes map[string]providers.Schema) *states.State {
	wrapper := states.NewState().SyncWrapper()
	for _, resource := range resources {
		instance := resource.Address.Instance(addrs.NoKey).Absolute(addrs.RootModuleInstance)
		providerConfig := addrs.RootModuleInstance.ProviderConfigDefault(providerSource)
		typeSchema := resourceTypes[resource.Address.Type]

		// we need to use type stored in the instance state as we modify the attributes to make it writable to HCL
		objectSrc, err := resource.InstanceState.Encode(resource.InstanceState.Value.Type(), uint64(typeSchema.Version))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		wrapper.SetResourceInstanceCurrent(instance, objectSrc, providerConfig)
	}
	return wrapper.Lock()

	//tfstate := &terraform.State{
	//	Version:   terraform.StateVersion,
	//	TFVersion: terraform.VersionString(), //nolint
	//	Serial:    1,
	//}
	//outputs := map[string]*terraform.OutputState{}
	//for _, r := range resources {
	//	for k, v := range r.Outputs {
	//		outputs[k] = v
	//	}
	//}
	//tfstate.Modules = []*terraform.ModuleState{
	//	{
	//		Path:      []string{"root"},
	//		Resources: map[string]*terraform.ResourceState{},
	//		Outputs:   outputs,
	//	},
	//}
	//for _, resource := range resources {
	//	resourceState := &terraform.ResourceState{
	//		Type:     resource.InstanceInfo.Type,
	//		Primary:  resource.InstanceState,
	//		Provider: "provider." + resource.Provider,
	//	}
	//	tfstate.Modules[0].Resources[resource.InstanceInfo.Type+"."+resource.ResourceName] = resourceState
	//}
	//return tfstate
}

func PrintTfState(resources []Resource, providerSource addrs.Provider, resourceTypes map[string]providers.Schema) ([]byte, error) {
	state := NewTfState(resources, providerSource, resourceTypes)
	var buf bytes.Buffer
	file := statefile.New(state, "", 1)
	err := statefile.Write(file, &buf)
	return buf.Bytes(), err
}

func RefreshResources(resources []*Resource, provider *providerwrapper.ProviderWrapper, slowProcessingResources [][]*Resource) ([]*Resource, error) {
	refreshedResources := []*Resource{}
	input := make(chan *Resource, len(resources))
	var wg sync.WaitGroup
	poolSize := 15
	for i := range resources {
		wg.Add(1)
		input <- resources[i]
	}
	close(input)

	for i := 0; i < poolSize; i++ {
		go RefreshResourceWorker(input, &wg, provider)
	}

	spInputs := []chan *Resource{}
	for i, resourceGroup := range slowProcessingResources {
		spInputs = append(spInputs, make(chan *Resource, len(resourceGroup)))
		for j := range resourceGroup {
			spInputs[i] <- resourceGroup[j]
		}
		close(spInputs[i])
	}

	for i := 0; i < len(spInputs); i++ {
		wg.Add(len(slowProcessingResources[i]))
		go RefreshResourceWorker(spInputs[i], &wg, provider)
	}

	wg.Wait()
	for _, r := range resources {
		if r.InstanceState != nil && r.ImportID != "" {
			refreshedResources = append(refreshedResources, r)
		} else {
			log.Printf("ERROR: Unable to refresh resource %s", r.Address.String())
		}
	}

	for _, resourceGroup := range slowProcessingResources {
		for i := range resourceGroup {
			r := resourceGroup[i]
			if r.InstanceState != nil && r.ImportID != "" {
				refreshedResources = append(refreshedResources, r)
			} else {
				log.Printf("ERROR: Unable to refresh resource %s", r.Address.String())
			}
		}
	}
	return refreshedResources, nil
}

func RefreshResourcesByProvider(providersMapping *ProvidersMapping, providerWrapper *providerwrapper.ProviderWrapper) error {
	allResources := providersMapping.ShuffleResources()
	slowProcessingResources := make(map[ProviderGenerator][]*Resource)
	regularResources := []*Resource{}
	for i := range allResources {
		resource := allResources[i]
		if resource.SlowQueryRequired {
			provider := providersMapping.MatchProvider(resource)
			if slowProcessingResources[provider] == nil {
				slowProcessingResources[provider] = []*Resource{}
			}
			slowProcessingResources[provider] = append(slowProcessingResources[provider], resource)
		} else {
			regularResources = append(regularResources, resource)
		}
	}

	var spResourcesList [][]*Resource
	for p := range slowProcessingResources {
		spResourcesList = append(spResourcesList, slowProcessingResources[p])
	}

	refreshedResources, err := RefreshResources(regularResources, providerWrapper, spResourcesList)
	if err != nil {
		return err
	}

	providersMapping.SetResources(refreshedResources)
	return nil
}

func RefreshResourceWorker(input chan *Resource, wg *sync.WaitGroup, provider *providerwrapper.ProviderWrapper) {
	for r := range input {
		log.Println("Refreshing state...", r.ImportID)
		r.Refresh(provider)
		wg.Done()
	}
}

func IgnoreKeys(resourcesTypes []string, p *providerwrapper.ProviderWrapper) map[string][]string {
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
		if a.ImportID == e.ImportID {
			return true
		}
	}
	return false
}
