package squadcast

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadcastProvider struct {
	terraformutils.Provider
	accesstoken  string
	refreshtoken string
	region       string
	teamID     	 string
	teamName 	 string
}

type AccessToken struct {
	Type         string `json:"type"`
	AccessToken  string `json:"access_token"`
	IssuedAt     int64  `json:"issued_at"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

// Meta holds the status of the request information
type Meta struct {
	Meta AppError `json:"meta,omitempty"`
}

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"error_message,omitempty"`
}

const (
	UserAgent = "terraformer-squadcast"
)

func (p *SquadcastProvider) Init(args []string) error {

	if refreshToken := os.Getenv("SQUADCAST_REFRESH_TOKEN"); refreshToken != "" {
		p.refreshtoken = os.Getenv("SQUADCAST_REFRESH_TOKEN")
	}
	if args[0] != "" {
		p.refreshtoken = args[0]
	}
	if p.refreshtoken == "" {
		return errors.New("required refresh Token missing")
	}

	if region := os.Getenv("SQUADCAST_REGION"); region != "" {
		p.region = os.Getenv("SQUADCAST_REGION")
	}
	if args[1] == "" {
		return errors.New("required region missing")
	}
	p.region = args[1]
	p.GetAccessToken()

	if args[2] != "" {
		p.teamName = args[2]
		p.GetTeamID()
	}

	return nil
}

func (p *SquadcastProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	// SetArgs are used for fetching details within other files in the terraformer code.
	p.Service.SetArgs(map[string]interface{}{
		"access_token":  p.accesstoken,
		"region":        p.region,
		"team_id":       p.teamID,
		"team_name":     p.teamName,
	})
	return nil
}

// @desc GetConfig is used to send details to provider block of terraform-provider-squadcast

func (p *SquadcastProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"region":        cty.StringVal(p.region),
		"refresh_token": cty.StringVal(p.refreshtoken),
	})
}

func (p *SquadcastProvider) GetProviderData(...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"squadcast": map[string]interface{}{
				"region": p.region,
			},
		},
	}
}

func (p *SquadcastProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *SquadcastProvider) GetName() string {
	return "squadcast"
}

func (p *SquadcastProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"user":              &UserGenerator{},
		"service":           &ServiceGenerator{},
		"squad":             &SquadGenerator{},
		"team":              &TeamGenerator{},
		"team_member":       &TeamMemberGenerator{},
		"team_roles":        &TeamRolesGenerator{},
		"escalation_policy": &EscalationPolicyGenerator{},
		"runbook":           &RunbookGenerator{},
		"slo":               &SLOGenerator{},
	}
}


func (p *SquadcastProvider) GetAccessToken() {
	host := GetHost(p.region)
	ctx := context.Background()

	url := fmt.Sprintf("https://auth.%s/oauth/access-token", host)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,url , nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Refresh-Token", p.refreshtoken)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		Data AccessToken `json:"data"`
		*Meta
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode > 299 {
		log.Fatal(err)
	}

	p.accesstoken = response.Data.AccessToken
}

func (p *SquadcastProvider) GetTeamID() {
	host := GetHost(p.region)
	ctx := context.Background()

	url := fmt.Sprintf("https://api.%s/v3/teams/by-name?name=%s", host, url.QueryEscape(p.teamName))
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	accessToken := fmt.Sprintf("Bearer %s", p.accesstoken)

	req.Header.Set("Authorization", accessToken)
	req.Header.Set("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		Data Team `json:"data"`
		*Meta
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode > 299 {
		log.Fatal(err)
	}

	p.teamID = response.Data.ID
}