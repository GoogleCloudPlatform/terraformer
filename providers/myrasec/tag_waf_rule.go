package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

// TagWafRuleGenerator
type TagWafRuleGenerator struct {
	MyrasecService
}

// createTagWafRuleResources
func (g *TagWafRuleGenerator) createTagWafRuleResources(api *mgo.API, tag mgo.Tag, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}

	for {
		params["page"] = strconv.Itoa(page)

		waf, err := api.ListTagWAFRules(tag.ID, params)
		if err != nil {
			return err
		}

		for _, w := range waf {
			r := terraformutils.NewResource(
				strconv.Itoa(w.ID),
				fmt.Sprintf("%s_%s_%d", tag.Name, w.Name, w.ID),
				"myrasec_tag_waf_rule",
				"myrasec",
				map[string]string{
					"tag_id": strconv.Itoa(tag.ID),
				},
				[]string{},
				map[string]any{},
			)
			g.Resources = append(g.Resources, r)
		}

		if len(waf) < pageSize {
			break
		}
		page++
	}

	return nil
}

// InitResources
func (g *TagWafRuleGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Tag, *sync.WaitGroup) error{
		g.createTagWafRuleResources,
	}

	err = createResourcesPerTag(api, funcs, &wg, "WAF")
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
