package authlete

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	authlete "github.com/authlete/openapi-for-go"
)

type AuthleteService struct { // nolint
	terraformutils.Service
}

func (s *AuthleteService) getClient() *authlete.APIClient {

	cnf := authlete.NewConfiguration()
	cnf.UserAgent = "terraformer-authlete"
	cnf.Servers[0].URL = s.GetArgs()["api_server"].(string)

	return authlete.NewAPIClient(cnf)
}
