package opsgenie

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type OpsgenieService struct { //nolint
	terraformutils.Service
}

func (s *OpsgenieService) UserClient() (*user.Client, error) {
	return user.NewClient(&client.Config{ApiKey: s.GetArgs()["api-key"].(string)})
}

func (s *OpsgenieService) TeamClient() (*team.Client, error) {
	return team.NewClient(&client.Config{ApiKey: s.GetArgs()["api-key"].(string)})
}
