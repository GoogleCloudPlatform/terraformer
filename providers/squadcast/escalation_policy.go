package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type EscalationPolicyGenerator struct {
	SquadcastService
}

type EscalationPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var getEscalationPolicyResponse struct {
	Data *[]EscalationPolicy `json:"data"`
}

func (g *EscalationPolicyGenerator) createResources(policies []EscalationPolicy) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, policy := range policies {
		resourceList = append(resourceList, terraformutils.NewResource(
			policy.ID,
			"policy_"+(policy.ID),
			"squadcast_escalation_policy",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resourceList
}

func (g *EscalationPolicyGenerator) InitResources() error {
	teamID := g.Args["team_id"].(string)
	if teamID == "" {
		return errors.New("--team-name is required")
	}
	getEscalationPolicyURL := fmt.Sprintf("/v3/escalation-policies?owner_id=%s", teamID)
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
