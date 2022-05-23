package myrasec

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// IPFilterGenerator
//
type IPFilterGenerator struct {
	MyrasecService
}

//
// createIPFilterResources
//
func (g *IPFilterGenerator) createIPFilterResources(api *mgo.API, domainId int, vhost mgo.VHost) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	page := 1
	pageSize := 250
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}

	for {
		params["page"] = strconv.Itoa(page)

		filters, err := api.ListIPFilters(domainId, vhost.Label, params)
		if err != err {
			return nil, err
		}

		for _, f := range filters {
			r := terraformutils.NewResource(
				strconv.Itoa(f.ID),
				fmt.Sprintf("%s_%d", vhost.Label, f.ID),
				"myrasec_ip_filter",
				"myrasec",
				map[string]string{
					"subdomain_name": vhost.Label,
				},
				[]string{},
				map[string]interface{}{},
			)
			resources = append(resources, r)
		}
		if len(filters) < pageSize {
			break
		}
		page++
	}
	return resources, nil
}

//
// InitResources
//
func (g *IPFilterGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost) ([]terraformutils.Resource, error){
		g.createIPFilterResources,
	}

	res, err := createResourcesPerSubDomain(api, funcs)
	if err != nil {
		return nil
	}

	g.Resources = res

	return nil
}
