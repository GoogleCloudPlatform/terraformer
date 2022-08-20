package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type QueryGenerator struct {
	HoneycombService
}

func (g *QueryGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	ctx := context.TODO()

	for _, dataset := range g.datasets {
		triggers, err := client.Triggers.List(ctx, dataset.Slug)
		if err != nil {
			return fmt.Errorf("unable to list Honeycomb triggers for dataset %s: %v", dataset.Slug, err)
		}

		for _, trigger := range triggers {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				trigger.QueryID,
				trigger.QueryID,
				"honeycombio_query",
				"honeycombio",
				map[string]string{
					"dataset": dataset.Name,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	boards, err := client.Boards.List(context.TODO())
	if err != nil {
		return err
	}

	for _, board := range boards {
		for _, query := range board.Queries {
			_, datasetSelected := g.datasetMap[query.Dataset]
			if datasetSelected {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					query.QueryID,
					query.QueryID,
					"honeycombio_query",
					"honeycombio",
					map[string]string{
						"dataset": query.Dataset,
					},
					[]string{"caption", "query_annotation_id"},
					map[string]interface{}{},
				))
			}
		}
	}

	return nil
}

// PostGenerateHook to format any generated query resource's QuerySpec JSON as a heredoc
// func (g *QueryGenerator) PostConvertHook() error {
// 	for i, resource := range g.Resources {
// 		if resource.InstanceInfo.Type != "honeycombio_query" {
// 			continue
// 		}
// 		if _, exist := resource.Item["query_json"]; exist {
// 			queryJSON := resource.Item["query_json"].(string)
// 			unquotedStr, _ := strconv.Unquote(queryJSON)
// 			fmt.Println(queryJSON)
// 			g.Resources[i].Item["query_json"] = `<<EOH
// ` + unquotedStr + `
// EOH`
// 		}
// 	}
// 	return nil
// }
