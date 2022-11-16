package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DeduplicationRulesGenerator struct {
	SCService
}

type DeduplicationRules struct {
	ID        string               `json:"id"`
	ServiceID string               `json:"service_id"`
	Rules     []*DeduplicationRule `json:"rules"`
}

type DeduplicationRule struct {
	ID string `json:"rule_id"`
}

func (g *DeduplicationRulesGenerator) createResources(deduplicationRules DeduplicationRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range deduplicationRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			fmt.Sprintf("deduplication_rule_%s", rule.ID),
			"squadcast_deduplication_rules",
			g.GetProviderName(),
			map[string]string{
				"team_id":    g.Args["team_id"].(string),
				"service_id": deduplicationRules.ServiceID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *DeduplicationRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		getServicesURL := "/v3/services"
		responseService, err := Request[[]Service](getServicesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			getDeduplicationRulesURL := fmt.Sprintf("/v3/services/%s/deduplication-rules", service.ID)
			response, err := Request[DeduplicationRules](getDeduplicationRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		getDeduplicationRulesURL := fmt.Sprintf("/v3/services/%s/deduplication-rules", g.Args["service_id"])
		response, err := Request[DeduplicationRules](getDeduplicationRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
