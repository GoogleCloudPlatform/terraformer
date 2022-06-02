package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// MaintenanceGenerator
//
type MaintenanceGenerator struct {
	MyrasecService
}

//
// createMaintenanceResources
//
func (g *MaintenanceGenerator) createMaintenanceResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}

	for {
		params["page"] = strconv.Itoa(page)

		maintenance, err := api.ListMaintenances(domainId, vhost.Label, params)
		if err != nil {
			return err
		}

		for _, m := range maintenance {
			r := terraformutils.NewResource(
				strconv.Itoa(m.ID),
				fmt.Sprintf("%s_%d", vhost.Label, m.ID),
				"myrasec_maintenance",
				"myrasec",
				map[string]string{
					"subdomain_name": vhost.Label,
				},
				[]string{},
				map[string]interface{}{},
			)
			g.Resources = append(g.Resources, r)
		}
		if len(maintenance) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// InitResources
//
func (g *MaintenanceGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createMaintenanceResources,
	}

	err = createResourcesPerSubDomain(api, funcs, &wg, false)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
