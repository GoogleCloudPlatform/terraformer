package yandex

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type DiskGenerator struct {
	YandexService
}

func (g *DiskGenerator) loadDisks(sdk *ycsdk.SDK, folderID string) ([]*compute.Disk, error) {
	disks := []*compute.Disk{}
	pageToken := ""
	for {
		resp, err := sdk.Compute().Disk().List(context.Background(), &compute.ListDisksRequest{
			FolderId:  folderID,
			PageSize:  defaultPageSize,
			PageToken: pageToken,
		})

		if err != nil {
			return nil, err
		}

		disks = append(disks, resp.GetDisks()...)

		if resp.GetNextPageToken() == "" {
			break
		}

	}
	return disks, nil

}

func (g *DiskGenerator) InitResources() error {
	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials: ycsdk.OAuthToken(g.Args["token"].(string)),
	})
	if err != nil {
		return err
	}

	result, err := g.loadDisks(sdk, g.Args["folder_id"].(string))
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result)

	return nil
}

func (g *DiskGenerator) createResources(disks []*compute.Disk) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, disk := range disks {
		resources = append(resources, terraformutils.NewSimpleResource(
			disk.GetId(),
			disk.GetId(),
			"yandex_compute_disk",
			"yandex",
			[]string{}))
	}
	return resources
}
