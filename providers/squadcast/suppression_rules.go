package squadcast

import (
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

	getSuppressionRulesURL := fmt.Sprintf("/v3/services/%s/suppression-rules", g.Args["service_id"])
	response, err := Request[SuppressionRules](getSuppressionRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
