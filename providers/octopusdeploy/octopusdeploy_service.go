package octopusdeploy

import (
	"errors"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type OctopusDeployService struct { //nolint
	terraformutils.Service
}

func (s *OctopusDeployService) Client() (*octopusdeploy.Client, error) {
	octopusURL := s.Args["address"].(string)
	octopusAPIKey := s.Args["api_key"].(string)

	if octopusURL == "" || octopusAPIKey == "" {
		err := errors.New("Please make sure to set the env variables 'OCTOPUS_CLI_SERVER' and 'OCTOPUS_CLI_API_KEY'")
		return nil, err
	}

	httpClient := http.Client{}
	client := octopusdeploy.NewClient(&httpClient, octopusURL, octopusAPIKey)

	return client, nil
}
