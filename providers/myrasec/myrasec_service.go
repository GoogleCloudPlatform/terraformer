package myrasec

import (
	"errors"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

// MyrasecService ...
type MyrasecService struct {
	terraformutils.Service
}

// initializeAPI ...
func (s *MyrasecService) initializeAPI() (*mgo.API, error) {
	apiKey := os.Getenv("MYRASEC_API_KEY")
	apiSecret := os.Getenv("MYRASEC_API_SECRET")
	apiURL, urlPresent := os.LookupEnv("MYRASEC_API_BASE_URL")

	if apiKey == "" || apiSecret == "" {
		err := errors.New("missing API credentials")
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	api, err := mgo.New(apiKey, apiSecret)
	if urlPresent {
		api.BaseURL = apiURL
	}
	api.EnableCaching()
	api.SetCachingTTL(3600)

	return api, err
}
