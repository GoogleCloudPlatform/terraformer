package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	SCService
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
	req := TRequest{
		URL:             "/v3/teams",
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]Team](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
