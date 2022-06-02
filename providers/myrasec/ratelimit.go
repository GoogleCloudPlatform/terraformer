package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// RateLimitGenerator
//
type RatelimitGenerator struct {
	MyrasecService
}

//
// createRatelimitResources
//
func (g *RatelimitGenerator) createRatelimitResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}

	for {
		params["page"] = strconv.Itoa(page)

		ratelimits, err := api.ListRateLimits(domainId, vhost.Label, params)
		if err != nil {
			return err
		}

		for _, rl := range ratelimits {
			r := terraformutils.NewResource(
				strconv.Itoa(rl.ID),
				fmt.Sprintf("%s_%d", vhost.Label, rl.ID),
				"myrasec_ratelimit",
				"myrasec",
				map[string]string{
					"subdomain_name": rl.SubDomainName,
				},
				[]string{},
				map[string]interface{}{},
			)
			g.Resources = append(g.Resources, r)
		}
		if len(ratelimits) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// InitResources
//
func (g *RatelimitGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createRatelimitResources,
	}

	err = createResourcesPerSubDomain(api, funcs, &wg, true)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
