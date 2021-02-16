// Copyright 2021 The Terraformer Authors.
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

package kibana

import (
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gk "github.com/ewilde/go-kibana"
)

type KibanaService struct { //nolint
	terraformutils.Service
}

func (s *KibanaService) generateClient() *gk.KibanaClient {
	kibanaClient := gk.NewClient(gk.NewDefaultConfig())
	kibanaClient.SetAuth(getAuthForContainerVersion(kibanaClient.Config.KibanaVersion, kibanaClient.Config.KibanaType))
	kibanaClient.Search()
	return kibanaClient
}

var authForContainerVersion = map[string]map[gk.KibanaType]gk.AuthenticationHandler{
	gk.DefaultLogzioVersion: {
		gk.KibanaTypeLogzio:  createLogzAuthenticationHandler(),
		gk.KibanaTypeVanilla: &gk.NoAuthenticationHandler{},
	},
	gk.DefaultKibanaVersion6: {gk.KibanaTypeVanilla: &gk.NoAuthenticationHandler{}},
}

func getAuthForContainerVersion(version string, kibanaType gk.KibanaType) gk.AuthenticationHandler {
	handler, ok := authForContainerVersion[version]
	if !ok {
		handler = authForContainerVersion[gk.DefaultKibanaVersion6]
	}

	if kibanaType == gk.KibanaTypeLogzio {
		handler = authForContainerVersion[gk.DefaultLogzioVersion]
	}

	return handler[kibanaType]
}

func createLogzAuthenticationHandler() *gk.LogzAuthenticationHandler {
	uri := os.Getenv(gk.EnvKibanaUri)
	if uri == "" {
		uri = "https://app-eu.logz.io"
	}

	handler := gk.NewLogzAuthenticationHandler(nil)
	handler.Auth0Uri = "https://logzio.auth0.com"
	handler.LogzUri = uri
	handler.ClientId = os.Getenv(gk.EnvLogzClientId)
	handler.UserName = os.Getenv(gk.EnvKibanaUserName)
	handler.Password = os.Getenv(gk.EnvKibanaPassword)
	handler.MfaSecret = os.Getenv(gk.EnvLogzMfaSecret)

	return handler
}
