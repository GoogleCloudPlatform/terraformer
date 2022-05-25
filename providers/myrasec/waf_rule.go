package myrasec

import (
	"fmt"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

//
// WafRuleGenerator
//
type WafRuleGenerator struct {
	MyrasecService
}

//
// createWafRuleResources
//
func (g *WafRuleGenerator) createWafRuleResources(api *mgo.API, domainId int, vhost mgo.VHost) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	page := 1
	pageSize := 250
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}

	for {
		params["page"] = strconv.Itoa(page)
		if vhost.Label != "" {
			params["subDomain"] = vhost.Label
		}

		waf, err := api.ListWAFRules(domainId, params)
		if err != nil {
			return nil, err
		}

		for _, w := range waf {
			r := terraformutils.NewResource(
				strconv.Itoa(w.ID),
				fmt.Sprintf("%s_%d", w.SubDomainName, w.ID),
				"myrasec_waf_rule",
				"myrasec",
				map[string]string{
					"subdomain_name": w.SubDomainName,
				},
				[]string{},
				map[string]interface{}{},
			)
			resources = append(resources, r)
		}

		if len(waf) < pageSize {
			break
		}
		page++
	}

	return resources, nil
}

//
// InitResources
//
func (g *WafRuleGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, int, mgo.VHost) ([]terraformutils.Resource, error){
		g.createWafRuleResources,
	}

	res, err := createResourcesPerSubDomain(api, funcs, true)
	if err != nil {
		return err
	}

	g.Resources = res

	return nil
}
