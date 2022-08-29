package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type SuppressionRulesGenerator struct {
	SquadcastService
	serviceID string
	teamID    string
}

type SuppressionRules struct {
	ID    string             `json:"id"`
	Rules []*SuppressionRule `json:"rules"`
}

type SuppressionRule struct {
	ID string `json:"rule_id"`
}

var getSuppressionRuleResponse struct {
	Data *SuppressionRules `json:"data"`
}

func (g *SuppressionRulesGenerator) createResources(suppressionRules SuppressionRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range suppressionRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			"suppression_rule_"+(rule.ID),
			"squadcast_suppression_rules",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.teamID,
				"service_id": g.serviceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resourceList
}

func (g *SuppressionRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		return errors.New("--service-name is required")
	}
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	team, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}
	err = json.Unmarshal(team, &getTeamResponse)
	if err != nil {
		return err
	}
	g.teamID = getTeamResponse.Data.ID
	service, err := g.getServiceByName(g.teamID, g.Args["service_name"].(string))
	if err != nil {
		return err
	}
	g.serviceID = service.ID
	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/suppression-rules", g.serviceID))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &getSuppressionRuleResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getSuppressionRuleResponse.Data)

	return nil
}
