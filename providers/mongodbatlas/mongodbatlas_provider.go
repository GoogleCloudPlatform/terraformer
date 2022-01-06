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

package mongodbatlas

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type MongoDBAtlasProvider struct { //nolint
	terraformutils.Provider
	publicKey  string
	privateKey string
	orgID      string
}

func (p *MongoDBAtlasProvider) Init(args []string) error {
	if args[0] != "" {
		p.publicKey = args[0]
	} else {
		if publicKey := os.Getenv("MCLI_PUBLIC_API_KEY"); publicKey != "" {
			p.publicKey = publicKey
		} else {
			return errors.New("publicKey requirement")
		}
	}

	if args[1] != "" {
		p.privateKey = args[1]
	} else {
		if privateKey := os.Getenv("MCLI_PRIVATE_API_KEY"); privateKey != "" {
			p.privateKey = privateKey
		} else {
			return errors.New("privateKey requirement")
		}
	}

	if args[2] != "" {
		p.orgID = args[2]
	} else {
		if orgID := os.Getenv("MCLI_ORG_ID"); orgID != "" {
			p.orgID = orgID
		} else {
			return errors.New("orgID requirement")
		}
	}

	return nil
}

func (p *MongoDBAtlasProvider) GetName() string {
	return "mongodbatlas"
}

func (p *MongoDBAtlasProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"public_key":  cty.StringVal(p.publicKey),
		"private_key": cty.StringVal(p.privateKey),
	})
}

func (p *MongoDBAtlasProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"public_key":  p.publicKey,
		"private_key": p.privateKey,
		"org_id":      p.orgID,
	})
	return nil
}

// GetSupportedService return map of support service for MongoDBAtlas
func (p *MongoDBAtlasProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"project": &ProjectGenerator{},
		"cluster": &ClusterGenerator{},
		"team":    &TeamGenerator{},
	}
}

func (p MongoDBAtlasProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p MongoDBAtlasProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
