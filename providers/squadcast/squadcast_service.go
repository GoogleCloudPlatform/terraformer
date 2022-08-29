package squadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadcastService struct {
	terraformutils.Service
}

var ResponseService struct {
	Data *Service `json:"data"`
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
	url := fmt.Sprintf("https://api.%s%s", host, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	accessToken := fmt.Sprintf("Bearer %s", s.Args["access_token"])

	req.Header.Set("Authorization", accessToken)
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

func (s *SquadcastService) getServiceByName(teamID string, name string) (Service, error) {
	body, err := s.generateRequest(fmt.Sprintf("/v3/services/by-name?name=%s&owner_id=%s", url.QueryEscape(name), teamID))
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &ResponseService)
	if err != nil {
		log.Fatal(err)
	}

	return *ResponseService.Data, nil
}
