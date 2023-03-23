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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type FolderGenerator struct {
	SumoLogicService
}

func (g *FolderGenerator) createResources(folders []sumologic.Folder) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(folders))

	for i, folder := range folders {
		title := strcase.ToSnake(replaceSpaceAndDash(folder.Name))

		resource := terraformutils.NewSimpleResource(
			folder.Id,
			fmt.Sprintf("%s-%s", title, folder.Id),
			"sumologic_folder",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *FolderGenerator) InitResources() error {
	client := g.Client()

	var resources []terraformutils.Resource

	req := client.FolderManagementApi.GetPersonalFolder(g.AuthCtx())

	respBody, _, err := client.FolderManagementApi.GetPersonalFolderExecute(req)
	if err != nil {
		return err
	}

	personalFolder := *respBody
	folders := g.getAllChildFolders(personalFolder)

	resources = g.createResources(folders)
	g.Resources = resources
	return nil
}

func (g *FolderGenerator) getFolderOk(folderId string) (sumologic.Folder, bool) {
	client := g.Client()

	req := client.FolderManagementApi.GetFolder(g.AuthCtx(), folderId)
	folder, _, err := client.FolderManagementApi.GetFolderExecute(req)
	if err != nil {
		fmt.Println(err)
		return sumologic.Folder{}, false
	}

	return *folder, true
}

func (g *FolderGenerator) getAllChildFolders(rootFolder sumologic.Folder) []sumologic.Folder {
	var folders []sumologic.Folder

	for _, child := range rootFolder.Children {
		if child.ItemType == "Folder" {
			if folder, ok := g.getFolderOk(child.Id); ok {
				folders = append(folders, folder)
				folders = append(folders, g.getAllChildFolders(folder)...)
			}
		}
	}

	return folders
}
