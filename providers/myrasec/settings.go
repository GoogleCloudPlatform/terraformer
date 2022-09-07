package myrasec

import (
	"fmt"
	"strconv"
	"sync"

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
func (g *SettingsGenerator) createSettingResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	params := map[string]string{}

	s, err := api.ListSettings(domainId, vhost.Label, params)
	if err != nil {
		return err
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
	g.Resources = append(g.Resources, r)
	return nil
}

//
// InitResources
//
func (g *SettingsGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return nil
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createSettingResources,
	}

	err = createResourcesPerSubDomain(api, funcs, &wg, true)
	if err != nil {
		return nil
	}

	wg.Wait()

	return nil
}
