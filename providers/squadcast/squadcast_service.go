package squadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SCService struct {
	terraformutils.Service
}

const (
	UserAgent = "terraformer-squadcast"
)

func GetHost(region string) string {
	switch region {
	case "us":
		return "squadcast.com"
	case "eu":
		return "eu.squadcast.com"
	case "staging":
		return "squadcast.tech"
	default:
		return ""
	}
}

func Request[TRes any](url string, token string, region string, isAuthenticated bool) (*TRes, error) {
	ctx := context.Background()
	var URL string
	var req *http.Request
	var err error
	host := GetHost(region)
	if isAuthenticated {
		URL = fmt.Sprintf("https://api.%s%s", host, url)
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	} else {
		URL = fmt.Sprintf("https://auth.%s%s", host, url)
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
		req.Header.Set("X-Refresh-Token", token)
	}
	if err != nil {
		log.Fatal(err)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
