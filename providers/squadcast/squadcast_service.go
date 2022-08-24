package squadcast

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadcastService struct {
	terraformutils.Service
}

func GetHost(region string) string {
	switch region {
	case "us":
		return "squadcast.com"
	case "eu":
		return "eu.squadcast.com"
	default:
		return ""
	}
}

func (s *SquadcastService) generateRequest(uri string) ([]byte, error) {
	host := GetHost(s.Args["region"].(string))
	if host == "" {
		log.Fatal("unknown region")
	}
	ctx := context.Background()

	var url string

	api_endpoint := s.Args["api_endpoint"].(string)
	if api_endpoint == "" {
		url = fmt.Sprintf("https://api.%s%s", host, uri)
	} else {
		url = fmt.Sprintf("%s%s", api_endpoint, uri)
	}

	fmt.Println(api_endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	AccessToken := fmt.Sprintf("Bearer %s", s.Args["access_token"])

	req.Header.Set("Authorization", AccessToken)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
