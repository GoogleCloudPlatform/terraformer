// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	sdk, err := g.InitSDK()
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
