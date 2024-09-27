package squadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hasura/go-graphql-client"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SCService struct {
	terraformutils.Service
}

const (
	UserAgent = "terraformer-squadcast"
)

func getHost(region string) string {
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

type TRequest struct {
	URL             string
	AccessToken     string
	RefreshToken    string
	Region          string
	IsAuthenticated bool
	IsV2            bool
}

func Request[TRes any](request TRequest) (*TRes, *Meta, error) {
	ctx := context.Background()
	var URL string
	var req *http.Request
	var err error
	host := getHost(request.Region)
	if request.IsAuthenticated {
		if !request.IsV2 {
			URL = fmt.Sprintf("https://api.%s%s", host, request.URL)
			req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.AccessToken))
		} else {
			URL = fmt.Sprintf("https://platform-backend.%s%s", host, request.URL)
			req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.AccessToken))
		}
	} else {
		URL = fmt.Sprintf("https://auth.%s%s", host, request.URL)
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
		req.Header.Set("X-Refresh-Token", request.RefreshToken)
	}
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(req)

	var response struct {
		Data *TRes `json:"data"`
		Meta *Meta `json:"meta"`
	}
	if err != nil {
		return nil, nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, nil, err
	}
	if response.Meta != nil {
		if response.Meta.Meta.Status >= 400 {
			return nil, nil, fmt.Errorf("error: %s", response.Meta.Meta.Message)
		}
	}

	return response.Data, response.Meta, nil
}

var gqlClient *graphql.Client

func GraphQLRequest[TReq any](method string, token string, region string, payload *TReq, variables map[string]interface{}) (*TReq, error) {
	graphQLURL := fmt.Sprintf("https://api.%s/v3/graphql", getHost(region))
	bearerToken := fmt.Sprintf("Bearer %s", token)

	if gqlClient == nil {
		gqlClient = graphql.NewClient(graphQLURL, nil).WithRequestModifier(func(req *http.Request) {
			req.Header.Set("Authorization", bearerToken)
		})
	}
	if err := gqlClient.WithDebug(false).Query(context.Background(), payload, variables); err != nil {
		return nil, err
	}

	return payload, nil
}
