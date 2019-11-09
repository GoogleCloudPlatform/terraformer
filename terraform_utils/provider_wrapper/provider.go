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

package provider_wrapper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/configs/configschema"
	tfplugin "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/version"
)

// DefaultDataDir is the default directory for storing local data.
const DefaultDataDir = ".terraform"

// DefaultPluginVendorDir is the location in the config directory to look for
// user-added plugin binaries. Terraform only reads from this path if it
// exists, it is never created by terraform.
const DefaultPluginVendorDir = "terraform.d/plugins/" + pluginMachineName

// pluginMachineName is the directory name used in new plugin paths.
const pluginMachineName = runtime.GOOS + "_" + runtime.GOARCH

type ProviderWrapper struct {
	Provider     *tfplugin.GRPCProvider
	client       *plugin.Client
	rpcClient    plugin.ClientProtocol
	providerName string
	config       cty.Value
}

func NewProviderWrapper(providerName string, providerConfig cty.Value) (*ProviderWrapper, error) {
	p := &ProviderWrapper{}
	p.providerName = providerName
	p.config = providerConfig
	err := p.initProvider()
	return p, err
}

func (p *ProviderWrapper) Kill() {
	p.client.Kill()
}

func (p *ProviderWrapper) GetReadOnlyAttributes(resourceTypes []string) (map[string][]string, error) {
	r := p.Provider.GetSchema()

	if r.Diagnostics.HasErrors() {
		return nil, r.Diagnostics.Err()
	}
	readOnlyAttributes := map[string][]string{}
	for resourceName, obj := range r.ResourceTypes {
		if contains(resourceTypes, resourceName) {
			readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^id$")
			for k, v := range obj.Block.Attributes {
				if !v.Optional && !v.Required {
					if v.Type.IsListType() || v.Type.IsSetType() {
						readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^"+k+".(.*)")
					} else {
						readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^"+k+"$")
					}

				}
			}
			readOnlyAttributes[resourceName] = p.readObjBlocks(obj.Block.BlockTypes, readOnlyAttributes[resourceName], "-1")
		}
	}
	return readOnlyAttributes, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (p *ProviderWrapper) readObjBlocks(block map[string]*configschema.NestedBlock, readOnlyAttributes []string, parent string) []string {
	for k, v := range block {
		if len(v.BlockTypes) > 0 {
			readOnlyAttributes = p.readObjBlocks(v.BlockTypes, readOnlyAttributes, k)
		}
		fieldCount := 0
		for key, l := range v.Attributes {
			if !l.Optional && !l.Required {
				fieldCount++
				switch v.Nesting {
				case configschema.NestingList:
					if parent == "-1" {
						readOnlyAttributes = append(readOnlyAttributes, "^"+k+".[0-9]."+key+"($|\\.[0-9]|\\.#)")
					} else {
						readOnlyAttributes = append(readOnlyAttributes, "^"+parent+".(.*)."+key+"$")
					}
				case configschema.NestingSet:
					if parent == "-1" {
						readOnlyAttributes = append(readOnlyAttributes, "^"+k+".[0-9]."+key+"$")
					} else {
						readOnlyAttributes = append(readOnlyAttributes, "^"+parent+".(.*)."+key+"($|\\.(.*))")
					}
				case configschema.NestingMap:
					readOnlyAttributes = append(readOnlyAttributes, parent+"."+key)
				default:
					readOnlyAttributes = append(readOnlyAttributes, parent+"."+key+"$")
				}
			}
		}
		if fieldCount == len(v.Block.Attributes) && fieldCount > 0 {
			readOnlyAttributes = append(readOnlyAttributes, "^"+k)
		}
	}
	return readOnlyAttributes
}

func (p *ProviderWrapper) Refresh(info *terraform.InstanceInfo, state *terraform.InstanceState) (*terraform.InstanceState, error) {
	schema := p.Provider.GetSchema()
	impliedType := schema.ResourceTypes[info.Type].Block.ImpliedType()
	priorState, err := state.AttrsAsObjectValue(impliedType)
	if err != nil {
		return nil, err
	}
	resp := p.Provider.ReadResource(providers.ReadResourceRequest{
		TypeName:   info.Type,
		PriorState: priorState,
		Private:    []byte{},
	})

	if resp.Diagnostics.HasErrors() {
		// retry with different serialization mechanism
		priorState, err = gocty.ToCtyValue(state, impliedType)
		resp = p.Provider.ReadResource(providers.ReadResourceRequest{
			TypeName:   info.Type,
			PriorState: priorState,
			Private:    []byte{},
		})
		if resp.Diagnostics.HasErrors() {
			return nil, resp.Diagnostics.Err()
		}
	}

	if resp.NewState.IsNull() {
		msg := fmt.Sprintf("ERROR: Read resource response is null for resource %s", info.Id)
		return nil, errors.New(msg)
	}

	return terraform.NewInstanceStateShimmedFromValue(resp.NewState, int(schema.Provider.Version)), nil
}

func (p *ProviderWrapper) initProvider() error {
	providerFilePath, err := getProviderFileName(p.providerName)
	if err != nil {
		return err
	}
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Level:  hclog.Error,
		Output: os.Stderr,
	})
	p.client = plugin.NewClient(
		&plugin.ClientConfig{
			Cmd:              exec.Command(providerFilePath),
			HandshakeConfig:  tfplugin.Handshake,
			VersionedPlugins: tfplugin.VersionedPlugins,
			Managed:          true,
			Logger:           logger,
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
			AutoMTLS:         true,
		})
	p.rpcClient, err = p.client.Client()
	if err != nil {
		return err
	}
	raw, err := p.rpcClient.Dispense(tfplugin.ProviderPluginName)
	if err != nil {
		return err
	}

	p.Provider = raw.(*tfplugin.GRPCProvider)

	config, err := p.Provider.GetSchema().Provider.Block.CoerceValue(p.config)
	if err != nil {
		return err
	}
	p.Provider.Configure(providers.ConfigureRequest{
		TerraformVersion: version.Version,
		Config:           config,
	})

	return nil
}

func getProviderFileName(providerName string) (string, error) {
	defaultDataDir := os.Getenv("TF_DATA_DIR")
	if defaultDataDir == "" {
		defaultDataDir = DefaultDataDir
	}
	pluginPath := defaultDataDir + string(os.PathSeparator) + "plugins" + string(os.PathSeparator) + runtime.GOOS + "_" + runtime.GOARCH
	files, err := ioutil.ReadDir(pluginPath)
	if err != nil {
		pluginPath = os.Getenv("HOME") + string(os.PathSeparator) + "." + DefaultPluginVendorDir
		files, err = ioutil.ReadDir(pluginPath)
		if err != nil {
			return "", err
		}
	}
	providerFilePath := ""
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "terraform-provider-"+providerName) {
			providerFilePath = pluginPath + string(os.PathSeparator) + file.Name()
		}
	}
	return providerFilePath, nil
}

func GetProviderVersion(providerName string) string {
	providerFilePath, err := getProviderFileName(providerName)
	if err != nil {
		log.Println("Can't find provider file path. Ensure that you are following https://www.terraform.io/docs/configuration/providers.html#third-party-plugins.")
		return ""
	}
	t := strings.Split(providerFilePath, string(os.PathSeparator))
	providerFileName := t[len(t)-1]
	providerFileNameParts := strings.Split(providerFileName, "_")
	if len(providerFileNameParts) < 2 {
		log.Println("Can't find provider version. Ensure that you are following https://www.terraform.io/docs/configuration/providers.html#plugin-names-and-versions.")
		return ""
	}
	providerVersion := providerFileNameParts[1]
	return "~>" + providerVersion
}
