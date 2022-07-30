package googleworkspace

import (
	"context"
	"io/ioutil"
	"math/rand"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"golang.org/x/oauth2/google"
	directory "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/chromepolicy/v1"
	"google.golang.org/api/option"
)

type GoogleWorkspaceService struct {
	terraformutils.Service

	orgID string
}

func (s *GoogleWorkspaceService) getCredentialJson() ([]byte, error) {
	return ioutil.ReadFile(s.Args["credential_json_filepath"].(string))
}

func (s *GoogleWorkspaceService) setDefaults() {
	s.orgID = s.Args["org_id"].(string)
}

func (s *GoogleWorkspaceService) ChromePolicyClient() (*chromepolicy.Service, error) {
	s.setDefaults()

	credentialJson, err := s.getCredentialJson()
	if err != nil {
		return nil, err
	}

	client, err := chromepolicy.NewService(
		context.Background(),
		option.WithCredentialsJSON(credentialJson))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *GoogleWorkspaceService) DirectoryClient() (*directory.Service, error) {
	s.setDefaults()

	credentialJson, err := s.getCredentialJson()
	if err != nil {
		return nil, err
	}

	auth, err := google.JWTConfigFromJSON(credentialJson, directory.AdminDirectoryOrgunitScope)
	if err != nil {
		return nil, err
	}
	auth.Subject = s.Args["impersonated_user_email"].(string)

	ts := auth.TokenSource(context.Background())

	client, err := directory.NewService(context.Background(), option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return client, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (s *GoogleWorkspaceService) EnsureStringRandomness(input string) string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return input + "--" + string(b)
}
