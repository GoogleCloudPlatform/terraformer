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

package panos

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type PanosService struct { //nolint
	terraformutils.Service
	client interface{}
	vsys   string
}

func (p *PanosService) Initialize() error {
	if _, ok := p.Args["vsys"].(string); ok {
		p.vsys = p.Args["vsys"].(string)
	} else {
		return errors.New(p.GetName() + ": " + "vsys name not parsable")
	}

	c, err := Initialize()
	if err != nil {
		return err
	}

	p.client = c

	return nil
}
