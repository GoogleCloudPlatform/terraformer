package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TaggingRulesGenerator struct {
	SquadcastService
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
		getServicesURL := "/v3/services"
		responseService, err := Request[[]Service](getServicesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		for _, service := range *responseService {
			getTaggingRulesURL := fmt.Sprintf("/v3/services/%s/tagging-rules", service.ID)
			response, err := Request[TaggingRules](getTaggingRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
			if err != nil {
				return err
			}

			g.Resources = append(g.Resources, g.createResources(*response)...)
		}
	} else {
		getTaggingRulesURL := fmt.Sprintf("/v3/services/%s/tagging-rules", g.Args["service_id"])
		response, err := Request[TaggingRules](getTaggingRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
		if err != nil {
			return err
		}

		g.Resources = g.createResources(*response)
	}
	return nil
}
