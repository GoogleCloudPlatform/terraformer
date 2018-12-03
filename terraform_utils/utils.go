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
	meta := command.Meta{
		OverrideDataDir: os.Getenv("HOME") + "/.terraform.d",
		Ui:              cli.Ui(&cli.ConcurrentUi{Ui: &cli.BasicUi{Writer: os.Stdout}}),
	}
	c := command.RefreshCommand{Meta: meta}
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
	opReq.Type = backend.OperationTypeRefresh
	opReq.Module = mod
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
