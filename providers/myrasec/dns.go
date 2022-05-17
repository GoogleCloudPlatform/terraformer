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

func (g *DNSGenerator) createDNSRecordsResource(api *mgo.API, domain mgo.Domain, record mgo.DNSRecord) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	r := terraformutils.NewResource(
		strconv.Itoa(record.ID),
		fmt.Sprintf("%s_%d", record.Name, record.ID),
		"myrasec_dns_record",
		"myrasec",
		map[string]string{
			"domain_name": domain.Name,
			"name":        record.Name,
			"value":       record.Value,
			"record_type": record.RecordType,
			"ttl":         strconv.Itoa(record.TTL),
		},
		[]string{},
		map[string]interface{}{},
	)

	r.IgnoreKeys = append(r.IgnoreKeys, "^metadata")
	resources = append(resources, r)

	return resources, nil
}

func (g *DNSGenerator) InitDnsResources(api *mgo.API, domain mgo.Domain) ([]terraformutils.Resource, error) {
	funcs := []func(*mgo.API, mgo.Domain, mgo.DNSRecord) ([]terraformutils.Resource, error){
		g.createDNSRecordsResource,
	}

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

		for _, r := range records {
			for _, f := range funcs {
				tmpRes, err := f(api, domain, r)
				if err != nil {
					log.Println(err)
					return g.Resources, nil
				}
				g.Resources = append(g.Resources, tmpRes...)
			}
		}
		if len(records) < pageSize {
			break
		}
		page++
	}

	return g.Resources, nil
}

func (g *DNSGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		log.Println(err)
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain) ([]terraformutils.Resource, error){
		g.InitDnsResources,
	}

	res, err := createResourcesPerDomain(api, funcs)
	if err != nil {
		log.Println(err)
		return err
	}

	g.Resources = append(g.Resources, res...)

	return nil
}
