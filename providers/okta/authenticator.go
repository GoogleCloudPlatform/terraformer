// Copyright 2025 The Terraformer Authors.
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

package okta

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

type AuthenticatorGenerator struct {
	OktaService
}

func (g AuthenticatorGenerator) createResources(authenticators []okta.ListAuthenticators200ResponseInner) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, authenticator := range authenticators {
		instance := authenticator.GetActualInstance()
		if instance == nil {
			continue
		}

		var resourceID, resourceName string
		switch inst := instance.(type) {
		case *okta.AuthenticatorKeyPassword:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		case *okta.AuthenticatorKeyEmail:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		case *okta.AuthenticatorKeyPhone:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		case *okta.AuthenticatorKeyGoogleOtp:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		case *okta.AuthenticatorKeyOktaVerify:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		case *okta.AuthenticatorKeyWebauthn:
			resourceID = inst.GetId()
			resourceName = normalizeResourceNameWithRandom(inst.GetName(), true)
		default:
			continue
		}

		resources = append(resources, terraformutils.NewSimpleResource(
			resourceID,
			resourceName,
			"okta_authenticator",
			"okta",
			[]string{},
		))
	}
	return resources
}

func (g *AuthenticatorGenerator) InitResources() error {
	ctx, client, err := g.ClientV5()
	if err != nil {
		return err
	}

	authenticators, _, err := client.AuthenticatorAPI.ListAuthenticators(ctx).Execute()
	if err != nil {
		return err
	}

	g.Resources = g.createResources(authenticators)
	return nil
}
