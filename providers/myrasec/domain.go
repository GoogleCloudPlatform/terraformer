package myrasec

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

type DomainGenerator struct {
	MyrasecService
}

func (g *DomainGenerator) createDomainResource(api *mgo.API, domainID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	dID, err := strconv.Atoi(domainID)
	if err != nil {
		return resources, err
	}

	domain, err := api.GetDomain(dID)
	if err != nil {
		log.Println(err)
		return resources, err
	}

	pausedUntil := ""
	if domain.PausedUntil != nil {
		pausedUntil = domain.PausedUntil.Format(time.RFC3339)
	}
	d := terraformutils.NewResource(
		domainID,
		fmt.Sprintf("%s_%d", domain.Name, domain.ID),
		"myrasec_domain",
		"myrasec",
		map[string]string{
			"domain_id":    domainID,
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

func (g *DomainGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		log.Println(err)
		return err
	}

	funcs := []func(*mgo.API, string) ([]terraformutils.Resource, error){
		g.createDomainResource,
	}

	page := 1
	pageSize := 20
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		domains, err := api.ListDomains(params)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, d := range domains {
			for _, f := range funcs {
				tmpRes, err := f(api, strconv.Itoa(d.ID))
				if err != nil {
					log.Println(err)
					return err
				}
				g.Resources = append(g.Resources, tmpRes...)
			}
		}
		if len(domains) < pageSize {
			break
		}
		page++
	}

	return nil
}
