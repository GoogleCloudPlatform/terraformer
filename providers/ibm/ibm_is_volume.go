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
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// VolumeGenerator ...
type VolumeGenerator struct {
	IBMService
}

func (g VolumeGenerator) createVolumeResources(volID, volName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		volID,
		normalizeResourceName(volName, true),
		"ibm_is_volume",
		"ibm",
		[]string{})
	return resources
}

// InitResources ...
func (g *VolumeGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	vpcurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region)
	vpcoptions := &vpcv1.VpcV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, vpcurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	start := ""
	var allrecs []vpcv1.Volume
	for {
		options := &vpcv1.ListVolumesOptions{}
		if start != "" {
			options.Start = &start
		}
		volumes, response, err := vpcclient.ListVolumes(options)
		if err != nil {
			return fmt.Errorf("Error Fetching Volumes %s\n%s", err, response)
		}
		start = GetNext(volumes.Next)
		allrecs = append(allrecs, volumes.Volumes...)
		if start == "" {
			break
		}
	}

	for _, volume := range allrecs {
		g.Resources = append(g.Resources, g.createVolumeResources(*volume.ID, *volume.Name))
	}
	return nil
}
