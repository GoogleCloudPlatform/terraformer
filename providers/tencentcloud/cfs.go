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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cfs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs/v20190719"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type CfsGenerator struct {
	TencentCloudService
}

func (g *CfsGenerator) InitResources() error {
	args := g.GetArgs()
	region := args["region"].(string)
	credential := args["credential"].(common.Credential)
	profile := NewTencentCloudClientProfile()
	client, err := cfs.NewClient(&credential, region, profile)
	if err != nil {
		return err
	}

	request := cfs.NewDescribeCfsFileSystemsRequest()
	response, err := client.DescribeCfsFileSystems(request)
	if err != nil {
		return err
	}

	for _, instance := range response.Response.FileSystems {
		resource := terraformutils.NewResource(
			*instance.FileSystemId,
			*instance.FsName+"_"+*instance.FileSystemId,
			"tencentcloud_cfs_file_system",
			"tencentcloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
