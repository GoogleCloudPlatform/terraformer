package myrasec

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

// MyrasecProvider
type MyrasecProvider struct {
	terraformutils.Provider
}

// Init
func (p *MyrasecProvider) Init(args []string) error {
	return nil
}

// GetName
func (p *MyrasecProvider) GetName() string {
	return "myrasec"
}

// GetProviderData
func (p *MyrasecProvider) GetProviderData(_ ...string) map[string]any {
	return map[string]any{}
}

// GetResourceConnections
func (MyrasecProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetSupportedService
func (p *MyrasecProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"cache_setting":     &CacheSettingGenerator{},
		"dns_record":        &DNSGenerator{},
		"domain":            &DomainGenerator{},
		"error_page":        &ErrorPageGenerator{},
		"ip_filter":         &IPFilterGenerator{},
		"maintenance":       &MaintenanceGenerator{},
		"redirect":          &RedirectGenerator{},
		"settings":          &SettingsGenerator{},
		"tag":               &TagGenerator{},
		"tag_cache_setting": &TagCacheSettingGenerator{},
		"tag_setting":       &TagSettingGenerator{},
		"tag_waf_rule":      &TagWafRuleGenerator{},
		"waf_rule":          &WafRuleGenerator{},
	}
}

// InitService
func (p *MyrasecProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("myrasec: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	return nil
}
