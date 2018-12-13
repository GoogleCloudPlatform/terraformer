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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/config"
	tfplugin "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/plugin/discovery"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/cli"
)

type BaseResource struct {
	Tags map[string]string `json:"tags,omitempty"`
}

// Generate tfstate empty and populate with terraform refresh all data
func GenerateTfState(resources []TerraformResource) error {
	tfState := NewTfState(resources)
	firstState, err := json.MarshalIndent(tfState, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("terraform.tfstate", firstState, os.ModePerm); err != nil {
		return err
	}
	// set to terraform don't check lock.json
	err = os.Setenv("TF_SKIP_PROVIDER_VERIFY", "skip")
	if err != nil {
		return err
	}
	// use plugins from os.Getenv("HOME") + "/.terraform.d"
	c := command.RefreshCommand{Meta: command.Meta{
		OverrideDataDir: os.Getenv("HOME") + "/.terraform.d",
		Ui:              cli.Ui(&cli.ConcurrentUi{Ui: &cli.BasicUi{Writer: os.Stdout}}),
	}}
	path, _ := os.Getwd()
	mod, _ := c.Module(path)

	var conf *config.Config
	if mod != nil {
		conf = mod.Config()
	}

	b, err := c.Backend(&command.BackendOpts{
		Config: conf,
	})
	if err != nil {
		return err
	}

	opReq := c.Operation()
	opReq.Module = mod
	opReq.Type = backend.OperationTypeRefresh
	op, err := c.RunOperation(b, opReq)
	if err != nil {
		return err
	}
	if op.Err != nil {
		return err
	}
	return nil
}

func NewTfState(resources []TerraformResource) *terraform.State {
	tfstate := &terraform.State{
		Version:   terraform.StateVersion,
		TFVersion: terraform.VersionString(),
		Serial:    1,
	}
	tfstate.Modules = []*terraform.ModuleState{
		{
			Path:      []string{"root"},
			Resources: map[string]*terraform.ResourceState{},
		},
	}
	for _, resource := range resources {
		resourceState := &terraform.ResourceState{
			Type:     resource.ResourceType,
			Primary:  resource.InstanceState,
			Provider: "provider." + resource.Provider,
		}
		tfstate.Modules[0].Resources[resource.ResourceType+"."+resource.ResourceName] = resourceState
	}
	return tfstate
}

func PrintTfState(resources []TerraformResource) ([]byte, error) {
	state := NewTfState(resources)
	var buf bytes.Buffer
	err := terraform.WriteState(state, &buf)
	return buf.Bytes(), err
}

func GetProvider(providerName string) (terraform.ResourceProvider, error) {
	pluginPath := os.Getenv("HOME") + "/." + command.DefaultPluginVendorDir
	files, err := ioutil.ReadDir(pluginPath)
	if err != nil {
		return nil, err
	}
	providerFileName := ""
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "terraform-provider-"+providerName) {
			providerFileName = pluginPath + "/" + file.Name()
		}
	}
	client := tfplugin.Client(discovery.PluginMeta{Path: providerFileName})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense(tfplugin.ProviderPluginName)
	if err != nil {
		return nil, err
	}

	provider := raw.(terraform.ResourceProvider)
	err = provider.Configure(&terraform.ResourceConfig{})
	return provider, err
}

func RefreshResources(provider terraform.ResourceProvider, cloudResources []TerraformResource) []TerraformResource {
	refreshedResources := []TerraformResource{}
	input := make(chan TerraformResource, 100)
	output := make(chan *terraform.InstanceState, 100)
	var wg sync.WaitGroup
	done := make(chan bool)
	go func() {
		for _, r := range cloudResources {
			input <- r
		}
		close(input)
	}()
	go func() {
		tmp := append([]TerraformResource{}, cloudResources...)
		for state := range output {
			if state == nil || state.ID == "" {
				continue
			}
			for _, r := range tmp {
				if r.ID == state.ID {
					log.Println(state.ID, r)
					r.InstanceState = state
					refreshedResources = append(refreshedResources, r)
					break
				}
			}
		}
		done <- true
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RefreshResource(provider, input, output, &wg)
	}
	wg.Wait()
	close(output)
	<-done
	return refreshedResources
}

func RefreshResource(provider terraform.ResourceProvider, input chan TerraformResource, output chan *terraform.InstanceState, wg *sync.WaitGroup) {
	for r := range input {
		state, _ := provider.Refresh(r.InstanceInfo, r.InstanceState)
		output <- state
	}
	wg.Done()
}
