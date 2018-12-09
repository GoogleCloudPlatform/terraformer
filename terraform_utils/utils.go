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
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/config"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform/terraform"
)

type BaseResource struct {
	Tags map[string]string `json:"tags,omitempty"`
}

// Generate tfstate empty and populate with terraform refresh all data
func GenerateTfState(resources []TerraformResource) error {
	tfState := newTfState(resources)
	firstState, err := json.MarshalIndent(tfState, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("terraform.tfstate", firstState, os.ModePerm); err != nil {
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

// print and write HCL to file
func GenerateTf(resources []TerraformResource, resourceName string, provider map[string]interface{}) error {
	data, err := HclPrint(resources, provider)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(resourceName+".tf", data, os.ModePerm)
}

func newTfState(resources []TerraformResource) terraform.State {
	tfstate := terraform.State{
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
			Type: resource.ResourceType,
			Primary: &terraform.InstanceState{
				ID:         resource.ID,
				Attributes: resource.Attributes,
			},
			Provider: "provider." + resource.Provider,
		}
		tfstate.Modules[0].Resources[resource.ResourceType+"."+resource.ResourceName] = resourceState
	}
	return tfstate
}
