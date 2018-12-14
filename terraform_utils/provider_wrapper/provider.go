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
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/config/configschema"
	tfplugin "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

type ProviderWrapper struct {
	Provider     terraform.ResourceProvider
	client       *plugin.Client
	rpcClient    plugin.ClientProtocol
	providerName string
}

func NewProviderWrapper(providerName string) (*ProviderWrapper, error) {
	p := &ProviderWrapper{}
	p.providerName = providerName
	err := p.initProvider()
	return p, err
}

func (p *ProviderWrapper) Kill() {
	p.client.Kill()
}

func (p *ProviderWrapper) GetReadOnlyAttributes(resourceTypes []string) (map[string][]string, error) {
	schema, err := p.Provider.GetSchema(&terraform.ProviderSchemaRequest{
		ResourceTypes: resourceTypes,
	})
	if err != nil {
		return map[string][]string{}, err
	}
	readOnlyAttributes := map[string][]string{}
	for resourceName, obj := range schema.ResourceTypes {
		readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^id$")
		for k, v := range obj.Attributes {
			if !v.Optional && !v.Required {
				if v.Type.IsListType() {
					readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^"+k+".(.*)")
				} else {
					readOnlyAttributes[resourceName] = append(readOnlyAttributes[resourceName], "^"+k+"$")
				}

			}
		}
		readOnlyAttributes[resourceName] = p.readObjBlocks(obj.BlockTypes, readOnlyAttributes[resourceName], "-1")
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
						readOnlyAttributes = append(readOnlyAttributes, "^"+k+".[0-9]."+key+"$")
					} else {
						readOnlyAttributes = append(readOnlyAttributes, "^"+parent+".(.*)."+key+"$")
					}
				case configschema.NestingSet:
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
	return p.Provider.Refresh(info, state)
}

func (p *ProviderWrapper) initProvider() error {
	pluginPath := os.Getenv("HOME") + "/." + command.DefaultPluginVendorDir
	files, err := ioutil.ReadDir(pluginPath)
	if err != nil {
		return err
	}
	providerFileName := ""
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "terraform-provider-"+p.providerName) {
			providerFileName = pluginPath + "/" + file.Name()
		}
	}
	p.client = plugin.NewClient(
		&plugin.ClientConfig{
			Cmd:              exec.Command(providerFileName),
			HandshakeConfig:  tfplugin.Handshake,
			Managed:          true,
			Plugins:          tfplugin.PluginMap,
			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC, plugin.ProtocolNetRPC},
			Logger: hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Level:  hclog.Trace,
				Output: os.Stderr,
			}),
		})
	p.rpcClient, err = p.client.Client()
	if err != nil {
		return err
	}
	raw, err := p.rpcClient.Dispense(tfplugin.ProviderPluginName)
	if err != nil {
		return err
	}

	p.Provider = raw.(terraform.ResourceProvider)
	err = p.Provider.Configure(&terraform.ResourceConfig{})
	return err
}
