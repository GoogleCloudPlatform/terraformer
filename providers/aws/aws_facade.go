package aws

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
)

type AwsFacade struct { //nolint
	AWSService
	service terraformutils.ServiceGenerator
}

func (s *AwsFacade) SetProviderName(providerName string) {
	s.service.SetProviderName(providerName)
}

func (s *AwsFacade) SetVerbose(verbose bool) {
	s.service.SetVerbose(verbose)
}

func (s *AwsFacade) ParseFilters(rawFilters []string) {
	s.service.ParseFilters(rawFilters)
}

func (s *AwsFacade) ParseFilter(rawFilter string) []terraformutils.ResourceFilter {
	return s.service.ParseFilter(rawFilter)
}

func (s *AwsFacade) SetName(name string) {
	s.service.SetName(name)
}
func (s *AwsFacade) GetName() string {
	return s.service.GetName()
}

func (s *AwsFacade) InitialCleanup() {
	s.service.InitialCleanup()
}

func (s *AwsFacade) PostRefreshCleanup() {
	s.service.PostRefreshCleanup()
}

func (s *AwsFacade) GetArgs() map[string]interface{} {
	return s.service.GetArgs()
}
func (s *AwsFacade) SetArgs(args map[string]interface{}) {
	s.service.SetArgs(args)
}

func (s *AwsFacade) GetResources() []terraformutils.Resource {
	return s.service.GetResources()
}
func (s *AwsFacade) SetResources(resources []terraformutils.Resource) {
	s.service.SetResources(resources)
}

func (s *AwsFacade) InitResources() error {
	err := s.service.InitResources()
	if err == nil {
		return nil
	}
	message := err.Error()
	if strings.Contains(message, "no such host") || strings.Contains(message, "i/o timeout") ||
		strings.Contains(message, "x509: certificate is valid for") ||
		strings.Contains(message, "Unavailable Operation") { // skip not available AWS services
		return nil
	}
	return err
}

func (s *AwsFacade) PostConvertHook() error {
	return s.service.PostConvertHook()
}

func (s *AwsFacade) PopulateIgnoreKeys(providerWrapper *providerwrapper.ProviderWrapper) {
	s.service.PopulateIgnoreKeys(providerWrapper)
}
