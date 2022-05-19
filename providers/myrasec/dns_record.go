package myrasec

import (
	"fmt"
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

type DNSGenerator struct {
	MyrasecService
}

//
// createDnsResources
//
func (g *DNSGenerator) createDnsResources(api *mgo.API, domain mgo.Domain) ([]terraformutils.Resource, error) {
	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		records, err := api.ListDNSRecords(domain.ID, params)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for _, d := range records {
			r := terraformutils.NewResource(
				strconv.Itoa(d.ID),
				fmt.Sprintf("%s_%d", domain.Name, d.ID),
				"myrasec_dns_record",
				"myrasec",
				map[string]string{
					"domain_name": domain.Name,
					"name":        d.Name,
					"value":       d.Value,
					"record_type": d.RecordType,
					"ttl":         strconv.Itoa(d.TTL),
				},
				[]string{},
				map[string]interface{}{},
			)

			g.Resources = append(g.Resources, r)
		}
		if len(records) < pageSize {
			break
		}
		page++
	}

	return g.Resources, nil
}

//
// InitResources
//
func (g *DNSGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		log.Println(err)
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error){
		g.createDnsResources,
	}

	res, err := createResourcesPerDomain(api, funcs)
	if err != nil {
		log.Println(err)
		return err
	}

	g.Resources = append(g.Resources, res...)

	return nil
}
