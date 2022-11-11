package googleworkspace

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	directory "google.golang.org/api/admin/directory/v1"
)

type OrgUnitGenerator struct {
	GoogleWorkspaceService
}

func (g *OrgUnitGenerator) InitResources() error {
	orgUnitList, err := g.getAllOrgUnits()
	if err != nil {
		return err
	}

	g.Resources = g.createResources(orgUnitList)
	return nil
}

func (g OrgUnitGenerator) createResources(orgUnits []*directory.OrgUnit) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, orgUnit := range orgUnits {
		resourceName := g.EnsureStringRandomness("org_unit_" + orgUnit.Name)
		resources = append(
			resources,
			terraformutils.NewResource(
				orgUnit.OrgUnitId,
				resourceName,
				"googleworkspace_org_unit",
				"googleworkspace",
				map[string]string{},
				[]string{},
				map[string]interface{}{},
			),
		)
	}
	return resources
}
