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
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const KeyToken = "token"
const KeyFolderID = "folder_id"
const KeySaKeyFileOrContent = "sa_key_or_content"

type YandexProvider struct { //nolint
	terraformutils.Provider
	token              string
	saKeyFileOrContent string
	folderID           string
}

func (p *YandexProvider) Init(args []string) error {
	if ycToken, ok := os.LookupEnv("YC_TOKEN"); ok {
		p.token = ycToken
	}

	if saKeyFileOrContent, ok := os.LookupEnv("YC_SERVICE_ACCOUNT_KEY_FILE"); ok {
		p.saKeyFileOrContent = saKeyFileOrContent
	}

	if len(args) > 0 {
		//  first args is target folder ID
		p.folderID = args[0]
	} else {
		if os.Getenv("YC_FOLDER_ID") == "" {
			return errors.New("set YC_FOLDER_ID env var")
		}
		p.folderID = os.Getenv("YC_FOLDER_ID")
	}

	return nil
}

func (p *YandexProvider) GetName() string {
	return "yandex"
}

func (p *YandexProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (YandexProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *YandexProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"disk":     &DiskGenerator{},
		"instance": &InstanceGenerator{},
		"network":  &NetworkGenerator{},
		"subnet":   &SubnetGenerator{},
	}
}

func (p *YandexProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("yandex: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		KeyFolderID:           p.folderID,
		KeyToken:              p.token,
		KeySaKeyFileOrContent: p.saKeyFileOrContent,
	})
	return nil
}
