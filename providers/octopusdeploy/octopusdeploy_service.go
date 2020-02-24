package octopusdeploy

import (
	"errors"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type OctopusDeployService struct {
	terraform_utils.Service
}

func (s *OctopusDeployService) Client() (*octopusdeploy.Client, error) {
	octopusURL := s.Args["server"].(string)
	octopusAPIKey := s.Args["apikey"].(string)

	if octopusURL == "" || octopusAPIKey == "" {
		err := errors.New("Please make sure to set the env variables 'OCTOPUS_CLI_SERVER' and 'OCTOPUS_CLI_API_KEY'")
		return nil, err
	}

	httpClient := http.Client{}
	client := octopusdeploy.NewClient(&httpClient, octopusURL, octopusAPIKey)

	return client, nil
}
