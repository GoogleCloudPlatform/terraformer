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

package digitalocean

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DigitalOceanProvider struct { //nolint
	terraformutils.Provider
	token string
}

func (p *DigitalOceanProvider) Init(args []string) error {
	if os.Getenv("DIGITALOCEAN_TOKEN") == "" {
		return errors.New("set DIGITALOCEAN_TOKEN env var")
	}
	p.token = os.Getenv("DIGITALOCEAN_TOKEN")

	return nil
}

func (p *DigitalOceanProvider) GetName() string {
	return "digitalocean"
}

func (p *DigitalOceanProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (DigitalOceanProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *DigitalOceanProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"cdn":                &CDNGenerator{},
		"certificate":        &CertificateGenerator{},
		"database_cluster":   &DatabaseClusterGenerator{},
		"domain":             &DomainGenerator{},
		"droplet":            &DropletGenerator{},
		"droplet_snapshot":   &DropletSnapshotGenerator{},
		"firewall":           &FirewallGenerator{},
		"floating_ip":        &FloatingIPGenerator{},
		"kubernetes_cluster": &KubernetesClusterGenerator{},
		"loadbalancer":       &LoadBalancerGenerator{},
		"project":            &ProjectGenerator{},
		"ssh_key":            &SSHKeyGenerator{},
		"tag":                &TagGenerator{},
		"volume":             &VolumeGenerator{},
		"volume_snapshot":    &VolumeSnapshotGenerator{},
		"vpc":                &VPCGenerator{},
	}
}

func (p *DigitalOceanProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("digitalocean: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token": p.token,
	})
	return nil
}
