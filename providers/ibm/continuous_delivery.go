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
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

// DatabaseRedisGenerator ...
type ContinuousDeliveryGenerator struct {
	IBMService
}

// loadRedisDB ...
func (g ContinuousDeliveryGenerator) loadContinuousDelivery(cdID string, cdName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		cdID,
		normalizeResourceName(cdName, true),
		"ibm_resource_instance",
		"ibm",
		[]string{})
	return resources
}

// InitResources ...
func (g *ContinuousDeliveryGenerator) InitResources() error {
	region := g.Args["region"].(string)
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
		Region:        region,
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	serviceID, err := catalogClient.ResourceCatalog().FindByName("continuous-delivery", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	continuousDeliveryInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	for _, cd := range continuousDeliveryInstances {
		if cd.RegionID == region {
			g.Resources = append(g.Resources, g.loadContinuousDelivery(cd.ID, cd.Name))
		}
	}

	return nil
}
