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
			resourceList = append(resourceList, terraformutils.NewResource(
				fmt.Sprintf("%d", rs.ID),
				fmt.Sprintf("ger_ruleset_rules_ordering_%d", rs.ID),
				"squadcast_ger_ruleset_rules_ordering",
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
	return resourceList
}

func (g *GlobalEventRulesGenerator) InitResources() error {
	var allRules []GER
	page := 1
	pageSize := 100

	for {
		req := TRequest{
			URL:             fmt.Sprintf("/v3/global-event-rules?owner_id=%s&page_number=%d&page_size=%d", g.Args["team_id"].(string), page, pageSize),
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}

		response, meta, err := Request[[]GER](req)
		if err != nil {
			return err
		}

		allRules = append(allRules, *response...)
		if page*pageSize >= meta.TotalCount {
			break
		}

		page++
	}

	for i, rule := range allRules {
		for j, rs := range rule.Rulesets {
			req := TRequest{
				URL:             fmt.Sprintf("/v3/global-event-rules/%d/rulesets/%s/%s/rules?page_number=1&page_size=100", rule.ID, rs.AlertSourceVersion, rs.AlertSourceShortName),
				AccessToken:     g.Args["access_token"].(string),
				Region:          g.Args["region"].(string),
				IsAuthenticated: true,
			}
			resp, _, err := Request[[]GER_Ruleset_Rules](req)
			if err != nil {
				return err
			}
			allRules[i].Rulesets[j].Rules = *resp
		}
	}

	g.Resources = g.createResources(allRules)
	return nil
}

type AlertSource struct {
	Type      string `json:"type"`
	ShortName string `json:"shortName"`
	Version   string `json:"version"`
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
	alertSources, _, err := Request[[]AlertSource](req)
	if err != nil {
		return nil, err
	}
	for _, alertSourceData := range *alertSources {
		alertSourcesMap[alertSourceData.ShortName] = alertSourceData.Type
	}
	return alertSourcesMap, nil
}
