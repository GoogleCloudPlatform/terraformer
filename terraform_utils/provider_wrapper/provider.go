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
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/configs/configschema"
	tfplugin "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/terraform"
)

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
	return readOnlyAttributes, nil
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
	stateVal, _ := state.AttrsAsObjectValue(cty.EmptyObject)

	p.Provider.ReadResource(providers.ReadResourceRequest{
		TypeName:   info.Type,
		PriorState: stateVal,
	})
	return state, nil
}

func (p *ProviderWrapper) initProvider() error {
	pluginPath := command.DefaultDataDir + string(os.PathSeparator) + "plugins" + string(os.PathSeparator) + runtime.GOOS + "_" + runtime.GOARCH
	files, err := ioutil.ReadDir(pluginPath)
	if err != nil {
		pluginPath = os.Getenv("HOME") + string(os.PathSeparator) + "." + command.DefaultPluginVendorDir
		files, err = ioutil.ReadDir(pluginPath)
		if err != nil {
			return err
		}
	}
	providerFileName := ""
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "terraform-provider-"+p.providerName) {
			providerFileName = pluginPath + string(os.PathSeparator) + file.Name()
		}
	}
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Level:  hclog.Trace,
		Output: os.Stderr,
	})

	p.client = plugin.NewClient(
		&plugin.ClientConfig{
			Cmd:              exec.Command(providerFileName),
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

	// requestConfig := cty.ObjectVal(map[string]cty.Value{})
	// if cty.NilVal != p.config {
	// 	requestConfig = p.config
	// }
	// configPrepareResp := p.Provider.PrepareProviderConfig(providers.PrepareProviderConfigRequest{
	// 	Config: requestConfig,
	// })

	return nil
}
