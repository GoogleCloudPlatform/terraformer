package squadcast

import (
	"errors"
	"fmt"
	"log"
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
	teamID       string
	teamName     string
	serviceName  string
	serviceID    string
	scheduleName string
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

	if args[3] != "" {
		p.serviceName = args[3]
		p.GetServiceID()
	}

	if args[4] != "" {
		p.scheduleName = args[4]
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
		"access_token": p.accesstoken,
		"region":       p.region,
		"team_id":      p.teamID,
		"team_name":    p.teamName,
		"service_name": p.serviceName,
		"service_id":   p.serviceID,
		"schedule_name":   p.scheduleName,
	})
	return nil
}

// @desc: GetConfig is used to send details to provider block of terraform-provider-squadcast

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
		"user":                &UserGenerator{},
		"service":             &ServiceGenerator{},
		"squad":               &SquadGenerator{},
		"team":                &TeamGenerator{},
		"team_member":         &TeamMemberGenerator{},
		"team_roles":          &TeamRolesGenerator{},
		"escalation_policy":   &EscalationPolicyGenerator{},
		"runbook":             &RunbookGenerator{},
		"slo":                 &SLOGenerator{},
		"tagging_rules":       &TaggingRulesGenerator{},
		"routing_rules":       &RoutingRulesGenerator{},
		"deduplication_rules": &DeduplicationRulesGenerator{},
		"suppression_rules":   &SuppressionRulesGenerator{},
		"schedule":            &SchedulesGenerator{},
	}
}

func (p *SquadcastProvider) GetAccessToken() {
	host := GetHost(p.region)
	url := fmt.Sprintf("https://auth.%s/oauth/access-token", host)
	header := map[string]string{"X-Refresh-Token": p.refreshtoken}
	response, err := Request[AccessToken](url, header)

	if err != nil {
		log.Fatal(err)
	}
	p.accesstoken = response.AccessToken
}

func (p *SquadcastProvider) GetTeamID() {
	host := GetHost(p.region)
	url := fmt.Sprintf("https://api.%s/v3/teams/by-name?name=%s", host, url.QueryEscape(p.teamName))
	header := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", p.accesstoken)}
	
	response, err := Request[Team](url, header)
	if err != nil {
		log.Fatal(err)
	}
	p.teamID = response.ID
}

func (p *SquadcastProvider) GetServiceID() {
	host := GetHost(p.region)

	url := fmt.Sprintf("https://api.%s/v3/services/by-name?name=%s&owner_id=%s", host, url.QueryEscape(p.serviceName), p.teamID)
	header := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", p.accesstoken)}
	response, err := Request[Service](url, header)

	if err != nil {
		log.Fatal(err)
	}
	p.serviceID = response.ID
}
