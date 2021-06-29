package vault

import (
	"errors"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type Provider struct {
	terraformutils.Provider
	token   string
	address string
}

func (p *Provider) Init(args []string) error {

	if address := os.Getenv("VAULT_ADDR"); address != "" {
		p.address = os.Getenv("VAULT_ADDR")
	}

	if token := os.Getenv("VAULT_TOKEN"); token != "" {
		p.token = os.Getenv("VAULT_TOKEN")
	}

	if len(args) > 0 && args[0] != "" {
		p.address = args[0]
	}

	if len(args) > 1 && args[1] != "" {
		p.token = args[1]
	}

	return nil
}

func (p *Provider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"token":   cty.StringVal(p.token),
		"address": cty.StringVal(p.address),
	})
}

func (p *Provider) GetName() string {
	return "vault"
}

func (p *Provider) InitService(serviceName string, verbose bool) error {
	if service, ok := p.GetSupportedService()[serviceName]; ok {
		p.Service = service
		p.Service.SetName(serviceName)
		p.Service.SetVerbose(verbose)
		p.Service.SetProviderName(p.GetName())
		p.Service.SetArgs(map[string]interface{}{
			"token":   p.token,
			"address": p.address,
		})
		if err := service.(*ServiceGenerator).setVaultClient(); err != nil {
			return err
		}
		return nil
	}
	return errors.New(p.GetName() + ": " + serviceName + " not supported service")
}

func getSupportedMountServices() map[string]terraformutils.ServiceGenerator {
	services := make(map[string]terraformutils.ServiceGenerator)
	mapping := map[string][]string{
		"secret_backend":      {"ad", "aws", "azure", "consul", "gcp", "nomad", "pki", "rabbitmq", "terraform_cloud"},
		"secret_backend_role": {"ad", "aws", "azure", "consul", "database", "pki", "rabbitmq", "ssh"},
		"auth_backend":        {"gcp", "github", "jwt", "ldap", "okta"},
		"auth_backend_role":   {"alicloud", "approle", "aws", "azure", "cert", "gcp", "jwt", "kubernetes", "token"},
		"auth_backend_user":   {"ldap", "okta"},
		"auth_backend_group":  {"ldap", "okta"},
	}
	for resource, mountTypes := range mapping {
		for _, mountType := range mountTypes {
			services[fmt.Sprintf("%s_%s", mountType, resource)] =
				&ServiceGenerator{mountType: mountType, resource: resource}
		}
	}
	return services
}

func (p *Provider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	generators := getSupportedMountServices()
	generators["policy"] = &ServiceGenerator{resource: "policy"}
	generators["mount"] = &ServiceGenerator{resource: "mount"}
	generators["generic_secret"] = &ServiceGenerator{resource: "generic_secret", mountType: "kv"}
	return generators
}

func (Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (Provider) GetProviderData(_ ...string) map[string]interface{} {
	return map[string]interface{}{}
}
