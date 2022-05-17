package myrasec

import (
	"fmt"
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ms "github.com/Myra-Security-GmbH/myrasec-go/v2"
)

type ErrorPageGenerator struct {
	MyrasecService
}

func (*ErrorPageGenerator) createErrorPageResources(api *ms.API) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	params := map[string]string{
		"pageSize": strconv.Itoa(20),
	}

	pages, err := api.ListErrorPages(1001771, params)
	if err != nil {
		log.Println(err)
		return resources, err
	}

	for _, page := range pages {
		p := terraformutils.NewResource(
			strconv.Itoa(page.ID),
			fmt.Sprintf("%s-%s", "p", page.SubDomainName),
			"myrasec_error_page",
			"myrasecfdasfasdf",
			map[string]string{
				"subdomain_name": page.SubDomainName,
				"error_code":     strconv.Itoa(page.ErrorCode),
				"content":        page.Content,
			},
			[]string{},
			map[string]interface{}{},
		)
		resources = append(resources, p)
	}

	return resources, nil
}

func (g *ErrorPageGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	res, err := g.createErrorPageResources(api)

	g.Resources = res

	return nil
}
