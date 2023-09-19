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

package grafana

import (
	"crypto/tls"
	"crypto/x509"
	"net/url"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gapi "github.com/grafana/grafana-api-golang-client"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
)

type GrafanaService struct { //nolint
	terraformutils.Service
}

func (s *GrafanaService) buildClient() (*gapi.Client, error) {
	auth := strings.SplitN(s.Args["auth"].(string), ":", 2)
	cli := cleanhttp.DefaultClient()
	transport := cleanhttp.DefaultTransport()
	transport.TLSClientConfig = &tls.Config{}

	// TLS Config
	tlsKey := s.Args["tls_key"].(string)
	tlsCert := s.Args["tls_cert"].(string)
	caCert := s.Args["ca_cert"].(string)
	insecure := s.Args["insecure_skip_verify"].(bool)

	if caCert != "" {
		ca, err := os.ReadFile(caCert)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(ca)
		transport.TLSClientConfig.RootCAs = pool
	}

	if tlsKey != "" && tlsCert != "" {
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return nil, err
		}
		transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
	}

	if insecure {
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	cli.Transport = logging.NewTransport("Grafana", transport)
	cfg := gapi.Config{
		Client: cli,
		OrgID:  int64(s.Args["org_id"].(int)),
	}

	if len(auth) == 2 {
		cfg.BasicAuth = url.UserPassword(auth[0], auth[1])
	} else {
		cfg.APIKey = auth[0]
	}

	client, err := gapi.New(s.Args["url"].(string), cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
