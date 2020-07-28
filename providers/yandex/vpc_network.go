package yandex

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type NetworkGenerator struct {
	YandexService
}

func (g *NetworkGenerator) loadNetworks(sdk *ycsdk.SDK, folderID string) ([]*vpc.Network, error) {
	networks := []*vpc.Network{}
	pageToken := ""
	for {
		resp, err := sdk.VPC().Network().List(context.Background(), &vpc.ListNetworksRequest{
			FolderId:  folderID,
			PageSize:  defaultPageSize,
			PageToken: pageToken,
		})

		if err != nil {
			return nil, err
		}

		networks = append(networks, resp.GetNetworks()...)

		if resp.GetNextPageToken() == "" {
			break
		}

	}
	return networks, nil

}

func (g *NetworkGenerator) InitResources() error {
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(g.Args["token"].(string)),
	})
	if err != nil {
		return err
	}

	result, err := g.loadNetworks(sdk, g.Args["folder_id"].(string))
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result)

	return nil
}

func (g *NetworkGenerator) createResources(networks []*vpc.Network) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, network := range networks {
		resources = append(resources, terraformutils.NewSimpleResource(
			network.GetId(),
			network.GetId(),
			"yandex_vpc_network",
			"yandex",
			[]string{}))
	}
	return resources
}
