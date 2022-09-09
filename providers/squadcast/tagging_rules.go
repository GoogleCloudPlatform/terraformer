package squadcast

import (
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

	getTaggingRulesURL := fmt.Sprintf("/v3/services/%s/tagging-rules", g.Args["service_id"])
	response, err := Request[TaggingRules](getTaggingRulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
