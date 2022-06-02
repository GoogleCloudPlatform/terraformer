package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// ErrorPageGenerator
//
type ErrorPageGenerator struct {
	MyrasecService
}

//
// createErrorPageResources
//
func (g *ErrorPageGenerator) createErrorPageResources(api *mgo.API, domain mgo.Domain, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		pages, err := api.ListErrorPages(domain.ID, params)
		if err != nil {
			return err
		}

		for _, p := range pages {
			r := terraformutils.NewResource(
				strconv.Itoa(p.ID),
				fmt.Sprintf("%s_%d", p.SubDomainName, p.ID),
				"myrasec_error_page",
				"myrasec",
				map[string]string{
					"subdomain_name": p.SubDomainName,
					"error_code":     strconv.Itoa(p.ErrorCode),
					"content":        p.Content,
				},
				[]string{},
				map[string]interface{}{},
			)
			r.IgnoreKeys = append(r.IgnoreKeys, "^metadata")
			g.Resources = append(g.Resources, r)
		}
		if len(pages) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// InitResources
//
func (g *ErrorPageGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain, *sync.WaitGroup) error{
		g.createErrorPageResources,
	}
	err = createResourcesPerDomain(api, funcs, &wg)
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}
