package myrasec

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mgo "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

type TagGenerator struct {
	MyrasecService
}

func (g *TagGenerator) createTagResource(api *mgo.API, tag mgo.Tag, wg *sync.WaitGroup) error {
	defer wg.Done()

	t := terraformutils.NewResource(
		strconv.Itoa(tag.ID),
		fmt.Sprintf("%s_%d", tag.Name, tag.ID),
		"myrasec_tag",
		"myrasec",
		map[string]string{},
		[]string{},
		map[string]any{},
	)
	g.Resources = append(g.Resources, t)
	return nil
}

func (g *TagGenerator) InitResources() error {
	wg := sync.WaitGroup{}

	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	funcs := []func(*mgo.API, mgo.Tag, *sync.WaitGroup) error{
		g.createTagResource,
	}

	err = createResourcesPerTag(api, funcs, &wg, "ALL")
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

func createResourcesPerTag(api *mgo.API, funcs []func(*mgo.API, mgo.Tag, *sync.WaitGroup) error, wg *sync.WaitGroup, tagType string) error {
	page := 1
	pageSize := 250
	params := map[string]string{
		"pageSize": strconv.Itoa(pageSize),
		"page":     strconv.Itoa(page),
	}
	if tagType != "ALL" {
		params["type"] = tagType
	}

	for {
		params["page"] = strconv.Itoa(page)

		tags, err := api.ListTags(params)
		if err != nil {
			return err
		}

		wg.Add(len(tags) * len(funcs))
		for _, t := range tags {
			for _, f := range funcs {
				f(api, t, wg)
			}
		}
		if len(tags) < pageSize {
			break
		}
		page++
	}
	return nil
}
