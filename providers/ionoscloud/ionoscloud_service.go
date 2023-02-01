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

package ionoscloud

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	dbaas "github.com/ionos-cloud/sdk-go-dbaas-postgres"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type Service struct {
	terraformutils.Service
}

type Bundle struct {
	CloudAPIClient *ionoscloud.APIClient
	DBaaSApiClient *dbaas.APIClient
}

type clientType int

const (
	ionosClient clientType = iota
	dbaasClient
)

func (s *Service) generateClient() *Bundle {
	username := s.Args[helpers.UsernameArg].(string)
	password := s.Args[helpers.PasswordArg].(string)
	token := s.Args[helpers.TokenArg].(string)
	url := s.Args[helpers.URLArg].(string)

	cleanedURL := cleanURL(url)

	newConfig := ionoscloud.NewConfiguration(username, password, token, cleanedURL)

	if os.Getenv(helpers.IonosDebug) != "" {
		newConfig.Debug = true
	}

	newConfig.MaxRetries = helpers.MaxRetries
	newConfig.WaitTime = helpers.MaxWaitTime

	clients := map[clientType]interface{}{
		ionosClient: NewClientByType(username, password, token, cleanedURL, ionosClient),
		dbaasClient: NewClientByType(username, password, token, cleanedURL, dbaasClient),
	}
	return &Bundle{
		CloudAPIClient: clients[ionosClient].(*ionoscloud.APIClient),
		DBaaSApiClient: clients[dbaasClient].(*dbaas.APIClient),
	}
}

func NewClientByType(username, password, token, url string, clientType clientType) interface{} {
	switch clientType {
	case ionosClient:
		{
			newConfig := ionoscloud.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go/%s_os/%s_arch/%s", ionoscloud.Version, runtime.GOOS, runtime.GOARCH)
			return ionoscloud.NewAPIClient(newConfig)
		}
	case dbaasClient:
		{
			newConfig := dbaas.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go/%s_os/%s_arch/%s", ionoscloud.Version, runtime.GOOS, runtime.GOARCH)
			return dbaas.NewAPIClient(newConfig)
		}
	default:
		log.Printf("[ERROR] unknown client type %d", clientType)
	}
	return nil
}

// cleanURL makes sure trailing slash does not corrupt the state
func cleanURL(url string) string {
	length := len(url)
	if length > 1 && url[length-1] == '/' {
		url = url[:length-1]
	}

	return url
}
func CreateTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		DisableKeepAlives:     true,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   3,
		MaxConnsPerHost:       3,
	}
}
