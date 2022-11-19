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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// InstanceGenerator ...
type InstanceGenerator struct {
	IBMService
}

func (g InstanceGenerator) createInstanceResources(instanceID, instanceName, instanceImgID string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		instanceID,
		normalizeResourceName(instanceName, true),
		"ibm_is_instance",
		"ibm",
		map[string]string{
			"image": instanceImgID,
		},
		[]string{},
		map[string]interface{}{
			"keys": []string{},
		})

	// Deprecated parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^port_speed$",
		"^primary_network_interface.[0-9].port_speed$",
		"^primary_network_interface.[0-9].primary_ip.[0-9].address$",
		"^primary_network_interface.[0-9].primary_ip.[0-9].reserved_ip$",
	)
	return resource
}

func (g InstanceGenerator) createVPCVolumeAttachmentResource(instanceID, volumeAttachedID, volumeAttachedName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", instanceID, volumeAttachedID),
		normalizeResourceName(volumeAttachedName, true),
		"ibm_is_instance_volume_attachment",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^volume$",
		"^iops$",
	)

	return resource
}

func (g InstanceGenerator) createInstanceActionResource(instanceID, instanceStatus string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		instanceID,
		normalizeResourceName(fmt.Sprintf("%s_%s", instanceID, instanceStatus), true),
		"ibm_is_instance_action",
		"ibm",
		map[string]string{
			"instance": instanceID,
			"action":   getAction(instanceStatus),
		},
		[]string{},
		map[string]interface{}{
			"force_action": false,
		})

	return resource
}

// InitResources ...
func (g *InstanceGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("No API key set")
	}

	isURL := GetVPCEndPoint(region)
	iamURL := GetAuthEndPoint()
	vpcoptions := &vpcv1.VpcV1Options{
		URL: isURL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
			URL:    iamURL,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	start := ""
	var allrecs []vpcv1.Instance
	for {
		options := &vpcv1.ListInstancesOptions{}
		if start != "" {
			options.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			options.ResourceGroupID = &rg
		}
		instances, response, err := vpcclient.ListInstances(options)
		if err != nil {
			return fmt.Errorf("Error Fetching Instances %s\n%s", err, response)
		}
		start = GetNext(instances.Next)
		allrecs = append(allrecs, instances.Instances...)
		if start == "" {
			break
		}
	}

	for _, instance := range allrecs {
		g.Resources = append(g.Resources, g.createInstanceResources(*instance.ID, *instance.Name, *instance.Image.ID))

		listVPCInsVolOptions := &vpcv1.ListInstanceVolumeAttachmentsOptions{
			InstanceID: instance.ID,
		}

		volumeAtts, response, err := vpcclient.ListInstanceVolumeAttachments(listVPCInsVolOptions)
		if err != nil {
			return fmt.Errorf("fetching vpc Instance volume Attachments %s\n%s", err, response)
		}
		allrecs := []vpcv1.VolumeAttachment{}
		allrecs = append(allrecs, volumeAtts.VolumeAttachments...)

		for _, volumeAtt := range allrecs {
			g.Resources = append(g.Resources, g.createVPCVolumeAttachmentResource(*instance.ID, *volumeAtt.ID, *volumeAtt.Name))
		}

		g.Resources = append(g.Resources, g.createInstanceActionResource(*instance.ID, *instance.Status))
	}
	return nil
}

func (g *InstanceGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_is_instance_volume_attachment" {
			continue
		}
		for _, ri := range g.Resources {
			if ri.InstanceInfo.Type != "ibm_is_instance" {
				continue
			}
			if r.InstanceState.Attributes["instance"] == ri.InstanceState.Attributes["id"] {
				g.Resources[i].Item["instance"] = "${ibm_is_instance." + ri.ResourceName + ".id}"
			}
		}
	}

	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_is_instance_action" {
			continue
		}
		for _, ri := range g.Resources {
			if ri.InstanceInfo.Type != "ibm_is_instance" {
				continue
			}
			if r.InstanceState.Attributes["instance"] == ri.InstanceState.Attributes["id"] {
				g.Resources[i].Item["instance"] = "${ibm_is_instance." + ri.ResourceName + ".id}"
			}
		}
	}

	return nil
}
