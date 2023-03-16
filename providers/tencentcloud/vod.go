// Copyright 2022 The Terraformer Authors.
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

package tencentcloud

import (
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

type VodGenerator struct {
	TencentCloudService
}

func (g *VodGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := vod.NewClient(&credential, region, profile)
	client.WithHttpTransport(&LogRoundTripper{
		Verbose: g.Verbose,
	})
	if err != nil {
		return err
	}

	if err := g.DescribeSubAppIds(client); err != nil {
		return err
	}
	if err := g.DescribeImageSpriteTemplates(client); err != nil {
		return err
	}
	if err := g.DescribeAdaptiveDynamicStreamingTemplates(client); err != nil {
		return err
	}
	if err := g.DescribeSnapshotByTimeOffsetTemplates(client); err != nil {
		return err
	}
	if err := g.DescribeProcedureTemplates(client); err != nil {
		return err
	}
	if err := g.DescribeSuperPlayerConfigs(client); err != nil {
		return err
	}

	return nil
}

func (g *VodGenerator) DescribeSubAppIds(client *vod.Client) error {
	request := vod.NewDescribeSubAppIdsRequest()

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.SubAppIdInfo, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeSubAppIds(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.SubAppIdInfoSet...)
		if len(response.Response.SubAppIdInfoSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		if *instance.Name == "主应用" && *instance.SubAppIdName == "PrimaryApplication" {
			continue
		}
		resource := terraformutils.NewResource(
			*instance.Name+"#"+strconv.FormatUint(*instance.SubAppId, 10),
			*instance.Name+"_"+strconv.FormatUint(*instance.SubAppId, 10),
			"tencentcloud_vod_sub_application",
			"tencentcloud",
			map[string]string{
				"status": *instance.Status,
			},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *VodGenerator) DescribeImageSpriteTemplates(client *vod.Client) error {
	request := vod.NewDescribeImageSpriteTemplatesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vod_image_sprite_template") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		definition, err := strconv.Atoi(filters[i])
		if err != nil {
			return err
		}
		d := uint64(definition)
		request.Definitions = append(request.Definitions, &d)
	}

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.ImageSpriteTemplate, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeImageSpriteTemplates(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.ImageSpriteTemplateSet...)
		if len(response.Response.ImageSpriteTemplateSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			strconv.FormatUint(*instance.Definition, 10),
			strconv.FormatUint(*instance.Definition, 10),
			"tencentcloud_vod_image_sprite_template",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *VodGenerator) DescribeAdaptiveDynamicStreamingTemplates(client *vod.Client) error {
	request := vod.NewDescribeAdaptiveDynamicStreamingTemplatesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vod_adaptive_dynamic_streaming_template") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		definition, err := strconv.Atoi(filters[i])
		if err != nil {
			return err
		}
		d := uint64(definition)
		request.Definitions = append(request.Definitions, &d)
	}

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.AdaptiveDynamicStreamingTemplate, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeAdaptiveDynamicStreamingTemplates(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.AdaptiveDynamicStreamingTemplateSet...)
		if len(response.Response.AdaptiveDynamicStreamingTemplateSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			strconv.FormatUint(*instance.Definition, 10),
			strconv.FormatUint(*instance.Definition, 10),
			"tencentcloud_vod_adaptive_dynamic_streaming_template",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *VodGenerator) DescribeSnapshotByTimeOffsetTemplates(client *vod.Client) error {
	request := vod.NewDescribeSnapshotByTimeOffsetTemplatesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vod_snapshot_by_time_offset_template") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		definition, err := strconv.Atoi(filters[i])
		if err != nil {
			return err
		}
		d := uint64(definition)
		request.Definitions = append(request.Definitions, &d)
	}

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.SnapshotByTimeOffsetTemplate, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeSnapshotByTimeOffsetTemplates(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.SnapshotByTimeOffsetTemplateSet...)
		if len(response.Response.SnapshotByTimeOffsetTemplateSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			strconv.FormatUint(*instance.Definition, 10),
			strconv.FormatUint(*instance.Definition, 10),
			"tencentcloud_vod_snapshot_by_time_offset_template",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *VodGenerator) DescribeProcedureTemplates(client *vod.Client) error {
	request := vod.NewDescribeProcedureTemplatesRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vod_procedure_template") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		request.Names = append(request.Names, &filters[i])
	}

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.ProcedureTemplate, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeProcedureTemplates(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.ProcedureTemplateSet...)
		if len(response.Response.ProcedureTemplateSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.Name,
			*instance.Name,
			"tencentcloud_vod_procedure_template",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
func (g *VodGenerator) DescribeSuperPlayerConfigs(client *vod.Client) error {
	request := vod.NewDescribeSuperPlayerConfigsRequest()
	filters := make([]string, 0)
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("tencentcloud_vod_super_player_config") {
			filters = append(filters, filter.AcceptableValues...)
		}
	}

	for i := range filters {
		request.Names = append(request.Names, &filters[i])
	}

	var offset uint64 = 0
	var limit uint64 = 50
	allInstances := make([]*vod.PlayerConfig, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := client.DescribeSuperPlayerConfigs(request)
		if err != nil {
			return err
		}
		allInstances = append(allInstances, response.Response.PlayerConfigSet...)
		if len(response.Response.PlayerConfigSet) < int(limit) {
			break
		}

		offset += limit
	}

	for _, instance := range allInstances {
		resource := terraformutils.NewResource(
			*instance.Name,
			*instance.Name,
			"tencentcloud_vod_super_player_config",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
