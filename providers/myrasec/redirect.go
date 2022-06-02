package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// RedirectGenerator
//
type RedirectGenerator struct {
	MyrasecService
}

//
// createRedirectResources
//
func (g *RedirectGenerator) createRedirectResources(api *mgo.API, domainId int, vhost mgo.VHost, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		redirects, err := api.ListRedirects(domainId, vhost.Label, params)
		if err != nil {
			return err
		}

		for _, redirect := range redirects {
			r := terraformutils.NewResource(
				strconv.Itoa(redirect.ID),
				fmt.Sprintf("%s_%d", redirect.SubDomainName, redirect.ID),
				"myrasec_redirect",
				"myrasec",
				map[string]string{
					"subdomain_name": redirect.SubDomainName,
				},
				[]string{},
				map[string]interface{}{},
			)
			g.Resources = append(g.Resources, r)
		}
		if len(redirects) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// InitResources
//
func (g *RedirectGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error{
		g.createRedirectResources,
	}
	err = createResourcesPerSubDomain(api, funcs, &wg, true)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
