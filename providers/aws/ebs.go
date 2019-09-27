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

package aws

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/hashcode"

	"github.com/aws/aws-sdk-go/service/ec2"
)

var ebsAllowEmptyValues = []string{"tags."}

type EbsGenerator struct {
	AWSService
}

func (g EbsGenerator) volumeAttachmentId(device, volumeID, instanceID string) string {
	return fmt.Sprintf("vai-%d", hashcode.String(fmt.Sprintf("%s-%s-%s-", device, instanceID, volumeID)))
}

func (g *EbsGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)

	err := svc.DescribeVolumesPages(&ec2.DescribeVolumesInput{}, func(volumes *ec2.DescribeVolumesOutput, lastPage bool) bool {
		for _, volume := range volumes.Volumes {

			isRootDevice := false // Let's leave root device configuration to be done in ec2_instance resources

			for _, attachment := range volume.Attachments {
				instances, _ := svc.DescribeInstances(&ec2.DescribeInstancesInput{
					InstanceIds: []*string{attachment.InstanceId},
				})
				for _, reservation := range instances.Reservations {
					for _, instance := range reservation.Instances {
						if aws.StringValue(instance.RootDeviceName) == aws.StringValue(attachment.Device) {
							isRootDevice = true
						}
					}
				}
			}

			if !isRootDevice {
				g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
					aws.StringValue(volume.VolumeId),
					aws.StringValue(volume.VolumeId),
					"aws_ebs_volume",
					"aws",
					ebsAllowEmptyValues,
				))

				for _, attachment := range volume.Attachments {
					if aws.StringValue(attachment.State) == ec2.VolumeAttachmentStateAttached {

						attachmentId := g.volumeAttachmentId(
							aws.StringValue(attachment.Device),
							aws.StringValue(attachment.VolumeId),
							aws.StringValue(attachment.InstanceId))
						g.Resources = append(g.Resources, terraform_utils.NewResource(
							attachmentId,
							aws.StringValue(attachment.InstanceId)+":"+aws.StringValue(attachment.Device),
							"aws_volume_attachment",
							"aws",
							map[string]string{
								"device_name": aws.StringValue(attachment.Device),
								"volume_id":   aws.StringValue(attachment.VolumeId),
								"instance_id": aws.StringValue(attachment.InstanceId),
							},
							[]string{},
							map[string]interface{}{},
						))
					}
				}
			}
		}

		return !lastPage
	})
	if err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
	return nil
}
