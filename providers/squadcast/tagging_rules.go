package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type TaggingRulesGenerator struct {
	SquadcastService
	serviceID string
	teamID    string
}

type TaggingRules struct {
	ID        string         `json:"id"`
	ServiceID string         `json:"service_id"`
	Rules     []*TaggingRule `json:"rules"`
}
type TaggingRule struct {
	IsBasic    bool   `json:"is_basic"`
	Expression string `json:"expression"`
}

var responseTaggingRule struct {
	Data *TaggingRules `json:"data"`
}

func (g *TaggingRulesGenerator) createResources(taggingRule TaggingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range taggingRule.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			g.teamID+"_"+rule.Expression,
			"tagging_rule_"+(rule.Expression),
			"squadcast_tagging_rules",
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

func (g *TaggingRulesGenerator) InitResources() error {
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
	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/tagging-rules", g.serviceID))
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body, &responseTaggingRule)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseTaggingRule.Data)

	return nil
}
