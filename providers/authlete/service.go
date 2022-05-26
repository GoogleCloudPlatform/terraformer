package authlete

import (
	"context"
	"errors"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	authlete "github.com/authlete/openapi-for-go"
)

type ServiceGenerator struct {
	AuthleteService
}

func (s *ServiceGenerator) InitResources() error {

	if s.Resources == nil {
		s.Resources = []terraformutils.Resource{}
	}

	authleteClient := s.getClient()

	auth := context.WithValue(context.Background(), authlete.ContextBasicAuth, authlete.BasicAuth{
		UserName: s.GetArgs()["service_owner_key"].(string),
		Password: s.GetArgs()["service_owner_secret"].(string),
	})

	end := int32(0)
	total := int32(10)
	for end < total {
		listServices, _, err := authleteClient.ServiceManagementApi.ServiceGetListApi(auth).Start(end).End(total).Execute()
		if err != nil {
			return errors.New("could not fetch the service list:  " + err.Error())
		}
		total = *listServices.TotalCount
		end = *listServices.End
		services := listServices.GetServices()
		s.Resources = append(s.Resources, mapInstances(services)...)
	}

	return nil
}

func mapInstances(services []authlete.Service) []terraformutils.Resource {

	result := []terraformutils.Resource{}
	for _, service := range services {
		newResource := terraformutils.NewResource(
			strconv.FormatInt(service.GetApiKey(), 10),
			strconv.FormatInt(service.GetApiKey(), 10),
			"authlete_service",
			"authlete",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		result = append(result, newResource)
	}
	return result
}
