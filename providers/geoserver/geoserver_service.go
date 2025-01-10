// Copyright 2018 The Terraformer Authors.
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

package geoserver

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gs "github.com/camptocamp/go-geoserver/client"
)

type GeoServerService struct { //nolint
	terraformutils.Service
}

// GeoserverClient creates a Geoserver client scoped to the global API
func (s *GeoServerService) GeoserverClient() *gs.Client {
	tspt := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.Args["insecure"].(bool),
		},
	}

	client := &gs.Client{
		URL:      s.Args["geoserverURL"].(string),
		Username: s.Args["user"].(string),
		Password: s.Args["password"].(string),
		HTTPClient: &http.Client{
			Transport: tspt,
		},
	}

	log.Printf("[INFO] Geoserver Client configured")

	return client
}

// Client creates a Geoserver client scoped to the global API
func (s *GeoServerService) GwcClient() *gs.Client {
	tspt := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.Args["insecure"].(bool),
		},
	}

	client := &gs.Client{
		URL:      s.Args["geowebcacheURL"].(string),
		Username: s.Args["user"].(string),
		Password: s.Args["password"].(string),
		HTTPClient: &http.Client{
			Transport: tspt,
		},
	}

	log.Printf("[INFO] GeoWebCache Client configured")

	return client
}
