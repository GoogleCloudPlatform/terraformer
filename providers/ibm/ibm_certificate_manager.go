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
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/certificatemanager"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
)

type CMGenerator struct {
	IBMService
}

func (g CMGenerator) loadCM(cmID, cmGuID string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		cmID,
		cmGuID,
		"ibm_resource_instance",
		"ibm",
		[]string{})
	return resources
}

func (g CMGenerator) loadImportedCM(cmID, certificateID, cisInstance string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		cmID,
		certificateID,
		"ibm_certificate_manager_import",
		"ibm",
		map[string]string{
			"dns_provider_instance_crn": cisInstance,
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g CMGenerator) loadOrderedCM(cmID, certificateID, cisInstance string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		cmID,
		certificateID,
		"ibm_certificate_manager_order",
		"ibm",
		map[string]string{
			"dns_provider_instance_crn": cisInstance,
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g *CMGenerator) InitResources() error {

	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	var cisInstance string
	var cisID string
	cis := g.Args["cis"]
	if cis != nil {
		cisInstance = cis.(string)
	}

	// Client creation
	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	certManagementClient, err := certificatemanager.New(sess)
	if err != nil {
		return err
	}

	// Get ServiceID of certificate manager service
	serviceID, err := catalogClient.ResourceCatalog().FindByName("cloudcerts", true)
	if err != nil {
		return err
	}

	serviceID2, err := catalogClient.ResourceCatalog().FindByName("internet-svcs", true)
	if err != nil {
		return err
	}

	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}

	query2 := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID2[0].ID,
	}

	// Get all Certificate manager instances
	cmInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}

	// Get all CIS instances
	cisInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query2)
	if err != nil {
		return err
	}
	for _, cis := range cisInstances {
		if cisInstance == cis.Name {
			cisID = cis.Guid
		}
	}

	// Get all certificates associated with a certificate manager instance
	for _, cmInstance := range cmInstances {

		g.Resources = append(g.Resources, g.loadCM(cmInstance.ID, cmInstance.Guid))

		// For each instance get associated certificates
		certificateList, err := certManagementClient.Certificate().ListCertificates(cmInstance.ID)
		if err != nil {
			return err
		}

		for _, cert := range certificateList {
			// Get certificate info
			certificatedata, err := certManagementClient.Certificate().GetCertData(cert.ID)
			if err != nil {
				return err
			}

			var dependsOn []string
			dependsOn = append(dependsOn,
				"ibm_resource_instance."+terraformutils.TfSanitize(cmInstance.Guid))

			if certificatedata.Imported {
				g.Resources = append(g.Resources, g.loadImportedCM(cert.ID, cert.ID, cisID, dependsOn))
			} else {
				g.Resources = append(g.Resources, g.loadOrderedCM(cert.ID, cert.ID, cisID, dependsOn))
			}
		}
	}

	return nil
}
