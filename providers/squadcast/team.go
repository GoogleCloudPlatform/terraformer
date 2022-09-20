package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	SquadcastService
}

type Team struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Members []*TeamMember `json:"members"`
	Roles   []*TeamRole   `json:"roles"`
}

func (g *TeamGenerator) createResources(teams []Team) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, team := range teams {
		resourceList = append(resourceList, terraformutils.NewSimpleResource(
			team.ID,
			fmt.Sprintf("team_%s", team.Name),
			"squadcast_team",
			g.GetProviderName(),
			[]string{},
		))
	}
	return resourceList
}

func (g *TeamGenerator) InitResources() error {
	getTeamsURL := "/v3/teams"
	response, err := Request[[]Team](getTeamsURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
