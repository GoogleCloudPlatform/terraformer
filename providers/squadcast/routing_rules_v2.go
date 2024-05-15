package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RoutingRuleGenerator struct {
	SCService
}

func (g *RoutingRuleGenerator) createResources(routingRules RoutingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range routingRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("routing_rule_v2_%s", rule.ID),
			"squadcast_routing_rule_v2",
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

func (g *RoutingRuleGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		req := TRequest{
			URL:             "/v3/services",
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		responseService, err := Request[[]Service](req)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			req := TRequest{
				URL:             fmt.Sprintf("/v3/services/%s/routing-rules", service.ID),
				AccessToken:     g.Args["access_token"].(string),
				Region:          g.Args["region"].(string),
				IsAuthenticated: true,
			}
			response, err := Request[RoutingRules](req)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		req := TRequest{
			URL:             fmt.Sprintf("/v3/services/%s/routing-rules", g.Args["service_id"]),
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		response, err := Request[RoutingRules](req)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
