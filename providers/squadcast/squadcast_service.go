package squadcast

import (
	"context"
	"errors"
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

func (s *SquadcastService) generateRequest(uri string) ([]byte, error) {
	var host string
	switch s.Args["region"] {
	case "us":
		host = "squadcast.com"
	case "eu":
		host = "eu.squadcast.com"
	case "internal":
		host = "squadcast.xyz"
	case "staging":
		host = "squadcast.tech"
	case "dev":
		host = "localhost"
	default:
		return nil, errors.New("unknown region")
	}

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.%s%s",host,uri), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Args["access_token"]))
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
