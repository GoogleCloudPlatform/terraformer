package myrasec

import (
	"fmt"
	"strconv"

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
func (g *ErrorPageGenerator) createErrorPageResources(api *mgo.API, domain mgo.Domain) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

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
			return nil, err
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
			resources = append(resources, r)
		}
		if len(pages) < pageSize {
			break
		}
		page++
	}
	return resources, nil
}

//
// InitResources
//
func (g *ErrorPageGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error){
		g.createErrorPageResources,
	}
	res, err := createResourcesPerDomain(api, funcs)
	if err != nil {
		return err
	}

	g.Resources = res

	return nil
}
