package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RoutingRulesGenerator struct {
	SquadcastService
}

type RoutingRules struct {
	ID        string         `json:"id"`
	ServiceID string         `json:"service_id"`
	Rules     []*RoutingRule `json:"rules"`
}

type RoutingRule struct {
	ID string `json:"rule_id"`
}

func (g *RoutingRulesGenerator) createResources(routingRules RoutingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range routingRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("routing_rule_%s", rule.ID),
			"squadcast_routing_rules",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.Args["team_id"].(string),
				"service_id": routingRules.ServiceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *RoutingRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		getServicesURL := "/v3/services"
		responseService, err := Request[[]Service](getServicesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			getRoutingRulesURL := fmt.Sprintf("/v3/services/%s/routing-rules", service.ID)
			response, err := Request[RoutingRules](getRoutingRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		getRoutingRulesURL := fmt.Sprintf("/v3/services/%s/routing-rules", g.Args["service_id"])
		response, err := Request[RoutingRules](getRoutingRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
