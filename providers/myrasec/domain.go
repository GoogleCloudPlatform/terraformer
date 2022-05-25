package myrasec

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// DomainGenerator
//
type DomainGenerator struct {
	MyrasecService
}

//
// createDomainResource
//
func (g *DomainGenerator) createDomainResource(api *mgo.API, domain mgo.Domain) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	d := terraformutils.NewResource(
		strconv.Itoa(domain.ID),
		fmt.Sprintf("%s_%d", domain.Name, domain.ID),
		"myrasec_domain",
		"myrasec",
		map[string]string{},
		[]string{},
		map[string]interface{}{},
	)

	d.IgnoreKeys = append(d.IgnoreKeys, "^metadata")
	resources = append(resources, d)

	return resources, nil
}

//
// InitResources
//
func (g *DomainGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error){
		g.createDomainResource,
	}

	res, err := createResourcesPerDomain(api, funcs)
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, res...)

	return nil
}

//
// createResourcesPerDomain
//
func createResourcesPerDomain(api *mgo.API, funcs []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error)) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		domains, err := api.ListDomains(params)
		if err != nil {
			return nil, err
		}

		for _, d := range domains {
			for _, f := range funcs {
				tmpRes, err := f(api, d)
				if err != nil {
					return nil, err
				}
				resources = append(resources, tmpRes...)
			}
		}
		if len(domains) < pageSize {
			break
		}
		page++
	}
	return resources, nil
}

//
// createResourcesPerSubDomain
//
func createResourcesPerSubDomain(api *mgo.API, funcs []func(*mgo.API, int, mgo.VHost) ([]terraformutils.Resource, error), onDomainLevel bool) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		domains, err := api.ListDomains(params)
		if err != nil {
			return nil, err
		}

		for _, d := range domains {
			// try to load data for ALL-{domainId}.
			if onDomainLevel {
				for _, f := range funcs {
					tmpRes, err := f(api, d.ID, mgo.VHost{
						Label: fmt.Sprintf("ALL-%d.", d.ID),
					})
					if err == nil {
						resources = append(resources, tmpRes...)
					}
				}

			}
			res, err := createResourcesPerVHost(api, d, funcs)
			if err != nil {
				return nil, err
			}
			resources = append(resources, res...)
		}
		if len(domains) < pageSize {
			break
		}
		page++
	}
	return resources, nil
}

//
// createResourcesPerVHost
//
func createResourcesPerVHost(api *mgo.API, domain mgo.Domain, funcs []func(*mgo.API, int, mgo.VHost) ([]terraformutils.Resource, error)) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		api.ListAllSubdomains(params)
		vhosts, err := api.ListAllSubdomainsForDomain(domain.ID, params)
		if err != nil {
			return nil, err
		}

		for _, v := range vhosts {
			for _, f := range funcs {
				tmpRes, err := f(api, domain.ID, v)
				if err != nil {
					return nil, err
				}
				resources = append(resources, tmpRes...)
			}
		}
		if len(vhosts) < pageSize {
			break
		}
		page++
	}
	return resources, nil
}
