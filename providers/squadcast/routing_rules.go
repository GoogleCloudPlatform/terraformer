package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type RoutingRulesGenerator struct {
	SquadcastService
	serviceID string
	teamID    string
}

type RoutingRules struct {
	ID    string         `json:"id"`
	Rules []*RoutingRule `json:"rules"`
}

type RoutingRule struct {
	ID string `json:"rule_id"`
}

var getRoutingRuleResponse struct {
	Data *RoutingRules `json:"data"`
}

func (g *RoutingRulesGenerator) createResources(routingRules RoutingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range routingRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			"routing_rule_"+(rule.ID),
			"squadcast_routing_rules",
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

func (g *RoutingRulesGenerator) InitResources() error {
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
	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/routing-rules", g.serviceID))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &getRoutingRuleResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getRoutingRuleResponse.Data)

	return nil
}
