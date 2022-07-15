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

package ibm

import (
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

// ContainerNlbDnsGenerator ...
type ContainerNlbDnsGenerator struct {
	IBMService
}

func (g ContainerNlbDnsGenerator) loadNlbDns(clusterName string, nlbIPs []interface{}) terraformutils.Resource {
	resources := terraformutils.NewResource(
		clusterName,
		normalizeResourceName(clusterName, true),
		"ibm_container_nlb_dns",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"nlb_ips": nlbIPs,
		})
	return resources
}

// InitResources ...
func (g *ContainerNlbDnsGenerator) InitResources() error {

	clusterName := g.Args["cluster"].(string)

	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}

	client, err := containerv2.New(sess)
	if err != nil {
		return err
	}

	nlbData, err := client.NlbDns().GetNLBDNSList(clusterName)
	if err != nil {
		return err
	}

	for _, data := range nlbData {
		g.Resources = append(g.Resources, g.loadNlbDns(data.Nlb.Cluster, data.Nlb.NlbIPArray))
	}

	return nil
}
