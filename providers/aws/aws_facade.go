package aws

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
)

type AwsFacade struct {
	AWSService
	service terraform_utils.ServiceGenerator
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

func (s *AwsFacade) ParseFilter(rawFilter string) []terraform_utils.ResourceFilter {
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

func (s *AwsFacade) GetResources() []terraform_utils.Resource {
	return s.service.GetResources()
}
func (s *AwsFacade) SetResources(resources []terraform_utils.Resource) {
	s.service.SetResources(resources)
}

func (s *AwsFacade) InitResources() error {
	err := s.service.InitResources()
	if err == nil {
		return nil
	} else {
		message := err.Error()
		if strings.Contains(message, "no such host") || strings.Contains(message, "i/o timeout") { // skip not available AWS services
			return nil
		} else {
			return err
		}
	}
}

func (s *AwsFacade) PostConvertHook() error {
	return s.service.PostConvertHook()
}

func (s *AwsFacade) PopulateIgnoreKeys(providerWrapper *provider_wrapper.ProviderWrapper) {
	s.service.PopulateIgnoreKeys(providerWrapper)
}
