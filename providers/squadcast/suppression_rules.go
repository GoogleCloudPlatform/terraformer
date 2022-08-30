package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SuppressionRulesGenerator struct {
	SquadcastService
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
				"team_id":    g.Args["team_id"].(string),
				"service_id": g.Args["service_id"].(string),
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

	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/suppression-rules", g.Args["service_id"]))
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
