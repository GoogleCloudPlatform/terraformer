package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TaggingRulesGenerator struct {
	SCService
}

type TaggingRules struct {
	ID        string         `json:"id"`
	ServiceID string         `json:"service_id"`
	Rules     []*TaggingRule `json:"rules"`
}

type TaggingRule struct {
	ID string `json:"rule_id"`
}

func (g *TaggingRulesGenerator) createResources(taggingRule TaggingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range taggingRule.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("tagging_rule_%s", rule.ID),
			"squadcast_tagging_rules",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.Args["team_id"].(string),
				"service_id": taggingRule.ServiceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *TaggingRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		req := TRequest{
			URL:             "/v3/services",
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		responseService, _, err := Request[[]Service](req)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			req := TRequest{
				URL:             fmt.Sprintf("/v3/services/%s/tagging-rules", service.ID),
				AccessToken:     g.Args["access_token"].(string),
				Region:          g.Args["region"].(string),
				IsAuthenticated: true,
			}
			response, _, err := Request[TaggingRules](req)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		req := TRequest{
			URL:             fmt.Sprintf("/v3/services/%s/tagging-rules", g.Args["service_id"]),
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		response, _, err := Request[TaggingRules](req)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
