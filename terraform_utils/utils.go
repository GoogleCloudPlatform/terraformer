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

func NewAwsRegionResource(region string) map[string]interface{} {
	return map[string]interface{}{
		"aws": map[string]interface{}{
			"region": region,
		},
	}
}

func GenerateTfState(resources []TerraformResource) error {
	tfstate := NewTfState(resources)
	firstState, _ := json.MarshalIndent(tfstate, "", "  ")
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
	log.Println(op.State)
	return nil
}

func GenerateTf(resources []TerraformResource, resourceName, region string) error {
	data, err := HclPrint(resources, region, "aws")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(resourceName+".tf", data, os.ModePerm)
}

func NewTfState(resources []TerraformResource) map[string]interface{} {
	tfstate := map[string]interface{}{
		"version":           1,
		"terraform_version": terraform.VersionString(),
		"serial":            1,
		"modules": []map[string]interface{}{
			{
				"path":      []string{"root"},
				"resources": map[string]interface{}{},
			},
		},
	}
	for _, resource := range resources {
		resourceState := map[string]interface{}{
			"type": resource.ResourceType,
			"primary": map[string]interface{}{
				"id": resource.ID,
			},
			"provider": "provider." + resource.Provider,
		}
		tfstate["modules"].([]map[string]interface{})[0]["resources"].(map[string]interface{})[resource.ResourceType+"."+resource.ResourceName] = resourceState
	}
	return tfstate
}
