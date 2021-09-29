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

package okta

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

type OktaProvider struct { //nolint
	terraformutils.Provider
	orgName  string
	baseURL  string
	apiToken string
}

func (p *OktaProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"okta": map[string]interface{}{
				"version": providerwrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (p *OktaProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"alerts": {"alert_notification_endpoints": []string{"alert_notification_endpoints", "id"}},
	}
}

func (p *OktaProvider) Init(args []string) error {
	orgName := os.Getenv("OKTA_ORG_NAME")
	if orgName == "" {
		return errors.New("set OKTA_ORG_NAME env var")
	}
	p.orgName = orgName

	baseURL := os.Getenv("OKTA_BASE_URL")
	if baseURL == "" {
		return errors.New("set OKTA_BASE_URL env var")
	}
	p.baseURL = baseURL

	apiToken := os.Getenv("OKTA_API_TOKEN")
	if apiToken == "" {
		return errors.New("set OKTA_API_TOKEN env var")
	}
	p.apiToken = apiToken

	return nil
}

func (p *OktaProvider) GetName() string {
	return "okta"
}

func (p *OktaProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " is not a supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetVerbose(verbose)
	p.Service.SetArgs(map[string]interface{}{
		"org_name":  p.orgName,
		"base_url":  p.baseURL,
		"api_token": p.apiToken,
	})
	return nil
}

func (p *OktaProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"okta_app_three_field":           &AppThreeFieldGenerator{},
		"okta_app_swa":                   &AppSWAGenerator{},
		"okta_app_secure_password_store": &AppSecurePasswordStoreGenerator{},
		"okta_app_basic_auth":            &AppBasicAuthGenerator{},
		"okta_app_auto_login":            &AppAutoLoginGenerator{},
		"okta_app_bookmark":              &AppBookmarkGenerator{},
		"okta_app_saml":                  &AppSamlGenerator{},
		"okta_app_oauth":                 &AppOAuthGenerator{},
		"okta_idp_oidc":                  &IdpOIDCGenerator{},
		"okta_idp_saml":                  &IdpSAMLGenerator{},
		"okta_idp_social":                &IdpSocialGenerator{},
		"okta_factor":                    &FactorGenerator{},
		"okta_network_zone":              &NetworkZoneGenerator{},
		"okta_trusted_origin":            &TrustedOriginGenerator{},
		"okta_user":                      &UserGenerator{},
		"okta_template_sms":              &SMSTemplateGenerator{},
		"okta_user_type":                 &UserTypeGenerator{},
		"okta_group":                     &GroupGenerator{},
		"okta_group_rule":                &GroupRuleGenerator{},
		"okta_event_hook":                &EventHookGenerator{},
		"okta_inline_hook":               &EventHookGenerator{},
		"okta_policy_password":           &PasswordPolicyGenerator{},
		"okta_policy_rule_password":      &PasswordPolicyRuleGenerator{},
		"okta_policy_signon":             &SignOnPolicyGenerator{},
		"okta_policy_rule_signon":        &SignOnPolicyRuleGenerator{},
		"okta_policy_mfa":                &MFAPolicyGenerator{},
		"okta_policy_rule_mfa":           &MFAPolicyRuleGenerator{},
		"okta_auth_server":               &AuthorizationServerGenerator{},
		"okta_auth_server_scope":         &AuthorizationServerScopeGenerator{},
		"okta_auth_server_claim":         &AuthorizationServerClaimGenerator{},
		"okta_auth_server_policy":        &AuthorizationServerPolicyGenerator{},
		"okta_user_schema":               &UserSchemaPropertyGenerator{},
		"okta_app_user_schema":           &AppUserSchemaPropertyGenerator{},
	}
}
