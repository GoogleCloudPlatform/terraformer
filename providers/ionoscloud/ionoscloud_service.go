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
	certificateManager "github.com/ionos-cloud/sdk-go-cert-manager"
	containerRegistry "github.com/ionos-cloud/sdk-go-container-registry"
	dataPlatform "github.com/ionos-cloud/sdk-go-dataplatform"
	dbaasMongo "github.com/ionos-cloud/sdk-go-dbaas-mongo"
	dbaasPgSQL "github.com/ionos-cloud/sdk-go-dbaas-postgres"
	dns "github.com/ionos-cloud/sdk-go-dns"
	logging "github.com/ionos-cloud/sdk-go-logging"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type Service struct {
	terraformutils.Service
}

type Bundle struct {
	CloudAPIClient              *ionoscloud.APIClient
	DBaaSPgSQLApiClient         *dbaasPgSQL.APIClient
	DBaaSMongoAPIClient         *dbaasMongo.APIClient
	CertificateManagerAPIClient *certificateManager.APIClient
	ContainerRegistryAPIClient  *containerRegistry.APIClient
	DataPlatformAPIClient       *dataPlatform.APIClient
	DNSAPIClient                *dns.APIClient
	LoggingAPIClient            *logging.APIClient
}

type clientType int

const (
	ionosClient clientType = iota
	dbaasPgSQLClient
	dbaasMongoClient
	certificateManagerClient
	containerRegistryClient
	dataPlatformClient
	dnsClient
	loggingClient
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
		ionosClient:              NewClientByType(username, password, token, cleanedURL, ionosClient),
		dbaasPgSQLClient:         NewClientByType(username, password, token, cleanedURL, dbaasPgSQLClient),
		dbaasMongoClient:         NewClientByType(username, password, token, cleanedURL, dbaasMongoClient),
		certificateManagerClient: NewClientByType(username, password, token, cleanedURL, certificateManagerClient),
		containerRegistryClient:  NewClientByType(username, password, token, cleanedURL, containerRegistryClient),
		dataPlatformClient:       NewClientByType(username, password, token, cleanedURL, dataPlatformClient),
		dnsClient:                NewClientByType(username, password, token, cleanedURL, dnsClient),
		loggingClient:            NewClientByType(username, password, token, cleanedURL, loggingClient),
	}

	return &Bundle{
		CloudAPIClient:              clients[ionosClient].(*ionoscloud.APIClient),
		DBaaSPgSQLApiClient:         clients[dbaasPgSQLClient].(*dbaasPgSQL.APIClient),
		DBaaSMongoAPIClient:         clients[dbaasMongoClient].(*dbaasMongo.APIClient),
		CertificateManagerAPIClient: clients[certificateManagerClient].(*certificateManager.APIClient),
		ContainerRegistryAPIClient:  clients[containerRegistryClient].(*containerRegistry.APIClient),
		DataPlatformAPIClient:       clients[dataPlatformClient].(*dataPlatform.APIClient),
		DNSAPIClient:                clients[dnsClient].(*dns.APIClient),
		LoggingAPIClient:            clients[loggingClient].(*logging.APIClient),
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
	case dbaasPgSQLClient:
		{
			newConfig := dbaasPgSQL.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-dbaas-postgres/%s_os/%s_arch/%s", dbaasPgSQL.Version, runtime.GOOS, runtime.GOARCH)
			return dbaasPgSQL.NewAPIClient(newConfig)
		}
	case dbaasMongoClient:
		{
			newConfig := dbaasMongo.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-dbaas-mongo/%s_os/%s_arch/%s", dbaasMongo.Version, runtime.GOOS, runtime.GOARCH)
			return dbaasMongo.NewAPIClient(newConfig)
		}
	case certificateManagerClient:
		{
			newConfig := certificateManager.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-cert-manager/%s_os/%s_arch/%s", certificateManager.Version, runtime.GOOS, runtime.GOARCH)
			return certificateManager.NewAPIClient(newConfig)
		}
	case containerRegistryClient:
		{
			newConfig := containerRegistry.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-container-registry/%s_os/%s_arch/%s", containerRegistry.Version, runtime.GOOS, runtime.GOARCH)
			return containerRegistry.NewAPIClient(newConfig)
		}
	case dataPlatformClient:
		{
			newConfig := dataPlatform.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-dataplatform/%s_os/%s_arch/%s", dataPlatform.Version, runtime.GOOS, runtime.GOARCH)
			return dataPlatform.NewAPIClient(newConfig)
		}
	case dnsClient:
		{
			newConfig := dns.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-dns/%s_os/%s_arch/%s", dns.Version, runtime.GOOS, runtime.GOARCH)
			return dns.NewAPIClient(newConfig)
		}
	case loggingClient:
		{
			newConfig := logging.NewConfiguration(username, password, token, url)

			if os.Getenv(helpers.IonosDebug) != "" {
				newConfig.Debug = true
			}
			newConfig.MaxRetries = helpers.MaxRetries
			newConfig.WaitTime = helpers.MaxWaitTime
			newConfig.HTTPClient = &http.Client{Transport: CreateTransport()}
			newConfig.UserAgent = fmt.Sprintf(
				"terraformer_ionos-cloud-sdk-go-logging/%s_os/%s_arch/%s", logging.Version, runtime.GOOS, runtime.GOARCH)
			return logging.NewAPIClient(newConfig)
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
