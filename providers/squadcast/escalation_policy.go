package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type EscalationPolicyGenerator struct {
	SquadcastService
	teamID string
}

var responseEscalationPolicy struct {
	Data *[]EscalationPolicy `json:"data"`
}

type EscalationPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EscalationPolicies []EscalationPolicy

func (g *EscalationPolicyGenerator) createResources(policies EscalationPolicies) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, policy := range policies {
		resourceList = append(resourceList, terraformutils.NewResource(
			policy.ID,
			"policy_"+(policy.ID),
			"squadcast_escalation_policy",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.teamID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resourceList
}

func (g *EscalationPolicyGenerator) InitResources() error {
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}
	team, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}
	err = json.Unmarshal(team, &ResponseTeam)
	if err != nil {
		return err
	}
	g.teamID = ResponseTeam.Data.ID

	body, err := g.generateRequest(fmt.Sprintf("/v3/escalation-policies?owner_id=%s", g.teamID))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &responseEscalationPolicy)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseEscalationPolicy.Data)

	return nil
}
