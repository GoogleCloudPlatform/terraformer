package myrasec

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"

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
func (g *DomainGenerator) createDomainResource(api *mgo.API, domain mgo.Domain, wg *sync.WaitGroup) error {
	defer wg.Done()

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
	g.Resources = append(g.Resources, d)

	return nil
}

//
// InitResources
//
func (g *DomainGenerator) InitResources() error {
	var wg = sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Domain, *sync.WaitGroup) error{
		g.createDomainResource,
	}

	err = createResourcesPerDomain(api, funcs, &wg)
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

//
// createResourcesPerDomain
//
func createResourcesPerDomain(api *mgo.API, funcs []func(*mgo.API, mgo.Domain, *sync.WaitGroup) error, wg *sync.WaitGroup) error {

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
			return err
		}

		wg.Add(len(domains) * len(funcs))
		for _, d := range domains {
			for _, f := range funcs {
				f(api, d, wg)
			}
		}
		if len(domains) < pageSize {
			break
		}
		page++
	}
	return nil
}

func getWaitChannel() chan struct{} {
	return make(chan struct{}, runtime.NumCPU()/2)
}

//
// createResourcesPerSubDomain
//
func createResourcesPerSubDomain(api *mgo.API, funcs []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error, wg *sync.WaitGroup, onDomainLevel bool) error {
	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	waitChan := getWaitChannel()
	count := 0

	for {
		params["page"] = strconv.Itoa(page)

		domains, err := api.ListDomains(params)
		if err != nil {
			return err
		}

		wg.Add(len(domains))
		for _, d := range domains {
			// try to load data for ALL-{domainId}.
			if onDomainLevel {
				wg.Add(len(funcs))
				for _, f := range funcs {
					go f(api, d.ID, mgo.VHost{
						Label: fmt.Sprintf("ALL-%d.", d.ID),
					}, wg)
				}

			}
			waitChan <- struct{}{}
			count++
			go func(count int, d mgo.Domain) {
				createResourcesPerVHost(api, d, funcs, wg)
				<-waitChan
			}(count, d)
		}
		if len(domains) < pageSize {
			break
		}
		page++
	}
	return nil
}

//
// createResourcesPerVHost
//
func createResourcesPerVHost(api *mgo.API, domain mgo.Domain, funcs []func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	waitChan := getWaitChannel()
	count := 0

	for {
		params["page"] = strconv.Itoa(page)

		vhosts, err := api.ListAllSubdomainsForDomain(domain.ID, params)
		if err != nil {
			return err
		}

		wg.Add(len(vhosts) * len(funcs))
		for _, v := range vhosts {
			for _, f := range funcs {
				waitChan <- struct{}{}
				count++
				go func(count int, v mgo.VHost, f func(*mgo.API, int, mgo.VHost, *sync.WaitGroup) error) {
					f(api, domain.ID, v, wg)
					<-waitChan
				}(count, v, f)
			}
		}
		if len(vhosts) < pageSize {
			break
		}
		page++
	}
	return nil
}
