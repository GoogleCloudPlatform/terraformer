package terraform_utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func NewGcpRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"google": map[string]interface{}{
			"region":  region,
			"project": "waze-development",
		},
	}
}

func NewAwsRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"aws": map[string]interface{}{
			"region": region,
		},
	}
}

func GenerateTfState(resources []TerraformResource) error {
	tfState := NewTfState(resources)
	firstState, _ := json.MarshalIndent(tfState, "", "  ")
	ioutil.WriteFile("terraform.tfstate", firstState, os.ModePerm)
	oldRegion := os.Getenv("AWS_DEFAULT_REGION")
	defer os.Setenv("AWS_DEFAULT_REGION", oldRegion)
	os.Setenv("AWS_DEFAULT_REGION", "eu-west-1")
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
		log.Println(err)
	}

	opReq := c.Operation()
	opReq.Type = backend.OperationTypeRefresh
	opReq.Module = mod
	op, err := c.RunOperation(b, opReq)
	if err != nil {
		log.Println(err)
	}
	if op.Err != nil {
		log.Println(op.Err)
	}
	return nil
}

func GenerateTf(resources []TerraformResource, resourceName, region, provider string) error {
	data, err := HclPrint(resources, region, provider)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(resourceName+".tf", data, os.ModePerm)
}

func NewTfState(resources []TerraformResource) terraform.State {
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
