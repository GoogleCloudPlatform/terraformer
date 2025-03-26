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

type InstanceGenerator struct {
	YandexService
}

func (g *InstanceGenerator) loadInstances(sdk *ycsdk.SDK, folderID string) ([]*compute.Instance, error) {
	instances := []*compute.Instance{}
	pageToken := ""
	for {
		resp, err := sdk.Compute().Instance().List(context.Background(), &compute.ListInstancesRequest{
			FolderId:  folderID,
			PageSize:  defaultPageSize,
			PageToken: pageToken,
		})

		if err != nil {
			return nil, err
		}

		instances = append(instances, resp.GetInstances()...)

		if resp.GetNextPageToken() == "" {
			break
		}

	}
	return instances, nil

}

func (g *InstanceGenerator) InitResources() error {
	sdk, err := g.InitSDK()
	if err != nil {
		return err
	}

	result, err := g.loadInstances(sdk, g.Args["folder_id"].(string))
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result)

	return nil
}

func (g *InstanceGenerator) createResources(instances []*compute.Instance) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, instance := range instances {
		resources = append(resources, terraformutils.NewSimpleResource(
			instance.GetId(),
			instance.GetId(),
			"yandex_compute_instance",
			"yandex",
			[]string{}))
	}
	return resources
}
