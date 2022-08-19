package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type EscalationPolicyGenerator struct {
	SquadcastService
	teamID string
}

type EscalationPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EscalationPolicies []EscalationPolicy

var getEscalationPolicyResponse struct {
	Data *[]EscalationPolicy `json:"data"`
}

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
	teamName := g.Args["team_name"].(string)
	if len(teamName) == 0 {
		return errors.New("--team-name is required")
	}

	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", escapedTeamName)

	team, err := g.generateRequest(getTeamURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(team, &getTeamResponse)
	if err != nil {
		return err
	}

	g.teamID = getTeamResponse.Data.ID

	getEscalationPolicyURL := fmt.Sprintf("/v3/escalation-policies?owner_id=%s", g.teamID)
	body, err := g.generateRequest(getEscalationPolicyURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &getEscalationPolicyResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getEscalationPolicyResponse.Data)

	return nil
}
