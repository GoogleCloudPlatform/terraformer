package googleworkspace

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

type GoogleWorkspaceProvider struct {
	terraformutils.Provider
	orgID                  string
	credentialJsonFilepath string
	impersonatedUserEmail  string
}

func (p *GoogleWorkspaceProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"customer_id": p.orgID,
			},
		},
		"terraform": map[string]interface{}{
			"required_providers": []map[string]interface{}{{
				p.GetName(): map[string]interface{}{
					"source":  "yohan460/googleworkspace",
					"version": providerwrapper.GetProviderVersion(p.GetName()),
				},
			}},
		},
	}
}

func (p *GoogleWorkspaceProvider) Init(args []string) error {
	orgID := os.Getenv("GOOGLEWORKSPACE_CUSTOMER_ID")
	if orgID == "" {
		return errors.New("set GOOGLEWORKSPACE_CUSTOMER_ID env var")
	}
	p.orgID = orgID

	credentialJsonFilepath := os.Getenv("GOOGLEWORKSPACE_CREDENTIALS")
	if credentialJsonFilepath == "" {
		return errors.New("set GOOGLEWORKSPACE_CREDENTIALS env var")
	}
	p.credentialJsonFilepath = credentialJsonFilepath

	impersonatedUserEmail := os.Getenv("GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL")
	if impersonatedUserEmail == "" {
		return errors.New("set GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL env var")
	}
	p.impersonatedUserEmail = impersonatedUserEmail

	return nil
}

func (p *GoogleWorkspaceProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"alerts": {"alert_notification_endpoints": []string{"alert_notification_endpoints", "id"}},
	}
}

func (p *GoogleWorkspaceProvider) GetName() string {
	return "googleworkspace"
}

func (p *GoogleWorkspaceProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	fullname := p.GetName() + "_" + serviceName
	if _, isSupported = p.GetSupportedService()[fullname]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " is not a supported service")
	}
	p.Service = p.GetSupportedService()[fullname]
	p.Service.SetName(fullname)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetVerbose(verbose)
	p.Service.SetArgs(map[string]interface{}{
		"org_id":                   p.orgID,
		"credential_json_filepath": p.credentialJsonFilepath,
		"impersonated_user_email":  p.impersonatedUserEmail,
	})
	return nil
}

func (p *GoogleWorkspaceProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"googleworkspace_org_unit": &OrgUnitGenerator{},
	}
}
