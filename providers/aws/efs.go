// Copyright 2020 The Terraformer Authors.
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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

var efsAllowEmptyValues = []string{"tags."}

type EfsGenerator struct {
	AWSService
}

func (g *EfsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := efs.New(config)
	if err := g.loadFileSystem(svc); err != nil {
		return err
	}
	if err := g.loadAccessPoint(svc); err != nil {
		return err
	}
	if err := g.loadMountTarget(svc); err != nil {
		return err
	}
	return nil
}

func (g *EfsGenerator) loadFileSystem(svc *efs.Client) error {
	p := efs.NewDescribeFileSystemsPaginator(svc.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{}))
	for p.Next(context.Background()) {
		for _, fileSystem := range p.CurrentPage().FileSystems {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				aws.StringValue(fileSystem.FileSystemId),
				aws.StringValue(fileSystem.FileSystemId),
				"aws_efs_file_system",
				"aws",
				efsAllowEmptyValues))

			targetsResponse, err := svc.DescribeMountTargetsRequest(&efs.DescribeMountTargetsInput{
				FileSystemId: fileSystem.FileSystemId,
			}).Send(context.Background())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			for _, mountTarget := range targetsResponse.MountTargets {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					aws.StringValue(mountTarget.MountTargetId),
					aws.StringValue(mountTarget.MountTargetId),
					"aws_efs_mount_target",
					"aws",
					efsAllowEmptyValues))
			}

			policyResponse, err := svc.DescribeFileSystemPolicyRequest(&efs.DescribeFileSystemPolicyInput{
				FileSystemId: fileSystem.FileSystemId,
			}).Send(context.Background())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			escapedPolicy := g.escapeAwsInterpolation(aws.StringValue(policyResponse.Policy))
			g.Resources = append(g.Resources, terraformutils.NewResource(
				aws.StringValue(fileSystem.FileSystemId),
				aws.StringValue(fileSystem.FileSystemId),
				"aws_efs_file_system_policy",
				"aws",
				map[string]string{
					"file_system_id": aws.StringValue(fileSystem.FileSystemId),
					"policy": fmt.Sprintf(`<<POLICY
%s
POLICY`, escapedPolicy),
				},
				efsAllowEmptyValues,
				map[string]interface{}{}))
		}
	}
	return p.Err()
}

func (g *EfsGenerator) loadMountTarget(svc *efs.Client) error {
	p := efs.NewDescribeFileSystemsPaginator(svc.DescribeFileSystemsRequest(&efs.DescribeFileSystemsInput{}))
	for p.Next(context.Background()) {
		for _, fileSystem := range p.CurrentPage().FileSystems {
			id := aws.StringValue(fileSystem.FileSystemId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				id,
				"aws_efs_file_system",
				"aws",
				efsAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *EfsGenerator) loadAccessPoint(svc *efs.Client) error {
	p := efs.NewDescribeAccessPointsPaginator(svc.DescribeAccessPointsRequest(&efs.DescribeAccessPointsInput{}))
	for p.Next(context.Background()) {
		for _, fileSystem := range p.CurrentPage().AccessPoints {
			id := aws.StringValue(fileSystem.AccessPointId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				id,
				id,
				"aws_efs_access_point",
				"aws",
				efsAllowEmptyValues))
		}
	}
	return p.Err()
}
