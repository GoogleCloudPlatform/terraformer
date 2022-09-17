package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SuppressionRulesGenerator struct {
	SquadcastService
}

type SuppressionRules struct {
	ID        string             `json:"id"`
	ServiceID string             `json:"service_id"`
	Rules     []*SuppressionRule `json:"rules"`
}

type SuppressionRule struct {
	ID string `json:"rule_id"`
}

func (g *SuppressionRulesGenerator) createResources(suppressionRules SuppressionRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range suppressionRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("suppression_rule_%s", rule.ID),
			"squadcast_suppression_rules",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.Args["team_id"].(string),
				"service_id": suppressionRules.ServiceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *SuppressionRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		getServicesURL := "/v3/services"
		responseService, err := Request[[]Service](getServicesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			getSuppressionRulesURL := fmt.Sprintf("/v3/services/%s/suppression-rules", service.ID)
			response, err := Request[SuppressionRules](getSuppressionRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		getSuppressionRulesURL := fmt.Sprintf("/v3/services/%s/suppression-rules", g.Args["service_id"])
		response, err := Request[SuppressionRules](getSuppressionRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
