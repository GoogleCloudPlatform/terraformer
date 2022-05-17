package myrasec

import (
	"fmt"
	"log"
	"strconv"
	"time"

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

	pausedUntil := ""
	if domain.PausedUntil != nil {
		pausedUntil = domain.PausedUntil.Format(time.RFC3339)
	}
	d := terraformutils.NewResource(
		strconv.Itoa(domain.ID),
		fmt.Sprintf("%s_%d", domain.Name, domain.ID),
		"myrasec_domain",
		"myrasec",
		map[string]string{
			"domain_id":    strconv.Itoa(domain.ID),
			"name":         domain.Name,
			"auto_update":  strconv.FormatBool(domain.AutoUpdate),
			"paused":       strconv.FormatBool(domain.Paused),
			"paused_until": pausedUntil,
		},
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
		log.Println(err)
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error){
		g.createDomainResource,
	}

	res, err := createResourcesPerDomain(api, funcs)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
			return nil, err
		}

		for _, d := range domains {
			for _, f := range funcs {
				tmpRes, err := f(api, d)
				if err != nil {
					log.Println(err)
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
