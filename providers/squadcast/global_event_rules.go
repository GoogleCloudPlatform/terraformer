package squadcast

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GlobalEventRulesGenerator struct {
	SCService
}

type GER struct {
	ID       uint          `json:"id"`
	Rulesets []GER_Ruleset `json:"rulesets"`
}
type GER_Ruleset struct {
	ID                   uint `json:"id"`
	AlertSourceName      string
	AlertSourceShortName string `json:"alert_source_shortname"`
	AlertSourceVersion   string `json:"alert_source_version"`

	Rules []GER_Ruleset_Rules
}
type GER_Ruleset_Rules struct {
	ID     uint `json:"id"`
	GER_ID uint `json:"global_event_rule_id"`
}

func (g *GlobalEventRulesGenerator) createResources(ger []GER) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	alertSourcesMap, err := g.getAlertSources()
	if err != nil {
		log.Fatal(err)
	}
	for _, rule := range ger {
		resourceList = append(resourceList, terraformutils.NewResource(
			fmt.Sprintf("%d", rule.ID),
			fmt.Sprintf("ger_%d", rule.ID),
			"squadcast_ger",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
		for _, rs := range rule.Rulesets {
			resourceList = append(resourceList, terraformutils.NewResource(
				fmt.Sprintf("%d", rs.ID),
				fmt.Sprintf("ger_ruleset_%d", rs.ID),
				"squadcast_ger_ruleset",
				g.GetProviderName(),
				map[string]string{
					"ger_id":                 fmt.Sprintf("%d", rule.ID),
					"alert_source":           alertSourcesMap[rs.AlertSourceShortName],
					"alert_source_shortname": rs.AlertSourceShortName,
					"alert_source_version":   rs.AlertSourceVersion,
				},
				[]string{},
				map[string]interface{}{},
			))

			for _, r := range rs.Rules {
				resourceList = append(resourceList, terraformutils.NewResource(
					fmt.Sprintf("%d", r.ID),
					fmt.Sprintf("ger_ruleset_rule_%d", r.ID),
					"squadcast_ger_ruleset_rule",
					g.GetProviderName(),
					map[string]string{
						"ger_id":                 fmt.Sprintf("%d", rule.ID),
						"alert_source":           alertSourcesMap[rs.AlertSourceShortName],
						"alert_source_shortname": rs.AlertSourceShortName,
						"alert_source_version":   rs.AlertSourceVersion,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}
	return resourceList
}

func (g *GlobalEventRulesGenerator) InitResources() error {
	req := TRequest{
		URL:             fmt.Sprintf("/v3/global-event-rules?owner_id=%s", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, err := Request[[]GER](req)
	if err != nil {
		return err
	}

	ger := *response

	for i, rule := range ger {
		for j, rs := range rule.Rulesets {
			req := TRequest{
				URL:             fmt.Sprintf("/v3/global-event-rules/%d/rulesets/%s/%s/rules", rule.ID, rs.AlertSourceVersion, rs.AlertSourceShortName),
				AccessToken:     g.Args["access_token"].(string),
				Region:          g.Args["region"].(string),
				IsAuthenticated: true,
			}
			resp, err := Request[[]GER_Ruleset_Rules](req)
			if err != nil {
				return err
			}
			ger[i].Rulesets[j].Rules = *resp
		}
	}

	g.Resources = g.createResources(*response)
	return nil
}

type AlertSource struct {
	ID             string `json:"_id"`
	Type           string `json:"type"`
	Heading        string `json:"heading"`
	SupportDocURL  string `json:"supportDoc"`
	DisplayKeyOnly bool   `json:"displayKeyOnly"`
	ShortName      string `json:"shortName"`
	Version        string `json:"version"`

	IsValid      bool `json:"isValid"`
	IsPrivate    bool `json:"isPrivate"`
	IsDeprecated bool `json:"deprecated"`
}

func (g *GlobalEventRulesGenerator) getAlertSources() (map[string]string, error) {
	alertSourcesMap := make(map[string]string, 0)
	req := TRequest{
		URL:             "/v2/public/integrations",
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
		IsV2:            true,
	}
	alertSources, err := Request[[]AlertSource](req)
	if err != nil {
		return nil, err
	}
	for _, alertSourceData := range *alertSources {
		alertSourcesMap[alertSourceData.ShortName] = alertSourceData.Type
	}
	return alertSourcesMap, nil
}
