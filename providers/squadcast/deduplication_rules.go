package squadcast

import (
	"encoding/json"
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

var getDeduplicationRuleResponse struct {
	Data *DeduplicationRules `json:"data"`
}

func (g *DeduplicationRulesGenerator) createResources(deduplicationRules DeduplicationRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range deduplicationRules.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			rule.ID,
			"deduplication_rule_"+(rule.ID),
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
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/deduplication-rules", g.Args["service_id"]))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &getDeduplicationRuleResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getDeduplicationRuleResponse.Data)

	return nil
}
