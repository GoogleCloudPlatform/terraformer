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
	kp "github.com/IBM/keyprotect-go-client"
)

type KPGenerator struct {
	IBMService
}

func (g KPGenerator) createResources(keysList []kp.Key) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, key := range keysList {
		resources = append(resources, terraformutils.NewSimpleResource(
			key.CRN,
			key.ID,
			"ibm_kp_key",
			"ibm",
			[]string{}))
	}
	return resources
}

func (g *KPGenerator) InitResources() error {
	region := envFallBack([]string{"IC_REGION"}, "us-south")
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
	client.Config.InstanceID = "e63cc5b3-2594-4e13-b1b1-e8263095ff4b"

	output, err := client.GetKeys(context.Background(), 100, 0)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output.Keys)
	return nil
}
