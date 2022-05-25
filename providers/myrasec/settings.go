package myrasec

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// SettingGenerator
//
type SettingsGenerator struct {
	MyrasecService
}

//
// createSettingResources
//
func (g *SettingsGenerator) createSettingResources(api *mgo.API, domainId int, vhost mgo.VHost) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	params := map[string]string{}

	s, err := api.ListSettings(domainId, vhost.Label, params)
	if err != nil {
		return nil, err
	}

	r := terraformutils.NewResource(
		strconv.Itoa(vhost.ID),
		fmt.Sprintf("%s_%d", vhost.Label, vhost.ID),
		"myrasec_settings",
		"myrasec",
		map[string]string{
			"subdomain_name": vhost.Label,
			"only_https":     strconv.FormatBool(s.OnlyHTTPS),
		},
		[]string{},
		map[string]interface{}{},
	)
	resources = append(resources, r)
	return resources, nil
}

//
// InitResources
//
func (g *SettingsGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return nil
	}

	funcs := []func(*mgo.API, int, mgo.VHost) ([]terraformutils.Resource, error){
		g.createSettingResources,
	}

	res, err := createResourcesPerSubDomain(api, funcs, false)
	if err != nil {
		return nil
	}

	g.Resources = res

	return nil
}
