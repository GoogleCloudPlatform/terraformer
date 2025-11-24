package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

// TagCacheSettingGenerator
type TagCacheSettingGenerator struct {
	MyrasecService
}

// createTagCacheSettingResources
func (g *TagCacheSettingGenerator) createTagCacheSettingResources(api *mgo.API, tag mgo.Tag, wg *sync.WaitGroup) error {
	defer wg.Done()

	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}

	for {
		params["page"] = strconv.Itoa(page)

		settings, err := api.ListTagCacheSettings(tag.ID, params)

		if err != nil {
			return err
		}

		for _, s := range settings {
			r := terraformutils.NewResource(
				strconv.Itoa(s.ID),
				fmt.Sprintf("%s_%d", tag.Name, s.ID),
				"myrasec_tag_cache_setting",
				"myrasec",
				map[string]string{
					"tag_id": strconv.Itoa(tag.ID),
				},
				[]string{},
				map[string]any{},
			)
			g.Resources = append(g.Resources, r)
		}
		if len(settings) < pageSize {
			break
		}
		page++
	}
	return nil
}

// InitResources
func (g *TagCacheSettingGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Tag, *sync.WaitGroup) error{
		g.createTagCacheSettingResources,
	}
	err = createResourcesPerTag(api, funcs, &wg, "CACHE")
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}
