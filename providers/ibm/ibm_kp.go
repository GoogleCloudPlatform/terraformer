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
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	kp "github.com/IBM/keyprotect-go-client"
)

type KPGenerator struct {
	IBMService
}

func (g KPGenerator) loadKP(kpID, kpName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		kpID,
		kpID,
		"ibm_resource_instance",
		"ibm",
		[]string{})
	return resources
}

func (g KPGenerator) loadkPKeys(kpKeyCRN, kpKeyID string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		kpKeyCRN,
		kpKeyID,
		"ibm_kms_key",
		"ibm",
		[]string{})
	return resources
}

func (g *KPGenerator) InitResources() error {
	region := envFallBack([]string{"IC_REGION"}, "us-south")
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
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

	serviceID, err := catalogClient.ResourceCatalog().FindByName("kms", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	kpInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}
	kpurl := fmt.Sprintf("https://%s.kms.cloud.ibm.com", region)
	options := kp.ClientConfig{
		BaseURL: envFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kpurl),
		APIKey:  os.Getenv("IC_API_KEY"),
		Verbose: kp.VerboseFailOnly,
	}

	client, err := kp.New(options, kp.DefaultTransport())
	if err != nil {
		return err
	}
	for _, kp := range kpInstances {
		g.Resources = append(g.Resources, g.loadKP(kp.ID, kp.Name))
		client.Config.InstanceID = kp.Guid

		output, err := client.GetKeys(context.Background(), 100, 0)
		if err != nil {
			return err
		}
		for _, key := range output.Keys {
			g.Resources = append(g.Resources, g.loadkPKeys(key.CRN, key.ID))
		}

	}

	return nil
}
