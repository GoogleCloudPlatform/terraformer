package squadcast

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DeduplicationRulesGenerator struct {
	SquadcastService
}

type DeduplicationRules struct {
	ID    string               `json:"id"`
	Rules []*DeduplicationRule `json:"rules"`
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
				"service_id": g.Args["service_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *DeduplicationRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		return errors.New("--service-name is required")
	}

	getDeduplicationRulesURL := fmt.Sprintf("/v3/services/%s/deduplication-rules", g.Args["service_id"])
	response, err := Request[DeduplicationRules](getDeduplicationRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
