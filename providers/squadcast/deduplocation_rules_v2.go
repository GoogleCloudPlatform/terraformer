package squadcast

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DeduplicationRuleGenerator struct {
	SCService
}

func (g *DeduplicationRuleGenerator) createResources(deduplicationRulesV2 DeduplicationRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range deduplicationRulesV2.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("deduplication_rule_v2_%s", rule.ID),
			"squadcast_deduplication_rule_v2",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.Args["team_id"].(string),
				"service_id": deduplicationRulesV2.ServiceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *DeduplicationRuleGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		getServicesURL := "/v3/services"
		if strings.TrimSpace(g.Args["team_id"].(string)) != "" {
			getServicesURL = fmt.Sprintf("/v3/services?owner_id=%s", g.Args["team_id"].(string))
		}
		req := TRequest{
			URL:             getServicesURL,
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
				URL:             fmt.Sprintf("/v3/services/%s/deduplication-rules", service.ID),
				AccessToken:     g.Args["access_token"].(string),
				Region:          g.Args["region"].(string),
				IsAuthenticated: true,
			}
			response, _, err := Request[DeduplicationRules](req)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		req := TRequest{
			URL:             fmt.Sprintf("/v3/services/%s/deduplication-rules", g.Args["service_id"]),
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		response, _, err := Request[DeduplicationRules](req)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
