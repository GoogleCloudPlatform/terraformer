package squadcast

import (
	"encoding/json"
	"errors"
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

var getTaggingRuleResponse struct {
	Data *TaggingRules `json:"data"`
}

func (g *TaggingRulesGenerator) createResources(taggingRule TaggingRules) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, rule := range taggingRule.Rules {
		resourceList = append(resourceList, terraformutils.NewResource(
			g.Args["team_id"].(string)+"_"+rule.ID,
			"tagging_rule_"+(rule.ID),
			"squadcast_tagging_rules",
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

func (g *TaggingRulesGenerator) InitResources() error {
	if len(g.Args["service_name"].(string)) == 0 {
		return errors.New("--service-name is required")
	}
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	body, err := g.generateRequest(fmt.Sprintf("/v3/services/%s/tagging-rules", g.Args["service_id"]))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &getTaggingRuleResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getTaggingRuleResponse.Data)

	return nil
}
