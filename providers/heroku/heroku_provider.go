// Copyright 2019 The Terraformer Authors.
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

package heroku

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type HerokuProvider struct { //nolint
	terraformutils.Provider
	apiKey string
	team   string
}

// Init the Provider for imports. args are defined in cmd/provider_cmd_heroku.go
func (p *HerokuProvider) Init(args []string) error {
	if len(args) > 0 {
		p.apiKey = args[0]
	}
	if len(args) > 1 {
		p.team = args[1]
	}
	return nil
}

func (p *HerokuProvider) GetName() string {
	return "heroku"
}

func (p *HerokuProvider) GetSource() string {
	return "heroku/heroku"
}

func (p *HerokuProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"heroku": map[string]interface{}{},
		},
	}
}

func (HerokuProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *HerokuProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"account_feature":   &AccountFeatureGenerator{},
		"app":               &AppGenerator{},
		"pipeline":          &PipelineGenerator{},
		"pipeline_coupling": &PipelineCouplingGenerator{},
		"team_collaborator": &TeamCollaboratorGenerator{},
		"team_member":       &TeamMemberGenerator{},
	}
}

func (p *HerokuProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("heroku: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key": p.apiKey,
		"team":    p.team,
	})
	return nil
}
