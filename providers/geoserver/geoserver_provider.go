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

package geoserver

import (
	"errors"
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type GeoServerProvider struct { //nolint
	terraformutils.Provider
	geoserverURL    string
	geowebcacheURL  string
	user            string
	password        string
	insecure        bool
	targetWorkspace string
	targetDatastore string
}

func (p *GeoServerProvider) Init(args []string) error {
	log.Println("Init Provider from ARGS")

	p.geoserverURL = args[0]
	p.geowebcacheURL = args[1]
	p.user = args[2]
	p.password = args[3]
	p.insecure, _ = strconv.ParseBool(args[4])
	p.targetWorkspace = args[5]
	p.targetDatastore = args[6]
	return nil
}

func (p *GeoServerProvider) GetName() string {
	return "geoserver"
}

func (p *GeoServerProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *GeoServerProvider) GetConfig() cty.Value {
	log.Println("GetConfig")
	return cty.ObjectVal(map[string]cty.Value{
		"url":      cty.StringVal(p.geoserverURL),
		"gwc_url":  cty.StringVal(p.geowebcacheURL),
		"username": cty.StringVal(p.user),
		"password": cty.StringVal(p.password),
		"insecure": cty.BoolVal(p.insecure),
	})
}

func (p *GeoServerProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}

func (p *GeoServerProvider) InitService(serviceName string, verbose bool) error {
	log.Println("InitService")
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"geoserverURL":    p.geoserverURL,
		"geowebcacheURL":  p.geowebcacheURL,
		"user":            p.user,
		"password":        p.password,
		"insecure":        p.insecure,
		"targetWorkspace": p.targetWorkspace,
		"targetDatastore": p.targetDatastore,
	})
	return nil
}

func (p *GeoServerProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	log.Println("GetSupportedService")
	return map[string]terraformutils.ServiceGenerator{
		"workspaces":   &WorkspacesGenerator{},
		"datastores":   &DatastoresGenerator{},
		"featuretypes": &FeatureTypesGenerator{},
	}
}

func (GeoServerProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}
