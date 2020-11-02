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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

var workspacesAllowEmptyValues = []string{"tags."}

type WorkspacesGenerator struct {
	AWSService
}

func (g *WorkspacesGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := workspaces.New(config)
	if err := g.loadWorkspaces(svc); err != nil {
		return err
	}
	if err := g.loadWorkspacesIPGroup(svc); err != nil {
		return err
	}
	return nil
}

func (g *WorkspacesGenerator) loadWorkspaces(svc *workspaces.Client) error {
	p := workspaces.NewDescribeWorkspacesPaginator(svc.DescribeWorkspacesRequest(&workspaces.DescribeWorkspacesInput{}))
	for p.Next(context.Background()) {
		for _, workspace := range p.CurrentPage().Workspaces {
			directoryID := aws.StringValue(workspace.DirectoryId)
			workspaceID := aws.StringValue(workspace.WorkspaceId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				directoryID,
				directoryID,
				"aws_workspaces_directory",
				"aws",
				workspacesAllowEmptyValues))
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				workspaceID,
				workspaceID,
				"aws_workspaces_workspace",
				"aws",
				workspacesAllowEmptyValues))
		}
	}
	return p.Err()
}

func (g *WorkspacesGenerator) loadWorkspacesIPGroup(svc *workspaces.Client) error {
	var nextToken *string
	for {
		response, err := svc.DescribeIpGroupsRequest(&workspaces.DescribeIpGroupsInput{NextToken: nextToken}).Send(context.Background())
		if err != nil {
			return err
		}
		for _, ipGroup := range response.Result {
			groupID := aws.StringValue(ipGroup.GroupId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				groupID,
				groupID,
				"aws_workspaces_ip_group",
				"aws",
				workspacesAllowEmptyValues))
		}
		nextToken = response.NextToken
		if nextToken == nil {
			break
		}
	}
	return nil
}
