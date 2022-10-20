package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// CacheSettingGenerator
//
type CacheSettingGenerator struct {
	MyrasecService
}

//
// createCacheSettingResources
//
func (g *CacheSettingGenerator) createCacheSettingResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		settings, err := api.ListCacheSettings(domainId, vhost.Label, params)

		if err != nil {
			return err
		}

		for _, s := range settings {
			r := terraformutils.NewResource(
				strconv.Itoa(s.ID),
				fmt.Sprintf("%s_%d", vhost.Label, s.ID),
				"myrasec_cache_setting",
				"myrasec",
				map[string]string{
					"subdomain_name": vhost.Label,
				},
				[]string{},
				map[string]interface{}{},
			)
			r.IgnoreKeys = append(r.IgnoreKeys, "^Metadata")
			g.Resources = append(g.Resources, r)
		}
		if len(settings) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// InitResources
//
func (g *CacheSettingGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createCacheSettingResources,
	}
	err = createResourcesPerSubDomain(api, funcs, &wg, true)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
