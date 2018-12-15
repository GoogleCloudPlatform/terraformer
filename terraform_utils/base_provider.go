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

type ProviderGenerator interface {
	Init(args []string) error
	GenerateFiles()
	GetService() ServiceGenerator
	GetName() string
	InitService(serviceName string) error
	GenerateOutputPath() error
	RegionResource() map[string]interface{}
	CurrentPath() string
}

type Provider struct {
	Service ServiceGenerator
}

func (p *Provider) Init(args []string) error {
	panic("implement me")
}

func (p *Provider) GetName() string {
	panic("implement me")
}

func (p *Provider) InitService(serviceName string) error {
	panic("implement me")
}

func (p *Provider) GenerateOutputPath() error {
	panic("implement me")
}

func (p *Provider) GenerateFiles() {
	panic("implement me")
}

func (p *Provider) GetService() ServiceGenerator {
	return p.Service
}

func GetName() string {
	panic("implement me")
}
