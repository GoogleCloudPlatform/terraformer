package squadcast

import (
	"context"
	"encoding/json"
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

// @desc function used for GetAccessToken, GetTeamID, GetServiceID
func Request [TRes any] (url string, header map[string]string) (*TRes, error) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(req)

	var response struct {
		Data *TRes `json:"data"`
		*Meta
	}

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

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// @desc function used for other APIs
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
