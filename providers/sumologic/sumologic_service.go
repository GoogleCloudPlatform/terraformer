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
	"context"
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type SumoLogicService struct { //nolint
	terraformutils.Service
}

func (s *SumoLogicService) Client() *sumologic.APIClient {
	log.Printf("%s access id:%s, baseUrl:%s\n",
		s.GetProviderName(), s.GetArgs()["accessId"].(string), s.GetArgs()["baseUrl"].(string))

	config := sumologic.NewConfiguration()
	config.Servers = []sumologic.ServerConfiguration{
		{
			URL:         strings.TrimSuffix(s.GetArgs()["baseUrl"].(string), "/"),
			Description: "Base URL of the API Server",
			Variables:   make(map[string]sumologic.ServerVariable),
		},
	}
	client := sumologic.NewAPIClient(config)
	return client
}

func (s *SumoLogicService) AuthCtx() context.Context {
	auth := context.WithValue(context.Background(),
		sumologic.ContextBasicAuth,
		sumologic.BasicAuth{
			UserName: s.GetArgs()["accessId"].(string),
			Password: s.GetArgs()["accessKey"].(string),
		},
	)
	return auth
}
